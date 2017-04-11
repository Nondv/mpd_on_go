package mpd

import (
	"bufio"
	"fmt"
	"net"
)

type Client struct {
	Host       string
	Port       int
	Version    string
	connection net.Conn
	reader     *bufio.Reader
}

func Connect(host string, port int) (Client, error) {
	client := Client{Host: host, Port: port}
	err := client.Connect()
	if err != nil {
		return client, err
	}
	return client, nil
}

func (client *Client) Connect() error {
	client.reset()
	var err error
	client.connection, err = net.Dial("tcp", client.address())
	if err != nil {
		return err
	}

	response, err := client.readline()

	if err != nil {
		client.reset()
		return err
	}

	fmt.Sscanf(response, "OK MPD %s\n", &client.Version)
	return nil
}

func (client *Client) readline() (string, error) {
	if client.reader == nil {
		client.reader = bufio.NewReader(client.connection)
	}
	return client.reader.ReadString('\n')
}

func (s *Client) writeline(data string) (int, error) {
	return fmt.Fprint(s.connection, data+"\n")
}

func (client *Client) reset() {
	if client.connection != nil {
		client.connection.Close()
		client.connection = nil
	}

	if client.reader != nil {
		client.reader = nil
	}
}

func (client *Client) address() string {
	return fmt.Sprintf("%s:%d", client.Host, client.Port)
}
