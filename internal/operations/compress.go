package operations

type Compressor interface {
	Compress(hypertableName string)
	Uncompress(hypertableName string)
}
