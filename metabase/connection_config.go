package metabase

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

type metabaseConfig struct {
	Url           *string `cty:"url"`
	Token         *string `cty:"token"`
	User          *string `cty:"user"`
	Password      *string `cty:"password"`
	TlsSkipVerify *bool   `cty:"tls_skip_verify"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"url": {
		Type: schema.TypeString,
	},
	"token": {
		Type: schema.TypeString,
	},
	"user": {
		Type: schema.TypeString,
	},
	"password": {
		Type: schema.TypeString,
	},
	"tls_skip_verify": {
		Type: schema.TypeBool,
	},
}

func ConfigInstance() interface{} {
	return &metabaseConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) metabaseConfig {
	if connection == nil || connection.Config == nil {
		return metabaseConfig{}
	}
	config, _ := connection.Config.(metabaseConfig)
	return config
}
