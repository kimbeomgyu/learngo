package main

import (
	"fmt"

	"github.com/tuckersGo/goWeb/web9/cipher"
	"github.com/tuckersGo/goWeb/web9/lzw"
)

// Component interface
type Component interface {
	Operator(string)
}

var sentData string
var receiveData string

// SendComponent struct
type SendComponent struct{}

// Operator method
func (s *SendComponent) Operator(data string) {
	// Send data
	sentData = data
}

// ReadComponent struct
type ReadComponent struct{}

// Operator method
func (s *ReadComponent) Operator(data string) {
	receiveData = data
}

// ZipComponent struct
type ZipComponent struct {
	com Component
}

// Operator method
func (s *ZipComponent) Operator(data string) {
	zipData, err := lzw.Write([]byte(data))
	if err != nil {
		panic(err)
	}
	s.com.Operator(string(zipData))
}

// UnzipComponent struct
type UnzipComponent struct {
	com Component
}

// Operator method
func (s *UnzipComponent) Operator(data string) {
	unzipData, err := lzw.Read([]byte(data))
	if err != nil {
		panic(err)
	}
	s.com.Operator(string(unzipData))
}

// EncryptComponent struct
type EncryptComponent struct {
	key string
	com Component
}

// Operator is method
func (s *EncryptComponent) Operator(data string) {
	encryptData, err := cipher.Encrypt([]byte(data), s.key)
	if err != nil {
		panic(err)
	}

	s.com.Operator(string(encryptData))
}

// DecryptComponent struct
type DecryptComponent struct {
	key string
	com Component
}

// Operator is method
func (s *DecryptComponent) Operator(data string) {
	decryptData, err := cipher.Decrypt([]byte(data), s.key)
	if err != nil {
		panic(err)
	}

	s.com.Operator(string(decryptData))
}

func main() {
	sender := &EncryptComponent{
		key: "abcde",
		com: &ZipComponent{
			com: &SendComponent{},
		},
	}
	sender.Operator("Hello World")

	fmt.Println(sentData)

	receiver := &UnzipComponent{
		com: &DecryptComponent{
			key: "abcde",
			com: &ReadComponent{},
		},
	}
	receiver.Operator(sentData)

	fmt.Println(receiveData)
}
