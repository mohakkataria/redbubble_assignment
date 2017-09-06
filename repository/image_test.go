package repository

import (
	"testing"
	"github.com/mohakkataria/redbubble_assignment/models"
	"fmt"
)

func TestImage_FindByMake(t *testing.T) {
	imageRepository := Image{Images:[]models.Image{{Make:"test", Model:"Test", ID:1, Thumbnail:"url"}}}
	actual := imageRepository.FindByMake("test", 2)
	fmt.Println(actual)
	if len(actual) == 1 && actual[0].Make != "test" {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", "test", actual[0].Make)
	}
}

func TestImage_FindByMakeAndModel(t *testing.T) {
	imageRepository := Image{Images:[]models.Image{{Make:"test", Model:"Test", ID:1, Thumbnail:"url"}}}
	actual := imageRepository.FindByMakeAndModel("test", "test", 3)
	fmt.Println(actual)
	if len(actual) == 1 && (actual[0].Make != "test" || actual[0].Model != "Test") {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", "test", actual[0].Make)
	}
}

func TestImage_Find(t *testing.T) {
	imageRepository := Image{Images:[]models.Image{{Make:"test", Model:"Test", ID:1, Thumbnail:"url"}}}
	actual := imageRepository.Find(2)
	fmt.Println(actual)
	if len(actual) == 1 && (actual[0].Make != "test" || actual[0].Model != "Test") {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", "test", actual[0].Make)
	}
}

func TestImage_GetAllMakes(t *testing.T) {
	imageRepository := Image{Images:[]models.Image{{Make:"test", Model:"Test", ID:1, Thumbnail:"url"}}}
	actual := imageRepository.GetAllMakes()
	fmt.Println(actual)
	if len(actual) == 1 && actual[0] != "test"  {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", "test", actual[0])
	}
}

func TestImage_GetAllModelsByMake(t *testing.T) {
	imageRepository := Image{Images:[]models.Image{{Make:"test", Model:"Test", ID:1, Thumbnail:"url"}}}
	actual := imageRepository.GetAllModelsByMake("test")
	fmt.Println(actual)
	if len(actual) == 1 && actual[0] != "Test"  {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", "Test", actual[0])
	}
}

