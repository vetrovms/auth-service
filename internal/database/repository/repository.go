package repository

import (
	"auth/internal/models"
	"context"

	"gorm.io/gorm"
)

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

// UserExistsByEmail Перевірка існування користувача за id.
func (r *Repo) UserExistsById(ctx context.Context, id int) (bool, error) {
	var exists bool
	err := r.db.WithContext(ctx).Model(models.User{}).Select("count(*) > 0").Where("id = ?", id).Find(&exists).Error
	return exists, err
}

// GetUserByEmail Повертає користувача за email.
func (r *Repo) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := r.db.WithContext(ctx).Where("email = ?", email).Find(&user).Error
	return &user, err
}

// GetClientById Повертає клієнта за client_id.
func (r *Repo) GetClientByClientId(ctx context.Context, client_id string) (*models.Client, error) {
	var client models.Client
	err := r.db.WithContext(ctx).Where("client_id", client_id).Find(&client).Error
	return &client, err
}

// Save Зберігає користувача.
func (r *Repo) Save(ctx context.Context, user models.User) error {
	return r.db.WithContext(ctx).Save(&user).Error
}
