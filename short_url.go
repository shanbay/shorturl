package main

import (
	"bytes"
	"fmt"
)

const (
	//VERSION is SDK version
	VERSION = "0.1"

	DEFAULT_ALPHABET   = "asdfghjklzxcvbnmqwertyui"
	DEFAULT_BLOCK_SIZE = uint(24)
	MIN_LENGTH         = 5
	ONE                = uint64(1)
)

type URLEncoder struct {
	alphabet   []byte
	block_size uint
}

type URLEncoderConfig struct {
	alphabet   string
	block_size uint
}

func NewURLEncoder(config *URLEncoderConfig) *URLEncoder {
	alphabet := []byte(DEFAULT_ALPHABET)
	block_size := DEFAULT_BLOCK_SIZE
	if config.alphabet != "" {
		alphabet = []byte(config.alphabet)
	}
	if config.block_size != 0 {
		block_size = config.block_size
	}
	url_encoder := &URLEncoder{
		alphabet:   alphabet,
		block_size: block_size,
	}
	return url_encoder
}

func getBit(n uint64, pos uint) int {
	if (n & (ONE << pos)) != 0 {
		return 1
	}
	return 0
}

func (encoder *URLEncoder) encode(n uint64) uint64 {
	var i uint = 0
	var j uint = encoder.block_size - 1
	for {
		if i >= j {
			break
		}
		if getBit(n, i) != getBit(n, j) {
			n ^= ((ONE << i) | (ONE << j))
		}
		i++
		j--
	}
	return n
}

func (encoder *URLEncoder) enbase(x uint64) string {
	n := uint64(len(encoder.alphabet))
	result := []byte{}
	for {
		ch := encoder.alphabet[x%n]
		result = append(result, ch)
		x = x / n
		if x == 0 && len(result) >= MIN_LENGTH {
			break
		}
	}
	revResult := []byte{}
	for i := len(result) - 1; i >= 0; i-- {
		revResult = append(revResult, result[i])
	}
	return string(revResult)
}

func (encoder *URLEncoder) debase(x string) uint64 {
	n := uint64(len(encoder.alphabet))
	result := uint64(0)
	bits := []byte(x)
	for _, bitValue := range bits {
		result = result*n + uint64(bytes.IndexByte(encoder.alphabet, bitValue))
	}
	return result
}

func (encoder *URLEncoder) EncodeURL(n uint64) string {
	return encoder.enbase(encoder.encode(n))
}

func (encoder *URLEncoder) DecodeURL(n string) uint64 {
	return encoder.encode(encoder.debase(n))
}

func main() {
	encoder := NewURLEncoder(&URLEncoderConfig{})
	for {
		var opt int
		var x string
		var n uint64
		fmt.Println("input 0: encode, 1: decode")
		fmt.Scanf("%d", &opt)
		if opt == 0 {
			fmt.Scanf("%d", &n)
			fmt.Println(encoder.EncodeURL(n))
		} else {
			fmt.Scanf("%s", &x)
			fmt.Println(encoder.DecodeURL(x))
		}
	}
}
