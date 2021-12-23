package mud

type Object interface {
}

type BaseObject struct {
	Name string
}

func newBaseObject() Object {
	return &BaseObject{
		Name: "Test Name",
	}
}
