package compress_test

import (
	"fmt"
	"testing"

	"github.com/akyoto/compress"
)

func TestByteCounts(t *testing.T) {
	fmt.Println("Len:", len(data), "bytes")
	counts := compress.ByteCounts(data)

	for i, count := range counts {
		fmt.Printf("%d: %d\n", i, count)
	}
}
