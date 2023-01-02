<<<<<<< HEAD
package postgress_test

import (
	"database/sql"
	"errors"
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

	rows = sqlmock.NewRows([]string{"id", "first_name", "last_name", "birthday", "email", "password_hash"}).
		AddRow(-1, u.FirstName, u.LastName, u.Birthday, u.Email, u.PasswordHash)
	mockConn.Mock.ExpectQuery(query).WithArgs(u.ID).WillReturnError(sql.ErrNoRows)
	newUser, err = postgress.NewUserDB(mockConn).GetByID(u.ID)
	assert.Error(t, err)
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

	mockConn.Mock.ExpectExec(query).
		WithArgs(u.FirstName, u.LastName, u.Birthday, u.Email, u.PasswordHash).
		WillReturnResult(sqlmock.NewErrorResult(errors.New("duplicate key value violates unique constraint \"users_email_key\"")))
	gottenId, err = postgress.NewUserDB(mockConn).Create(u)
	assert.Error(t, err)
=======
package postgress

import (
	"log"
	"regexp"
	"testing"
	"time"
	"zerg-team-student-information-service/internal/models"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var (
	uD = models.UserData{
		FirstName: "Ivan",
		LastName:  "Ivanov",
		Birthday:  time.Time{}.AddDate(2000, 2, 2).String(),
		Email:     "test@mail.ru",
	}
	u = models.User{
		UserID:   1000,
		UserData: uD,
	}
)

func TestGetByEmail_Success(t *testing.T) {
	var passwordHash string
	mockConn, err := NewMockConnect()
	if err != nil {
		log.Fatal(err)
	}
	dbconn := &PostgresConnect{
		conn: mockConn.conn,
	}
	repo := NewUserDB(dbconn)

	query := regexp.QuoteMeta("SELECT * FROM users WHERE email=$1")

	rows := sqlmock.NewRows(
		[]string{
			"id",
			"first_name",
			"last_name",
			"birthday",
			"email",
			"password_hash",
		}).AddRow(u.UserID, u.FirstName, u.LastName, u.Birthday, u.Email, passwordHash)

	u.SetPasswordHash(passwordHash)

	mockConn.Mock.ExpectQuery(query).WithArgs(u.Email).WillReturnRows(rows)

	user, err := repo.GetByEmail(u.Email)

	assert.NoError(t, err)
	assert.Equal(t, u, user)
>>>>>>> development
}
