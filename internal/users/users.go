package users

import (
	"zerg-team-student-information-service/internal/storage/db"
)

type User struct {
	Id           int
	FirstName    string
	LastName     string
	Birthday     string
	Email        string
	PasswordHash string
}

func (u *User) GetById(dbConn db.DBConnect) error {
	row := dbConn.GetConn().QueryRow("SELECT * FROM users WHERE id=$1", u.Id)
	err := row.Scan(&u.Id, &u.FirstName, &u.LastName, &u.Birthday, &u.Email, &u.PasswordHash)
	return err
}

func (u *User) Insert(dbConn db.DBConnect) error {
	sqlStatement := "INSERT INTO users (first_name, last_name, birthday, email, password_hash) "
	sqlStatement += "VALUES ($1, $2, $3, $4, $5) RETURNING id"
	row := dbConn.GetConn().QueryRow(sqlStatement, u.FirstName, u.LastName, u.Birthday, u.Email, u.PasswordHash)
	err := row.Scan(&u.Id)
	return err
}
