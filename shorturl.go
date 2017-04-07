package shorturl

import (
	"fmt"
	"strings"
)

const (
	//VERSION is SDK version
	VERSION = "0.1"

	default_alphabet   = "asdfghjklURLEncoderConfigqwertyui"
	default_block_size = uint(24)
	min_length         = 5
	one                = uint64(1)
)

type urlEncoder struct {
	alphabet   string
	block_size uint
}

type URLEncoderConfig struct {
	alphabet   string
	block_size uint
}

func NewURLEncoder(config *URLEncoderConfig) *urlEncoder {
	alphabet := default_alphabet
	block_size := default_block_size
	if config.alphabet != "" {
		alphabet = config.alphabet
	}
	if config.block_size != 0 {
		block_size = config.block_size
	}
	url_encoder := &urlEncoder{
		alphabet:   alphabet,
		block_size: block_size,
	}
	return url_encoder
}

func getBit(n uint64, pos uint) int {
	if (n & (one << pos)) != 0 {
		return 1
	}
	return 0
}

func (encoder *urlEncoder) encode(n uint64) uint64 {
	for i, j := uint(0), uint(encoder.block_size-1); i < j; i, j = i+1, j-1 {
		if getBit(n, i) != getBit(n, j) {
			n ^= ((one << i) | (one << j))
		}
	}
	return n
}

func (encoder *urlEncoder) enbase(x uint64) string {
	n := uint64(len(encoder.alphabet))
	result := []byte{}
	for {
		ch := encoder.alphabet[x%n]
		result = append(result, ch)
		x = x / n
		if x == 0 && len(result) >= min_length {
			break
		}
	}
	revResult := []byte{}
	for i := len(result) - 1; i >= 0; i-- {
		revResult = append(revResult, result[i])
	}
	return string(revResult)
}

func (encoder *urlEncoder) debase(x string) uint64 {
	n := uint64(len(encoder.alphabet))
	result := uint64(0)
	bits := []byte(x)
	for _, bitValue := range bits {
		result = result*n + uint64(strings.IndexByte(encoder.alphabet, bitValue))
	}
	return result
}

func (encoder *urlEncoder) EncodeURL(n uint64) string {
	return encoder.enbase(encoder.encode(n))
}

func (encoder *urlEncoder) DecodeURL(n string) uint64 {
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
