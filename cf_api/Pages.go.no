package cfapi

import (
	"go-debug/env"
	"go-debug/output"
)

func loadPagesCommands() {
	// Pages
	Cmds.Add(*PagesMainCommand)
}

type Pages struct {
	Request CFRequest
}

// Pages API url base
var pagesBaseURL string

func PagesCommand(c *CFCommand) {

	if env.CLOUDFLARE_ACCOUNT_ID == "" {
		output.Errorf("No account ID provided, \n Run 'cf_api setup' to set up your account ID, or run 'cf_api setup --config' to generate a config file\n")
		output.Exit("Exiting...")
	}
	pagesBaseURL = cfbaseurl + "client/v4/accounts/" + env.CLOUDFLARE_ACCOUNT_ID + "/pages/projects"

	pages := Pages{}
	flags := c.Flags

	switch c.CMD {
	case "list":
		pages.List(flags)
	case "delete":
		pages.Delete(flags)
	case "get":
		pages.Get(flags)
	default:
		output.Error("Invalid command")
		output.Exit("Exiting...")
	}

}

func (p *Pages) List(m flagsMap) {
	cf := CFRequest{
		url:    pagesBaseURL,
		method: GET,
	}
	CreateRequest(&cf)
}

func (p *Pages) Delete(m flagsMap) {
	// Values
	exists, projectName := flagExists(m, "-project-name")
	if !exists || len(projectName) < 1 {
		output.Error("No project name provided")
		output.Exit("Please provide a project name using: cf-cli pages delete -project-name <project_name>")
	}

	cf := CFRequest{
		url:    pagesBaseURL + "/" + projectName,
		method: DEL,
	}
	// Const body
	cf.body.HasBody = false

	CreateRequest(&cf)
}

func (p *Pages) Get(m flagsMap) {
	// Values
	exists, projectName := flagExists(m, "-project-name")
	if !exists || len(projectName) < 1 {
		output.Error("No project name provided")
		output.Exit("Please provide a project name using: cf-cli pages get -project-name <project_name>")
	}

	cf := CFRequest{
		url:    pagesBaseURL + "/" + projectName,
		method: GET,
	}
	// Const body
	cf.body.HasBody = false

	CreateRequest(&cf)
}

// func (p *Pages) Create(m flagsMap) {

// 	var (
// 		useFile = true
// 	)
// 	// Values
// 	exists, file := flagExists(m, "-file")
// 	if !exists || len(file) < 1 {
// 		useFile = false
// 	}
// 	if useFile {
// 		// Read file
// 		// TODO
// 	} else {
// 		// Values
// 		exists, projectName := flagExists(m, "-name")
// 		if !exists {
// 			output.Error("No project name provided")
// 			output.Exit("Please provide a project name using:  -name <project_name>")
// 		}

// 		type BuildConfig struct {
// 			BuildCaching      bool   `json:"build_caching"`
// 			BuildCommand      string `json:"build_command"`
// 			DestinationDir    string `json:"destination_dir"`
// 			RootDir           string `json:"root_dir"`
// 			WebAnalyticsTag   string `json:"web_analytics_tag"`
// 			WebAnalyticsToken string `json:"web_analytics_token"`
// 		}

// 		type Project struct {
// 			BuildConfig         BuildConfig            `json:"build_config"`
// 			CanonicalDeployment map[string]interface{} `json:"canonical_deployment"`
// 			DeploymentConfigs   map[string]interface{} `json:"deployment_configs"`
// 			LatestDeployment    map[string]interface{} `json:"latest_deployment"`
// 			Name                string                 `json:"name"`
// 			ProductionBranch    string                 `json:"production_branch"`
// 		}

// 		project := Project{
// 			BuildConfig: BuildConfig{
// 				BuildCaching:      true,
// 				BuildCommand:      buildCommand,
// 				DestinationDir:    destinationDir,
// 				RootDir:           rootDir,
// 				WebAnalyticsTag:   webAnalyticsTag,
// 				WebAnalyticsToken: webAnalyticsToken,
// 			},
// 			CanonicalDeployment: map[string]interface{}{},
// 			DeploymentConfigs:   deploymentConfigs,
// 			LatestDeployment:    map[string]interface{}{},
// 			Name:                projectName,
// 			ProductionBranch:    "main",
// 		}

// 		body, err := toJsonReflect(project)
// 		if err != nil {
// 			output.Errorf("Error: %s", err)
// 			output.Exit("Exiting...")
// 		}

// 		cf := CFRequest{
// 			url:    pagesBaseURL,
// 			method: POST,
// 		}
// 		cf.body.HasBody = true
// 		cf.body.Body = body
// 		CreateRequest(&cf)
// 	}
// }

// func (p *Pages) UpdateDeploymentConfigs(accountID, projectName string) {
// 	type EnvVar struct {
// 		Value string `json:"value,omitempty"`
// 		Type  string `json:"type,omitempty"`
// 	}

// 	type DeploymentConfig struct {
// 		CompatibilityDate  string             `json:"compatibility_date"`
// 		CompatibilityFlags []string           `json:"compatibility_flags"`
// 		EnvVars            map[string]*EnvVar `json:"env_vars"`
// 	}

// 	type UpdatePayload struct {
// 		DeploymentConfigs map[string]DeploymentConfig `json:"deployment_configs"`
// 	}

// 	payload := UpdatePayload{
// 		DeploymentConfigs: map[string]DeploymentConfig{
// 			"production": {
// 				CompatibilityDate: "2022-01-01",
// 				CompatibilityFlags: []string{
// 					"url_standard",
// 				},
// 				EnvVars: map[string]*EnvVar{
// 					"BUILD_VERSION": {
// 						Value: "3.3",
// 					},
// 					"delete_this_env_var": nil,
// 					"secret_var": {
// 						Type:  "secret_text",
// 						Value: "A_CMS_API_TOKEN",
// 					},
// 				},
// 			},
// 		},
// 	}

// 	body, err := toJsonReflect(payload)
// 	if err != nil {
// 		log.Fatalf("Error marshalling update payload to JSON: %v", err)
// 	}

// 	cf := CFRequest{
// 		url:    fmt.Sprintf("https://api.cloudflare.com/client/v4/accounts/%s/pages/projects/%s", accountID, projectName),
// 		method: "PATCH",
// 		body:   body,
// 	}
// 	CreateRequest(&cf)
// }
