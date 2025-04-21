package errors


type Error interface {
	error
	Code() Code
}

type Code struct {
	code uint16
}

type CodeI interface {
	CodeDesc() CodeDesc
}

type GroupI interface {

}

type CodeDesc struct {

}
