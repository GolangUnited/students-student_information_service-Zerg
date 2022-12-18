package postgress

import (
	"database/sql"
	"fmt"
	"github.com/lib/pq"
	"zerg-team-student-information-service/internal/models"
	"zerg-team-student-information-service/internal/storage"
)

type MentorDB struct {
	dbConn storage.DBConnect
}

func NewMentorDB(dbConn storage.DBConnect) *MentorDB {
	return &MentorDB{dbConn: dbConn}
}

func (m *MentorDB) GetAll() ([]models.Mentor, error) {
	var mentors []models.Mentor
	var mentorsIDs []int

	rows, err := m.dbConn.GetConn().Query(`SELECT mentors.id, first_name, last_name, birthday, email
	                                    FROM mentors LEFT JOIN users ON users.id = mentors.id`)
	defer rows.Close()
	if err != nil {
		return mentors, fmt.Errorf("mentors select error: %w", err)
	}
	for rows.Next() {
		var mentor models.Mentor
		err = rows.Scan(&mentor.MentorID, &mentor.FirstName, &mentor.LastName, &mentor.Birthday, &mentor.Email)
		if err != nil {
			return mentors, fmt.Errorf("mentors scan error: %w", err)
		}
		mentors = append(mentors, mentor)
		mentorsIDs = append(mentorsIDs, mentor.MentorID)
	}

	mentorsGroups := make(map[int][]models.Group)
	rows, err = m.dbConn.GetConn().Query(`SELECT id, mentor FROM groups WHERE mentor = any ($1)`, pq.Array(mentorsIDs))
	if err != nil && err != sql.ErrNoRows {
		return mentors, fmt.Errorf("select groups error: %w", err)
	}
	for rows.Next() {
		var group models.Group
		err = rows.Scan(&group.GroupID, &group.MentorID)
		if err != nil {
			return mentors, fmt.Errorf("scan group error: %w", err)
		}
		mentorsGroups[group.MentorID] = append(mentorsGroups[group.MentorID], group)
	}
	for i, mentor := range mentors {
		for _, group := range mentorsGroups[mentor.MentorID] {
			mentors[i].Groups = append(mentors[i].Groups, group)
		}
	}

	return mentors, nil
}

func (m *MentorDB) GetByID(id int) (models.Mentor, error) {
	var mentorModel models.Mentor
	mentorModel.MentorID = id
	row := m.dbConn.GetConn().QueryRow(`SELECT users.first_name, users.last_name, users.birthday, users.email
	                                   FROM mentors LEFT JOIN users ON users.id = mentors.id
	                                   WHERE  mentors.id = $1`, id)
	err := row.Scan(
		&mentorModel.FirstName,
		&mentorModel.LastName,
		&mentorModel.Birthday,
		&mentorModel.Email,
	)
	if err != nil {
		return mentorModel, fmt.Errorf("mentor scan error: %w", err)
	}

	rows, err := m.dbConn.GetConn().Query(`SELECT id, mentor FROM groups WHERE mentor = $1`, id)
	defer rows.Close()
	if err != nil && err != sql.ErrNoRows {
		return mentorModel, fmt.Errorf("select groups error: %w", err)
	}
	for rows.Next() {
		var group models.Group
		err = rows.Scan(&group.GroupID, &group.MentorID)
		if err != nil {
			return mentorModel, fmt.Errorf("scan group error: %w", err)
		}
		mentorModel.Groups = append(mentorModel.Groups, group)
	}

	return mentorModel, nil
}

func (m *MentorDB) DeleteByID(id int) error {
	res, err := m.dbConn.GetConn().Exec("DELETE FROM mentors WHERE id = $1", id)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (m *MentorDB) Add(id int) error {
	_, err := m.dbConn.GetConn().Exec("INSERT INTO mentors (id) VALUES ($1)", id)
	return err
}
