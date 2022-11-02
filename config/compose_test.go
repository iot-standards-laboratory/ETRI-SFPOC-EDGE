package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateCompose(t *testing.T) {
	assert := assert.New(t)

	err := CreateCompose()
	assert.NoError(err)
}
