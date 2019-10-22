package compress_test

import (
	"fmt"
	"testing"

	"github.com/akyoto/assert"
	"github.com/akyoto/compress"
)

func TestABC(t *testing.T) {
	compressed := compress.ABCCompress(data)
	decompressed := compress.ABCDecompress(compressed)
	assert.DeepEqual(t, data, decompressed)

	fmt.Println("Initial:", len(data), "bytes")
	fmt.Println("Compressed:", len(compressed), "bytes")
}
