package mud

type Amount int

type Container struct {
	Contents map[string]Contents
}

type Contents struct {
	Object Object
	Amount Amount
}

func NewContainer() Object {
	return &Container{
		Contents: make(map[string]Contents),
	}
}
