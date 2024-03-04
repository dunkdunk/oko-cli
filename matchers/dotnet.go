package matchers

import (
	"regexp"
	"strings"
)

// DotNetModel represents a .NET data model
type DotNetModel struct {
	Name       string   `json:"name"`
	Parameters []string `json:"parameters"`
}

// FindDotNetModelsInFile finds data models in a single .NET file
func FindDotNetModelsInFile(content string) ([]DotNetModel, error) {
	var models []DotNetModel

	// Define patterns for .NET data models
	patterns := []*regexp.Regexp{
		// Entity Framework pattern for Entity Framework models
		regexp.MustCompile(`public\s+class\s+(\w+)\s*{`),
		// Dapper pattern for Dapper models
		regexp.MustCompile(`class\s+(\w+)\s*{`),
	}

	// Match patterns
	for _, pattern := range patterns {
		matches := pattern.FindAllStringSubmatch(content, -1)
		for _, match := range matches {
			model := DotNetModel{Name: match[1], Parameters: extractDotNetParameters(content)}
			models = append(models, model)
		}
	}

	return models, nil
}

// extractParameters extracts parameter names from constructor
func extractDotNetParameters(content string) []string {
	var parameters []string

	// Pattern to extract parameter names from constructor
	re := regexp.MustCompile(`(?m)^\s*public\s+(\w+)\((.*?)\)`)
	matches := re.FindAllStringSubmatch(content, -1)
	for _, match := range matches {
		params := strings.Split(match[2], ",")
		for _, param := range params {
			param = strings.TrimSpace(param)
			if param != "" {
				parameters = append(parameters, param)
			}
		}
	}

	return parameters
}

