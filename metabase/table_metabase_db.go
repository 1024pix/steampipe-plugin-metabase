package metabase

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableMetabaseDb() *plugin.Table {
	return &plugin.Table{
		Name:        "metabase_db",
		Description: "List databases created in Metabase.",
		List: &plugin.ListConfig{
			Hydrate: listDatabases,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"id"}),
			Hydrate:    getDatabase,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_INT, Transform: transform.FromField("Id"), Description: "ID of database."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of database."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "Description of database."},
			{Name: "is_full_sync", Type: proto.ColumnType_BOOL, Description: "Is fully synchronize."},
			{Name: "is_sample", Type: proto.ColumnType_BOOL, Description: "Is database is sample."},
			{Name: "cache_field_values_schedule", Type: proto.ColumnType_STRING, Description: "Cache field."},
			{Name: "metadata_sync_schedule", Type: proto.ColumnType_STRING, Description: "Synchroniez schedule value."},
			{Name: "caveats", Type: proto.ColumnType_STRING, Description: "Warning about this table."},
			{Name: "engine", Type: proto.ColumnType_STRING, Description: "Engine of database."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "When database was created in Metabase."},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Description: "When database was updated in Metabase."},
			{Name: "native_permissions", Type: proto.ColumnType_STRING, Description: "Ability to write native/SQL."},
			{Name: "points_of_interest", Type: proto.ColumnType_STRING, Description: "Description of why this table is interest."},
		},
	}
}

func listDatabases(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(d)

	if err != nil {
		plugin.Logger(ctx).Error("metabase_db.listDatabases", "connection_error", err)
		return nil, err
	}

	request := client.DatabaseApi.ListDatabases(context.Background())

	dbList, resp, err := client.DatabaseApi.ListDatabasesExecute(request)

	err = manageError("metabase_db.listDatabases", ctx, resp, err)

	if err != nil {
		return nil, err
	}

	for _, db := range dbList.Data {
		d.StreamListItem(ctx, db)
	}

	return nil, nil
}

func getDatabase(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(d)

	if err != nil {
		plugin.Logger(ctx).Error("metabase_db.getDatabase", "connection_error", err)
		return nil, err
	}

	quals := d.EqualsQuals
	id := quals["id"].GetInt64Value()

	request := client.DatabaseApi.GetDatabase(context.Background(), int32(id))

	db, resp, err := client.DatabaseApi.GetDatabaseExecute(request)

	err = manageError("metabase_db.getDatabase", ctx, resp, err)

	if err != nil {
		return nil, err
	}

	return db, nil
}
