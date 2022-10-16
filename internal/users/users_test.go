package users_test

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
	"zerg-team-student-information-service/internal/storage/db"
	"zerg-team-student-information-service/internal/users"
)

func createTestUser() users.User {
	var user users.User
	user.Id = 1
	user.FirstName = "John"
	user.LastName = "Doe"
	user.Birthday = "1999-01-01"
	user.Email = "test@example.com"
	user.PasswordHash = "hash_string"

	return user
}

func TestUser_GetById(t *testing.T) {
	mockConn, err := db.NewMockConnect()
	assert.NoError(t, err)

	u := createTestUser()

	query := `SELECT \* FROM users WHERE id=\$1`
	rows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "birthday", "email", "password_hash"}).
		AddRow(u.Id, u.FirstName, u.LastName, u.Birthday, u.Email, u.PasswordHash)

	mockConn.Mock.ExpectQuery(query).WithArgs(u.Id).WillReturnRows(rows)

	var newUser users.User
	newUser.Id = u.Id
	err = newUser.GetById(mockConn)
	assert.NoError(t, err)
	assert.Equal(t, u, newUser)
}

func TestUser_Insert(t *testing.T) {
	mockConn, err := db.NewMockConnect()
	assert.NoError(t, err)

	u := createTestUser()
	newUserId := 223
	query := `INSERT INTO users \(first_name, last_name, birthday, email, password_hash\) `
	query += `VALUES \(\$1, \$2, \$3, \$4, \$5\) RETURNING id`
	mockConn.Mock.ExpectExec(query).WithArgs(u.FirstName, u.LastName, u.Birthday, u.Email, u.PasswordHash).WillReturnResult(sqlmock.NewResult(int64(newUserId), 1))

	err = u.Insert(mockConn)
	assert.NoError(t, err)
	assert.Equal(t, newUserId, u.Id)
}
