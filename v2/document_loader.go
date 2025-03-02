package v2

import (
	"encoding/json"
	"os"
)

// DocumentLoader is responsible for loading data from a specific type of source.
type DocumentLoader interface {
	Load() (Documents, error)
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

	var unmarshalledObjects []interface{}
	if unmarshalErr := json.Unmarshal(file, &unmarshalledObjects); unmarshalErr != nil {
		return nil, unmarshalErr
	}

	var documents Documents
	for _, unmarshalledObject := range unmarshalledObjects {
		marshalledObject, marshalErr := json.Marshal(unmarshalledObject)
		if marshalErr != nil {
			return nil, marshalErr
		}
		documents = append(documents, Document(marshalledObject))
	}

	return documents, nil
}
