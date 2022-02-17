package main

type Persons []*Person

func (p Persons) Len() int {
	return len(p)
}

func (p Persons) Less(i, j int) bool {
	return p[i].fatRate > p[j].fatRate
}

func (p Persons) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
