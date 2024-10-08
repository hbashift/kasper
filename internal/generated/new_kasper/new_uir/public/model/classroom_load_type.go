//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import "errors"

type ClassroomLoadType string

const (
	ClassroomLoadType_Practice   ClassroomLoadType = "practice"
	ClassroomLoadType_Lectures   ClassroomLoadType = "lectures"
	ClassroomLoadType_Laboratory ClassroomLoadType = "laboratory"
	ClassroomLoadType_Exam       ClassroomLoadType = "exam"
)

func (e *ClassroomLoadType) Scan(value interface{}) error {
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
	case "practice":
		*e = ClassroomLoadType_Practice
	case "lectures":
		*e = ClassroomLoadType_Lectures
	case "laboratory":
		*e = ClassroomLoadType_Laboratory
	case "exam":
		*e = ClassroomLoadType_Exam
	default:
		return errors.New("jet: Invalid scan value '" + enumValue + "' for ClassroomLoadType enum")
	}

	return nil
}

func (e ClassroomLoadType) String() string {
	return string(e)
}
