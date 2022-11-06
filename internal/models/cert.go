package models

import "time"

type Cert struct {
	StudentID int
	CertID    string
	Issued    time.Time
}
