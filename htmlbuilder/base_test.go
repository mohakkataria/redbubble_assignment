package htmlbuilder

import (
	"github.com/mohakkataria/redbubble/htmlbuilder"
	"github.com/mohakkataria/redbubble_assignment/models"
	"testing"
	//"fmt"
	//"io/ioutil"
)

func TestBase_Output(t *testing.T) {
	indexPage := models.Page{Title: "Test", Navigation: []string{"test"}, Gallery: []models.Image{{Make: "test", Model: "test", Thumbnail: "url", ID: 1}}}
	outputPath := "./../output_directory"
	indexPageBuilder := htmlbuilder.Base{Page: indexPage, Dir: outputPath}
	textOutput := indexPageBuilder.Output()

	/*text, err := ioutil.ReadFile(outputPath+"/Test.html") // just pass the file name
	if err != nil {
		t.Errorf("File read error", err)
		fmt.Print(err)
	}*/

	if textOutput != "<h1>Test</h1><ul><li><a href=\"test.html\">test</a></li></ul><ul><li><img src=\"url\"/></li></ul>" {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", "<h1>Test</h1><ul><li><a href=\"test.html\">test</a></li></ul><ul><li><img src=\"url\"/></li></ul>", textOutput)
	}
}
