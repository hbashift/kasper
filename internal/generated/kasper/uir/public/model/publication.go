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

type Publication struct {
	PublicationID uuid.UUID `sql:"primary_key"`
	Name          string
	OutputData    string
	NumOfPages    int32
	CoAuthors     string
	TypeID        int32
}
