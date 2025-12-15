package metabase

import (
	"context"

	"github.com/1024pix/go-metabase/metabase"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableMetabaseGroup() *plugin.Table {
	return &plugin.Table{
		Name:        "metabase_permission_group",
		Description: "List of group created in Metabase.",
		List: &plugin.ListConfig{
			Hydrate: listPermissionsGroup,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"id"}),
			Hydrate:    getPermissionGroup,
		},
		Columns: []*plugin.Column{
			// Key column cannot be a pointer. Transform helps us to manage them
			{Name: "id", Type: proto.ColumnType_INT, Transform: transform.FromField("Id"), Description: "ID of the group."},
			{Name: "member_count", Type: proto.ColumnType_STRING, Description: "Number of members."},
			{Name: "members", Type: proto.ColumnType_JSON, Hydrate: hydrateGroupMembers, Transform: transform.FromValue(), Description: "List of users in the group."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the group."},
		},
	}
}

func listPermissionsGroup(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(d)

	if err != nil {
		plugin.Logger(ctx).Error("metabase_permission_group.listPermissionsGroup", "connection_error", err)
		return nil, err
	}

	request := client.PermissionsApi.ListPermissionsGroups(context.Background())

	permissions, resp, err := client.PermissionsApi.ListPermissionsGroupsExecute(request)

	err = manageError("metabase_permission_group.listPermissionsGroup", ctx, resp, err)

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
	id := int32(quals["id"].GetInt64Value())

	request := client.PermissionsApi.GetPermissionsGroup(context.Background(), id)

	permissions, resp, err := client.PermissionsApi.GetPermissionsGroupExecute(request)

	err = manageError("metabase_permission_group.getPermissionGroup", ctx, resp, err)

	if err != nil {
		return nil, err
	}

	return permissions, nil
}

func hydrateGroupMembers(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	group := h.Item.(metabase.PermissionGroup)

	if group.Id == nil {
		return nil, nil
	}

	client, err := connect(d)
	if err != nil {
		plugin.Logger(ctx).Error("metabase_permission_group.hydrateGroupMembers", "connection_error", err)
		return nil, err
	}

	request := client.PermissionsApi.GetPermissionsGroup(context.Background(), *group.Id)

	permissionsGroup, resp, err := client.PermissionsApi.GetPermissionsGroupExecute(request)

	err = manageError("metabase_permission_group.hydrateGroupMembers", ctx, resp, err)

	if err != nil {
		return nil, err
	}

	return permissionsGroup.Members, nil
}
