package env

import (
	"os"
)

// Main CF API Details
var (
	CLOUDFLARE_ACCOUNT_EMAIL string // Your accounts email
	CLOUDFLARE_ACCOUNT_ID    string // Your accounts ID
	CLOUDFLARE_API_KEY       string // Your accounts API KEY
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

	if UseEnv {
		CLOUDFLARE_ACCOUNT_EMAIL = GetEnv("CLOUDFLARE_ACCOUNT_EMAIL")
		CLOUDFLARE_ACCOUNT_ID = GetEnv("CLOUDFLARE_ACCOUNT_ID")
		CLOUDFLARE_API_KEY = GetEnv("CLOUDFLARE_API_KEY")
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

	return value
}

func SetupEnvConfig() {

	configFile err := os.Stat("env.conf")
	if err != nil {
		if os.IsNotExist(err) {
			createConfig()
		}
	}

}

createConfig() {
	// Create a new file
	f, err := os.Create("env.conf")
	

}