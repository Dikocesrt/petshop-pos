package repository

import (
	"context"
	"petshop-pos/pkg/exception"
)

type TenantRepository interface {
	FindIDByName(ctx context.Context, name string) (string, *exception.Exception)
}