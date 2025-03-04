package v1

import (
	"encoding/json"
	"os"
)

// DocumentLoader is responsible for loading data from a specific type of source.
type DocumentLoader interface {
	Load() (Documents, error)
}

// TestDocumentLoader always loads a specific documents for test purpose.
type TestDocumentLoader struct{}

func (t TestDocumentLoader) Load() (Documents, error) {
	return []Document{}, nil
}

// JSONDocumentLoader implements the DocumentLoader interface to load data from a JSON file.
// The JSON file must contain an array of JSON objects, each representing a Document.
type JSONDocumentLoader struct {
	FilePath string
}

func (j JSONDocumentLoader) Load() (Documents, error) {
	file, readFileErr := os.ReadFile(j.FilePath)
	if readFileErr != nil {
		return nil, readFileErr
	}

	var jsonArray []interface{}
	if unmarshalErr := json.Unmarshal(file, &jsonArray); unmarshalErr != nil {
		return nil, unmarshalErr
	}

	var documents Documents
	for _, jsonObject := range jsonArray {
		marshalledObject, marshalErr := json.Marshal(jsonObject)
		if marshalErr != nil {
			return nil, marshalErr
		}
		documents = append(documents, Document(marshalledObject))
	}

	return documents, nil
}
