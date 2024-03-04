package main

import (
	"encoding/json"
	"fmt"
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
			model := DotNetModel{Name: match[1], Parameters: extractParameters(content)}
			models = append(models, model)
		}
	}

	return models, nil
}

// extractParameters extracts parameter names from constructor
func extractParameters(content string) []string {
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

func main() {
	// Example usage
	dotNetContent := `
using System;

public class User
{
    public int Id { get; set; }
    public string Username { get; set; }
    public string Email { get; set; }

    public User(int id, string username, string email)
    {
        Id = id;
        Username = username;
        Email = email;
    }
}
`

	models, err := FindDotNetModelsInFile(dotNetContent)
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
