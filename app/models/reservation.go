package models

import (
	"github.com/revel/revel"
)

type Reservation struct {
	ReservationId int64
	PrivateId     string
	Email         string
	Name          string
	EventId       int64
}

func (r Reservation) Validate(v *revel.Validation) {

}
