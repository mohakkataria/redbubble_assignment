package models

// Image struct is the representation of Image parsed from the given data provider. In case of the assignment, this
// is the XML received.
type Image struct {
	ID        int
	Make      string
	Model     string
	Thumbnail string
}
