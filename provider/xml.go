package provider

import (
	"encoding/xml"
	"github.com/mohakkataria/redbubble_assignment/models"
)

// XML struct contains the raw byte XML and its parsed format
type XML struct {
	XML       []byte
	parsedXML models.XML
}

// ConvertToImages is used to convert the unparsed xml to slice of Images after parsing the XML
// to all the metadata received from the fetched URL
func (xp XML) ConvertToImages() []models.Image {
	err := xml.Unmarshal(xp.XML, &xp.parsedXML)
	if err != nil {
		panic(err)
	}
	images := []models.Image{}
	works := xp.parsedXML.Work

	for _, work := range works {
		image := models.Image{}
		image.ID = work.ID
		for _, url := range work.URLs {
			if url.Type == "small" {
				image.Thumbnail = url.URL
			}
		}
		image.Make = work.EXIF.Make
		image.Model = work.EXIF.Model
		images = append(images, image)
	}

	return images

}
