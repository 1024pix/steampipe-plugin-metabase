package metabase

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableMetabaseUser() *plugin.Table {
	return &plugin.Table{
		Name:        "metabase_user",
		Description: "List of users of metabase.",
		List: &plugin.ListConfig{
			Hydrate: listUsers,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"id"}),
			Hydrate:    getUser,
		},
		Columns: []*plugin.Column{
			// Key column cannot be a pointer. Transform helps us to manage them
			{Name: "id", 			Type: proto.ColumnType_INT, 	Transform: transform.FromField("Id"), 	Description: "ID of the user."},
			{Name: "email", 		Type: proto.ColumnType_STRING, 	                                     	Description: "User login email."},
			{Name: "common_name", 	Type: proto.ColumnType_STRING, 	                                     	Description: "User common name."},
			{Name: "first_name", 	Type: proto.ColumnType_STRING, 	                                     	Description: "User first name."},
			{Name: "last_name", 	Type: proto.ColumnType_STRING, 	                                     	Description: "User last name."},
			{Name: "locale", 		Type: proto.ColumnType_STRING, 	                                     	Description: "Locale string chosen by user."},
			{Name: "last_login", 	Type: proto.ColumnType_TIMESTAMP, 										Description: "Date and time of user last login."},
			{Name: "is_active", 	Type: proto.ColumnType_BOOL, 										 	Description: "Is user activated by an admin."},
			{Name: "updated_at", 	Type: proto.ColumnType_TIMESTAMP, 										Description: "Date and time of last change in user profile."},
			{Name: "date_joined", 	Type: proto.ColumnType_TIMESTAMP, 										Description: "Date and time of user creation (or first connexion ?)."},
			{Name: "group_ids", Type: proto.ColumnType_STRING, 											 	Description: "Array of group IDs (type incorrect !)."},
			{Name: "is_superuser", 	Type: proto.ColumnType_BOOL, 										 	Description: "True if user id admin."},
			{Name: "personal_collection_id", Type: proto.ColumnType_INT, 									Description: "Link with user collection."},
			{Name: "is_qbnewb", 	Type: proto.ColumnType_BOOL, 										 	Description: "Unknown usage."},
			{Name: "jwt_attributes", Type: proto.ColumnType_STRING, 									 	Description: "Unknown usage and type."},
			// {Name: "login_attributes", Type: proto.ColumnType_INT, Transform: transform.FromField("Id"), Description: "ID of the group."},
			{Name: "tenant_id", 	Type: proto.ColumnType_STRING, 										 	Description: "Unknown usage and type."},
			{Name: "sso_source", 	Type: proto.ColumnType_STRING, 										 	Description: "Name of SSO if uses an SSO."},
		},
	}
}

func listUsers(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(d)

	if err != nil {
		plugin.Logger(ctx).Error("metabase_user.listUsers", "connection_error", err)
		return nil, err
	}

	users, resp, err := client.ListUsers(context.Background())

	err = manageError("metabase_user.listUsers", ctx, resp, err)

	if err != nil {
		return nil, err
	}

	for _, user := range users.Data {
		d.StreamListItem(ctx, user)
	}

	return nil, nil
}

func getUser(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(d)

	if err != nil {
		plugin.Logger(ctx).Error("metabase_user.getUser", "connection_error", err)
		return nil, err
	}

	quals := d.EqualsQuals
	id := quals["id"].GetInt64Value()

	db, resp, err := client.GetUser(context.Background(), int32(id))

	err = manageError("metabase_user.getUser", ctx, resp, err)

	if err != nil {
		return nil, err
	}

	return db, nil
}