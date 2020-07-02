package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"

	"github.com/spf13/cobra"
)

const (
	repo string = "https://github.com/TChi91/rest-vue"
)

var (
	pathsNeeded        = []string{"git", "python3", "npm"}
	pathsNeededWindows = []string{"git", "python", "npm", "pip"}
)

func init() {
	rootCmd.AddCommand(newCmd)
}

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "create new project",
	Long:  `with new you can create new project`,
	Run: func(cmd *cobra.Command, args []string) {
		missing, err := checkPaths()
		if err != nil {
			for key, value := range missing {
				fmt.Println(key, ":", value)
			}
			fmt.Println("")
			return
		}

		var projectName string
		fmt.Print("Project Name: ")
		if _, err := fmt.Scanf("%s", &projectName); err != nil {
			fmt.Printf("%s\n", err)
			return
		}

		fmt.Println("This may take some while ...")
		err = cloneProject(projectName)
		must(err)

		err = changeDirectory(projectName)
		must(err)

		opsys := runtime.GOOS
		switch opsys {
		case "windows":
			err = installRequirmentsWindows("npm")
		case "linux":
			err = installRequirments("npm")
		}
		must(err)

		fmt.Println("")
		fmt.Println("")
		fmt.Println("✓ All Done")

	},
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func checkPaths() (map[string]error, error) {
	missing := make(map[string]error)
	opsys := runtime.GOOS
	switch opsys {
	case "windows":
		for _, path := range pathsNeededWindows {
			_, err := exec.LookPath(path)
			if err != nil {
				missing[path] = err
			}
		}
	case "linux":
		for _, path := range pathsNeeded {
			_, err := exec.LookPath(path)
			if err != nil {
				missing[path] = err
			}
		}
	}

	if len(missing) != 0 {
		return missing, errors.New("Missing Dependencies")
	}
	return nil, nil

}

func cloneProject(projectName string) error {
	command := exec.Command("git", "clone", repo, projectName)

	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	must(err)
	return nil

}

func changeDirectory(dir string) error {
	err := os.Chdir(dir)
	must(err)
	return nil
}

func installRequirments(packageManager string) error {
	frontCmd := fmt.Sprintf("%v install; %v run build", packageManager, packageManager)
	backCmd := fmt.Sprint("python3 -m venv venv; source venv/bin/activate; pip install -r requirements.txt; python manage.py migrate")
	GitInit := fmt.Sprintf("rm -rf .git; git init") // commands to add: ; git add .; git commit -m \"Initial commit\"

	execGitInit := exec.Command("bash", "-c", GitInit)
	execBackCmd := exec.Command("bash", "-c", backCmd)
	execFrontCmd := exec.Command("bash", "-c", frontCmd)

	execFrontCmd.Stdout = os.Stdout
	execFrontCmd.Stderr = os.Stderr

	execGitInit.Stdout = os.Stdout
	execGitInit.Stderr = os.Stderr

	execBackCmd.Stdout = os.Stdout
	execBackCmd.Stderr = os.Stderr

	must(execGitInit.Start())
	must(execBackCmd.Start())
	must(execFrontCmd.Start())

	must(execGitInit.Wait())
	must(execBackCmd.Wait())
	must(execFrontCmd.Wait())

	return nil

}

func installRequirmentsWindows(packageManager string) error {

	frontCmd := fmt.Sprintf("%v install && %v run build", packageManager, packageManager)
	backCmd := fmt.Sprint("pip install virtualenv && python -m venv venv && .\\venv\\Scripts\\activate && pip install -r requirements.txt && python manage.py migrate")

	execBackCmd := exec.Command("cmd", "/C", backCmd)
	execFrontCmd := exec.Command("cmd", "/C", frontCmd)

	execFrontCmd.Stdout = os.Stdout
	execFrontCmd.Stderr = os.Stderr

	execBackCmd.Stdout = os.Stdout
	execBackCmd.Stderr = os.Stderr

	must(execBackCmd.Start())
	must(execFrontCmd.Start())

	must(execBackCmd.Wait())
	must(execFrontCmd.Wait())

	return nil

}
