package v1

import (
	"encoding/json"
	"os"
)

// DocumentLoader is responsible for loading data from a specific type of source.
type DocumentLoader interface {
	Load() (Documents, error)
}

// StringDocumentLoader always loads a specific documents for test purpose.
type StringDocumentLoader struct {
	Strings []string
}

func (t StringDocumentLoader) Load() (Documents, error) {
	var documents Documents
	for _, param := range t.Strings {
		documents = append(documents, Document(param))
	}
	return documents, nil
}

// JSONFileDocumentLoader implements the DocumentLoader interface to load data from a JSON file.
// The JSON file must contain an array of JSON objects, each representing a Document.
type JSONFileDocumentLoader struct {
	FilePath string
}

func (j JSONFileDocumentLoader) Load() (Documents, error) {
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
