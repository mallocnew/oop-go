package abstract_factory

type Car struct {
	name  string
	price int
}

func (c *Car) getName() string {
	return c.name
}

func (c *Car) getPrice() int {
	return c.price
}

type Bus struct {
	name  string
	price int
}

func (b *Bus) getName() string {
	return b.name
}

func (b *Bus) getPrice() int {
	return b.price
}
