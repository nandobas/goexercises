package services

type Car struct{
	color,
	chassis,
	factory,
	model string
	year int
}

type carMod func(c *Car)
type CarBuilder struct {
	actions []carMod
}

func (cb *CarBuilder) Color(color string)*CarBuilder{
	cb.actions = append(cb.actions, func(c *Car) {
		c.color = color
	})
	return cb
}

func (cb *CarBuilder) Chassis(chassis string)*CarBuilder{
	cb.actions = append(cb.actions, func(c *Car) {
		c.chassis = chassis
	})
	return cb
}

func (cb *CarBuilder) FactoryAndModel(factory, model string)*CarBuilder{
	cb.actions = append(cb.actions, func(c *Car) {
		c.factory = factory
		c.model = model
	})
	return cb
}

func (cb *CarBuilder) Year(year int)*CarBuilder{
	cb.actions = append(cb.actions, func(c *Car) {
		c.year = year
	})
	return cb
}

func (cb *CarBuilder) Build() *Car{
	c := Car{}
	for _, act := range cb.actions {
		act(&c)
	}
	return &c
}