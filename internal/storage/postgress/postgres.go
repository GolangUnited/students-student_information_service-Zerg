package postgress

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"zerg-team-student-information-service/internal/storage"
)

type PGConfig struct {
	Host     string
	Port     int
	Username string
	DBName   string
	Password string
	SSLMode  string
}

func (cfg *PGConfig) CreateConnectString() string {
	return fmt.Sprintf("host=%s port=%v user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode)
}

type PostgresConnect struct {
	conn *sql.DB
}

func NewPostgresConnect(cfg storage.DbConfig) (*PostgresConnect, error) {
	conn, err := sql.Open("postgres", cfg.CreateConnectString())

	if err != nil {
		return nil, err
	}

	return &PostgresConnect{
		conn: conn,
	}, nil
}

func (db *PostgresConnect) CheckConn() error {

	err := db.conn.Ping()

	if err != nil {
		return err
	}

	return nil
}

func (db *PostgresConnect) GetConn() *sql.DB {
	return db.conn
}

func (db *PostgresConnect) Close() error {
	return db.conn.Close()
}
