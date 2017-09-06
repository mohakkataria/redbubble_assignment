package provider

import "github.com/mohakkataria/redbubble_assignment/models"


// Provider Interface to support multiple other sources and use cases, if any in future
type Provider interface {
	ConvertToImages()  []models.Image
}
