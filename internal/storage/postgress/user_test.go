package postgress_test

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
	"zerg-team-student-information-service/internal/models"
	"zerg-team-student-information-service/internal/storage/postgress"
)

func TestCreateUser(t *testing.T) {
	dbConnect, _ := postgress.NewMockConnect()
	defer dbConnect.Close()

	user := models.User{1, "Ivan", "Sidorov", "30.01.1985",
		"ivan.sidorov@mail.ru", "sst67a8dfge638"}
	dbConnect.Mock.ExpectExec("INSERT INTO users").WithArgs(user.FirstName, user.LastName,
		user.Birthday, user.Email, user.PasswordHash).WillReturnResult(sqlmock.NewResult(5, 1))

	db := postgress.NewUserDB(dbConnect)
	userId, err := db.Create(user)

	assert.Equal(t, nil, err)
	assert.Equal(t, int64(5), userId)
}
