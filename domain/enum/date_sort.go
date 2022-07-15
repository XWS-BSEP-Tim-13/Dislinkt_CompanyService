package enum

type DateSort int

const (
	LATEST DateSort = iota
	OLDEST
	NONE_DATE
)
