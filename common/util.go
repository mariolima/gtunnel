package common

import (
	"encoding/base64"
	"encoding/binary"
	"math/rand"
	"net"
	"time"
	"unicode/utf16"
)

// Credit to https://gist.github.com/ammario/ipint.go
func IpToInt32(ip net.IP) uint32 {
	return binary.BigEndian.Uint32(ip)
}

func Int32ToIP(i uint32) net.IP {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, i)
	return ip
}

const TunnelIDSize = 8
const ClientIDSize = 8

// Credit to: https://www.calhoun.io/creating-random-strings-in-go/
func GenerateString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func B64Utf16le(s string) string {
	encoded := utf16.Encode([]rune(s))
	b := convertUTF16ToLittleEndianBytes(encoded)
	return base64.StdEncoding.EncodeToString([]byte(b))
}

func convertUTF16ToLittleEndianBytes(u []uint16) []byte {
	b := make([]byte, 2*len(u))
	for index, value := range u {
		binary.LittleEndian.PutUint16(b[index*2:], value)
	}
	return b
}

type TaskTCPScanOpts struct {
	Ipnet     string `json:"ipnet"`
	Port      string `json:"port"`
	MaxThread int    `json:"maxthread"`
}
