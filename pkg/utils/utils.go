package utils

import "math/rand"

var (
	colors = []string{
		"870058",
		"A4303F",
		"F2D0A4",
		"FFECCC",
		"C8D6AF",
		"C9E4CA",
		"87BBA2",
		"55828B",
		"3B6064",
		"364958",
		"DDE0BD",
		"D0D1AC",
		"B6A39E",
		"948B89",
		"726E75",
		"E3D8F1",
		"DABECA",
		"BF8B85",
		"AD8A64",
		"5D5F71",
		"8DA1B9",
		"95ADB6",
		"CBB3BF",
		"DBC7BE",
		"EF959C",
		"3C91E6",
		"9FD356",
		"FA824C",
		"F8F32B",
		"012622",
		"0B4F6C",
		"ECE5F0",
		"E98A15",
		"59114D",
		"F2E94E",
		"DA3E52",
		"96E6B3",
		"7E6B8F",
		"A3D9FF",
		"F9D4BB",
		"F0E2A3",
		"C1D37F",
		"E2D58B",
		"5DD9C1",
		"ACFCD9",
		"B084CC",
		"665687",
		"190933",
		"2292A4",
		"C08552",
		"DAB49D",
		"58BC82",
		"8FE388",
		"4E0250",
	}
)

type MyColors struct {
	list []string
}

func NewMyColors() *MyColors {
	return &MyColors{
		list: colors,
	}
}

func (c *MyColors) GetRandomColorUniq() string {
	if len(c.list) == 0 {
		return "C0C0C0"
	}
	i := rand.Intn(len(c.list))
	color := c.list[i]
	c.list = append(c.list[:i], c.list[i+1:]...)
	return color
}
