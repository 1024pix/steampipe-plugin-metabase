package metabase

import (
	"context"
	"errors"
	"os"

	"github.com/1024pix/go-metabase/metabase"
	"github.com/1024pix/go-metabase/metabaseutil"
	goauth "github.com/grokify/goauth/metabase"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func connect(ctx context.Context, d *plugin.QueryData) (*metabase.APIClient, error) {
	// get metabase client from cache
	cacheKey := "metabase"
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(*metabase.APIClient), nil
	}

	token := os.Getenv("METABASE_TOKEN")
	user := os.Getenv("METABASE_USER")
	password := os.Getenv("METABASE_PASSWORD")

	metabaseConfig := GetConfig(d.Connection)

	if metabaseConfig.Token != nil {
		token = *metabaseConfig.Token
	}

	if metabaseConfig.User != nil {
		user = *metabaseConfig.User
	}

	if metabaseConfig.Password != nil {
		password = *metabaseConfig.Password
	}

	if len(token) == 0 && len(user) == 0 && len(password) == 0 {
		return nil, errors.New("'token' or 'user/password' must be set in the connection configuration. Edit your connection configuration file or set the METABASE_TOKEN or METABASE_USER/METABASE_PASSWORD environment variable and then restart Steampipe")
	} else if len(token) == 0 &&
		(len(user) == 0 && len(password) != 0) {
		return nil, errors.New("'user' must be set in the connection configuration. Edit your connection configuration file or set the METABASE_USER environment variable and then restart Steampipe")
	} else if len(token) == 0 &&
		(len(user) != 0 && len(password) == 0) {
		return nil, errors.New("'password' must be set in the connection configuration. Edit your connection configuration file or set the METABASE_PASSWORD environment variable and then restart Steampipe")
	}

	config := goauth.Config{
		BaseURL:       *metabaseConfig.Url,
		SessionID:     token,
		Username:      user,
		Password:      password,
		TLSSkipVerify: *metabaseConfig.TlsSkipVerify,
	}

	client, _, err := metabaseutil.NewApiClient(config)

	if err != nil {
		return nil, err
	}

	// Save to cache
	d.ConnectionManager.Cache.Set(cacheKey, client)

	return client, nil
}
