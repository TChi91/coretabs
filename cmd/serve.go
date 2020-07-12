package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/TChi91/coretabs/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start your servers",
	Long: `Start you Front-End and Back-End servers with 
one command only: coretabs serve.`,
	Run: func(cmd *cobra.Command, args []string) {
		var config config.AppConfig

		err := viper.Unmarshal(&config)
		if err != nil {
			log.Fatalf("unable to decode into struct, %v", err)
		}

		var server int
		fmt.Print(`What server you want to start:
(1) ==> Front-End server
(2) ==> Back-End server: `)
		if _, err := fmt.Scanf("%d", &server); err != nil {
			fmt.Printf("%s\n", err)
			return
		}
		switch server {
		case 1:
			frontEndServer(config)

		case 2:
			backEndServer(config)

		}

	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

func backEndServer(config config.AppConfig) error {
	backCmd := fmt.Sprintf("source venv/bin/activate; %v", config.BackEnd.Server)

	execBackCmd := exec.Command("bash", "-c", backCmd)

	execBackCmd.Stdout = os.Stdout
	execBackCmd.Stderr = os.Stderr

	must(execBackCmd.Start())

	must(execBackCmd.Wait())

	return nil

}

func frontEndServer(config config.AppConfig) error {

	execfrontCmd := exec.Command("bash", "-c", config.FrontEnd.Server)

	execfrontCmd.Stdout = os.Stdout
	execfrontCmd.Stderr = os.Stderr

	err := execfrontCmd.Start()
	if err != nil {
		fmt.Println(err)
	}
	err = execfrontCmd.Wait()
	if err != nil {
		fmt.Println(err)
	}
	return nil

}
