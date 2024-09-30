package repositories_test

import (
	"fmt"
	"testing"
	"time"
	"vastfeel-backend/internal/models"
	"vastfeel-backend/internal/repositories"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestShouldCreateUser(t *testing.T){
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	user := &models.User{
		Username: "test",
		Email: "test",
		Password: "test",
		Role: "guest",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	rows := mock.NewRows([]string{"id"}).AddRow(1)

	mock.ExpectQuery("INSERT INTO users").WithArgs(user.Username, user.Email, user.Password, user.Role).WillReturnRows(rows)

	repo := repositories.NewUserRepository(db)

	if err := repo.Create(user); err != nil {
		t.Errorf("error was not expected while creating user: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestShouldCreateUserOnFailure(t *testing.T){
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	user := &models.User{
		Username: "test",
		Email: "test",
		Password: "test",
		Role: "guest",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mock.ExpectQuery("INSERT INTO users").WithArgs(user.Username, user.Email, user.Password, user.Role).WillReturnError(fmt.Errorf("Error create user"))

	repo := repositories.NewUserRepository(db)

	if err := repo.Create(user); err == nil {
		t.Errorf("error was expected while creating user")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestShouldUpdateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	user := &models.User{
		ID: 1,
		Username: "test",
		Email: "test",
		Password: "test",
		Role: "test",
		UpdatedAt: time.Now(),
	}

	mock.ExpectExec("UPDATE users").WithArgs(user.Username, user.Email, user.Password, user.Role, user.UpdatedAt, user.ID).WillReturnResult(sqlmock.NewResult(1, 1))

	repo := repositories.NewUserRepository(db)
	
	if err := repo.Update(user); err != nil {
		t.Errorf("error was not expected while updating user: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestShouldUpdateUserOnFailure(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	user := &models.User{
		ID: 1,
		Username: "test",
		Email: "test",
		Password: "test",
		Role: "test",
		UpdatedAt: time.Now(),
	}

	mock.ExpectExec("UPDATE users").WithArgs(user.Username, user.Email, user.Password, user.Role, user.UpdatedAt, user.ID).WillReturnError(fmt.Errorf("Error update user"))

	repo := repositories.NewUserRepository(db)

	if err := repo.Update(user); err == nil {
		t.Errorf("error was expected while updating user")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestShouldGetAllUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "username", "email", "password", "role", "created_at", "updated_at"}).
		AddRow(1,"Alam", "email@email.com", "password231", "admin", time.Now(), time.Now()).
		AddRow(2,"Bark", "test@tes.com", "password123", "guest", time.Now(), time.Now())

	mock.ExpectQuery("^SELECT \\* FROM users$").WillReturnRows(rows)

	repo := repositories.NewUserRepository(db)

	users, err := repo.GetAllUser()
	if err != nil {
		t.Errorf("error was not expected while getting all user: %s", err)
	}

	assert.NoError(t, err)
	assert.Len(t, users, 2)
	assert.Equal(t, "Alam", users[0].Username)
	assert.Equal(t, "Bark", users[1].Username)
	

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestShouldGetAllUserOnFailure(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("^SELECT \\* FROM users$").WillReturnError(fmt.Errorf("Errors Get All"))

	repo := repositories.NewUserRepository(db)

	if _, err := repo.GetAllUser(); err == nil {
		t.Errorf("error was expected while getting all user")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestShouldGetUserById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	
	user := &models.User{
		ID: 1,
		Username: "test",
		Email: "test",
		Password: "test",
		Role: "guest",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	rows := sqlmock.NewRows([]string{"id", "username", "email", "password", "role", "created_at", "updated_at"}).AddRow(user.ID, user.Username, user.Email, user.Password, user.Role, user.CreatedAt, user.UpdatedAt)

	mock.ExpectQuery("^SELECT \\* FROM users WHERE id=\\$1").WithArgs(1).WillReturnRows(rows)

	repo := repositories.NewUserRepository(db)

	if _, err := repo.GetByID(1); err != nil {
		t.Errorf("error was not expected while getting user by id: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestShouldGetUserByIdOnFailure(t *testing.T){
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("^SELECT \\* FROM users WHERE id=\\$1").WithArgs(1).WillReturnError(fmt.Errorf("Errors Get User By Id"))

	repo := repositories.NewUserRepository(db)

	if _, err := repo.GetByID(1); err == nil {
		t.Errorf("error was expected while getting user by id")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestShouldDeleteUser(t *testing.T){
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec("^DELETE FROM users WHERE id=\\$1").WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))

	repo := repositories.NewUserRepository(db)

	if err := repo.Delete(1); err != nil {
		t.Errorf("error was not expected while deleting user: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestShouldDeleteUserOnFailure(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec("^DELETE FROM users WHERE id=\\$1").WithArgs(1).WillReturnError(fmt.Errorf("Errors Delete User"))

	repo := repositories.NewUserRepository(db)

	if err := repo.Delete(1); err == nil {
		t.Errorf("error was expected while deleting user")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}