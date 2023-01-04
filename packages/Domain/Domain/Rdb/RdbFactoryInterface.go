package Rdb

type RdbFactoryInterface interface {
  NewRdbFactory(string) RdbInterface
}