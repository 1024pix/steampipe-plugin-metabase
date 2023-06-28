package metabase

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableMetabaseGroup() *plugin.Table {
	return &plugin.Table{
		Name:        "metabase_db_group",
		Description: "List of group created in Metabase.",
		List: &plugin.ListConfig{
			Hydrate: listDatabaseGroup,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"id"}),
			Hydrate:    getDatabaseGroup,
		},
		Columns: []*plugin.Column{
			// Key column cannot be a pointer. Transform helps us to manage them
			{Name: "id", Type: proto.ColumnType_INT, Transform: transform.FromField("Id"), Description: "ID of groud."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of group"},
			{Name: "member_count", Type: proto.ColumnType_STRING, Description: "Number of member."},
		},
	}
}

func listDatabaseGroup(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)

	if err != nil {
		plugin.Logger(ctx).Error("metabase_db.listDatabaseGroup", "connection_error", err)
		return nil, err
	}

	quals := d.EqualsQuals
	dbId := quals["db_id"].GetInt64Value()

	request := client.DatabaseApi.GetDatabase(context.Background(), int32(dbId))
	request = request.Include("tables")

	db, resp, err := client.DatabaseApi.GetDatabaseExecute(request)

	if err != nil {
		plugin.Logger(ctx).Error("metabase_db.listDatabaseGroup", err)

		return nil, err
	} else if resp.StatusCode >= 300 {
		err = fmt.Errorf("HTTP code = %d", resp.StatusCode)
		plugin.Logger(ctx).Debug("metabase_db.listDatabaseGroup", "http-response", resp)
		plugin.Logger(ctx).Error("metabase_db.listDatabaseGroup", err)

		return nil, err
	}

	for _, table := range db.Tables {
		d.StreamListItem(ctx, table)
	}

	return nil, nil
}

func getDatabaseGroup(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)

	if err != nil {
		plugin.Logger(ctx).Error("metabase_db.getDatabaseGroup", "connection_error", err)
		return nil, err
	}

	quals := d.EqualsQuals
	id := quals["id"].GetInt64Value()
	dbId := quals["db_id"].GetInt64Value()

	request := client.DatabaseApi.GetDatabase(context.Background(), int32(dbId))
	request = request.Include("tables")

	db, resp, err := client.DatabaseApi.GetDatabaseExecute(request)

	if err != nil {
		plugin.Logger(ctx).Error("metabase_db.getDatabaseGroup", err)
		return nil, err
	} else if resp.StatusCode >= 300 {
		err = fmt.Errorf("HTTP code = %d", resp.StatusCode)
		plugin.Logger(ctx).Error("metabase_db.getDatabaseGroup", err)
		return nil, err
	}

	plugin.Logger(ctx).Error("metabase_db.getDatabaseGroup", "coucouc")

	for _, table := range db.Tables {
		if *table.Id == id {
			return table, nil
		}
	}

	return nil, nil
}
