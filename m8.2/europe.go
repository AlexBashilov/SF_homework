package main

type EuropeCars struct {
	brand       string
	model       string
	dim         Dimensions
	maxSpeed    int
	enginePower int
}

func (e *EuropeCars) Brand() string {
	return e.brand
}

func (e *EuropeCars) Model() string {
	return e.model
}

func (e *EuropeCars) Dimensions() Dimensions {
	return e.dim
}

func (e *EuropeCars) MaxSpeed() int {
	return e.maxSpeed
}

func (e *EuropeCars) EnginePower() int {
	return e.enginePower
}

func EuCarCreate(brand, model string, maxSpeed, enginePower int, dim Dimensions) *AmCars {
	return &AmCars{brand, model, dim, maxSpeed, enginePower}
}
