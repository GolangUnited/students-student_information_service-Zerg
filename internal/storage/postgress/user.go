package postgress

import (
	"zerg-team-student-information-service/internal/models"
	"zerg-team-student-information-service/internal/storage"
)

type UserDb struct {
	dbConn storage.DBConnect
}

func NewUserDb(dbConn storage.DBConnect) *UserDb {
	return &UserDb{dbConn: dbConn}
}

func (u *UserDb) GetById(id int) (models.User, error) {
	var userModel models.User
	row := u.dbConn.GetConn().QueryRow("SELECT * FROM users WHERE id=$1", id)
	err := row.Scan(&userModel.Id, &userModel.FirstName, &userModel.LastName, &userModel.Birthday, &userModel.Email, &userModel.PasswordHash)
	return userModel, err
}

func (u *UserDb) Create(user models.User) (int64, error) {
	sqlStatement := "INSERT INTO users (first_name, last_name, birthday, email, password_hash) "
	sqlStatement += "VALUES ($1, $2, $3, $4, $5) RETURNING id"
	result, err := u.dbConn.GetConn().Exec(sqlStatement, user.FirstName, user.LastName, user.Birthday, user.Email, user.PasswordHash)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}
