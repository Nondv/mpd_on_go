package mpd

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"strings"
)

func (client *Client) Stats() (map[string]string, error) {
	return client.ExecuteAndParseMap("stats")
}

func (client *Client) CurrentSong() (map[string]string, error) {
	return client.ExecuteAndParseMap("currentsong")
}

func (client *Client) Status() (map[string]string, error) {
	return client.ExecuteAndParseMap("status")
}

func (client *Client) ExecuteAndParseMap(command string) (map[string]string, error) {
	response, err := client.Execute(command)
	if err != nil {
		return nil, err
	}
	dataLines := lines(response)
	if isMpdError(dataLines[len(dataLines)-1]) {
		return nil, errors.New("mpd error")
	}

	return parseAsMap(dataLines[:len(dataLines)-1]), nil
}

func parseAsMap(lines []string) map[string]string {
	result := make(map[string]string)
	for i := 0; i < len(lines); i++ {
		var key, val string
		fmt.Sscanf(lines[i], "%s %s", &key, &val)
		key = key[:len(key)-1] // colon
		result[key] = val
	}

	return result
}

func (client *Client) Execute(command string) (string, error) {
	_, err := client.writeline(command)
	if err != nil {
		return "", err
	}
	var buffer bytes.Buffer
	for {
		line, err := client.readline()
		if err != nil {
			return buffer.String(), err
		}
		buffer.WriteString(line)
		if isEndOfResponse(line) {
			break
		}
	}

	return buffer.String(), nil
}

func isEndOfResponse(line string) bool {
	return line == "OK\n" || isMpdError(line)
}

func isMpdError(line string) bool {
	return strings.HasPrefix(line, "ACK")
}

func lines(str string) []string {
	linesNumber := strings.Count(str, "\n")
	result := make([]string, linesNumber)
	scanner := bufio.NewScanner(strings.NewReader(str))
	for i := 0; i < linesNumber; i++ {
		scanner.Scan()
		result[i] = scanner.Text()
	}
	return result
}
