package models

import "time"

type Cert struct {
	StudentID int       `json:"student_id"`
	CertID    string    `json:"cert_id"`
	Issued    time.Time `json:"issued"`
}
