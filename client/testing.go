package client

import (
	"testing"

	"github.com/cloudquery/cq-provider-sdk/cqproto"
	"github.com/cloudquery/cq-provider-sdk/provider"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func DOTestHelper(t *testing.T, table *schema.Table) {
	cfg := ``
	providertest.TestResource(t, providertest.ResourceTestCase{
		Provider: &provider.Provider{
			Name:      "digitalocean_test_provider",
			Version:   "development",
			Configure: Configure,
			Config: func(f cqproto.ConfigFormat) provider.Config {
				return NewConfig(f)
			},
			ResourceMap: map[string]*schema.Table{
				"test_resource": table,
			},
		},
		Config: cfg,
	})
}
