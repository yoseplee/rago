package v1

import (
	"encoding/json"
	"os"
)

// Loader is an interface that defines the behavior of a data loader.
type Loader interface {
	Load() (Documents, error)
}

type JSONLoader struct {
	FilePath string
}

func (J JSONLoader) Load() (Documents, error) {
	jsonFile, fileReadErr := os.ReadFile(J.FilePath)
	if fileReadErr != nil {
		return nil, fileReadErr
	}

	var unmarshalledObjects []interface{}
	if unmarshalErr := json.Unmarshal(jsonFile, &unmarshalledObjects); unmarshalErr != nil {
		return nil, unmarshalErr
	}

	var documents Documents
	for _, unmarshalledObject := range unmarshalledObjects {
		marshalledObject, marshalErr := json.Marshal(unmarshalledObject)
		if marshalErr != nil {
			return nil, marshalErr
		}

		documents = append(documents, string(marshalledObject))
	}

	return documents, nil
}
