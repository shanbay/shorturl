package shorturl

import (
	"fmt"
	"os"
	"testing"
)

func TestEncodeURL(t *testing.T) {
	file, err := os.Open("test_data.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	encoder := NewURLEncoder(&URLEncoderConfig{alphabet: "mn6j2c4rv8bpygw95z7hsdaetxuk3fq"})
	var number uint64
	var alpha string
	for {
		_, err := fmt.Fscanf(file, "%d:%s", &number, &alpha)
		if err != nil {
			break
		}
		resultNumber := encoder.DecodeURL(alpha)
		resultString := encoder.EncodeURL(number)
		if resultString != alpha {
			t.Errorf("result was incorrect, got: %s, want: %s.", resultString, alpha)
		}
		if resultNumber != number {
			t.Errorf("result was incorrect, got: %d, want: %d.", resultNumber, number)
		}
	}
}
