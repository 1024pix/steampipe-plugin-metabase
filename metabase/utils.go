package metabase

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/1024pix/go-metabase/metabase"
	"github.com/1024pix/go-metabase/metabaseutil"
	goauth "github.com/grokify/goauth/metabase"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func connect(d *plugin.QueryData) (*metabase.APIClient, error) {
	// get metabase client from cache
	cacheKey := "metabase"
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(*metabase.APIClient), nil
	}

	token := os.Getenv("METABASE_TOKEN")
	user := os.Getenv("METABASE_USER")
	password := os.Getenv("METABASE_PASSWORD")

	url := ""
	tLSSkipVerify := false

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

	if metabaseConfig.Url != nil {
		url = *metabaseConfig.Url
	}

	if metabaseConfig.TlsSkipVerify != nil {
		tLSSkipVerify = *metabaseConfig.TlsSkipVerify
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
		BaseURL:       url,
		SessionID:     token,
		Username:      user,
		Password:      password,
		TLSSkipVerify: tLSSkipVerify,
	}

	client, _, err := metabaseutil.NewApiClient(config)

	if err != nil {
		return nil, err
	}

	// Save to cache
	d.ConnectionManager.Cache.Set(cacheKey, client)

	return client, nil
}

func manageError(methodCallStack string, ctx context.Context, resp *http.Response, err error) error {
	if err != nil {
		plugin.Logger(ctx).Error("metabase_db_xxx.manageError", err)

		return err
	} else if resp.StatusCode >= 300 {
		err = fmt.Errorf("HTTP code = %d", resp.StatusCode)
		plugin.Logger(ctx).Debug(fmt.Sprintf("%s.manageError", methodCallStack), "http-response", resp)
		plugin.Logger(ctx).Error(fmt.Sprintf("%s.manageError", methodCallStack), err)

		return err
	}

	return nil
}
