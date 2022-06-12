package test2

type A struct {
	_struct struct{} `codec:",omitempty"`
	Dummy1  int      `codec:"dummy1"`
}
