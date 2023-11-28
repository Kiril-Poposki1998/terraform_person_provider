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
			"person_id": {
				Type:        schema.TypeString,
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
	id := d.Get("person_id").(string)
	postBody, _ := json.Marshal(map[string]string{
		"name":    name,
		"surname": surname,
	})
	body := bytes.NewBuffer(postBody)
	resp, err := http.Post("http://localhost:8080/api/person", "application/json", body)
	if err != nil {
		return err
	} else if resp.StatusCode != 200 {
		return errors.New("The server responded with a status different than 200")
	}
	d.SetId(id)
	return nil
}

func updatePerson(d *schema.ResourceData, m any) error {
	name := d.Get("name").(string)
	surname := d.Get("surname").(string)
	id := d.Get("person_id").(string)
	postBody, _ := json.Marshal(map[string]string{
		"name":    name,
		"surname": surname,
	})
	body := bytes.NewBuffer(postBody)
	req, err := http.NewRequest("PUT", "http://localhost:8080/api/person/"+id, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
    resp, err := client.Do(req)
	if err != nil {
		return err
	} else if resp.StatusCode < 200 && resp.StatusCode > 300{
		return errors.New("The server responded with a status different than 200")
	}
	d.SetId(id)
	return nil
}

func readPerson(d *schema.ResourceData, m any) error {
	return nil
}

func deletePerson(d *schema.ResourceData, m any) error {
	name := d.Get("name").(string)
	surname := d.Get("surname").(string)
	id := d.Get("person_id").(string)
	postBody, _ := json.Marshal(map[string]string{
		"name":    name,
		"surname": surname,
	})
	body := bytes.NewBuffer(postBody)
	req, err := http.NewRequest("DELETE", "http://localhost:8080/api/person/"+id, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
    resp, err := client.Do(req)
	if err != nil {
		return err
	} else if resp.StatusCode != 200 {
		return errors.New("The server responded with a status different than 200")
	}
	return nil
}
