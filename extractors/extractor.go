package extractors

type IExtractor interface {
	Extract()
}

type Extractor struct {
	Name string
}
