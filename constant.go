package main

const maxVisibleLogLines = 10

type ResourceSet struct {
	Wood  int
	Stone int
	Gold  int
}

type BuildingSet struct {
	Houses int
	Mines  int
	Farms  int
}

type model struct {
	input        string
	log          []string
	res          ResourceSet
	build        BuildingSet
	pop          int
	popMax       int
	scrollOffset int
	lang         string
}
