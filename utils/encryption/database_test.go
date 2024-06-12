package encryption

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDatabaseConnection_Connect(t *testing.T) {
	dbConnection := DatabaseConnection{
		Host:     "127.0.0.1",
		Port:     5432,
		User:     "postgres",
		Password: "",
	}
	_, err := dbConnection.Connect()
	assert.NoError(t, err)
}