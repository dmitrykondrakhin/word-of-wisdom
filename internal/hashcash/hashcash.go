package hashcash

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	dateFormat      = "060102"
	randomStringLen = 8
)

var hasher = sha1.New()

type Hash struct {
	bits         uint
	date         string
	randomString string
}

func New(bits uint) *Hash {
	return &Hash{
		bits: bits,
		date: time.Now().Format(dateFormat),
		randomString: func() string {
			randomString, err := getRandomString()
			if err != nil {
				return ""
			}

			return randomString
		}(),
	}
}

func (h *Hash) GetHeader(token string) (string, error) {
	counter := 0
	var header string
	for {
		header = fmt.Sprintf("1:%d:%s:%s::%s:%x", h.bits, h.date, token, h.randomString, counter)
		if h.checkZeros(header) {
			return header, nil
		}
		counter++
	}
}

func (h *Hash) Check(header string) bool {
	if h.checkDate(header) {
		return h.checkZeros(header)
	}

	return false
}

func getRandomString() (string, error) {
	buf := make([]byte, randomStringLen)
	_, err := rand.Read(buf)
	if err != nil {
		return "", err
	}

	randomString := base64.StdEncoding.EncodeToString(buf)

	return randomString[:randomStringLen], nil
}

func (h *Hash) checkZeros(header string) bool {
	hasher.Reset()
	hasher.Write([]byte(header))
	sum := hasher.Sum(nil)
	sumUint64 := binary.BigEndian.Uint64(sum)
	sumBits := strconv.FormatUint(sumUint64, 2)
	zeroes := 64 - len(sumBits)

	return uint(zeroes) >= h.bits
}

func (h *Hash) checkDate(header string) bool {
	fields := strings.Split(header, ":")
	if len(fields) != 7 {
		return false
	}
	date, err := time.Parse(dateFormat, fields[2])
	if err != nil {
		return false
	}
	duration := time.Since(date)
	return duration.Hours() <= 48
}
