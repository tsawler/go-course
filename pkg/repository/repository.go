package repository

import "tsawler/go-course/pkg/models"

type DatabaseRepo interface {
	AllUsers() ([]models.User, error)
}
