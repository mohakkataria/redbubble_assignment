// Package htmlbuilder contains the functionality to create the necessary HTML files from a Page
// which describes the necessary HTHML elements of Title, name, navigation and Thumbnails to show
package htmlbuilder

import (
	"github.com/mohakkataria/redbubble_assignment/models"
	"os"
	"fmt"
)


// Base struct contains the Page to render and the output path to render it to
type Base struct {
	Page models.Page
	Dir string
}


// Output opens the file for write and generates the html and writes it to the file.
func (b Base) Output() {
	fileName := b.Page.Title + ".html"
	path := b.Dir + "/" + fileName
	var _, err = os.Stat(path)

	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if err != nil {
			fmt.Println(err.Error())
		}
		_, err = file.WriteString(b.render())
		err = file.Sync()
		defer file.Close()
	} else {
		var file, err = os.OpenFile(path, os.O_RDWR, 0644)
		if err != nil {
			fmt.Println(err.Error())
		}
		_, err = file.WriteString(b.render())
		err = file.Sync()
		defer file.Close()
	}
}


// render generates the HTML from Page properties and returns it as a string
func (b Base) render() string {

	html := "<h1>" + b.Page.Title + "</h1>"

	// render navigation
	html += "<ul>";
	for _, navItem := range b.Page.Navigation  {
		html += "<li><a href=\"" + navItem + ".html" + "\">" + navItem + "</a></li>";
	}
	html += "</ul>"

	// render gallery

	html += "<ul>";
	for _, image := range b.Page.Gallery {
		html += "<li><img src=\"" + image.Thumbnail + "\"/></li>"
	}
	html += "</ul>"

	return html
}
