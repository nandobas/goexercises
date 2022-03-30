package services

type Person struct {
	name, position string
}

func (p *Person) GetName() string{
	return p.name
}

func (p *Person) GetPosition() string{
	return p.position
}


type personMod func(*Person)
type PersonBuilder struct {
	actions []personMod
}

func (b *PersonBuilder) MyName(name string) *PersonBuilder {
	b.actions = append(b.actions, func(p *Person) {
		p.name = name
	})
	return b
}

func (b *PersonBuilder) Build() *Person {
	p := Person{}
	for _, a := range b.actions {
		a(&p)
	}
	return &p
}

func (b *PersonBuilder) MyWork(position string) *PersonBuilder {
	b.actions = append(b.actions, func(p *Person) {
		p.position = position
	})
	return b
}