package cmd

import (
	"bufio"
    "fmt"
	"os"
	
	"github.com/spf13/cobra"
	"github.com/pkg/browser"
)

// morningCmd represents the morning command
var morningCmd = &cobra.Command{
	Use:   "morning",
	Short: "Open a list of urls",
	Long: `Url opens a list of urls defined in a txt file`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Good morning ! Dustin is opening your tabs")
		// Open the file.
		f, _ := os.Open("/home/soniamanoubi/Bureau/links.txt")
		// Create a new Scanner for the file.
		scanner := bufio.NewScanner(f)
		// Loop over all lines in the file and print them.
		for scanner.Scan() {
		  line := scanner.Text()
		  browser.OpenURL(line)
		}
	},
}

func init() {
	rootCmd.AddCommand(morningCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// morningCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// morningCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
