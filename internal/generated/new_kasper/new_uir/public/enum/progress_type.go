//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package enum

import "github.com/go-jet/jet/v2/postgres"

var ProgressType = &struct {
	Intro      postgres.StringExpression
	Ch1        postgres.StringExpression
	Ch2        postgres.StringExpression
	Ch3        postgres.StringExpression
	Ch4        postgres.StringExpression
	Ch5        postgres.StringExpression
	Ch6        postgres.StringExpression
	End        postgres.StringExpression
	Literature postgres.StringExpression
	Abstract   postgres.StringExpression
}{
	Intro:      postgres.NewEnumValue("intro"),
	Ch1:        postgres.NewEnumValue("ch. 1"),
	Ch2:        postgres.NewEnumValue("ch. 2"),
	Ch3:        postgres.NewEnumValue("ch. 3"),
	Ch4:        postgres.NewEnumValue("ch. 4"),
	Ch5:        postgres.NewEnumValue("ch. 5"),
	Ch6:        postgres.NewEnumValue("ch. 6"),
	End:        postgres.NewEnumValue("end"),
	Literature: postgres.NewEnumValue("literature"),
	Abstract:   postgres.NewEnumValue("abstract"),
}
