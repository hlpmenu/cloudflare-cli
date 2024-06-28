package cfapi

import (
	"bytes"
	"encoding/json"
	"go-debug/env"
	"go-debug/output"
	"io"
	"net/http"
	"strings"
)

var c = http.Client{}

const (
	POST = "POST"
	GET  = "GET"
	PUT  = "PUT"
	DEL  = "DELETE"
)

type AvailableMethods []string

var availableMethods = AvailableMethods{POST, GET, PUT, DEL}

type CFRequest struct {
	url         string
	body        CFRequestBody
	contentType string
	UseApiKey   bool
	method      string
}

type CFRequestBody struct {
	Body    []byte
	HasBody bool
}
type CFCommand struct {
	CMD   string
	Flags map[string]string
}

var cfbaseurl = "https://api.cloudflare.com/"

// SendRequest sends a request to the Cloudflare API
func sendRequest(req *http.Request) (*http.Response, error) {
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil

}

func ConstRequest(cf *CFRequest) *http.Request {
	var (
		bodyReader io.Reader
		hdrs       = make(http.Header)
	)

	// Check if the request type has a body
	if cf.body.HasBody {
		bodyReader = bytes.NewBuffer(cf.body.Body)
	} else {
		bodyReader = nil
	}

	if cf.contentType == "json" || cf.contentType == "" || cf.contentType == "application/json" {
		hdrs.Set("Content-Type", "application/json")
	}

	// Check if we are using an API key or token
	if cf.UseApiKey {
		if env.CLOUDFLARE_API_KEY == "" || env.CLOUDFLARE_ACCOUNT_EMAIL == "" {
			hdrs.Set("X-Auth-Key", env.CLOUDFLARE_API_KEY)
			hdrs.Set("X-Auth-Email", env.CLOUDFLARE_ACCOUNT_EMAIL)
		} else {
			output.Errorf("Error: API Key or Email not set \n Please set CLOUDFLARE_API_KEY and CLOUDFLARE_ACCOUNT_EMAIL in your environment variables or use the API Token instead.")
			output.Info("Trying to use API Token instead...")
			if env.CLOUDFLARE_API_TOKEN != "" {
				hdrs.Set("Authorization", "Bearer "+env.CLOUDFLARE_API_TOKEN)
				output.Info("Found API Token")
			} else {
				output.Errorf("Error: API Token not set \n Please set CLOUDFLARE_API_TOKEN in your environment variables.")
				output.Exit("Exiting...")
			}
		}
	} else {
		if env.CLOUDFLARE_API_TOKEN != "" {
			hdrs.Set("Authorization", "Bearer "+env.CLOUDFLARE_API_TOKEN)
		} else {
			output.Errorf("Error: API Token not set \n Please set CLOUDFLARE_API_TOKEN in your environment variables.")
			output.Info("Trying to use API Key instead...")
			if env.CLOUDFLARE_API_KEY != "" && env.CLOUDFLARE_ACCOUNT_EMAIL != "" {
				hdrs.Set("X-Auth-Key", env.CLOUDFLARE_API_KEY)
				hdrs.Set("X-Auth-Email", env.CLOUDFLARE_ACCOUNT_EMAIL)
				output.Info("Found API Key")
			} else {
				output.Errorf("Error: API Key or Email not set \n Please set CLOUDFLARE_API_KEY and CLOUDFLARE_ACCOUNT_EMAIL in your environment variables.")
				output.Exit("Exiting...")
			}
		}
	}

	method := strings.ToUpper(cf.method)

	// Check if method is in availableMethods
	methodExists := false
	for _, m := range availableMethods {
		if method == m {
			methodExists = true
			break
		}
	}
	if !methodExists {
		output.Errorf("Error: Method %s not found in available methods", method)
		output.Exit("Exiting...")
	}

	// Set the URL
	url := cf.url
	if !strings.Contains(url, "https://api.cloudflare.com") {
		output.Errorf("Error: URL %s is not a Cloudflare API URL", url)
	}

	req, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		output.Errorf("Error: %s", err)
	}

	return req

}
func CreateRequest(cf *CFRequest) {
	req := ConstRequest(cf)
	res, err := sendRequest(req)
	if err != nil {
		output.Errorf("Error: %s", err)
		output.Exit("Exiting...")
	}
	logresponse(res)

}

func logresponse(res *http.Response) {
	body, err := io.ReadAll(res.Body)
	if err != nil {
		output.Errorf("Error: %s", err)
	}
	defer res.Body.Close()
	bodyS := string(body)
	if len(bodyS) > 1 {
		output.Errorf("Error, Suspicious response: %s", bodyS)
	}
	status := res.Status

	if status != "200 OK" {
		output.Errorf("Error: %s", status)
	} else {
		output.Successf("Status: %s", status)
	}
	output.Infof("Response: %s", bodyS)

}

func toJson(m map[string]string) []byte {
	json, err := json.Marshal(m)
	if err != nil {
		output.Errorf("Error: %s", err)
	}
	return json
}

// Exist in map
func flagExists(m map[string]string, key string) (bool, string) {
	_, ok := m[key]
	return ok, m[key]
}
