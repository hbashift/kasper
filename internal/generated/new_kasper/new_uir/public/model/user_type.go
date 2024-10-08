//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import "errors"

type UserType string

const (
	UserType_Admin      UserType = "admin"
	UserType_Student    UserType = "student"
	UserType_Supervisor UserType = "supervisor"
)

func (e *UserType) Scan(value interface{}) error {
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
	case "admin":
		*e = UserType_Admin
	case "student":
		*e = UserType_Student
	case "supervisor":
		*e = UserType_Supervisor
	default:
		return errors.New("jet: Invalid scan value '" + enumValue + "' for UserType enum")
	}

	return nil
}

func (e UserType) String() string {
	return string(e)
}
