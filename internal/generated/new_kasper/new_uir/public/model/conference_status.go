//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import "errors"

type ConferenceStatus string

const (
	ConferenceStatus_Registered ConferenceStatus = "registered"
	ConferenceStatus_Performed  ConferenceStatus = "performed"
)

func (e *ConferenceStatus) Scan(value interface{}) error {
	var enumValue string
	switch val := value.(type) {
	case string:
		enumValue = val
	case []byte:
		enumValue = string(val)
	default:
		return errors.New("jet: Invalid scan value for AllTypesEnum enum. Enum value has to be of type string or []byte")
	}

	switch enumValue {
	case "registered":
		*e = ConferenceStatus_Registered
	case "performed":
		*e = ConferenceStatus_Performed
	default:
		return errors.New("jet: Invalid scan value '" + enumValue + "' for ConferenceStatus enum")
	}

	return nil
}

func (e ConferenceStatus) String() string {
	return string(e)
}
