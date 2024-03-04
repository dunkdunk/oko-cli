package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"github.com/dunkdunk/oko-cli/matchers"

	"github.com/spf13/cobra"
)

// modelsCmd represents the models command
var modelsCmd = &cobra.Command{
	Use:   "models",
	Short: "Extract data models from a web application.",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("models called")
		// If the source flag is set, walk the source directory
		// Otherwise, walk the current directory
		source, _ := cmd.Flags().GetString("source")
		if source != "" {
			walk(source)
		}	else {
			walk(".")
		}
	},
}

func init() {
	rootCmd.AddCommand(modelsCmd)
	modelsCmd.Flags().String("source", "s", "A source directory to extract models from.")
	modelsCmd.Flags().BoolP("extract-sensitive", "", false, "Help message for toggle")
}

func walk(path string) error {
	return filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			// Determine the file extension
			ext := filepath.Ext(path)

			// Read the file content
			content, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}

			// Match patterns
			// Pass the file content and extension to the DetectModels function
			detectedModel, err := matchers.DetectModels(string(content), ext)
			if err != nil {
				return err
			}
			
			fmt.Println(detectedModel)
		}
		
		return nil
	})
}