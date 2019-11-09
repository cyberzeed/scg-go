package line

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBroadcastMessage(t *testing.T) {
	accessToken := ""
	secret := ""
	bot := NewBot(accessToken, secret)

	_, err := bot.BroadcastMessage("TestBroadcastMessage")

	assert.Nil(t, err)
}
