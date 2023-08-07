package main

type SM struct { ///С А Н Т И М Е Т Р Ы
	length int
	width  int
	height int
}

func (s *SM) Length() int { ///С А Н Т И М Е Т Р Ы
	return s.length
}

func (s *SM) Width() int { ///С А Н Т И М Е Т Р Ы
	return s.width
}

func (s *SM) Height() int { ///С А Н Т И М Е Т Р Ы
	return s.height
}

func CreateSM(length, width, height int) *SM { ///С А Н Т И М Е Т Р Ы
	return &SM{length, width, height}
}
