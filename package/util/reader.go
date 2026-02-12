package util

import (
	"crypto/rand"
	"encoding/binary"
	"io"
)

// ReadData reads data chunk from conn
func ReadData(conn io.Reader) ([]byte, error) {
	messageLenBuffer := make([]byte, 4)
	var err error
	_, err = io.ReadFull(conn, messageLenBuffer)
	if err != nil {
		return nil, err
	}

	messageLen := int(binary.LittleEndian.Uint32(messageLenBuffer))
	message := make([]byte, messageLen)
	_, err = io.ReadFull(conn, message)
	if err != nil {
		return nil, err
	}
	return message, nil
}

// WriteData writes data chunk to conn
func WriteData(conn io.Writer, data []byte) error {
	messageLenBuffer := make([]byte, 4)
	binary.LittleEndian.PutUint32(messageLenBuffer, uint32(len(data)))
	var (
		err error
		n   int
	)
	for totalWrote := 0; totalWrote < 4; totalWrote += n {
		n, err = conn.Write(messageLenBuffer[totalWrote:])
		if err != nil {
			return err
		}
	}

	for totalWrote := 0; totalWrote < len(data); totalWrote += n {
		n, err = conn.Write(data[totalWrote:])
		if err != nil {
			return err
		}
	}
	return nil
}

// RandomBytes fills dst with bytes from rand.Reader
func RandomBytes(dst []byte) error {
	_, err := rand.Reader.Read(dst)
	if err != nil {
		return err
	}
	return nil
}
