package metabase

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableMetabaseDbTable() *plugin.Table {
	return &plugin.Table{
		Name:        "metabase_db_table",
		Description: "List of tables of databases created in Metabase.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.AllColumns([]string{"db_id"}),
			Hydrate:    listDatabaseTable,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"id", "db_id"}),
			Hydrate:    getDatabaseTable,
		},
		Columns: []*plugin.Column{
			// Key column cannot be a pointer. Transform helps us to manage them
			{Name: "id", Type: proto.ColumnType_INT, Transform: transform.FromField("Id"), Description: "ID of table."},
			{Name: "db_id", Type: proto.ColumnType_INT, Transform: transform.FromField("DbId"), Description: "ID of database."},
			{Name: "entity_type", Type: proto.ColumnType_STRING, Description: "???"},
			{Name: "schema", Type: proto.ColumnType_STRING, Description: "Database schema."},
			{Name: "show_in_getting_started", Type: proto.ColumnType_BOOL, Description: "If table is show on start."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of table."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "Description of table."},
			{Name: "caveats", Type: proto.ColumnType_STRING, Description: "???"},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "When table was created in Metabase."},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Description: "When table was updated in Metabase."},
			{Name: "visibility_type", Type: proto.ColumnType_STRING, Description: "???"},
			{Name: "display_name", Type: proto.ColumnType_STRING, Description: "Display name of table."},
			{Name: "points_of_interest", Type: proto.ColumnType_STRING, Description: "???"},
		},
	}
}

func listDatabaseTable(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(d)

	if err != nil {
		plugin.Logger(ctx).Error("metabase_db_table.listDatabaseTable", "connection_error", err)
		return nil, err
	}

	quals := d.EqualsQuals
	dbId := quals["db_id"].GetInt64Value()

	request := client.DatabaseApi.GetDatabase(context.Background(), int32(dbId))
	request = request.Include("tables")

	db, resp, err := client.DatabaseApi.GetDatabaseExecute(request)

	err = manageError("metabase_db_table.listDatabaseTable", ctx, resp, err)

	if err != nil {
		return nil, err
	}

	for _, table := range db.Tables {
		d.StreamListItem(ctx, table)
	}

	return nil, nil
}

func getDatabaseTable(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(d)

	if err != nil {
		plugin.Logger(ctx).Error("metabase_db_table.getDatabaseTable", "connection_error", err)
		return nil, err
	}

	quals := d.EqualsQuals
	id := quals["id"].GetInt64Value()
	dbId := quals["db_id"].GetInt64Value()

	request := client.DatabaseApi.GetDatabase(context.Background(), int32(dbId))
	request = request.Include("tables")

	db, resp, err := client.DatabaseApi.GetDatabaseExecute(request)

	err = manageError("metabase_db_table.getDatabaseTable", ctx, resp, err)

	if err != nil {
		return nil, err
	}

	for _, table := range db.Tables {
		if *table.Id == id {
			return table, nil
		}
	}

	return nil, nil
}
