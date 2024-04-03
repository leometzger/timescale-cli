package domain

type OptionFlag uint8

const (
	OptionFlagNotDefined OptionFlag = iota
	OptionFlagTrue
	OptionFlagFalse
)
