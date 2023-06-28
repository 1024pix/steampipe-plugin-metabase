package metabase

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

type DbFeature struct {
	Id      int64  `json:"db_id"`
	Feature string `json:"feature"`
}

func tableMetabaseDbFeature() *plugin.Table {
	return &plugin.Table{
		Name:        "metabase_db_feature",
		Description: "Features of one database.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.AllColumns([]string{"db_id"}),
			Hydrate:    listDatabaseFeature,
		},

		Columns: []*plugin.Column{
			{Name: "db_id", Type: proto.ColumnType_INT, Transform: transform.FromQual("db_id"), Description: "ID of the database."},
			{Name: "feature", Type: proto.ColumnType_STRING, Description: "Feature of the database."},
		},
	}
}

func listDatabaseFeature(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(d)

	if err != nil {
		plugin.Logger(ctx).Error("metabase_db_feature.listDatabaseFeature", "connection_error", err)
		return nil, err
	}

	quals := d.EqualsQuals
	dbId := quals["db_id"].GetInt64Value()

	request := client.DatabaseApi.GetDatabase(context.Background(), int32(dbId))

	db, resp, err := client.DatabaseApi.GetDatabaseExecute(request)

	err = manageError("metabase_db_feature.listDatabaseFeature", ctx, resp, err)

	if err != nil {
		return nil, err
	}

	for _, feature := range db.Features {
		d.StreamListItem(ctx, DbFeature{
			Id:      db.Id,
			Feature: feature,
		})
	}

	return nil, nil
}
