package compress_test

import (
	"fmt"
	"testing"

	"github.com/akyoto/assert"
	"github.com/akyoto/compress"
)

func TestNoop(t *testing.T) {
	compressed := compress.NoopCompress(data)
	decompressed := compress.NoopDecompress(compressed)
	assert.DeepEqual(t, data, decompressed)

	fmt.Println("Initial:", len(data), "bytes")
	fmt.Println("Compressed:", len(compressed), "bytes")
}
