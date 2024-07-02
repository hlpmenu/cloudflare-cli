// Description: D1 API commands
// Commands
// - list
// - create
// - delete
// - query
// - raw
// - exec
// - get-id

package cfapi

import (
	"fmt"
	"go-debug/cmd/commands"
	"go-debug/env"
	"go-debug/output"
	"io"
	"log"
	"net/http"
)

func loadD1Commands() {
	// D1
	Cmds.Add(*D1MainCommand)
}

type D1 struct {
	Request CFRequest
}

// D1 Api url base
var d1baseurl string

func D1Commdand(c *CFCommand) {

	if env.CLOUDFLARE_ACCOUNT_ID == "" {
		output.Errorf("No account ID provided, \n Run 'cf_api setup' to set up your account ID, or run 'cf_api setup --config' to generate a config file\n")
		output.Exit("Exiting...")
	}
	d1baseurl = cfbaseurl + "client/v4/accounts/" + env.CLOUDFLARE_ACCOUNT_ID + "/d1/database"

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
	case "get-id":
		d1.GetID(flags)
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
	exists, name := flagExists(m, "-name")
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
	exists, id := flagExists(m, "-id")
	if !exists || len(id) < 1 {

		output.Error("No database name provided")
		output.Exit("Please provide a database name using: cf-cli d1 create -name <name>")

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
	exists, id := flagExists(m, "-id")
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
	output.Errorf("Error: Query not implemented in the CLI. Use cf-api exec -db 'id' -sql 'SELECT * ...' to run raw queries")
}

func (d *D1) Raw(m flagsMap) {

	var (
		sql, id string
		exists  bool
	)

	log.Printf("Flags: %v", m)
	for k, v := range m {
		log.Printf("Key: %s, Value: %s", k, v)
	}

	// Values
	fmt.Println(env.DB_ID)
	if env.DB_ID == "" {

		exists, id = flagExists(m, "-db")
		if !exists || len(id) < 1 {
			output.Error("No database id provided")
			output.Exit("Please provide a database id using: cf-cli d1 create -db <db_id>")
		}
	} else {
		id = env.DB_ID
	}

	exists, sql = flagExists(m, "-sql")
	if !exists || len(sql) < 1 {

		fileFlagExist, file := flagExists(m, "-file") // -file
		if len(file) < 1 && fileFlagExist {
			log.Printf("File: %s", len(file))
			output.Errorf("Flagexist %s", fileFlagExist)
			output.Error("Invalid file provided")
		} else if fileFlagExist {
			log.Printf("File: %s", file)
			sql = parseSqlFile(file)
		} else {
			output.Error("No SQL query provided")
			output.Exit("Please provide a SQL query using: cf-cli d1 create -sql 'SELECT * FROM ...' \n or use the -file flag to provide a file with the SQL query")

		}
	}

	log.Printf("\033[31mID: %s, SQL: %s\033[0m", id, sql)

	log.Printf("ID: %s, SQL: %s", id, sql)

	url := fmt.Sprintf("%s/%s/raw", d1baseurl, id)

	cf := CFRequest{
		url:    url,
		method: POST,
	}
	// Const body
	cf.body.HasBody = true
	log.Printf("sql: %s", sql)
	// data := map[string]string{
	// 	"sql": sql,
	// }

	json := []byte(fmt.Sprintf(`{"sql": "%s"}`, sql))

	// Earlier tests had this:
	/*
		Method 1:
			// Const body
			cf.body.HasBody = true
			data := map[string]string{
				"sql": sql,
			}
			json := toJson(data)
		Method 2:

			query := map[string]string{
				"sql": "SELECT email FROM users",
			}
			body, _ := json.Marshal(query)


	*/

	cf.body.Body = []byte(json)
	CreateRequest(&cf)

}

// Convinence function to execute a raw request
func (d *D1) Exec(m flagsMap) {

	d.Raw(m)

}

// Convinence

// A convenience function that returns the ID of a D1 db by its name.
// Removes the annoyance of having to run wrangler d1 list to find the ID
func (d *D1) GetID(m flagsMap) {

	// Values
	exists, name := flagExists(m, "-name")
	if !exists {
		output.Error("No database name provided")
		output.Exit("Please provide a database name using: cf-cli d1 get-id -name <name>")
	}

	cf := CFRequest{
		url:    d1baseurl,
		method: GET,
	}

	req, err := http.NewRequest(cf.method, cf.url, nil)
	if err != nil {
		output.Errorf("Error: %s", err)
	}
	res, err := c.Do(req)
	if err != nil {
		output.Errorf("Error: %s", err)
		output.Exit("Exiting...")
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()

	// Call the function to find the UUID
	id, err := findD1DatabaseUUID(body, name)
	if err != nil {
		output.Errorf("Error: %s", err)
		return
	}
	if id == "" {
		output.Errorf("Error: Database not found")
		return
	}

	output.Successf("🌟The Id for the DB: %s, is:  %s 🌟", name, id)
	// Parse body to get id

}

// Commands for D1

var D1MainCommand = &commands.Command{
	Name:        "d1",
	Description: "Interact with your D1 databases",
	SubCommands: []commands.SubCommand{
		*D1ListCommand,
		*D1CreateCommand,
		*D1DeleteCommand,
		*D1GetCommand,
		*D1ExecCommand,
	},
	Run: func(m map[string]string) {

	},
}

var D1ListCommand = &commands.SubCommand{
	Name:        "list",
	Description: "List all D1 databases",
	Flags: []commands.Flag{
		{
			Name:     "-name",
			HasValue: true,
		},
	},
	Run: func(m map[string]string) {
		D1Commdand(&CFCommand{CMD: "list", Flags: m})
	},
}

var D1CreateCommand = &commands.SubCommand{
	Name:        "create",
	Description: "Create a new D1 database",
	Flags: []commands.Flag{
		{
			Name:     "-name",
			HasValue: true,
		},
	},
	Run: func(m map[string]string) {
		D1Commdand(&CFCommand{CMD: "create", Flags: m})
	},
}
var D1DeleteCommand = &commands.SubCommand{
	Name:        "delete",
	Description: "Delete a D1 database",
	Flags: []commands.Flag{
		{
			Name:     "-id",
			HasValue: true,
		},
		// Add Id flag
	},
	Run: func(m map[string]string) {
		D1Commdand(&CFCommand{CMD: "delete", Flags: m})
	},
}

var D1GetCommand = &commands.SubCommand{
	Name:        "get",
	Description: "get a D1 database",
	Flags: []commands.Flag{
		{
			Name:     "-id",
			HasValue: true,
		},
		// Add Id flag
	},
	Run: func(m map[string]string) {
		D1Commdand(&CFCommand{CMD: "get", Flags: m})
	},
}

var D1ExecCommand = &commands.SubCommand{
	Name:        "exec",
	Description: "Exec a raw query on a D1 database",
	Flags: []commands.Flag{
		{
			Name:     "-sql",
			HasValue: true,
		},
		{
			Name:     "-db",
			HasValue: true,
		},
		{
			Name:     "-file",
			HasValue: true,
		},
		// Add Id flag
	},
	Run: func(m map[string]string) {
		log.Printf("Flags: %v", m)
		D1Commdand(&CFCommand{CMD: "exec", Flags: m})
	},
}
var D1GetIDCommand = &commands.SubCommand{
	Name:        "get-id",
	Description: "Exec a raw query on a D1 database",
	Flags: []commands.Flag{
		{
			Name:     "-name",
			HasValue: true,
		},
		// Add Id flag
	},
	Run: func(m map[string]string) {
		D1Commdand(&CFCommand{CMD: "get-id", Flags: m})
	},
}
