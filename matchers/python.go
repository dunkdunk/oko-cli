package matchers

import (
	"regexp"
	"strings"
)

// PythonModel represents a Python data model
type PythonModel struct {
	Name       string   `json:"name"`
	Parameters []string `json:"parameters"`
}

// FindPythonModelsInFile finds data models in a single Python file
func FindPythonModelsInFile(content string) ([]PythonModel, error) {
	var models []PythonModel

	// Define patterns for Python data models
	patterns := []*regexp.Regexp{
		// Django pattern for Django models
		regexp.MustCompile(`class\s+(\w+)\s*\(\s*models\.Model\s*\)\s*:`),
		// SQLAlchemy pattern for SQLAlchemy models
		regexp.MustCompile(`class\s+(\w+)\s*\(\s*Base\s*\)\s*:`),
		// Peewee pattern for Peewee models
		regexp.MustCompile(`class\s+(\w+)\s*\(\s*Model\s*\)\s*:`),
		// Flask-SQLAlchemy pattern for Flask-SQLAlchemy models
		regexp.MustCompile(`class\s+(\w+)\s*\(.*db\.Model\s*\)\s*:`),
		// MongoEngine pattern for MongoEngine models
		regexp.MustCompile(`class\s+(\w+)\s*\(.*Document\s*\)\s*:`),
	}

	// Match patterns
	for _, pattern := range patterns {
		matches := pattern.FindAllStringSubmatch(content, -1)
		for _, match := range matches {
			model := PythonModel{Name: match[1], Parameters: extractPythonParameters(content)}
			models = append(models, model)
		}
	}

	return models, nil
}

// extractParameters extracts parameter names from class declaration
func extractPythonParameters(content string) []string {
	var parameters []string

	// Pattern to extract parameter names from class declaration
	re := regexp.MustCompile(`(?m)^\s*def\s+__init__\(self(.*?)\):`)
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

