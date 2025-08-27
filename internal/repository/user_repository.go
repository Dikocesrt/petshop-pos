package repository

import (
	"petshop-pos/internal/entity"
	"petshop-pos/pkg/exception"
)

type UserRepository interface {
	FindByUsername(username string) (*entity.User, *exception.Exception)
	FindByID(id string) (*entity.User, *exception.Exception)
	FindByUsernameAndTenant(username, tenantName string) (*entity.User, *exception.Exception)
}