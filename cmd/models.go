package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"

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
	// Read patterns from JSON file
	patternsFile, err := os.Open("./lib/patterns/data-models.json")
	if err != nil {
		fmt.Println("Error opening patterns file:", err)
		return err
	}
	defer patternsFile.Close()

	var patterns map[string]string
	err = json.NewDecoder(patternsFile).Decode(&patterns)
	if err != nil {
		fmt.Println("Error decoding patterns:", err)
		return err
	}

	return filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			// Determine the file extension
			ext := filepath.Ext(path)

			// Check if a pattern exists for the file extension
			pattern, ok := patterns[ext]
			if !ok {
				return nil // Skip files with unsupported extensions
			}

			// Read the file content
			content, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}

			// Match patterns
			re := regexp.MustCompile(pattern)
			matches := re.FindAllStringSubmatch(string(content), -1)
			for _, match := range matches {
				fmt.Printf("Found data model in %s: %s\n", path, match[1])
			}
		}
		
		return nil
	})
}