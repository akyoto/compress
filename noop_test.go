package compress_test

import (
	"testing"

	"github.com/akyoto/assert"
	"github.com/akyoto/compress"
)

var data = []byte("Hello World")

func TestNoop(t *testing.T) {
	compressed := compress.NoopCompress(data)
	decompressed := compress.NoopDecompress(compressed)
	assert.DeepEqual(t, data, decompressed)
}
