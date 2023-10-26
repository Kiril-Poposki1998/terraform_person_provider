package personprovider

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePerson() *schema.Resource {
	return &schema.Resource{
		Create: createPerson,
		Update: updatePerson,
		Read:   readPerson,
		Delete: deletePerson,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the person",
			},
			"surname": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Surname of the person",
			},
			"SSID": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Social security number of the person (please don't enter a valid one)",
			},
		},
	}
}

func createPerson(d *schema.ResourceData, m any) error {
	name := d.Get("name").(string)
	surname := d.Get("surname").(string)
	url := d.Get("server_url").(string)
	postBody, _ := json.Marshal(map[string]string{
		"name":    name,
		"surname": surname,
	})
	body := bytes.NewBuffer(postBody)
	resp, err := http.Post(url+"/api/person", "application/json", body)
	if err != nil {
		return err
	} else if resp.StatusCode != 200 {
		return errors.New("The server responsed with a status different than 200")
	}
	return nil
}

func updatePerson(d *schema.ResourceData, m any) error {
	// name := d.Get("name").(string)
	// surname := d.Get("surname").(string)

	return nil
}

func readPerson(d *schema.ResourceData, m any) error {
	// name := d.Get("name").(string)
	// surname := d.Get("surname").(string)
	return nil
}

func deletePerson(d *schema.ResourceData, m any) error {
	// name := d.Get("name").(string)
	// surname := d.Get("surname").(string)
	return nil
}
