package matchers

import (
	"regexp"
	"strings"
)

// RubyModel represents a Ruby data model
type RubyModel struct {
	Name       string   `json:"name"`
	Parameters []string `json:"parameters"`
}

// FindRubyModelsInFile finds data models in a single Ruby file
func FindRubyModelsInFile(content string) ([]RubyModel, error) {
	var models []RubyModel

	// Define patterns for Ruby data models
	patterns := []*regexp.Regexp{
		// ActiveRecord pattern for Rails 5+
		regexp.MustCompile(`class\s+(\w+)\s+<\s+ApplicationRecord\s*\{`),
		// ActiveRecord pattern for Rails 4.x and earlier
		regexp.MustCompile(`class\s+(\w+)\s+<\s+ActiveRecord::Base\s*\{`),
		// Mongoid pattern for MongoDB
		regexp.MustCompile(`class\s+(\w+)\s*include\s*Mongoid::Document\s*\{`),
		// Sequel pattern for SQL databases
		regexp.MustCompile(`class\s+(\w+)\s*<\s*Sequel::Model\s*\{`),
		// DataMapper pattern for ORM
		regexp.MustCompile(`class\s+(\w+)\s*<\s*DataMapper::Resource\s*\{`),
		// ROM pattern for Ruby Object Mapper
		regexp.MustCompile(`class\s+(\w+)\s*<\s*ROM::Struct\s*\{`),
		// Hanami pattern for Hanami framework
		regexp.MustCompile(`class\s+(\w+)\s*<\s*Hanami::Entity\s*\{`),
		// Sinatra-activerecord pattern for Sinatra with ActiveRecord
		regexp.MustCompile(`class\s+(\w+)\s*<\s*ActiveRecord::Base\s*\{`),
		// RethinkDB pattern for RethinkDB ORM
		regexp.MustCompile(`class\s+(\w+)\s*<\s*RethinkDB::Document\s*\{`),
		// Mongomapper pattern for MongoDB ORM
		regexp.MustCompile(`class\s+(\w+)\s*<\s*MongoMapper::Document\s*\{`),
		// Dynamoid pattern for Amazon DynamoDB ORM
		regexp.MustCompile(`class\s+(\w+)\s*<\s*Dynamoid::Document\s*\{`),
		// Ohm pattern for Redis ORM
		regexp.MustCompile(`class\s+(\w+)\s*<\s*Ohm::Model\s*\{`),
		// ROM pattern for ROM-sql
		regexp.MustCompile(`class\s+(\w+)\s*<\s*ROM::Relation\s*\{`),
		// Neo4j pattern for Neo4j ORM
		regexp.MustCompile(`class\s+(\w+)\s*<\s*Neo4j::ActiveNode\s*\{`),
		// ActiveRecord pattern for Sinatra with ActiveRecord
		regexp.MustCompile(`class\s+(\w+)\s*<\s*ActiveRecord::Base\s*\{`),
		// Ruby-on-rest pattern for Ruby-on-Rest framework
		regexp.MustCompile(`class\s+(\w+)\s*<\s*RubyOnRest::Model\s*\{`),
		// NoBrainer pattern for NoBrainer ORM
		regexp.MustCompile(`class\s+(\w+)\s*<\s*NoBrainer::Document\s*\{`),
	}

	// Match patterns
	for _, pattern := range patterns {
		matches := pattern.FindAllStringSubmatch(content, -1)
		for _, match := range matches {
			model := RubyModel{Name: match[1], Parameters: extractRubyParameters(content)}
			models = append(models, model)
		}
	}

	return models, nil
}

// extractParameters extracts parameter names from class declaration
func extractRubyParameters(content string) []string {
	var parameters []string

	// Pattern to extract parameter names from class declaration
	re := regexp.MustCompile(`(?m)^ *attr_(?:reader|accessor|writer) *(?::\w+(?:, *:?\w+)*)?`)
	matches := re.FindAllString(content, -1)
	for _, match := range matches {
		parts := strings.Split(match, "(")
		if len(parts) > 1 {
			paramStr := parts[1]
			paramStr = strings.TrimRight(paramStr, ")")
			params := strings.Split(paramStr, ",")
			for _, param := range params {
				param = strings.TrimSpace(param)
				parameters = append(parameters, param)
			}
		}
	}

	return parameters
}
