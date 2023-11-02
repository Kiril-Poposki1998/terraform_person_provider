package personprovider

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePerson() *schema.Resource {
	return &schema.Resource{
		Create: createPerson,
		Update: updatePerson,
		Read:   readPerson,
		Delete: deletePerson,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Id for the person in the database",
			},
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
		},
	}
}

func createPerson(d *schema.ResourceData, m any) error {
	name := d.Get("name").(string)
	surname := d.Get("surname").(string)
	id := d.Get("id").(int)
	postBody, _ := json.Marshal(map[string]string{
		"name":    name,
		"surname": surname,
	})
	body := bytes.NewBuffer(postBody)
	resp, err := http.Post("http://localhost:8080/api/person", "application/json", body)
	if err != nil {
		return err
	} else if resp.StatusCode != 200 {
		return errors.New("The server responsed with a status different than 200")
	}
	d.SetId(strconv.Itoa(id))
	return nil
}

func updatePerson(d *schema.ResourceData, m any) error {
	name := d.Get("name").(string)
	surname := d.Get("surname").(string)
	id := d.Get("id").(int)
	postBody, _ := json.Marshal(map[string]string{
		"name":    name,
		"surname": surname,
	})
	body := bytes.NewBuffer(postBody)
	resp, err := http.Post("http://localhost:8080/api/person/"+d.Id(), "application/json", body)
	if err != nil {
		return err
	} else if resp.StatusCode != 200 {
		return errors.New("The server responsed with a status different than 200")
	}
	d.SetId(strconv.Itoa(id))
	return nil
}

func readPerson(d *schema.ResourceData, m any) error {
	return nil
}

func deletePerson(d *schema.ResourceData, m any) error {
	name := d.Get("name").(string)
	surname := d.Get("surname").(string)
	postBody, _ := json.Marshal(map[string]string{
		"name":    name,
		"surname": surname,
	})
	body := bytes.NewBuffer(postBody)
	id := d.Id()
	resp, err := http.NewRequest("DELETE", "http://localhost:8080/api/person/"+id, body)
	if err != nil {
		return err
	} else if resp.Response.StatusCode != 200 {
		return errors.New("The server responsed with a status different than 200")
	}
	return nil
}
