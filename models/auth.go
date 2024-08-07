package models

import (
	"gorm.io/gorm"
	"time"
)

type AuthUser struct {
	gorm.Model
	Password    string     `gorm:"size:128;not null;column:password"`
	LastLogin   *time.Time `gorm:"default:null;column:last_login"` // Optional field
	Username    string     `gorm:"size:50;not null;unique;column:username"`
	IsSuperuser bool       `gorm:"not null;default:false;column:is_superuser"`
	IsStaff     bool       `gorm:"not null;default:false;column:is_staff"`
	IsActive    bool       `gorm:"not null;default:true;column:is_active"`
	IsTest      bool       `gorm:"not null;default:false;column:is_test"`
	DateJoined  time.Time  `gorm:"not null;default:CURRENT_TIMESTAMP;column:date_joined"`
	IsStudent   bool       `gorm:"not null;default:true;column:is_student"`
	IsTeacher   bool       `gorm:"not null;default:false;column:is_teacher"`
}

type UserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
