package main

type Person struct {
	name    string
	fatRate float64
}

/**
注册
*/
func (p *Person) register(name string, fatRate float64) {
	p.name = name
	p.fatRate = fatRate
}
