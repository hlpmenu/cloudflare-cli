package cfapi

import (
	"go-debug/env"
	"go-debug/output"
)

type D1 struct {
	Request CFRequest
}

type flagsMap map[string]string

// D1 Api url base
var d1baseurl string

func D1Commdand(c *CFCommand) {

	if env.CLOUDFLARE_ACCOUNT_ID == "" {
		output.Errorf("No account ID provided, \n Run 'cf_api setup' to set up your account ID, or run 'cf_api setup --config' to generate a config file\n")
		output.Exit("Exiting...")
	}
	d1baseurl = cfbaseurl + "client/v4/accounts" + env.CLOUDFLARE_ACCOUNT_ID + "/d1/database"

	d1 := D1{}
	flags := c.Flags

	switch c.CMD {
	case "list":
		d1.List(flags)
	case "create":
		d1.Create(flags)
	case "delete":
		d1.Delete(flags)
	case "query":
		d1.Query(flags)
	case "raw":
		d1.Raw(flags)
	case "exec":
		d1.Exec(flags)
	default:
		output.Error("Invalid command")
		output.Exit("Exiting...")
	}

}

func (d *D1) List(m flagsMap) {
	cf := CFRequest{
		url:    d1baseurl,
		method: GET,
	}
	CreateRequest(&cf)

}

func (d *D1) Create(m flagsMap) {

	// Values
	exists, name := exists(m, "-name")
	if !exists {
		output.Error("No database name provided")
		output.Exit("Please provide a database name using: cf-cli d1 create -name <name>")
	}

	cf := CFRequest{
		url:    d1baseurl,
		method: POST,
	}
	// Const body
	cf.body.HasBody = true
	data := map[string]string{
		"name": name,
	}
	json := toJson(data)
	cf.body.Body = json
	CreateRequest(&cf)

}

func (d *D1) Delete(m flagsMap) {
	// Values
	exists, id := exists(m, "-id")
	if !exists || len(id) < 1 {
		output.Error("No database id provided")
		output.Exit("Please provide a database id using: cf-cli d1 create -id <db_id>")
	}

	cf := CFRequest{
		url:    d1baseurl + "/" + id,
		method: DEL,
	}
	// Const body
	cf.body.HasBody = false

	CreateRequest(&cf)

}

func (d *D1) GET(m flagsMap) {

	// Values
	exists, id := exists(m, "-id")
	if !exists || len(id) < 1 {
		output.Error("No database id provided")
		output.Exit("Please provide a database id using: cf-cli d1 create -id <db_id>")
	}

	cf := CFRequest{
		url:    d1baseurl + "/" + id,
		method: DEL,
	}
	// Const body
	cf.body.HasBody = false

	CreateRequest(&cf)

}

func (d *D1) Query(m flagsMap) {
	// Not implemented in the CLI
}

func (d *D1) Raw(m flagsMap) {

}

// Convinence function to execute a raw request
func (d *D1) Exec(m flagsMap) {

	d.Raw(m)

}
