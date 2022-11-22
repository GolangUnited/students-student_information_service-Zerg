package postgress_test

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
	"zerg-team-student-information-service/internal/models"
	"zerg-team-student-information-service/internal/storage/postgress"
)

var userModel = models.User{ID: 7, FirstName: "Ivan", LastName: "Sidorov", Birthday: "30.01.1985",
	Email: "ivan.sidorov@mail.ru", PasswordHash: "sst67a8dfge638"}

var columns = []string{"id", "first_name", "last_name", "birthday", "email", "password_hash"}

func TestCreateUser(t *testing.T) {
	dbConnect, _ := postgress.NewMockConnect()
	defer dbConnect.Close()

	userId := int64(5)
	dbConnect.Mock.ExpectExec("INSERT INTO users").
		WithArgs(userModel.FirstName, userModel.LastName,
			userModel.Birthday, userModel.Email, userModel.PasswordHash).
		WillReturnResult(sqlmock.NewResult(userId, 1))

	db := postgress.NewUserDB(dbConnect)
	id, err := db.Create(userModel)

	assert.Equal(t, nil, err)
	assert.Equal(t, userId, id)
}

func TestGetByEmail(t *testing.T) {
	dbConnect, _ := postgress.NewMockConnect()
	defer dbConnect.Close()

	dbConnect.Mock.ExpectQuery(`SELECT \* FROM users WHERE email=(.+)`).
		WithArgs(userModel.Email).
		WillReturnRows(sqlmock.NewRows(columns).AddRow(userModel.ID, userModel.FirstName, userModel.LastName,
			userModel.Birthday, userModel.Email, userModel.PasswordHash))

	db := postgress.NewUserDB(dbConnect)
	user, err := db.GetByEmail(userModel.Email)

	assert.Equal(t, nil, err)
	assert.EqualValues(t, userModel, user)
}

func TestGetByID(t *testing.T) {
	dbConnect, _ := postgress.NewMockConnect()
	defer dbConnect.Close()

	dbConnect.Mock.ExpectQuery(`SELECT \* FROM users WHERE id=(.+)`).
		WithArgs(userModel.ID).
		WillReturnRows(sqlmock.NewRows(columns).AddRow(userModel.ID, userModel.FirstName, userModel.LastName,
			userModel.Birthday, userModel.Email, userModel.PasswordHash))

	db := postgress.NewUserDB(dbConnect)
	user, err := db.GetByID(userModel.ID)

	assert.Equal(t, nil, err)
	assert.EqualValues(t, userModel, user)
}
