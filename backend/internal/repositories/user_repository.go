package repositories

import (
	"database/sql"
	"time"
	"vastfeel-backend/internal/models"
)

type UserRepository interface {
	Create(user *models.User) error
	GetByID(id int) (*models.User, error)
	GetAllUser() ([]models.User, error)
	Update(user *models.User) error
	Delete(id int) error
}

type UserRepositoryImpl struct {
	DB *sql.DB
}

// Create implements UserRepository.
func (u *UserRepositoryImpl) Create(user *models.User) error {
	query := `INSERT INTO users (username, email, password, role) VALUES ($1, $2, $3, $4) RETURNING id`
	return u.DB.QueryRow(query, user.Username, user.Email, user.Password, user.Role).Scan(&user.ID)
}

// Delete implements UserRepository.
func (u *UserRepositoryImpl) Delete(id int) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := u.DB.Exec(query, id)
	return err
}

// GetAllUser implements UserRepository.
func (u *UserRepositoryImpl) GetAllUser() ([]models.User, error) {
	users := []models.User{}
	query := `SELECT * FROM users`
	rows, err := u.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		user := models.User{}
		err = rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// GetByID implements UserRepository.
func (u *UserRepositoryImpl) GetByID(id int) (*models.User, error) {
	users := &models.User{}
	query := `SELECT * FROM users WHERE id = $1`
	err := u.DB.QueryRow(query, id).Scan(&users.ID, &users.Username, &users.Email, &users.Password, &users.Role, &users.CreatedAt, &users.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return users, nil
}

// Update implements UserRepository.
func (u *UserRepositoryImpl) Update(user *models.User) error {
	query := `UPDATE users SET username = $1, email = $2, password = $3, role = $4, updated_at = $5 WHERE id = $6`
	_, err := u.DB.Exec(query, user.Username, user.Email, user.Password, user.Role, time.Now(), user.ID)
	return err
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &UserRepositoryImpl{DB: db}
}
