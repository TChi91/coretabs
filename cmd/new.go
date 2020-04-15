package cmd

import (
	"io"
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var djangoVue string = "https://github.com/yaseralnajjar/django-vue-template"

func copyAndCapture(w io.Writer, r io.Reader) ([]byte, error) {
	var out []byte
	buf := make([]byte, 1024, 1024)
	for {
		n, err := r.Read(buf[:])
		if n > 0 {
			d := buf[:n]
			out = append(out, d...)
			_, err := w.Write(d)
			if err != nil {
				return out, err
			}
		}
		if err != nil {
			// Read returns io.EOF at the end of file, which is not an error for us
			if err == io.EOF {
				err = nil
			}
			return out, err
		}
	}
}

func init() {
	rootCmd.AddCommand(newCmd)
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

		command := exec.Command("git", "clone", djangoVue)
		//command := exec.Command("date")
		command.Stdout = os.Stdout
		command.Stderr = os.Stderr

		err = command.Run()
		if err != nil {
			log.Fatal("Error:", err)
		}

	},
}
