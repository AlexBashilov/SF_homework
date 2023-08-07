package main

type Dimensions interface {
	Length() int
	Width() int
	Height() int
}

type Auto interface {
	Brand() string
	Model() string
	Dimensions() Dimensions
	MaxSpeed() int
	EnginePower() int
}
