package compress_test

import (
	"testing"

	"github.com/akyoto/assert"
	"github.com/akyoto/compress"
)

func TestNoop(t *testing.T) {
	data := []byte("Hello World")
	compressed := compress.NoopCompress(data)
	decompressed := compress.NoopDecompress(compressed)
	assert.DeepEqual(t, data, decompressed)
}
