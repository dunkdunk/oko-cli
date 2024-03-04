package main

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

// NodeJSModel represents a Node.js data model
type NodeJSModel struct {
	Name       string   `json:"name"`
	Parameters []string `json:"parameters"`
}

// FindNodeJSModelsInFile finds data models in a single Node.js file
func FindNodeJSModelsInFile(content string) ([]NodeJSModel, error) {
	var models []NodeJSModel

	// Define patterns for Node.js data models
	patterns := []*regexp.Regexp{
		// Mongoose pattern for Mongoose models
		regexp.MustCompile(`const\s+(\w+)\s*=\s*mongoose\.model\s*\(\s*'(\w+)'\s*,`),
		// Sequelize pattern for Sequelize models
		regexp.MustCompile(`const\s+(\w+)\s*=\s*sequelize\.define\s*\(\s*'(\w+)'\s*,`),
	}

	// Match patterns
	for _, pattern := range patterns {
		matches := pattern.FindAllStringSubmatch(content, -1)
		for _, match := range matches {
			model := NodeJSModel{Name: match[2], Parameters: extractParameters(content)}
			models = append(models, model)
		}
	}

	return models, nil
}

// extractParameters extracts parameter names from constructor
func extractParameters(content string) []string {
	var parameters []string

	// Pattern to extract parameter names from constructor
	re := regexp.MustCompile(`(?m)^\s*constructor\s*\((.*?)\)`)
	matches := re.FindAllStringSubmatch(content, -1)
	for _, match := range matches {
		params := strings.Split(match[1], ",")
		for _, param := range params {
			param = strings.TrimSpace(param)
			if param != "" {
				parameters = append(parameters, param)
			}
		}
	}

	return parameters
}

func main() {
	// Example usage
	nodeJSContent := `
const mongoose = require('mongoose');

const userSchema = new mongoose.Schema({
    username: { type: String, required: true },
    email: { type: String, required: true }
});

const User = mongoose.model('User', userSchema);
`

	models, err := FindNodeJSModelsInFile(nodeJSContent)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	jsonData, err := json.MarshalIndent(models, "", "    ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	fmt.Println(string(jsonData))
}
