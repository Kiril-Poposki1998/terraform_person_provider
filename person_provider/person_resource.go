package personprovider

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/buger/jsonparser"
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
		},
	}
}

func createPerson(d *schema.ResourceData, m any) error {
	name := d.Get("name").(string)
	surname := d.Get("surname").(string)
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
	resp, err = http.Get("http://localhost:8080/api/person")
	if err != nil {
		errors.New("Cannot contact server")
	}
	resp_body, _ := io.ReadAll(resp.Body)
	jsonparser.ArrayEach(resp_body, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		name_bytes, _, _, _ := jsonparser.Get(value, "name")
		surname_bytes, _, _, _ := jsonparser.Get(value, "surname")
		id_bytes, _, _, _ := jsonparser.Get(value, "Id")
		if name == string(name_bytes) && surname == string(surname_bytes) {
			d.SetId(string(id_bytes))
		}
	})
	return nil
}

func updatePerson(d *schema.ResourceData, m any) error {
	name := d.Get("name").(string)
	surname := d.Get("surname").(string)
	postBody, _ := json.Marshal(map[string]string{
		"name":    name,
		"surname": surname,
	})
	body := bytes.NewBuffer(postBody)
	req, err := http.NewRequest("PUT", "http://localhost:8080/api/person/"+d.Id(), body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	} else if resp.StatusCode < 200 && resp.StatusCode > 300 {
		return errors.New("The server responded with a status different than 200")
	}
	return nil
}

func readPerson(d *schema.ResourceData, m any) error {
	name := d.Get("name").(string)
	surname := d.Get("surname").(string)
	resp, err := http.Get("http://localhost:8080/api/person")
	if err != nil {
		errors.New("Cannot contact server")
	}
	resp_body, _ := io.ReadAll(resp.Body)
	var resource_exists bool = false
	jsonparser.ArrayEach(resp_body, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		name_bytes, _, _, _ := jsonparser.Get(value, "name")
		surname_bytes, _, _, _ := jsonparser.Get(value, "surname")
		id_bytes, _, _, _ := jsonparser.Get(value, "Id")
		if name == string(name_bytes) && surname == string(surname_bytes) {
			d.SetId(string(id_bytes))
			resource_exists = true
		}
	})
	if resource_exists {
		return nil
	}
	d.SetId("")
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
	req, err := http.NewRequest("DELETE", "http://localhost:8080/api/person/"+d.Id(), body)
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
