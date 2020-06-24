package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/janeczku/go-spinner"
	"github.com/spf13/cobra"
)

const (
	repo string = "https://github.com/TChi91/rest-vue"
)

var (
	pathsNeeded = []string{"git", "pip3"}
)

func init() {
	rootCmd.AddCommand(newCmd)
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
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
	// fmt.Println("cloning project started ....")
	command := exec.Command("git", "clone", repo, projectName)

	// command.Stdout = os.Stdout
	// command.Stderr = os.Stderr

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
	// fmt.Println("starting virtual env ....")

	frontCmd := fmt.Sprintf("%v install; %v run build", packageManager, packageManager)
	backCmd := fmt.Sprint("python3 -m venv venv; source venv/bin/activate; pip install -r requirements.txt; python manage.py migrate")

	execBackCmd := exec.Command("bash", "-c", backCmd)
	execFrontCmd := exec.Command("bash", "-c", frontCmd)

	// execFrontCmd.Stdout = os.Stdout
	// execFrontCmd.Stderr = os.Stderr

	// execBackCmd.Stdout = os.Stdout
	// execBackCmd.Stderr = os.Stderr

	must(execBackCmd.Start())

	must(execFrontCmd.Start())

	must(execBackCmd.Wait())

	must(execFrontCmd.Wait())

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
		fmt.Print("Project Name: ")
		if _, err := fmt.Scanf("%s", &projectName); err != nil {
			fmt.Printf("%s\n", err)
			return
		}
		fmt.Print("What package manager you want to use?\nyarn or npm (default yarn): ")

		var packageManager string

		reader := bufio.NewReader(os.Stdin)
		packageManager, _ = reader.ReadString('\n')

		if packageManager != "npm" {
			packageManager = "yarn"
		}
		fmt.Println("This may take some while ...")
		s := spinner.StartNew("cloning project ...")
		s.SetCharset([]string{"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷"})
		err = cloneProject(projectName)
		must(err)
		s.Stop()
		fmt.Println("✓ Cloning: Completed")

		err = changeDirectory(projectName)
		must(err)

		s = spinner.StartNew("Installing Dependencies ...")
		s.SetCharset([]string{"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷"})
		err = installRequirments(packageManager)
		must(err)
		s.Stop()
		fmt.Println("✓ Installing Dependencies: Completed")

		fmt.Println("All Done ✓")

	},
}
