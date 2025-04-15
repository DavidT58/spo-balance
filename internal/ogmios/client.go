// internal/ogmios/client.go
package ogmios

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/davidt58/spo-balance/pkg/models"
	"github.com/gorilla/websocket"
)

type Client struct {
	conn     *websocket.Conn
	endpoint string
}

func NewClient(endpoint string) (*Client, error) {
	conn, _, err := websocket.DefaultDialer.Dial(endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Ogmios: %w", err)
	}

	return &Client{
		conn:     conn,
		endpoint: endpoint,
	}, nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}

func (c *Client) QueryAddressBalance(address string) (*models.Balance, error) {
	// This is where you would implement the actual query to Ogmios
	// You'll need to construct the proper JSON request according to Ogmios API

	_, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Example request structure (you'll need to adjust based on Ogmios documentation)
	request := map[string]interface{}{
		"type":        "jsonwsp/request",
		"version":     "1.0",
		"servicename": "ogmios",
		"methodname":  "queryUtxo",
		"args": map[string]interface{}{
			"addresses": []string{address},
		},
	}

	requestBytes, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	// Send the request
	if err := c.conn.WriteMessage(websocket.TextMessage, requestBytes); err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	// Read the response
	// This is simplified - in reality you'd need proper handling of async responses
	_, message, err := c.conn.ReadMessage()
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	// Parse the response
	var response map[string]interface{}
	if err := json.Unmarshal(message, &response); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	// Process the response to extract balance
	// This is a placeholder - you'll need to extract the actual balance data
	// based on Ogmios response format
	balance := &models.Balance{
		Lovelace: 0, // Extract from response
		Assets:   make(map[string]uint64),
	}

	return balance, nil
}
