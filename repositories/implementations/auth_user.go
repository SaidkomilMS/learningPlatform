package implementations

import (
	"context"
	"gorm.io/gorm"
	"learningPlatform/models"
	"learningPlatform/repositories/interfaces"
	"time"
)

type GormUserRepository struct {
	DB *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) interfaces.UserRepository {
	return &GormUserRepository{DB: db}
}

func (repo *GormUserRepository) FindByID(ctx context.Context, id uint) (*models.AuthUser, error) {
	var user models.AuthUser
	result := repo.DB.First(&user, id)
	return &user, result.Error
}

func (repo *GormUserRepository) GetActiveUserByUsername(ctx context.Context, username string) (*models.AuthUser, error) {
	var user models.AuthUser
	if err := repo.DB.WithContext(ctx).Where("username = ? and is_active = True", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *GormUserRepository) Create(ctx context.Context, user *models.AuthUser) error {
	return repo.DB.Create(user).Error
}

func (repo *GormUserRepository) Update(ctx context.Context, user *models.AuthUser) error {
	return repo.DB.Save(user).Error
}

func (repo *GormUserRepository) Delete(ctx context.Context, id uint) error {
	return repo.DB.Delete(&models.AuthUser{}, id).Error
}

func (repo *GormUserRepository) UpdateLastLogin(ctx context.Context, userID uint) error {
	return repo.DB.Model(&models.AuthUser{}).Where("id = ?", userID).Update("last_login", time.Now()).Error
}
