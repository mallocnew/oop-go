package abstract_factory

// package main

import "fmt"

type ICarFactory interface {
	makeCar() ICar
	makeBus() IBus
}

func GetCarFactory(company string) (ICarFactory, error) {
	if company == "benz" {
		return &Benz{}, nil
	}
	if company == "lixiang" {
		return &LiAuto{}, nil
	}
	return nil, fmt.Errorf("Error factory: %s", company)
}
