package main

type Round struct {
	rad  int
	area float32
}

type Cat struct {
	name  string
	color string
}

type BlackCat struct {
	Cat
}

func NewCat(lname string, lcolor string) *Cat {
	return &Cat{
		name:  lname,
		color: lcolor,
	}
}

func NewBlackCat(color string) *BlackCat {
	cat := &BlackCat{}
	cat.color = color
	return cat
}

func (r Round) getArea() float32 {
	return 3.14 * float32(r.rad) * float32(r.rad)
}
