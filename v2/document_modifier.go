package v2

// DocumentModifier is responsible for modifying data, such as splitting, mapping, marshalling.
type DocumentModifier interface {
	Modify(Documents) (Documents, error)
}

type DocumentModifiers []DocumentModifier

func (dm DocumentModifiers) ApplyAll(documents Documents) (Documents, error) {
	var modifiedDocuments = documents
	for _, modifier := range dm {
		var modifyErr error
		modifiedDocuments, modifyErr = modifier.Modify(modifiedDocuments)
		if modifyErr != nil {
			return nil, modifyErr
		}
	}
	return modifiedDocuments, nil
}
