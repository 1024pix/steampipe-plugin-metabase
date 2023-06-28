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

type PermissionType int

const (
	PermissionData PermissionType = iota
	PermissionDownload
)

var SubPermissionColum = []*plugin.Column{
	// Key column cannot be a pointer. Transform helps us to manage them
	{Name: "group_id", Type: proto.ColumnType_INT, Transform: transform.FromField("GroupId"), Description: "ID of group."},
	{Name: "db_id", Type: proto.ColumnType_INT, Transform: transform.FromField("DbId"), Description: "ID of database"},
	{Name: "schema_name", Type: proto.ColumnType_STRING, Transform: transform.FromField("SchemaName"), Description: "Name of schema."},
	{Name: "table_id", Type: proto.ColumnType_INT, Transform: transform.FromField("TableId"), Description: "Table id."},
	{Name: "level_access", Type: proto.ColumnType_STRING, Transform: transform.FromField("LevelAccess"), Description: "Level access of table."},
}

type PermissionSchema struct {
	GroupId     int
	DbId        int
	SchemaName  string
	TableId     int
	LevelAccess string
}

func tableMetabasePermissionData() *plugin.Table {
	return &plugin.Table{
		Name:        "metabase_permission_data",
		Description: "List of permissions for data in Metabase.",
		List: &plugin.ListConfig{
			Hydrate: listPermissionData,
		},
		Columns: SubPermissionColum,
	}
}

func listPermissionData(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	return listSubPermission("listPermissionData", ctx, d, PermissionData)
}

func listSubPermission(methodCallStack string, ctx context.Context, d *plugin.QueryData, permissionType PermissionType) (interface{}, error) {
	client, err := connect(d)
	methodCallStack = fmt.Sprintf("%s.listSubPermission", methodCallStack)

	if err != nil {
		plugin.Logger(ctx).Error(methodCallStack, "connection_error", err)
		return nil, err
	}

	request := client.PermissionsApi.GetPermissionsGraph(context.Background())

	permission, resp, err := client.PermissionsApi.GetPermissionsGraphExecute(request)

	err = manageError(methodCallStack, ctx, resp, err)

	if err != nil {
		return nil, err
	}

	permissions, err := createPermissionData(methodCallStack, ctx, permission.Groups, permissionType)

	if err == nil {
		for _, perm := range permissions {
			d.StreamListItem(ctx, perm)
		}
	}

	return nil, nil
}

func createPermissionData(methodCallStack string, ctx context.Context, groups *map[string]map[string]metabase.PermissionGraphData, permissionType PermissionType) ([]PermissionSchema, error) {
	var permissions []PermissionSchema
	methodCallStack = fmt.Sprintf("%s.createPermissionData", methodCallStack)

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

			var d *map[string]map[string]interface{}

			switch permissionType {
			case PermissionData:
				d = extractSchemasAndTable(methodCallStack, ctx, data.Data)
			default:
				d = extractSchemasAndTable(methodCallStack, ctx, data.Download)
			}

			if d == nil {
				continue
			}

			perms, err := extractGranularPermissions(methodCallStack, ctx, gId, dId, d)

			if err != nil {
				return nil, err
			}

			permissions = append(permissions, perms...)
		}
	}

	return permissions, nil
}

func extractSchemasAndTable(methodCallStack string, ctx context.Context, data map[string]interface{}) *map[string]map[string]interface{} {
	var result *map[string]map[string]interface{}
	methodCallStack = fmt.Sprintf("%s.extractSchemasAndTable", methodCallStack)

	ds, ok := data["schemas"]

	if ok && ds != nil {
		switch v := ds.(type) {
		case string:
			result = nil
		default:
			theRootMap, ok := v.(map[string]interface{})

			if plugin.Logger(ctx).IsDebug() {
				plugin.Logger(ctx).Error(methodCallStack, fmt.Sprintf("ok: %t value: %+v", ok, theRootMap))
			}

			if ok {
				convertMap := make(map[string]map[string]interface{})

				for key, value := range theRootMap {
					theChildMap, ok := value.(map[string]interface{})

					if ok {
						convertMap[key] = theChildMap
					}
				}

				if plugin.Logger(ctx).IsDebug() {
					plugin.Logger(ctx).Debug(methodCallStack, fmt.Sprintf("convert map: %+v", convertMap))
				}

				result = &convertMap
			}
		}
	}

	return result
}

func extractGranularPermissions(methodCallStack string, ctx context.Context, groupId int, dbId int, data *map[string]map[string]interface{}) ([]PermissionSchema, error) {
	var permissions []PermissionSchema
	methodCallStack = fmt.Sprintf("%s.extractSchemasAndTable", methodCallStack)

	for schemaName, tableList := range *data {
		for tableId, levelAccess := range tableList {
			tId, errTID := strconv.Atoi(tableId)

			if errTID != nil {
				err := fmt.Errorf("Table id is not a integer '%s'", tableId)
				plugin.Logger(ctx).Error(methodCallStack, err)
				return nil, err
			}

			var la string

			switch v := levelAccess.(type) {
			case string:
				la = fmt.Sprint(v)
			default:
				la = ""
			}

			permissions = append(permissions, PermissionSchema{
				GroupId:     groupId,
				DbId:        dbId,
				SchemaName:  schemaName,
				TableId:     tId,
				LevelAccess: la,
			})
		}
	}

	return permissions, nil
}
