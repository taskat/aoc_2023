package types

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64 | uintptr
}

type String interface {
	~string
}

type Summable interface {
	Number | String
}

type Pair[FIRST, SECOND any] struct {
	First  FIRST
	Second SECOND
}
