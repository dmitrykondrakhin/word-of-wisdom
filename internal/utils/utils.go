package utils

import (
	"encoding/binary"
	"io"
	"net"
)

func Read(conn net.Conn) ([]byte, error) {
	var length uint64
	err := binary.Read(conn, binary.BigEndian, &length)
	if err != nil {
		return []byte{}, err
	}

	msg := make([]byte, length)
	_, err = io.ReadFull(conn, msg)
	if err != nil {
		return []byte{}, err
	}

	return msg, nil
}

func Write(conn net.Conn, msg []byte) error {
	err := binary.Write(conn, binary.BigEndian, uint64(len(msg)))
	if err != nil {
		return err
	}

	_, err = conn.Write(msg)
	if err != nil {
		return err
	}

	return nil
}
