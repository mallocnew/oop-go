package abstract_factory

import "fmt"

func AbstractFactory() {
	benz, _ := GetCarFactory("benz")
	lixiang, _ := GetCarFactory("lixiang")

	benzCar := benz.makeCar()
	benzBus := benz.makeBus()

	lixiangCar := lixiang.makeCar()
	lixiangBus := lixiang.makeBus()

	printCarDetails(benzCar)
	printBusDetails(benzBus)

	printCarDetails(lixiangCar)
	printBusDetails(lixiangBus)
}

func printCarDetails(c ICar) {
	fmt.Printf("Car name: %s\n", c.getName())
	fmt.Printf("Car price: %d\n", c.getPrice())
}

func printBusDetails(b IBus) {
	fmt.Printf("Bus name: %s\n", b.getName())
	fmt.Printf("Bus price: %d\n", b.getPrice())
}
