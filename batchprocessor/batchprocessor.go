// Package batchprocessor enlists necessary functionality to read the xml from URL and
// create the necessary HTML pages as specified in the redbubble assignment.
package batchprocessor

import (
	"net/http"
	"log"
	"io/ioutil"
	"golang.org/x/sys/unix"
	"github.com/mohakkataria/redbubble_assignment/provider"
	"github.com/mohakkataria/redbubble_assignment/repository"
	"github.com/mohakkataria/redbubble_assignment/models"
	"github.com/mohakkataria/redbubble_assignment/htmlbuilder"
	"fmt"
)


// fetchURL makes a GET request to the required URL and reads the response body and returns it
func fetchURL(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return body
}


// buildHTML generates the HTML pages as specified in the assignment from the imagesProvider.
func buildHTML(provider provider.Provider, outputPath string) {
	images := provider.ConvertToImages()
	imageRepository := repository.Image{Images:images}
	allMakes := imageRepository.GetAllMakes()
	// build Index page
	indexPageTitle := "Index"
	firstTenImages := imageRepository.Find(10);
	indexPage := models.Page{Title:indexPageTitle, Navigation:allMakes, Gallery:firstTenImages}
	indexPageBuilder := htmlbuilder.Base{Page:indexPage, Dir:outputPath}
	indexPageBuilder.Output()
	fmt.Println(allMakes)
	fmt.Println(len(allMakes))
	// build make and model pages
	for _, make := range allMakes {
		if len(make) != 0 {
			modelsOfMake := imageRepository.GetAllModelsByMake(make)
			makePageNav := append(modelsOfMake, indexPageTitle)
			makePageGallery := imageRepository.FindByMake(make, 10)
			makePage := models.Page{Title:make, Navigation:makePageNav, Gallery:makePageGallery}
			fmt.Println(makePage)
			makePageBuilder := htmlbuilder.Base{Page:makePage, Dir:outputPath}
			makePageBuilder.Output()
			// build model pages of this make
			for _, model := range modelsOfMake {
				if len(model) != 0 {
					modelPageNav := []string{indexPageTitle, make}
					modelPageGallery := imageRepository.FindByMakeAndModel(make, model, 10)
					modelPage := models.Page{Title:model, Navigation:modelPageNav, Gallery:modelPageGallery}
					modelPageBuilder := htmlbuilder.Base{Page:modelPage, Dir:outputPath}
					modelPageBuilder.Output()
				}
			}
		}
	}
}


// Execute function is exposed to the main package for it to initiate the process of batch processing
// of images obtained from the URL and build the HTML. The url and outputPath are provided by the user or assumed
// default values as specified in the assignment
func Execute(url string, outputPath string) {

	if unix.Access(outputPath, unix.W_OK) != nil {
		log.Fatal("Cannot make directory writable", outputPath);
	}

	data := fetchURL(url)
	if len(data) == 0 {
		log.Fatal("Cannot retrieve content from the endpoint")
	}

	xmlProvider := provider.XML{
		XML: data,
	}

	buildHTML(xmlProvider, outputPath)
}
