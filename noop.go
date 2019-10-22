package compress

func NoopCompress(data []byte) []byte {
	return data
}

func NoopDecompress(data []byte) []byte {
	return data
}
