package mpd

import (
	"errors"
	"fmt"
	"strconv"
)

/*
 * TODO:
 *   * crossfade <SECONDS>
 *   * mixrampdb <deciBels>
 *   * mixrampdelay <SECONDS>
 *   * replay_gain_mode <MODE>
 *   * replay_gain_status
 */

func (client *Client) SetConsume(state bool) error {
	return client.executeStateSetCommand("consume", state)
}

func (client *Client) SetSingle(state bool) error {
	return client.executeStateSetCommand("single", state)
}

func (client *Client) SetRepeat(state bool) error {
	return client.executeStateSetCommand("repeat", state)
}

func (client *Client) SetRandom(state bool) error {
	return client.executeStateSetCommand("random", state)
}

func (client *Client) executeStateSetCommand(command string, state bool) error {
	if state == true {
		return client.ExecuteAndCheckMpdError(command + " 1")
	} else {
		return client.ExecuteAndCheckMpdError(command + " 0")
	}
}

func (client *Client) SetVolume(n int) error {
	if n < 0 || n > 100 {
		return errors.New("wrong value")
	}
	return client.ExecuteAndCheckMpdError(fmt.Sprintf("setvol %d", n))
}

func (client *Client) IncreaseVolume(n int) error {
	if n < 1 || n > 100 {
		return errors.New("Argument should be 1..100")
	}
	currentVolume, err := client.currentVolume()
	if err != nil {
		return err
	}
	value := currentVolume + n
	if value > 100 {
		value = 100
	}
	return client.SetVolume(value)
}

func (client *Client) DecreaseVolume(n int) error {
	if n < 1 || n > 100 {
		return errors.New("Argument should be 1..100")
	}
	currentVolume, err := client.currentVolume()
	if err != nil {
		return err
	}
	value := currentVolume - n
	if value < 0 {
		value = 0
	}
	return client.SetVolume(value)
}

func (client *Client) currentVolume() (int, error) {
	status, err := client.Status()
	if err != nil {
		return 0, err
	}

	result, err := strconv.ParseInt(status["volume"], 10, 0)
	if err != nil {
		return 0, err
	}
	return int(result), nil
}
