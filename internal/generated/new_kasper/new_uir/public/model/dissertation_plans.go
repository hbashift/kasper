//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import (
	"github.com/google/uuid"
)

type DissertationPlans struct {
	PlanID    uuid.UUID `sql:"primary_key"`
	StudentID uuid.UUID
	Semester  int32
	PlanText  *string
}
