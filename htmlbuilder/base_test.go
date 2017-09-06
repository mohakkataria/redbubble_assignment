package htmlbuilder

import (
	"github.com/mohakkataria/redbubble_assignment/models"
	"os"
	"testing"
)

func TestBase_Output(t *testing.T) {
	indexPage := models.Page{Title: "Test", Navigation: []string{"test"}, Gallery: []models.Image{{Make: "test", Model: "test", Thumbnail: "url", ID: 1}}}
	outputPath := "./../output_directory"
	indexPageBuilder := Base{Page: indexPage, Dir: outputPath}
	textOutput := indexPageBuilder.Output()

	if textOutput != "<h1>Test</h1><ul><li><a href=\"test.html\">test</a></li></ul><ul><li><img src=\"url\"/></li></ul>" {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", "<h1>Test</h1><ul><li><a href=\"test.html\">test</a></li></ul><ul><li><img src=\"url\"/></li></ul>", textOutput)
	}

	os.Remove(outputPath + "/Test.html")
}

func TestBase_Output2(t *testing.T) {
	indexPage := models.Page{Title: "Test", Navigation: []string{"test"}, Gallery: []models.Image{{Make: "test", Model: "test", Thumbnail: "url", ID: 1}}}
	outputPath := "./../output_directory"
	indexPageBuilder := Base{Page: indexPage, Dir: outputPath}
	textOutput := indexPageBuilder.Output()

	if textOutput != "<h1>Test</h1><ul><li><a href=\"test.html\">test</a></li></ul><ul><li><img src=\"url\"/></li></ul>" {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", "<h1>Test</h1><ul><li><a href=\"test.html\">test</a></li></ul><ul><li><img src=\"url\"/></li></ul>", textOutput)
	}
}
