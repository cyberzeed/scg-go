package line

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSeriesCommand(t *testing.T) {
	commands := []string{
		"series:0",
		"SERIES:0",
		"SeRiEs:0",
		" series  :  0 ",
	}

	for _, command := range commands {
		cmd := getBotCommand(command)
		assert.Equal(t, "series", cmd.command)
		assert.Equal(t, "0", cmd.argument)
	}
}

func TestGetRestaurantCommand(t *testing.T) {
	commands := []string{
		"restaurant:bangsue",
		"RESTAURANT:bangsue",
		"ReStAuRaNt:bangsue",
		"  restaurant  :  bangsue  ",
	}

	for _, command := range commands {
		cmd := getBotCommand(command)
		assert.Equal(t, "restaurant", cmd.command)
		assert.Equal(t, "bangsue", cmd.argument)
	}
}

func TestCallSeriesCommand(t *testing.T) {
	series := []int{3, 5, 9, 15, 23, 33, 45}

	for index, value := range series {
		cmd := &BotCommand{
			command:  "series",
			argument: fmt.Sprint(index),
		}

		result, err := seriesCommand(cmd)
		assert.Nil(t, err)
		assert.Equal(t, fmt.Sprint(value), result[0])
	}
}

func TestCallRestaurantCommand(t *testing.T) {
	cmd := &BotCommand{
		command:  "restaurant",
		argument: "bangsue",
	}

	result, err := restaurantCommand(cmd)
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(result), 0)
}

func TestExecuteSeriesCommand(t *testing.T) {
	series := []int{3, 5, 9, 15, 23, 33, 45}

	for index, value := range series {
		result, err := executeCommand(fmt.Sprintf("series:%v", index))
		assert.Nil(t, err)
		assert.Equal(t, fmt.Sprint(value), result[0])
	}
}

func TestExecuteRestaurantCommand(t *testing.T) {
	result, err := executeCommand("restaurant:bangsue")
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(result), 0)
}

func TestSkipExecuteCommand(t *testing.T) {
	result, err := executeCommand("test")
	assert.Nil(t, err)
	assert.Equal(t, "test", result[0])
}
