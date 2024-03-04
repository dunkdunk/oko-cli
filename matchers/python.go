package main

import (
	"encoding/json"
	"fmt"
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
			model := PythonModel{Name: match[1], Parameters: extractParameters(content)}
			models = append(models, model)
		}
	}

	return models, nil
}

// extractParameters extracts parameter names from class declaration
func extractParameters(content string) []string {
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

func main() {
	// Example usage
	pythonContent := `
from flask_sqlalchemy import SQLAlchemy

db = SQLAlchemy()

class User(db.Model):
    __tablename__ = 'users'
    id = db.Column(db.Integer, primary_key=True)
    username = db.Column(db.String(80), unique=True, nullable=False)
    email = db.Column(db.String(120), unique=True, nullable=False)
`

	models, err := FindPythonModelsInFile(pythonContent)
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
