package electronic

type androidPhone struct {
	brand string
	model string
	os    string
}

func (d *androidPhone) Brand() string { //brand string
	return d.brand
}

func (d *androidPhone) Model() string { //model string
	return d.model
}

func (d *androidPhone) Type() string {
	return "smartphone"
}

func (d *androidPhone) OS() string { //os string
	return d.os
}

func AndroidCreate(brand, model, os string) *androidPhone {
	return &androidPhone{brand, model, os}
}
