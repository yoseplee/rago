package v2

// DocumentLoader is responsible for loading data from a specific type of source.
type DocumentLoader interface {
	Load() (Documents, error)
}
