package metabase

import (
	"context"
	"fmt"

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
			{Name: "id", Type: proto.ColumnType_INT, Transform: transform.FromField("Id"), Description: "ID of groud."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of group"},
			{Name: "member_count", Type: proto.ColumnType_STRING, Description: "Number of member."},
		},
	}
}

func listPermissionGroup(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)

	if err != nil {
		plugin.Logger(ctx).Error("metabase_db.listPermissionGroup", "connection_error", err)
		return nil, err
	}

	request := client.PermissionsApi.GetPermissionsGroup(context.Background())

	permissions, resp, err := client.PermissionsApi.GetPermissionsGroupExecute(request)

	if err != nil {
		plugin.Logger(ctx).Error("metabase_db.listPermissionGroup", err)

		return nil, err
	} else if resp.StatusCode >= 300 {
		err = fmt.Errorf("HTTP code = %d", resp.StatusCode)
		plugin.Logger(ctx).Debug("metabase_db.listPermissionGroup", "http-response", resp)
		plugin.Logger(ctx).Error("metabase_db.listPermissionGroup", err)

		return nil, err
	}

	for _, permission := range permissions {
		d.StreamListItem(ctx, permission)
	}

	return nil, nil
}

func getPermissionGroup(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)

	if err != nil {
		plugin.Logger(ctx).Error("metabase_db.getPermissionGroup", "connection_error", err)
		return nil, err
	}

	quals := d.EqualsQuals
	id := quals["id"].GetInt64Value()

	request := client.PermissionsApi.GetPermissionsGroup(context.Background())

	permissions, resp, err := client.PermissionsApi.GetPermissionsGroupExecute(request)

	if err != nil {
		plugin.Logger(ctx).Error("metabase_db.getPermissionGroup", err)
		return nil, err
	} else if resp.StatusCode >= 300 {
		err = fmt.Errorf("HTTP code = %d", resp.StatusCode)
		plugin.Logger(ctx).Error("metabase_db.getPermissionGroup", err)
		return nil, err
	}

	for _, permission := range permissions {
		if int64(*permission.Id) == id {
			return permission, nil
		}
	}

	return nil, nil
}
