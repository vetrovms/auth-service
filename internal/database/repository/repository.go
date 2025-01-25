package repository

import (
	"auth/internal/models"
	"context"

	"gorm.io/gorm"
)

// IRepo Інтерфейс репозиторія.
type IRepo interface {
	UserExistsByEmail(ctx context.Context, email string) (bool, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	Save(ctx context.Context, user models.User) error
}

// Repo Репозиторій.
type Repo struct {
	db *gorm.DB
}

// NewRepo Конструктор репозиторія.
func NewRepo(db *gorm.DB) Repo {
	return Repo{
		db: db,
	}
}

// UserExistsByEmail Перевірка існування користувача за email.
func (r *Repo) UserExistsByEmail(ctx context.Context, email string) (bool, error) {
	var exists bool
	// r.db.WithContext(ctx).Exec("select pg_sleep(10);") // @debug
	err := r.db.WithContext(ctx).Model(models.User{}).Select("count(*) > 0").Where("email = ?", email).Find(&exists).Error
	return exists, err
}

// GetUserByEmail Повертає користувача за email.
func (r *Repo) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := r.db.WithContext(ctx).Where("email = ?", email).Find(&user).Error
	return &user, err
}

// Save Зберігає користувача.
func (r *Repo) Save(ctx context.Context, user models.User) error {
	return r.db.WithContext(ctx).Save(&user).Error
}
