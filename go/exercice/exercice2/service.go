package main

import (
)

type OperationService interface {
	Somme(int64, int64) (int64, error)
}

type operationService struct{}

func (operationService) Somme(nombre1 int64, nombre2 int64) (int64, error) {
	return nombre1 + nombre2, nil
}
