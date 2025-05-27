package main

type BuildingMeta struct {
	Name       string
	BaseCost   ResourceSet
	Multiplier float64
	Effect     func(m *model)
}

var buildingDefs = map[string]BuildingMeta{
	"house": {
		Name: "Haus",
		BaseCost: ResourceSet{
			Wood:  5,
			Stone: 2,
		},
		Multiplier: 1.15,
		Effect: func(m *model) {
			m.popMax++
			m.build.Houses++
		},
	},
}
