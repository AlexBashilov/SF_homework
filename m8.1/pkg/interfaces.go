package electronic

type Phone interface {
	Brand() string
	Model() string
	Type() string
}

type Smartphone interface {
	OS() string
	Phone
}

type StationPhone interface {
	ButtonsCount() int
	Phone
}
