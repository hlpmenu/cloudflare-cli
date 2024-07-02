package setenv

import (
	i "go-debug/cmd/interactive"
	"go-debug/env"
	output "go-debug/output"
	"os"
	"os/exec"
	"reflect"
	"strings"
)

var cli *i.CLI
var multiselect int

var userenv *envSetup

type envSetup struct {
	Auth    Auth
	D1      D1
	Workers Workers
	Pages   Pages
	Shell   Shell
	Dump    []string
}

type Workers struct {
	WorkerID string `env:"WORKER_ID"`
}

type Pages struct {
	PagesID string `env:"PAGES_ID"`
}

type D1 struct {
	Name string `env:"D1_NAME"`
	ID   string `env:"DB_ID"`
}

type Auth struct {
	Type      string
	Token     string `env:"CLOUDFLARE_API_TOKEN"`
	Key       string `env:"CLOUDFLARE_API_KEY"`
	Email     string `env:"CLOUDFLARE_ACCOUNT_EMAIL"`
	AccountID string `env:"CLOUDFLARE_ACCOUNT_ID"`
}

type Shell struct {
	Shell string
	OS    string
}

var AvailableShells = []string{"bash", "zsh", "fish", "powershell", "cmd", "tcsh"}
var AvailableOS = []string{"windows", "linux", "macos"}

type vars map[string]string

var definedvars vars

func (d vars) Add(key string, value string) {
	d[key] = value
}

func SetEnv() {
	userenv = &envSetup{}
	cli = i.NewCLI()

	cli.WaitForEnter("You will now be guided through setting environment variables. Press enter to continue...")

	// Set up authentication
	// Alt 1: API Token
	// Alt 2: API Key
	multiselect = cli.MultiChoiceQuestion(
		"Which authentication method would you like to set up?",
		[]string{"API Token", "API Key"},
	)
	if multiselect == 0 {
		// API Token
		userenv.Auth.Type = "token"
		token := cli.AskForInput("Enter your API Token")
		if token == "" {
			cli.Print("Error: Empty token not allowed when WIP")
			token = cli.Input("Enter your token again")
			// TODO: Add validity check
		}
		userenv.Auth.Token = strings.TrimSpace(token)

	} else if multiselect == 1 {
		// Type
		userenv.Auth.Type = "key"

		key := cli.AskForInput("Enter your API Key")
		if key == "" {
			cli.Print("Error: Empty key not allowed when WIP")
			key = cli.Input("Enter your key again")
			// TODO: Add validity check
		}
		userenv.Auth.Key = strings.TrimSpace(key)

		// Email
		email := cli.AskForInput("Enter your Account Email")
		userenv.Auth.Email = strings.TrimSpace(email)
		if userenv.Auth.Email == "" {
			cli.Print("Error: Empty email not allowed when WIP")
			email = cli.Input("Enter your email again")
		}
		userenv.Auth.Email = strings.TrimSpace(email)

	}

	// Set up Account ID
	cli.WaitForEnter("Now lets setup your Account ID. Press enter to continue...")
	accountID := cli.Input("Enter your Account ID")
	if accountID == "" {
		cli.Print("Error: Empty account ID not allowed when WIP")
		accountID = cli.Input("Enter your account ID again")
		// TODO: Add validity check
	}
	userenv.Auth.AccountID = strings.TrimSpace(accountID)

	cli.Success("All done!")

	// Services
setupServiceVars:
	for {
		reRun := services()
		switch reRun {
		case true:
			continue setupServiceVars
		case false:
			break setupServiceVars
		}
	}
	cli.Print("All done!")
	cli.Print("Setting environment variables...")
	setVars()
	cli.AutoPrint([]string{".", ".", "."})

	multiselect = cli.MultiChoiceQuestion(
		"Do you want to create a .env file?",
		[]string{"Yes", "No"},
	)
	if multiselect == 0 {
		createDotEnv()
		cli.Success("File created and environment variables set! Now you can run start creating!")
		cli.Leave("Goodbye!")
	}
	output.Success("Environment variables set! Now you can run start creating!")
	cli.Leave("Goodbye!")

}

func services() bool {
	multiselect = cli.MultiChoiceQuestion(
		"Do you want to set up additional services for this project?",
		[]string{"Workers", "Pages", "D1"},
	)
	if multiselect == 0 {
		// Workers
		cli.Print("Workers not implemented yet")
	}
	if multiselect == 1 {
		// Pages
		cli.Print("Pages not implemented yet")
	}
	if multiselect == 2 {
		// D1
		cli.Print("Lets set up D1")
		dbId := cli.Input("Enter your D1 Database ID")
		if dbId == "" {
			cli.Print("Error: Empty database ID not allowed when WIP")
			dbId = cli.Input("Enter your database ID again")
			// TODO: Add validity check
		}
		userenv.D1.ID = strings.TrimSpace(dbId)

		dbname := cli.Input("Enter your D1 Database Name")
		if dbname == "" {
			cli.Print("Error: Empty database name not allowed when WIP")
			dbname = cli.Input("Enter your database name again")
		}
		userenv.D1.Name = strings.TrimSpace(dbname)

	}
	cli.Print("Done!")
	multiselect = cli.MultiChoiceQuestion(
		"Do you want to set up additional services for this project?",
		[]string{"Yes", "No"},
	)
	return multiselect == 0
}

func createDotEnv() {
	var (
		shell, out string
	)

	userenv.Shell.OS = env.GetOS().String()

	cmd := exec.Command("sh", "-c", "echo $SHELL")
	o, err := cmd.Output()
	if err != nil {
		output.Errorf("Error: %s", err)
		output.Info("Please manually supply your shell")
		manuallyAssignShell()
	} else {
		out = string(o)
		shell = string(out)

		// Check if shell is valid
		var loopN int
		for _, s := range AvailableShells {
			if strings.Contains(shell, s) {
				userenv.Shell.Shell = s
				break
			}
			if loopN < len(AvailableShells) {
				loopN++
				continue
			} else {
				output.Error("Error: Invalid shell")
				cli.Print("Please manually assign your shell")
				manuallyAssignShell()
				break
			}
		}
	}
	var ext string
	switch userenv.Shell.Shell {
	case "bash":
		ext = ".sh"
	case "zsh":
		ext = ".zsh"
	case "fish":
		ext = ".fish"
	case "powershell":
		ext = ".ps1"
	case "cmd":
		ext = ".cmd"
	case "tcsh":
		ext = ".tcsh"
	}
	writeEnvFile(ext)

	cli.Print("Dotenv file created")

}

func manuallyAssignShell() {
	output.Error("Not implemented yet")
	output.Exit("Exiting...")
}

func setVars() {
	definedvars = make(vars)
	// Initialize a queue with the reflect.Value of userenv
	queue := []reflect.Value{reflect.ValueOf(userenv).Elem()}

	for len(queue) > 0 {
		// Pop the first element from the queue
		val := queue[0]
		queue = queue[1:]

		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)
			fieldType := val.Type().Field(i)

			// Check if the field has the "env" tag
			envTag := fieldType.Tag.Get("env")
			if envTag != "" && field.Kind() == reflect.String {
				// Set the environment variable
				os.Setenv(envTag, field.String())
				definedvars.Add(envTag, field.String())
			}

			// If the field is a struct, add it to the queue
			if field.Kind() == reflect.Struct {
				queue = append(queue, field)
			} else if field.Kind() == reflect.Ptr && field.Elem().Kind() == reflect.Struct {
				// If the field is a pointer to a struct, dereference it and add to the queue
				queue = append(queue, field.Elem())
			}
		}
	}
}
