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
	"mine": {
		Name: "Mine",
		BaseCost: ResourceSet{
			Stone: 5,
		},
		Multiplier: 1.15,
		Effect: func(m *model) {
			m.res.Gold++
		},
	},
	"field": {
		Name: "Field",
		BaseCost: ResourceSet{
			Gold: 1,
		},
		Multiplier: 1.15,
		Effect: func(m *model) {
			m.res.Wood++
		},
	},
	"forest": {
		Name: "Forest",
		BaseCost: ResourceSet{
			Wood: 3,
			Gold: 1,
		},
		Multiplier: 1.15,
		Effect: func(m *model) {
			m.res.Wood++
			m.res.Wood++
		},
	},
}
