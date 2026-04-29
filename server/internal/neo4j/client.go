package neo4j

import (
	"context"
	"fmt"
	"log"

	"github.com/neo4j/neo4j-go-driver/v6/neo4j"
)

type Client struct {
	driver neo4j.Driver
}

func New(uri, user, pass string) (*Client, error) {
	driver, err := neo4j.NewDriver(uri, neo4j.BasicAuth(user, pass, ""))
	if err != nil {
		return nil, fmt.Errorf("create neo4j driver: %w", err)
	}

	err = driver.VerifyConnectivity(context.Background())
	if err != nil {
		driver.Close(context.Background())
		return nil, fmt.Errorf("verify neo4j connection: %w", err)
	}

	log.Printf("neo4j connected: %s", uri)
	return &Client{driver: driver}, nil
}

func (c *Client) Check(ctx context.Context) error {
	_, err := neo4j.ExecuteQuery[*neo4j.EagerResult](
		ctx,
		c.driver,
		"RETURN 1 AS ping",
		nil,
		neo4j.EagerResultTransformer,
	)
	if err != nil {
		return fmt.Errorf("neo4j check: %w", err)
	}
	return nil
}

func (c *Client) Close(ctx context.Context) error {
	return c.driver.Close(ctx)
}
