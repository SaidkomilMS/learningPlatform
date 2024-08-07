package interfaces

import (
	"context"
	"learningPlatform/models"
)

type UserRepository interface {
	FindByID(ctx context.Context, id uint) (*models.AuthUser, error)
	GetActiveUserByUsername(ctx context.Context, username string) (*models.AuthUser, error)
	UpdateLastLogin(ctx context.Context, userID uint) error
	Create(ctx context.Context, user *models.AuthUser) error
	Update(ctx context.Context, user *models.AuthUser) error
	Delete(ctx context.Context, id uint) error
}
