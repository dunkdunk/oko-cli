package main

import (
	"errors"
)

// DetectedModel represents a detected model in a file
type DetectedModel struct {
	Language string      // Language of the detected model
	Models   interface{} // Detected models
}

// DetectModels detects models in the given file content based on the file extension
func DetectModels(fileContent string, fileExtension string) (DetectedModel, error) {
	var detectedModel DetectedModel

	// Call appropriate language-specific function based on file extension
	switch fileExtension {
	case ".rb":
		models, err := FindRubyModelsInFile(fileContent)
		if err != nil {
			return detectedModel, err
		}
		detectedModel = DetectedModel{Language: "Ruby", Models: models}
	case ".py":
		models, err := FindPythonModelsInFile(fileContent)
		if err != nil {
			return detectedModel, err
		}
		detectedModel = DetectedModel{Language: "Python", Models: models}
	case ".cs":
		models, err := FindDotNetModelsInFile(fileContent)
		if err != nil {
			return detectedModel, err
		}
		detectedModel = DetectedModel{Language: ".NET", Models: models}
	case ".js":
		models, err := FindNodeJSModelsInFile(fileContent)
		if err != nil {
			return detectedModel, err
		}
		detectedModel = DetectedModel{Language: "Node.js", Models: models}
	default:
		return detectedModel, errors.New("unsupported file extension: " + fileExtension)
	}

	return detectedModel, nil
}
