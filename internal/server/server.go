package server

import (
	"context"
	"fmt"
	"log/slog"
	"net"

	"github.com/dmitrykondrakhin/word-of-wisdom/internal/hashcash"
	"github.com/dmitrykondrakhin/word-of-wisdom/internal/usecase"
	"github.com/dmitrykondrakhin/word-of-wisdom/internal/utils"
)

type WordOfWisdomUseCase interface {
	GetQuote(ctx context.Context) (string, error)
}

type server struct {
	host         string
	port         string
	hashCashBits uint
	logger       *slog.Logger
	usecase      WordOfWisdomUseCase
}

const protocol = "tcp"
const randomStringLength = 8

func NewServer(host string, port string, hashCashBits uint, logger *slog.Logger, usecases *usecase.Usecases) *server {
	return &server{
		host:         host,
		port:         port,
		hashCashBits: hashCashBits,
		logger:       logger,
		usecase:      *usecases.WordOfWidsomUseCase,
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
			s.logger.Error(fmt.Sprintf("close listen error. %s", err.Error()))
		}
	}()

	go func() {
		<-ctx.Done()
		listen.Close()
		s.logger.Info("server was stopped")
	}()

	for {
		conn, err := listen.Accept()
		if err != nil {
			return err
		}

		err = s.handleRequest(ctx, conn)
		if err != nil {
			s.logger.Error(err.Error())
		}
	}
}

func (s *server) handleRequest(ctx context.Context, conn net.Conn) error {
	defer func() {
		err := conn.Close()
		if err != nil {
			s.logger.Error(fmt.Sprintf("close connection error. %s", err.Error()))
		}
	}()

	_, err := utils.Read(conn)
	if err != nil {
		return fmt.Errorf("read message err: %w", err)
	}

	token := utils.RandStringBytes(randomStringLength)

	err = utils.Write(conn, token)
	if err != nil {
		return fmt.Errorf("send token err: %w", err)
	}

	s.logger.Info("send token", "token", string(token))

	solution, err := utils.Read(conn)
	if err != nil {
		return fmt.Errorf("receive solution err: %w", err)
	}

	s.logger.Info("get hashcash header", "header", string(solution))

	if !hashcash.New(s.hashCashBits).Check(string(solution)) {
		return fmt.Errorf("wrong solution")
	}

	answer, err := s.usecase.GetQuote(ctx)
	if err != nil {
		return fmt.Errorf("get quote err: %w", err)
	}

	err = utils.Write(conn, []byte(answer))
	if err != nil {
		return fmt.Errorf("send answer err: %w", err)
	}

	return nil
}
