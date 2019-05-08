package metadata

type CardBin struct {
	Name     string
	Length   int
	CardType string
	Prefixes []int
}

type Response struct {
	Name   string
	Mobile string
	IdNo   string
	Bank   string
	Email  string
	Addr   string
}
