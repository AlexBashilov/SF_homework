package electronic

type applePhone struct {
	model string
	os    string
}

func (a *applePhone) Brand() string { //brand string
	return "apple"
}

func (a *applePhone) Model() string { //model string
	return a.model
}

func (a *applePhone) Type() string { //obj string
	return "smartphone"
}

func (a *applePhone) OS() string { //os string
	return a.os
}

func AppleCreate(model, os string) *applePhone {
	return &applePhone{model, os}
}
