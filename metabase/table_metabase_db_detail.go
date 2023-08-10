package metabase

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type DbDetail struct {
	DbID  int64
	Key   string
	Value string
}

func tableMetabaseDbDetail() *plugin.Table {
	return &plugin.Table{
		Name:        "metabase_db_detail",
		Description: "List of details of databases created in Metabase.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.AllColumns([]string{"db_id"}),
			Hydrate:    listDatabaseDetails,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"db_id", "key"}),
			Hydrate:    getDatabaseDetail,
		},
		Columns: []*plugin.Column{
			{Name: "db_id", Type: proto.ColumnType_INT, Description: "ID of database."},
			{Name: "key", Type: proto.ColumnType_STRING, Description: "Key of property of database."},
			{Name: "value", Type: proto.ColumnType_STRING, Description: "Value of property of database."},
		},
	}
}

func listDatabaseDetails(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(d)

	if err != nil {
		plugin.Logger(ctx).Error("metabase_db_detail.listDatabaseDetails", "connection_error", err)
		return nil, err
	}

	quals := d.EqualsQuals
	dbId := quals["db_id"].GetInt64Value()

	request := client.DatabaseApi.GetDatabase(context.Background(), int32(dbId))

	db, resp, err := client.DatabaseApi.GetDatabaseExecute(request)

	err = manageError("metabase_db_detail.listDatabaseDetails", ctx, resp, err)

	if err != nil {
		return nil, err
	}

	for key, value := range db.Details {
		d.StreamListItem(ctx, DbDetail{
			DbID:  dbId,
			Key:   key,
			Value: fmt.Sprintf("%v", value),
		})
	}

	return nil, nil
}

func getDatabaseDetail(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(d)

	if err != nil {
		plugin.Logger(ctx).Error("metabase_db_detail.getDatabaseDetail", "connection_error", err)
		return nil, err
	}

	quals := d.EqualsQuals
	dbDetailKey := quals["key"].GetStringValue()
	dbId := quals["db_id"].GetInt64Value()

	request := client.DatabaseApi.GetDatabase(context.Background(), int32(dbId))

	db, resp, err := client.DatabaseApi.GetDatabaseExecute(request)

	err = manageError("metabase_db_detail.getDatabaseDetail", ctx, resp, err)

	if err != nil {
		return nil, err
	}

	for key, value := range db.Details {
		if key == dbDetailKey {
			return DbDetail{
				DbID:  dbId,
				Key:   key,
				Value: fmt.Sprintf("%v", value),
			}, nil
		}
	}
	return nil, nil
}
