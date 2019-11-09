package line

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

var (
	apiEndpoint      = "https://scg-go-odjtm7kfta-uc.a.run.app"
	seriesRegexp     = regexp.MustCompile(`(?i)^\s*(series)\s*:\s*([0-9]+)\s*$`)
	restaurantRegexp = regexp.MustCompile(`(?i)^\s*(restaurant)\s*:\s*(.+)\s*$`)
)

// BotCommand struct of command line in chat bot
type BotCommand struct {
	command  string
	argument string
}

func getBotCommand(message string) *BotCommand {
	var result [][]string

	// parsing message if it is series command
	if seriesRegexp.MatchString(message) {
		result = seriesRegexp.FindAllStringSubmatch(message, -1)
	}

	// parsing message if it is restaurant command
	if restaurantRegexp.MatchString(message) {
		result = restaurantRegexp.FindAllStringSubmatch(message, -1)
	}

	// exit if not found command
	if len(result) == 0 {
		return nil
	}

	// return bot command
	return &BotCommand{
		command:  strings.Trim(strings.ToLower(result[0][1]), " "),
		argument: strings.Trim(result[0][2], " "),
	}
}

func executeCommand(message string) ([]string, error) {
	// if found bot command then jump to command executor function
	if botCommand := getBotCommand(message); botCommand != nil {
		switch botCommand.command {
		case "restaurant":
			return restaurantCommand(botCommand)
		case "series":
			return seriesCommand(botCommand)
		}
	}

	return []string{message}, nil
}

func seriesCommand(cmd *BotCommand) ([]string, error) {
	// call series API
	uri := fmt.Sprintf("%v/series/%v", apiEndpoint, url.PathEscape(cmd.argument))
	body, err := callCommandAPI(uri)
	if err != nil {
		return nil, err
	}

	// parsing JSON response
	var result map[string]int
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return []string{fmt.Sprint(result["value"])}, nil
}

func restaurantCommand(cmd *BotCommand) ([]string, error) {
	// call restaurant API
	uri := fmt.Sprintf("%v/restaurant/%v", apiEndpoint, url.PathEscape(cmd.argument))
	body, err := callCommandAPI(uri)
	if err != nil {
		return nil, err
	}

	// parsing JSON response
	var result map[string][]struct {
		Name    string `json:"name"`
		Address string `json:"formatted_address,omitempty"`
	}
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	// transform response to be message
	var messages []string
	for index, restaurant := range result["results"] {
		if index >= 10 {
			break
		}

		line := fmt.Sprintf("Name: %v\nAddress: %v", restaurant.Name, restaurant.Address)
		messages = append(messages, line)
	}
	return []string{strings.Join(messages, "\n\n")}, err
}

func callCommandAPI(uri string) ([]byte, error) {
	resp, err := http.Get(uri)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
