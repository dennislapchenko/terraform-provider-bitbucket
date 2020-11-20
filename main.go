package main

import (
	"terraform-provider-bitbucket/bitbucket"

	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: bitbucket.Provider})
}
