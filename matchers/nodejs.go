package matchers

import (
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
			model := NodeJSModel{Name: match[2], Parameters: extractNodeJSParameters(content)}
			models = append(models, model)
		}
	}

	return models, nil
}

// extractParameters extracts parameter names from constructor
func extractNodeJSParameters(content string) []string {
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
