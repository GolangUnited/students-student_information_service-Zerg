package models

import "time"

type Cert struct {
	CertID    int       `json:"cert_id"`
	CertCode  string    `json:"cert_code"`
	StudentID int       `json:"student_id"`
	Issued    time.Time `json:"issued"`
}
