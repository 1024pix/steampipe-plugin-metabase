package metabase

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableMetabasePermissionDownload() *plugin.Table {
	return &plugin.Table{
		Name:        "metabase_permission_download",
		Description: "List of permissions for download in Metabase.",
		List: &plugin.ListConfig{
			Hydrate: listPermissionDownload,
		},
		Columns: SubPermissionColum,
	}
}

func listPermissionDownload(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	return listSubPermission("listPermissionDownload", ctx, d, PermissionDownload)
}
