package models

type Embedding interface {
	Dimension() Dimension
	Vector() Vector
}
