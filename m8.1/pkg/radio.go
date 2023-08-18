package electronic

type radioPhone struct {
	brand        string
	model        string
	buttonscount int
}

func (r *radioPhone) Brand() string { //brand string
	return r.brand
}

func (r *radioPhone) Model() string { //model string
	return r.model
}

func (r *radioPhone) Type() string {
	return "station"
}

func (r *radioPhone) ButtonsCount() int { //buttonscount int
	return r.buttonscount
}

func RadioCreate(brand, model string, buttonscount int) *radioPhone {
	return &radioPhone{brand, model, buttonscount}
}
