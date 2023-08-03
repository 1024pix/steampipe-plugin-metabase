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
			Hydrate: listPermissionsDownload,
		},
		Columns: SubPermissionColum,
	}
}

func listPermissionsDownload(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	return listSubPermissions("listPermissionsDownload", ctx, d, PermissionDownload)
}
