package metabase

import (
	"context"
	"fmt"
	"strconv"

	"github.com/1024pix/go-metabase/metabase"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

type Permission struct {
	GroupId        int
	DbId           int
	DownloadNative *string
	DownloadSchema *string
	DataNative     *string
	DataSchema     *string
}

func tableMetabasePermission() *plugin.Table {
	return &plugin.Table{
		Name:        "metabase_permission",
		Description: "List of permissions in Metabase.",
		List: &plugin.ListConfig{
			Hydrate: listPermission,
		},
		Columns: []*plugin.Column{
			// Key column cannot be a pointer. Transform helps us to manage them
			{Name: "group_id", Type: proto.ColumnType_INT, Transform: transform.FromField("GroupId"), Description: "ID of the group."},
			{Name: "db_id", Type: proto.ColumnType_INT, Transform: transform.FromField("DbId"), Description: "ID of the database."},
			{Name: "download_native", Type: proto.ColumnType_STRING, Transform: transform.FromField("DownloadNative"), Description: "Type of download."},
			{Name: "download_schema", Type: proto.ColumnType_STRING, Transform: transform.FromField("DownloadSchema"), Description: "Schema that you can download."},
			{Name: "data_native", Type: proto.ColumnType_STRING, Transform: transform.FromField("DataNative"), Description: "Type of data."},
			{Name: "data_schema", Type: proto.ColumnType_STRING, Transform: transform.FromField("DataSchema"), Description: "Data that you can download."},
		},
	}
}

func listPermission(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(d)

	if err != nil {
		plugin.Logger(ctx).Error("metabase_permission.listPermission", "connection_error", err)
		return nil, err
	}

	request := client.PermissionsApi.GetPermissionsGraph(context.Background())

	permission, resp, err := client.PermissionsApi.GetPermissionsGraphExecute(request)

	err = manageError("metabase_permission.listPermission", ctx, resp, err)

	if err != nil {
		return nil, err
	}

	permissions, err := createPermission("metabase_permission.listPermission", ctx, permission.Groups)

	if err == nil {
		for _, perm := range permissions {
			d.StreamListItem(ctx, perm)
		}
	}

	return nil, nil
}

func createPermission(methodCallStack string, ctx context.Context, groups *map[string]map[string]metabase.PermissionGraphData) ([]Permission, error) {
	var permissions []Permission
	methodCallStack = fmt.Sprintf("%s.extractSchemasAndTable", methodCallStack)

	for groupId, database := range *groups {
		if plugin.Logger(ctx).IsDebug() {
			plugin.Logger(ctx).Debug(methodCallStack, "response", fmt.Sprintf("Group id = %s, data = %+v", groupId, database))
		}

		gId, errGID := strconv.Atoi(groupId)

		if errGID != nil {
			err := fmt.Errorf("Group id is not a integer '%s'", groupId)
			plugin.Logger(ctx).Error(methodCallStack, err)
			return nil, err
		}

		// If some rights are set then, there are removed. Api return nil.
		if database == nil {
			continue
		}

		for databaseId, data := range database {
			dId, errDID := strconv.Atoi(databaseId)

			if errDID != nil {
				err := fmt.Errorf("Database id is not a integer '%s'", databaseId)
				plugin.Logger(ctx).Error(methodCallStack, err)
				return nil, err
			}

			downloadNative, downloadSchema := extractNativeAndSchemas(data.Download)
			dataNative, dataSchema := extractNativeAndSchemas(data.Data)

			permissions = append(permissions, Permission{
				GroupId:        gId,
				DbId:           dId,
				DownloadNative: downloadNative,
				DownloadSchema: downloadSchema,
				DataNative:     dataNative,
				DataSchema:     dataSchema,
			})
		}
	}

	return permissions, nil
}

func extractNativeAndSchemas(data map[string]interface{}) (*string, *string) {
	var (
		native *string
		schema *string
	)

	dn, ok := data["native"]

	if ok {
		tmp := fmt.Sprint(dn)
		native = &tmp
	}

	ds, ok := data["schemas"]

	if ok && ds != nil {
		switch v := ds.(type) {
		case string:
			tmp := fmt.Sprint(v)
			schema = &tmp
		default:
			tmp := "limited"
			schema = &tmp
		}
	}

	return native, schema
}
