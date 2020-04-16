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
	var repo string = "https://github.com/yaseralnajjar/django-vue-template"
	command := exec.Command("git", "clone", repo)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
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

		djangoVue()

	},
}
