package client

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net"

	"github.com/dmitrykondrakhin/word-of-wisdom/internal/hashcash"
	"github.com/dmitrykondrakhin/word-of-wisdom/internal/utils"
)

const protocol = "tcp"

type client struct {
	host          string
	port          string
	repeatedCount int
	hashCashBits  uint
	logger        *slog.Logger
}

func NewClient(host string, port string, repeatedCount int, hashCashBits uint, logger *slog.Logger) *client {
	return &client{
		host:          host,
		port:          port,
		repeatedCount: repeatedCount,
		hashCashBits:  hashCashBits,
		logger:        logger,
	}
}

func (c *client) Run(ctx context.Context) error {
	for i := 0; i < c.repeatedCount; i++ {
		answer, err := c.GetAnswerFromTCPServer(ctx)
		if err != nil {
			c.logger.Error(err.Error())
		} else {
			c.logger.Info(string(answer))
		}
	}

	return nil
}

func (c *client) GetAnswerFromTCPServer(ctx context.Context) ([]byte, error) {
	var dialer net.Dialer

	conn, err := dialer.DialContext(ctx, protocol, fmt.Sprintf("%s:%s", c.host, c.port))
	if err != nil {
		return nil, fmt.Errorf("failed to dial: %w", err)
	}

	defer func() {
		err := conn.Close()
		if err != nil {
			log.Println("close connection error", err.Error())
		}
	}()

	err = utils.Write(conn, []byte(""))
	if err != nil {
		return nil, fmt.Errorf("request err: %w", err)
	}

	token, err := utils.Read(conn)
	if err != nil {
		return nil, fmt.Errorf("receive token err: %w", err)
	}

	solution, err := hashcash.New(c.hashCashBits).GetHeader(string(token))
	if err != nil {
		return nil, fmt.Errorf("calculate hashcash err: %w", err)
	}

	err = utils.Write(conn, []byte(solution))
	if err != nil {
		return nil, fmt.Errorf("send solution err: %w", err)
	}

	answer, err := utils.Read(conn)
	if err != nil {
		return nil, fmt.Errorf("receive answer err: %w", err)
	}

	return answer, nil
}
