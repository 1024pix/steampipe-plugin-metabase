package metabase

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableMetabaseDb() *plugin.Table {
	return &plugin.Table{
		Name:        "metabase_db",
		Description: "List databases created in Metabase.",
		List: &plugin.ListConfig{
			Hydrate: listDatabase,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"id"}),
			Hydrate:    getDatabase,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_INT, Transform: transform.FromField("Id"), Description: "ID of database."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of database."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "Description of database."},
			//Features                 []string `json:"features,omitempty"`
			{Name: "is_full_sync", Type: proto.ColumnType_BOOL, Description: "Is fully synchronize."},
			{Name: "is_sample", Type: proto.ColumnType_BOOL, Description: "Is database is sample."},
			{Name: "cache_field_values_schedule", Type: proto.ColumnType_STRING, Description: "Cache field."},
			{Name: "metadata_sync_schedule", Type: proto.ColumnType_STRING, Description: "Synchroniez schedule value."},
			{Name: "caveats", Type: proto.ColumnType_STRING, Description: "???."},
			{Name: "engine", Type: proto.ColumnType_STRING, Description: "Engine of database."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "When database was created in Metabase."},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Description: "When database was updated in Metabase."},
			{Name: "native_permissions", Type: proto.ColumnType_STRING, Description: "????."},
			{Name: "points_of_interest", Type: proto.ColumnType_STRING, Description: "????."},
			//Tables           []DatabaseTable  `json:"tables,omitempty"`
		},
	}
}

func listDatabase(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)

	if err != nil {
		plugin.Logger(ctx).Error("metabase_db.listDatabases", "connection_error", err)
		return nil, err
	}

	request := client.DatabaseApi.ListDatabases(context.Background())

	dbList, resp, err := client.DatabaseApi.ListDatabasesExecute(request)

	if err != nil {
		plugin.Logger(ctx).Error("metabase_db.listDatabases", err)
		return nil, err
	} else if resp.StatusCode >= 300 {
		err = fmt.Errorf("HTTP code = %d", resp.StatusCode)
		plugin.Logger(ctx).Error("metabase_db.listDatabases", err)
		return nil, err
	}

	for _, db := range dbList.Data {
		d.StreamListItem(ctx, db)
	}

	return nil, nil
}

func getDatabase(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)

	if err != nil {
		plugin.Logger(ctx).Error("metabase_db.getDatabase", "connection_error", err)
		return nil, err
	}

	quals := d.EqualsQuals
	id := quals["id"].GetInt64Value()

	request := client.DatabaseApi.GetDatabase(context.Background(), int32(id))

	db, resp, err := client.DatabaseApi.GetDatabaseExecute(request)

	if err != nil {
		plugin.Logger(ctx).Error("metabase_db.getDatabase", err)
		return nil, err
	} else if resp.StatusCode >= 300 {
		err = fmt.Errorf("HTTP code = %d", resp.StatusCode)
		plugin.Logger(ctx).Error("metabase_db.getDatabase", err)
		return nil, err
	}

	return db, nil
}
