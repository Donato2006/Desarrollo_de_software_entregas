package dao

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnect_InicializaDB(t *testing.T) {
	Connect()

	assert.NotNil(t, DB)
}
