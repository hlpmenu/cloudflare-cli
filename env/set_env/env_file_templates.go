package setenv

import (
	"bufio"
	"fmt"
	output "go-debug/output"
	"os"
)

func writeEnvFile(filetype string) {

	f, err := os.Create("env." + filetype)
	if err != nil {
		output.Errorf("Error: %s", err)
		output.Exit("Exiting...")
	}
	defer f.Close()

	// Todo: Add support for multiple shells
	// Now only bash is supported

	w := bufio.NewWriter(f)
	for k, v := range definedvars {
		if k != "" {
			c := writeForShell(k, v)
			_, err = w.Write(c)
			if err != nil {
				w.Reset(f)
				output.Errorf("Error: %s", err)
				return
			}
		}
	}
	w.Flush()

	// Set executable
	err = os.Chmod("env."+filetype, 0755)
	if err != nil {
		output.Errorf("Error: %s", err)
		output.Exit("Exiting...")
	}

	output.Success("File created: env." + filetype)

}

func writeForShell(k string, v string) []byte {
	var s string
	switch userenv.Shell.Shell {
	case "bash":
		s = fmt.Sprintf(`export %s="%s"`, k, v)
		s = s + "\n"
	case "fish":
		s = `set ` + k + ` = ` + v + "\n"
	case "zsh":
		s = `export ` + k + `=` + v + "\n"
	case "powershell":
		s = `$env:` + k + ` = ` + v + "\n"
	case "cmd":
		s = `set ` + k + `=` + v + "\n"
	case "tcsh":
		s = `setenv ` + k + ` ` + v + "\n"
	default:
		output.Error("Not implemented yet")
		output.Exit("Exiting...")
	}

	return []byte(s)
}
