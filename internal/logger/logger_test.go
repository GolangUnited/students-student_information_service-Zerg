package logger_test

import (
	"testing"
	"zerg-team-student-information-service/internal/logger"

	"github.com/stretchr/testify/assert"
)

func TestNewLogrusLogger(t *testing.T) {
	l := logger.NewLogrusLogger()
	assert.Implements(t, (*logger.Logger)(nil), l, "LogrusLogger not implementing Logger interface")
}
