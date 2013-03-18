package main

import (
	"thrift"
    "bigselect/api"
)

type BigSelectHandler struct {
}

func NewBigSelectHandler() *BigSelectHandler {
	return &BigSelectHandler{}
}

func (p * BigSelectHandler) Select(queues thrift.TList) (ret *api.Reservation, err error) {
    ret = NewReservation()
}

func (p * BigSelectHandler) Release(reservation *api.Reservation) (err error) {
}

func (p * BigSelectHandler) Push(queue string, data string) (err error) {
}

func (p * BigSelectHandler) Pop(reservation *api.Reservation) (ret string, err error) {
}

func (p * BigSelectHandler) Peek(reservation *api.Reservation, nitems int32) (ret thrift.TList, err error) {
}

