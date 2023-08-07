package main

type Duim struct { ///Д Ю Й М Ы
	length int
	width  int
	height int
}

func (d *Duim) Length() int { ///Д Ю Й М Ы
	return d.length
}

func (d *Duim) Width() int { ///Д Ю Й М Ы
	return d.width
}

func (d *Duim) Height() int { ///Д Ю Й М Ы
	return d.height
}

func CreateDuim(length, width, height int) *Duim { ///Д Ю Й М Ы
	return &Duim{length, width, height}
}
