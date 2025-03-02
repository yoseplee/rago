package v2

// Document represents loaded data.
type Document string

type Documents []Document

func (ds Documents) AsStrings() []string {
	var result []string
	for _, d := range ds {
		result = append(result, string(d))
	}
	return result
}
