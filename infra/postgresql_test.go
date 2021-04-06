package infra

import (
	"github.com/locpham24/go-coffee/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitDB(t *testing.T) {
	configs := config.Config{
		PostgreSQL: struct {
			Username string
			Password string
			Host     string
			Port     string
			Name     string
			Debug    bool
		}{Username: "default", Password: "secret", Host: "0.0.0.0", Port: "5432", Name: "default", Debug: false}}

	err := InitPostgreSQL(configs)
	assert.NoError(t, err)
}
