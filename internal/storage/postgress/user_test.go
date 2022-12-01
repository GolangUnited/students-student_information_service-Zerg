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

var u = models.User{
	ID:           1000,
	FirstName:    "Ivan",
	LastName:     "Ivanov",
	Birthday:     time.Time{}.AddDate(2000, 2, 2),
	Email:        "test@mail.ru",
	PasswordHash: "hash",
}

func TestGetByEmail_Success(t *testing.T) {
	mockConn, err := NewMockConnect()
	if err != nil {
		log.Fatal(err)
	}
	dbconn := &PostgresConnect{
		conn: mockConn.conn,
	}
	repo := NewUserDB(dbconn)

	query := regexp.QuoteMeta("SELECT * FROM users WHERE email=$1")

	rows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "birthday", "email", "password_hash"}).
		AddRow(u.ID, u.FirstName, u.LastName, u.Birthday, u.Email, u.PasswordHash)

	mockConn.Mock.ExpectQuery(query).WithArgs(u.Email).WillReturnRows(rows)

	user, err := repo.GetByEmail(u.Email)

	assert.NoError(t, err)
	assert.Equal(t, u, user)
}
