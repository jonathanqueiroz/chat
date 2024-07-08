package repositories

import (
	"github.com/jonathanqueiroz/chat/internal/app/models"
	"gorm.io/gorm"
)

// UserRepositoryInterface defines the methods that a UserRepository must have
type UserRepositoryInterface interface {
	Create(user *models.User) (*models.UserPublic, error)
	GetUsers() ([]*models.UserPublic, error)
	GetByID(id int) (*models.UserPublic, error)
	Update(id int, user *models.User) error
	Delete(userID int) error
}

type UserRepository struct {
	DB *gorm.DB
}

// NewUserRepository creates a new UserRepository
func NewUserRepository(db *gorm.DB) UserRepositoryInterface {
	return &UserRepository{
		DB: db,
	}
}

// Create creates a new user
func (r *UserRepository) Create(user *models.User) (*models.UserPublic, error) {
	if err := r.DB.Create(user).Error; err != nil {
		return nil, err
	}

	userPublic := &models.UserPublic{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Avatar:    user.Avatar,
		Bio:       user.Bio,
		CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	return userPublic, nil
}

// GetUsers gets all users
func (r *UserRepository) GetUsers() ([]*models.UserPublic, error) {
	users := []*models.User{}

	if err := r.DB.Order("username ASC").Find(&users).Error; err != nil {
		return nil, err
	}

	usersPublic := []*models.UserPublic{}

	for _, user := range users {
		usersPublic = append(usersPublic, &models.UserPublic{
			ID:        user.ID,
			Username:  user.Username,
			Email:     user.Email,
			Avatar:    user.Avatar,
			Bio:       user.Bio,
			CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return usersPublic, nil
}

// GetByID gets a user by ID
func (r *UserRepository) GetByID(id int) (*models.UserPublic, error) {
	user := &models.User{}
	if err := r.DB.First(&user, id).Error; err != nil {
		return nil, err
	}

	userPublic := &models.UserPublic{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Avatar:    user.Avatar,
		Bio:       user.Bio,
		CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	return userPublic, nil
}

// Update updates a user
func (r *UserRepository) Update(id int, user *models.User) error {
	return r.DB.Model(&models.User{}).Where("id = ?", id).Updates(user).Error
}

// Delete deletes a user
func (r *UserRepository) Delete(userID int) error {
	return r.DB.Delete(&models.User{}, userID).Error
}
