package repository

import (
	"petshop-pos/internal/entity"
	"petshop-pos/pkg/exception"

	"gorm.io/gorm"
)

type UserRepositoryIMPL struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepositoryIMPL {
    db.AutoMigrate(&entity.User{})
    return &UserRepositoryIMPL{db: db}
}

func (r *UserRepositoryIMPL) FindByUsername(username string) (*entity.User, *exception.Exception) {
    var user entity.User
    err := r.db.Where("username = ? AND deleted_at IS NULL", username).First(&user).Error
    if err != nil {
        return nil, exception.Internal("Failed to find user", err)
    }
    return &user, nil
}

func (r *UserRepositoryIMPL) FindByID(id string) (*entity.User, *exception.Exception) {
    var user entity.User
    err := r.db.Where("id = ? AND deleted_at IS NULL", id).First(&user).Error
    if err != nil {
        return nil, exception.Internal("Failed to find user", err)
    }
    return &user, nil
}

func (r *UserRepositoryIMPL) FindByUsernameAndTenant(username, tenantName string) (*entity.User, *exception.Exception) {
    var user entity.User
    err := r.db.Preload("Tenant").
        Joins("JOIN tenants ON users.tenant_id = tenants.id").
        Where("users.username = ? AND tenants.name = ?", username, tenantName).
        First(&user).Error
    
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            return nil, exception.NotFound("user not found")
        }
        return nil, exception.Internal("Failed to find user", err)
    }
    
    return &user, nil
}