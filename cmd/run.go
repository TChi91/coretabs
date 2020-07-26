package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"

	"github.com/TChi91/coretabs/config"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run your servers",
	Long: `Run you Front-End and Back-End servers with 
one command only: coretabs run.`,
	Run: func(cmd *cobra.Command, args []string) {
		config := config.NewConfig()

		// err := viper.Unmarshal(&config)
		// if err != nil {
		// 	log.Fatalf("unable to decode into struct, %v", err)
		// }

		var server int
		fmt.Print(`What server you want to start:
(1) ==> Front-End server
(2) ==> Back-End server: `)
		if _, err := fmt.Scanf("%d", &server); err != nil {
			fmt.Printf("%s\n", err)
			return
		}

		var port int

		switch server {
		case 1:
			err := checkFileToRunWith("package.json")
			if err != nil {
				return
			}

			fmt.Print(`Which port you want to use, (8080): `)

			inputPort, err := readPort(&port)
			if err != nil {
				return
			}

			if inputPort != 0 {
				config.FrontEnd.Port = port
				err = frontEndServer(config)
				if err != nil {
					return
				}
			} else {
				err = frontEndServer(config)
				if err != nil {
					return
				}
			}

		case 2:
			err := checkFileToRunWith("manage.py")
			if err != nil {
				return
			}

			fmt.Print(`Which port you want to use, (8000): `)

			inputPort, err := readPort(&port)
			if err != nil {
				return
			}

			if inputPort != 0 {
				config.BackEnd.Port = port
				err = backEndServer(config)
				if err != nil {
					return
				}
			} else {
				err = backEndServer(config)
				if err != nil {
					return
				}
			}

		}

	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}

func backEndServer(config config.AppConfig) error {
	backCmd := fmt.Sprintf("source venv/bin/activate; python manage.py runserver %v", config.BackEnd.Port)

	execBackCmd := exec.Command("bash", "-c", backCmd)

	if runtime.GOOS == "windows" {
		backCmd = fmt.Sprintf(".\\venv\\Scripts\\activate && python manage.py runserver %v", config.BackEnd.Port)
		execBackCmd = exec.Command("cmd", "/C", backCmd)
	}

	execBackCmd.Stdout = os.Stdout
	execBackCmd.Stderr = os.Stderr

	must(execBackCmd.Start())

	must(execBackCmd.Wait())

	return nil

}

func frontEndServer(config config.AppConfig) error {

	var frontCmd string
	var execFrontCmd *exec.Cmd
	OS := runtime.GOOS

	switch OS {
	case "windows":
		frontCmd = fmt.Sprintf("set PORT=%v && npm run serve", config.FrontEnd.Port)
		execFrontCmd = exec.Command("cmd", "/C", frontCmd)
	case "linux":
		frontCmd = fmt.Sprintf("PORT=%v npm run serve", config.FrontEnd.Port)
		execFrontCmd = exec.Command("bash", "-c", frontCmd)
	}

	execFrontCmd.Stdout = os.Stdout
	execFrontCmd.Stderr = os.Stderr

	err := execFrontCmd.Start()
	if err != nil {
		return err
	}
	err = execFrontCmd.Wait()
	if err != nil {
		return err

	}
	return nil

}

func readPort(port *int) (int, error) {
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

	if *port, err = strconv.Atoi(input); err != nil {
		return *port, err
	}

	return *port, nil
}

func checkFileToRunWith(name string) error {
	if _, err := os.Stat(name); os.IsNotExist(err) {
		err = fmt.Errorf("File not found: %v", name)
		fmt.Println(err)
		return err
	}
	return nil
}
