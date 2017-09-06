package repository

import (
	"github.com/mohakkataria/redbubble_assignment/models"
	"reflect"
)


// Image repository contains the slice of Images
type Image struct {
	Images []models.Image
}


// GetAllMakes returns all the unique non empty makes of the Images in the repository
func (i Image) GetAllMakes() []string {
	imageMakes := map[string]int{}
	for _, image := range i.Images {
		imageMakes[image.Make] = 1
	}
	makes := reflect.ValueOf(imageMakes).MapKeys()
	makesToReturn := []string{}
	for _, make := range makes {
		makesToReturn = append(makesToReturn, make.Interface().(string))
	}
	return makesToReturn
}

// GetAllModelsByMake returns all the unique non empty models of a given make of the Images in the repository
func (i Image) GetAllModelsByMake(make string) []string {
	imageModels := map[string]int{}
	for _, image := range i.Images {
		if image.Make == make {
			imageModels[image.Model] = 1
		}

	}
	models := reflect.ValueOf(imageModels).MapKeys()
	modelsToReturn := []string{}
	for _, model := range models {
		modelsToReturn = append(modelsToReturn, model.Interface().(string))
	}
	return modelsToReturn
}


// FindByMake returns given number of Images of given make from the repository
func (i Image) FindByMake(make string, limit int) []models.Image {
	images := []models.Image{}
	for _, image := range i.Images {
		if image.Make == make {
			images = append(images, image)
		}
	}

	if limit > 0 {
		length := 0
		if length = limit; limit > len(images) {
			length = len(images)
		}
		return images[:length];
	}

	return images
}

// FindByMakeAndModel returns given number of Images of given make and model from the repository
func (i Image) FindByMakeAndModel(make string, model string, limit int) []models.Image {
	images := []models.Image{}
	for _, image := range i.Images {
		if image.Make == make && image.Model == model {
			images = append(images, image)
		}
	}

	if limit > 0 {
		length := 0
		if length = limit; limit > len(images) {
			length = len(images)
		}
		return images[:length];
	}

	return images
}

// Find returns given number of Images from the repository
func (i Image) Find(limit int) []models.Image {
	if limit > 0 {
		length := 0
		if length = limit; limit > len(i.Images) {
			length = len(i.Images)
		}
		return i.Images[:length];

	}
	return i.Images
}
