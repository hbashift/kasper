//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package enum

import "github.com/go-jet/jet/v2/postgres"

var UserType = &struct {
	Administrator postgres.StringExpression
	Student       postgres.StringExpression
	Supervisor    postgres.StringExpression
}{
	Administrator: postgres.NewEnumValue("administrator"),
	Student:       postgres.NewEnumValue("student"),
	Supervisor:    postgres.NewEnumValue("supervisor"),
}
