package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// notesCmd represents the notes command
var notesCmd = &cobra.Command{
	Use:   "notes",
	Short: "A command to take quick notes",
	Long: `This command helps you take quick notes by opening a file .txt`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Homer is opening a new note")
		// create a folder for the notes
		_ = os.Mkdir("./notes", 0755)
		// open the editor of your choice => nano for me
		// the name of the file will be your argument
		// homer notes myfile will open a file called myfile.txt 
		editorCmd := exec.Command("nano", fmt.Sprintf("./notes/%v.txt", args[0]))
		editorCmd.Stdin = os.Stdin
		editorCmd.Stdout = os.Stdout
		editorCmd.Stderr = os.Stderr
    
		err := editorCmd.Run()
		if (err != nil) {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(notesCmd)
}
