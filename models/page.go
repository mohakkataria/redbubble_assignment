package models

// Page struct is representation of single HTML to be generated.
type Page struct {
	Title      string
	Navigation []string
	Gallery    []Image
}
