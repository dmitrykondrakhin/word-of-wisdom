package server

import (
	"context"
	"fmt"
	"log/slog"
	"math/rand"
	"net"

	"github.com/dmitrykondrakhin/word-of-wisdom/internal/hashcash"
	"github.com/dmitrykondrakhin/word-of-wisdom/internal/utils"
)

type server struct {
	host         string
	port         string
	hashCashBits uint
	logger       *slog.Logger
}

const protocol = "tcp"
const randomStringLength = 8

func NewServer(host string, port string, hashCashBits uint, logger *slog.Logger) *server {
	return &server{
		host:         host,
		port:         port,
		hashCashBits: hashCashBits,
		logger:       logger,
	}
}

func (s *server) Start(ctx context.Context) error {
	listen, err := net.Listen(protocol, s.host+":"+s.port)
	if err != nil {
		return err
	}

	defer func() {
		err := listen.Close()
		if err != nil {
			s.logger.Error(fmt.Sprintf("close connection error. %s", err.Error()))
		}
	}()

	for {
		conn, err := listen.Accept()
		if err != nil {
			return err
		}
		err = s.handleRequest(conn)
		if err != nil {
			s.logger.Error(err.Error())
		}
	}
}

func (s *server) handleRequest(conn net.Conn) error {
	_, err := utils.Read(conn)
	if err != nil {
		return fmt.Errorf("read message err: %w", err)
	}

	token := RandStringBytes(randomStringLength)
	err = utils.Write(conn, token)
	if err != nil {
		return fmt.Errorf("send token err: %w", err)
	}
	s.logger.Info("send token " + string(token))

	solution, err := utils.Read(conn)
	if err != nil {
		return fmt.Errorf("receive solution err: %w", err)
	}
	s.logger.Info("get hashcash header " + string(solution))

	if !hashcash.New(s.hashCashBits).Check(string(solution)) {
		return fmt.Errorf("check solution error: %w", err)
	}

	if err = utils.Write(conn, []byte("yohooo")); err != nil {
		return fmt.Errorf("send answer err: %w", err)
	}

	conn.Close()

	return nil
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) []byte {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	return b
}
