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
}
