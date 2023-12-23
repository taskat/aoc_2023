package types

type Number interface {
	Real | Complex
}

type Real interface {
	Integer | Float
}

type Complex interface {
	~complex64 | ~complex128
}

type Integer interface {
	SignedInteger | UnsignedInteger
}

type SignedInteger interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type UnsignedInteger interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | uintptr
}

type Float interface {
	~float32 | ~float64
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
