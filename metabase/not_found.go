package metabase

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func isNotFoundError(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData, err error) bool {
	_, ok := err.(*NotFoundError)

	return ok
}
