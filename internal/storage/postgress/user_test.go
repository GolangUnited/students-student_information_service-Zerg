package postgress_test

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
	"zerg-team-student-information-service/internal/models"
	"zerg-team-student-information-service/internal/storage/postgress"
)

func createTestUser() models.User {
	return models.User{
		ID:           1,
		FirstName:    "John",
		LastName:     "Doe",
		Birthday:     "1999-01-01",
		Email:        "test@example.com",
		PasswordHash: "hash_string",
	}
}

func TestUserDb_GetById(t *testing.T) {
	mockConn, err := postgress.NewMockConnect()
	assert.NoError(t, err)

	u := createTestUser()
	query := `SELECT \* FROM users WHERE id=\$1`
	rows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "birthday", "email", "password_hash"}).
		AddRow(u.ID, u.FirstName, u.LastName, u.Birthday, u.Email, u.PasswordHash)
	mockConn.Mock.ExpectQuery(query).WithArgs(u.ID).WillReturnRows(rows)

	newUser, err := postgress.NewUserDB(mockConn).GetByID(u.ID)
	assert.NoError(t, err)
	assert.Equal(t, u, newUser)
}

func TestUserDb_Create(t *testing.T) {
	mockConn, err := postgress.NewMockConnect()
	assert.NoError(t, err)

	u := createTestUser()
	wantedId := int64(223)
	query := `INSERT INTO users \(first_name, last_name, birthday, email, password_hash\) `
	query += `VALUES \(\$1, \$2, \$3, \$4, \$5\) RETURNING id`
	mockConn.Mock.ExpectExec(query).
		WithArgs(u.FirstName, u.LastName, u.Birthday, u.Email, u.PasswordHash).
		WillReturnResult(sqlmock.NewResult(wantedId, 1))

	gottenId, err := postgress.NewUserDB(mockConn).Create(u)
	assert.NoError(t, err)
	assert.Equal(t, wantedId, gottenId)
}
