package config_test

import (
	"testing"

	"github.com/chadsmith12/pacer/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestDbConfigString(t *testing.T) {
    dbConfig := config.DatabaseConfig {
        Username: "postgres",
        Password: "secretpassword123",
        Port: "5432",
        Host: "localhost",
        SSLMode: false,
        DatabaseName: "postgres",
    }

    actual := dbConfig.String()
    expected := "postgresql://postgres:secretpassword123@localhost:5432/postgres"

    assert.Equal(t, expected, actual)
}
