package metabase

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             "steampipe-plugin-metabase",
		DefaultTransform: transform.FromGo().NullIfZero(),
		// DefaultIgnoreConfig: &plugin.IgnoreConfig{
		// 	ShouldIgnoreErrorFunc: isNotFoundError,
		// },
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		TableMap: map[string]*plugin.Table{
			"metabase_db":                  tableMetabaseDb(),
			"metabase_db_feature":          tableMetabaseDbFeature(),
			"metabase_db_table":            tableMetabaseDbTable(),
			"metabase_db_detail":           tableMetabaseDbDetail(),
			"metabase_permission_group":    tableMetabaseGroup(),
			"metabase_permission":          tableMetabasePermission(),
			"metabase_permission_data":     tableMetabasePermissionData(),
			"metabase_permission_download": tableMetabasePermissionDownload(),
		},
	}
	return p
}
