package coredns

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"coredns": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ terraform.ResourceProvider = Provider()
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("COREDNS_ETCD_ENDPOINTS"); v == "" {
		t.Fatal("COREDNS_ETCD_ENDPOINTS must be set for acceptance tests")
	}

	if v := os.Getenv("COREDNS_ZONES"); v == "" {
		t.Fatal("COREDNS_ZONES must be set for acceptance tests")
	}
}
