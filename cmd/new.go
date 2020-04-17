package cmd

import (
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(newCmd)
}

func djangoVue() {
	repo := "https://github.com/yaseralnajjar/django-vue-template"

	command := exec.Command("git", "clone", repo)

	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		log.Fatal("Error:", err)
	}

}

func setProjectName(args []string) {
	if len(args) != 2 {
		log.Fatal("new command can be only with two args")
	}

	err := os.Rename("django-vue-template", args[1])
	if err != nil {
		log.Fatal("Error:", err)
	}
}

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "create new project",
	Long:  `with new you can create new project`,
	Run: func(cmd *cobra.Command, args []string) {
		_, err := exec.LookPath("git")
		if err != nil {
			log.Fatal(err)
		}

		switch args[0] {
		case "fullstack":
			switch len(args) {
			case 1:
				djangoVue()
			case 2:
				djangoVue()
				setProjectName(args)
			}

		}

	},
}
