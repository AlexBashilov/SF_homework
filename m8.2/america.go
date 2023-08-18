package main

type AmCars struct {
	brand       string
	model       string
	dim         Dimensions
	maxSpeed    int
	enginePower int
}

func (a *AmCars) Brand() string {
	return a.brand
}

func (a *AmCars) Model() string {
	return a.model
}

func (a *AmCars) Dimensions() Dimensions {
	return a.dim
}

func (a *AmCars) MaxSpeed() int {
	return a.maxSpeed
}

func (a *AmCars) EnginePower() int {
	return a.enginePower
}

func AmCarCreate(brand, model string, maxSpeed, enginePower int, dim Dimensions) *AmCars {
	return &AmCars{brand, model, dim, maxSpeed, enginePower}
}
