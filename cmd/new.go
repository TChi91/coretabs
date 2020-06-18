package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

const (
	repo string = "https://github.com/TChi91/rest-vue"
)

var (
	pathsNeeded = []string{"git", "pip3", "yarn"}
)

func init() {
	rootCmd.AddCommand(newCmd)
}

func checkPaths() (map[string]error, error) {
	missing := make(map[string]error)
	for _, path := range pathsNeeded {
		_, err := exec.LookPath(path)
		if err != nil {
			missing[path] = err
		}
	}
	if len(missing) != 0 {
		return missing, errors.New("missing dependeties")
	}
	return nil, nil

}

func cloneProject(projectName string) error {
	fmt.Println("cloning project started ....")
	command := exec.Command("git", "clone", repo, projectName)

	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		log.Fatal("Error:", err)
	}
	return nil

}

func changeDirectory(dir string) error {
	err := os.Chdir(dir)

	if err != nil {
		log.Fatal("Error:", err)
	}
	return nil
}

func installRequirments() error {
	fmt.Println("starting virtual env ....")
	command := exec.Command("bash", "-c",
		"yarn install; yarn build; pip3 install virtualenv; virtualenv venv; source venv/bin/activate; pip3 install -r requirements.txt; python3 manage.py migrate")

	// command := exec.Command("pip3", "install", "requirements.txt")

	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		log.Fatal("Error:", err)
	}
	return nil

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
		// var projectType string
		fmt.Print("Project Name: ")
		if _, err := fmt.Scanf("%s", &projectName); err != nil {
			fmt.Printf("%s\n", err)
			return
		}
		// fmt.Print("Project Type: ")

		// if _, err := fmt.Scanf("%s", &projectType); err != nil {
		// 	fmt.Printf("%s\n", err)
		// 	return
		// }

		err = cloneProject(projectName)
		if err != nil {
			log.Fatal(err)
		}

		err = changeDirectory(projectName)
		err = installRequirments()

	},
}
