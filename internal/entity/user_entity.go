package entity

import (
	"gorm.io/gorm"
)

type User struct {
    ID          string `gorm:"type:char(36);primaryKey;default:(UUID())"`
    Name        string `gorm:"type:varchar(255);not null"`
    Username    string `gorm:"type:varchar(255);not null;unique"`
    Password    string `gorm:"type:varchar(255);not null"`
    PhoneNumber string `gorm:"type:varchar(255)"`
    Role        string `gorm:"type:varchar(255);not null"`
    TenantID    string `gorm:"type:char(36);not null;index"`
    Tenant      Tenant `gorm:"foreignKey:TenantID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
    gorm.Model
}

func (User) TableName() string {
    return "users"
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
    if u.ID == "" {
        tx.Statement.SetColumn("id", gorm.Expr("UUID()"))
    }
    return nil
}