package mpd

/*
 * TODO:
 *   * seek <POS> <TIME>
 *   * seekid <ID> <TIME>
 *   * seekcur <TIME>
 */

import "fmt"

func (client *Client) Stop(id int) error {
	return client.ExecuteAndCheckMpdError("stop")
}

func (client *Client) PlayId(id int) error {
	return client.ExecuteAndCheckMpdError(fmt.Sprintf("playid %d", id))
}

func (client *Client) PlayPosition(pos int) error {
	return client.ExecuteAndCheckMpdError(fmt.Sprintf("play %d", pos))
}

func (client *Client) Play() error {
	return client.ExecuteAndCheckMpdError("pause 0")
}

func (client *Client) Pause() error {
	return client.ExecuteAndCheckMpdError("pause 1")
}

func (client *Client) Previous() error {
	return client.ExecuteAndCheckMpdError("previous")
}

func (client *Client) Next() error {
	return client.ExecuteAndCheckMpdError("next")
}
