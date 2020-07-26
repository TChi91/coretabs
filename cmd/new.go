package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	homedir "github.com/mitchellh/go-homedir"

	"github.com/spf13/cobra"
)

const (
	djVue   string = "https://github.com/TChi91/rest-vue"
	djReact string = "https://github.com/TChi91/drf-reactjs"
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

		//Create coretabs folder in the Home directory if doesn't exist
		wd, err := coretabsFolderIfNotExists()
		must(err)

		//change the working directory to §HOME/coretabs
		os.Chdir(wd)

		var projectName string
		fmt.Print("Project Name: ")
		if _, err := fmt.Scanf("%s", &projectName); err != nil {
			fmt.Printf("%s\n", err)
			return
		}

		var repo string
		fmt.Print(`Framework to use with DRF?
"r" for ReactJS, or "v" for Vuejs: `)

		// fmt.Scanf("%s", &repo)
		repo, err = fef(&repo)
		whichRepo(&repo)

		fmt.Println("This may take some while ...")
		err = cloneProject(projectName, repo)
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
		gitErrMsg := deleteCommitsHistory()
		if gitErrMsg != nil {
			fmt.Println(gitErrMsg)
		}

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

func cloneProject(projectName, repo string) error {
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

	execBackCmd := exec.Command("bash", "-c", backCmd)
	execFrontCmd := exec.Command("bash", "-c", frontCmd)

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

func deleteCommitsHistory() error {
	var GitInit string
	var execGitInit *exec.Cmd
	var errMsg error = nil
	var gitConfigFilePath string

	switch opsys := runtime.GOOS; opsys {
	case "windows":
		GitInit = fmt.Sprint(`RD /S /Q .git && git init && git add . && git commit -m "Init"`)
		home, _ := homedir.Dir()
		gitConfigFilePath = fmt.Sprint(home, "\\.gitconfig")
		if _, err := os.Stat(gitConfigFilePath); os.IsNotExist(err) {
			GitInit = fmt.Sprint("RD /S /Q .git && git init")
			errMsg = errors.New(`# Setting Up User Name and Email Address
			$ git config --global user.name "Your Name"
			$ git config --global user.email Your@email.com`)
		}
		execGitInit = exec.Command("cmd", "/C", GitInit)

	case "linux":
		GitInit = fmt.Sprint(`rm -rf .git; git init; git add .; git commit -m "Init"`)
		home, _ := homedir.Dir()
		gitConfigFilePath = fmt.Sprint(home, "/.gitconfig")
		if _, err := os.Stat(gitConfigFilePath); os.IsNotExist(err) {
			GitInit = fmt.Sprint("rm -rf .git; git init")
			errMsg = errors.New(`# Setting Up User Name and Email Address
			$ git config --global user.name "Your Name"
			$ git config --global user.email Your@email.com`)
		}
		execGitInit = exec.Command("bash", "-c", GitInit)
	}

	execGitInit.Stdout = os.Stdout
	execGitInit.Stderr = os.Stderr
	must(execGitInit.Start())
	must(execGitInit.Wait())
	return errMsg
}

//Create coretabs folder in the Home directory
func coretabsFolderIfNotExists() (string, error) {
	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	coretabsPath := filepath.Join(home, "coretabs")

	if _, err := os.Stat(coretabsPath); os.IsNotExist(err) {
		os.MkdirAll(coretabsPath, os.ModePerm)
	}
	return coretabsPath, nil
}

//Choose what template to install
func whichRepo(repo *string) {
	switch *repo {
	case "r", "react":
		*repo = djReact
	default:
		*repo = djVue
	}
}

// fef stand for FrontEndFramework
func fef(frontEndFramework *string) (string, error) {
	var input string
	var err error

	reader := bufio.NewReader(os.Stdin)

	if runtime.GOOS == "windows" {
		input, err = reader.ReadString('\r')
	} else {
		input, err = reader.ReadString('\n')

	}
	if err != nil {
		fmt.Println(err)
	}
	input = strings.Trim(input, "\r\n")
	if input == "" {
		input = "0"
	}
	*frontEndFramework = input

	return *frontEndFramework, nil
}
