package google

import (
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

// SharedConfigForRegion returns a common config setup needed for the sweeper
// functions for a given region
func SharedConfigForRegion(region string) (*transport_tpg.Config, error) {
	return acctest.SharedConfigForRegion(region)
}

func IsSweepableTestResource(resourceName string) bool {
	return acctest.IsSweepableTestResource(resourceName)
}
