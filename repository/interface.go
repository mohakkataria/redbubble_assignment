// Package repository contains the functionality to store and query the desired data store for parsed
// Images. We can make use of the interface to implement an SQL store or otherwise to use any kind of store
package repository

import "github.com/mohakkataria/redbubble_assignment/models"

// Repository interface determines the contract for every repository to follow to be used for the assignment
type Repository interface {
	GetAllMakes() []string
	GetAllModelsByMake(make string) []string
	FindByMake(make string, limit int) []models.Image
	FindByMakeAndModel(make string, model string, limit int) []models.Image
	Find(limit int) []models.Image
}
