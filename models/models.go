package models

type CardBin struct {
	Name     string
	Length   int
	CardType string
	Prefixes []int
}

type GenData struct {
	Name   string
	Mobile string
	IdNo   string
	Bank   string
	Email  string
	Addr   string
}
