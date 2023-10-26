package main

import (
	personprovider "github.com/Kiril-Poposki1998/terraform_person_provider/person_provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: personprovider.Provider,
	})
}
