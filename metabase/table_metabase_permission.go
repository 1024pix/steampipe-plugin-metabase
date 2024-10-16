package metabase

import (
	"context"
	"fmt"
	"strconv"
	"encoding/json"

  go_kit "github.com/turbot/go-kit/types"
	"github.com/1024pix/go-metabase/metabase"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

type Permission struct {
	GroupID       int
	DbId          int
	ViewData      *string
	CreateQueries interface{}
}

func tableMetabasePermission() *plugin.Table {
	return &plugin.Table{
		Name:        "metabase_permission",
		Description: "List of permissions in Metabase.",
		List: &plugin.ListConfig{
			Hydrate: listPermissions,
		},
		Columns: []*plugin.Column{
			{Name: "db_id", Type: proto.ColumnType_INT, Transform: transform.FromField("DbId"), Description: "ID of the database."},
			{Name: "group_id", Type: proto.ColumnType_INT, Description: "ID of the group."},
			{Name: "view_data", Type: proto.ColumnType_STRING, Transform: transform.FromField("ViewData"), Description: "Permission that determines what data people can see."},
			{Name: "create_queries", Type: proto.ColumnType_JSON, Transform: transform.FromField("CreateQueries"), Description: "Permission that specifies whether people can create new questions."},
		},
	}
}

func listPermissions(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(d)

	if err != nil {
		plugin.Logger(ctx).Error("metabase_permission.listPermissions", "connection_error", err)
		return nil, err
	}

	request := client.PermissionsApi.GetPermissionsGraph(context.Background())

	permission, resp, err := client.PermissionsApi.GetPermissionsGraphExecute(request)

	err = manageError("metabase_permission.listPermissions", ctx, resp, err)

	if err != nil {
		return nil, err
	}

	permissions, err := createPermission("metabase_permission.listPermissions", ctx, permission.Groups)

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

			var createQueries interface{}
			err := json.Unmarshal([]byte(go_kit.SafeString(data.CreateQueries)), &createQueries)
      if err != nil {
        return nil, err
      }

			permissions = append(permissions, Permission{
				GroupID:       gId,
				DbId:          dId,
				ViewData:      data.ViewData,
				CreateQueries: createQueries,
			})
		}
	}

	return permissions, nil
}
