package personprovider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"server_url": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("http://localhost:8080/", nil),
			},
			"session_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"database_source": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("root:root@tcp(localhost:3306)/person?charset=utf8mb4&parseTime=True&loc=Local", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"person": resourcePerson(),
		},
		DataSourcesMap: map[string]*schema.Resource{},
	}
}
