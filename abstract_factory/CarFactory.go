package abstract_factory

type Benz struct {
}

func (b *Benz) makeCar() ICar {
	return &Car{
		name:  "benz-car",
		price: 500000,
	}
}

func (b *Benz) makeBus() IBus {
	return &Bus{
		name:  "benz-bus",
		price: 1000000,
	}
}

type LiAuto struct{}

func (l *LiAuto) makeCar() ICar {
	return &Car{
		name:  "lixiang-car",
		price: 200000,
	}
}

func (l *LiAuto) makeBus() IBus {
	return &Bus{
		name:  "lixiang-bus",
		price: 400000,
	}
}
