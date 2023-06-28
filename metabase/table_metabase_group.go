package metabase

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableMetabaseGroup() *plugin.Table {
	return &plugin.Table{
		Name:        "metabase_permission_group",
		Description: "List of group created in Metabase.",
		List: &plugin.ListConfig{
			Hydrate: listPermissionGroup,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"id"}),
			Hydrate:    getPermissionGroup,
		},
		Columns: []*plugin.Column{
			// Key column cannot be a pointer. Transform helps us to manage them
			{Name: "id", Type: proto.ColumnType_INT, Transform: transform.FromField("Id"), Description: "ID of group."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of group"},
			{Name: "member_count", Type: proto.ColumnType_STRING, Description: "Number of member."},
		},
	}
}

func listPermissionGroup(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(d)

	if err != nil {
		plugin.Logger(ctx).Error("metabase_permission_group.listPermissionGroup", "connection_error", err)
		return nil, err
	}

	request := client.PermissionsApi.GetPermissionsGroup(context.Background())

	permissions, resp, err := client.PermissionsApi.GetPermissionsGroupExecute(request)

	err = manageError("metabase_permission_group.listPermissionGroup", ctx, resp, err)

	if err != nil {
		return nil, err
	}

	for _, permission := range permissions {
		d.StreamListItem(ctx, permission)
	}

	return nil, nil
}

func getPermissionGroup(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(d)

	if err != nil {
		plugin.Logger(ctx).Error("metabase_permission_group.getPermissionGroup", "connection_error", err)
		return nil, err
	}

	quals := d.EqualsQuals
	id := quals["id"].GetInt64Value()

	request := client.PermissionsApi.GetPermissionsGroup(context.Background())

	permissions, resp, err := client.PermissionsApi.GetPermissionsGroupExecute(request)

	err = manageError("metabase_permission_group.getPermissionGroup", ctx, resp, err)

	if err != nil {
		return nil, err
	}

	for _, permission := range permissions {
		if int64(*permission.Id) == id {
			return permission, nil
		}
	}

	return nil, nil
}
