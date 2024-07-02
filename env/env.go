package env

import (
	"fmt"
	"go-debug/output"
	"log"
	"os"
	"regexp"
	"runtime"
	"strings"
)

var allOptions = []string{"CLOUDFLARE_ACCOUNT_EMAIL", "CLOUDFLARE_ACCOUNT_ID", "CLOUDFLARE_API_KEY", "DB_NAME", "DB_ID", "PAGES_NAME", "PAGES_ID", "WORKERS_NAME", "WORKERS_ID"}

// Remove, used to avoid rewriting the usage of non capitalized var
var AllCFVars *[]string

// Main CF API Details
var (
	CLOUDFLARE_ACCOUNT_EMAIL string // Your accounts email
	CLOUDFLARE_ACCOUNT_ID    string // Your accounts ID
	CLOUDFLARE_API_KEY       string // Your accounts API KEY
	CLOUDFLARE_API_TOKEN     string // Your accounts API TOKEN
)

// For D1
var (
	DB_NAME string
	DB_ID   string
)

// For Pages
var (
	PAGES_NAME string
	PAGES_ID   string
)

// For Workers
var (
	WORKERS_NAME string
	WORKERS_ID   string
)

// UseEnv is a flag to use the environment variables
var UseEnv bool

func SetupEnv() {

	// Temp to avoid having to set up the config file or env variables
	AllCFVars = &allOptions
	lazyness()

	// Temp
	UseEnv = true

	if UseEnv {
		CLOUDFLARE_ACCOUNT_EMAIL = GetEnv("CLOUDFLARE_ACCOUNT_EMAIL")
		CLOUDFLARE_ACCOUNT_ID = GetEnv("CLOUDFLARE_ACCOUNT_ID")
		CLOUDFLARE_API_KEY = GetEnv("CLOUDFLARE_API_KEY")
		CLOUDFLARE_API_TOKEN = GetEnv("CLOUDFLARE_API_TOKEN")
		DB_NAME = GetEnv("DB_NAME")
		DB_ID = GetEnv("DB_ID")
		PAGES_NAME = GetEnv("PAGES_NAME")
		PAGES_ID = GetEnv("PAGES_ID")
		WORKERS_NAME = GetEnv("WORKERS_NAME")
		WORKERS_ID = GetEnv("WORKERS_ID")

	} else {
		SetupEnvConfig()
	}

}

func GetEnv(key string) string {
	value := os.Getenv(key)

	log.Printf("Getting env: %s\n, with value %s", key, value)

	return value
}

func SetupEnvConfig() {

	// Check if the config file exists
	_, err := os.Stat("config.conf")
	if err != nil {
		if os.IsNotExist(err) {
			createConfigFile()
		}
	}

	// Read the entire config file into memory
	content, err := os.ReadFile("config.conf")
	if err != nil {
		output.Error("Could not read config file")
		output.Exit("Exiting...")
		return
	}

	// Convert the content to a string for easier processing
	contentStr := string(content)

	// Iterate over all options and extract their values
	for _, key := range allOptions {
		// Regular expression to match both quoted and unquoted values
		re := regexp.MustCompile(fmt.Sprintf(`(?m)^%s\s*=\s*(?:"([^"]*)"|([^#\n]+))`, key))
		matches := re.FindStringSubmatch(contentStr)

		if matches != nil {
			var value string
			if matches[1] != "" {
				// Matched a quoted value
				value = matches[1]
			} else if matches[2] != "" {
				// Matched an unquoted value
				value = strings.TrimSpace(matches[2])
			}

			// Assign the value to the corresponding variable
			switch key {
			case "CLOUDFLARE_ACCOUNT_EMAIL":
				CLOUDFLARE_ACCOUNT_EMAIL = value
			case "CLOUDFLARE_ACCOUNT_ID":
				CLOUDFLARE_ACCOUNT_ID = value
			case "CLOUDFLARE_API_KEY":
				CLOUDFLARE_API_KEY = value
			case "DB_NAME":
				DB_NAME = value
			case "DB_ID":
				DB_ID = value
			case "PAGES_NAME":
				PAGES_NAME = value
			case "PAGES_ID":
				PAGES_ID = value
			case "WORKERS_NAME":
				WORKERS_NAME = value
			case "WORKERS_ID":
				WORKERS_ID = value
			}
		}
	}

}

func lazyness() {

	f, err := os.ReadFile("lazy")
	if err != nil {
		if os.IsNotExist(err) {
			output.Errorf("Error: File not found")
		} else {
			output.Error("Error reading file")
		}
		output.Exit("Exiting...")
		return
	}
	dbid := strings.TrimSpace(string(f))

	err = os.Setenv("DB_ID", dbid)
	if err != nil {
		output.Error("Error setting env")
		output.Exit("Exiting...")
		return
	}

}

type OS string

func (o OS) String() string {
	return string(o)
}

func GetOS() OS {
	rt := strings.ToLower(runtime.GOOS)
	o := OS(rt)
	return o
}
