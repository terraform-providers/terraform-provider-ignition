package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/terraform-providers/terraform-provider-ignition/ignition"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: ignition.Provider})
}
