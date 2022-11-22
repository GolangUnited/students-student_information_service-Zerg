package postgres

import (
	"zerg-team-student-information-service/internal/models"
	"zerg-team-student-information-service/internal/storage"
)

type UserDB struct {
	dbConn storage.DBConnect
}

func NewUserDB(dbConn storage.DBConnect) *UserDB {
	return &UserDB{dbConn: dbConn}
}

func (u *UserDB) GetByID(id int) (models.User, error) {
	var userModel models.User
	row := u.dbConn.GetConn().QueryRow("SELECT * FROM users WHERE id=$1", id)
	err := row.Scan(
		&userModel.ID,
		&userModel.FirstName,
		&userModel.LastName,
		&userModel.Birthday,
		&userModel.Email,
		&userModel.PasswordHash,
	)
	return userModel, err
}

func (u *UserDB) GetByEmail(email string) (models.User, error) {
	var userModel models.User
	row := u.dbConn.GetConn().QueryRow("SELECT * FROM users WHERE email=$1", email)
	err := row.Scan(
		&userModel.ID,
		&userModel.FirstName,
		&userModel.LastName,
		&userModel.Birthday,
		&userModel.Email,
		&userModel.PasswordHash,
	)
	return userModel, err
}

func (u *UserDB) Create(user models.User) (int64, error) {
	sqlStatement := "INSERT INTO users (first_name, last_name, birthday, email, password_hash) "
	sqlStatement += "VALUES ($1, $2, $3, $4, $5) RETURNING id"
	result, err := u.dbConn.GetConn().Exec(
		sqlStatement,
		user.FirstName,
		user.LastName,
		user.Birthday,
		user.Email,
		user.PasswordHash,
	)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}
