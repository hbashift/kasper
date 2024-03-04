//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/postgres"
)

var Marks = newMarksTable("public", "marks", "")

type marksTable struct {
	postgres.Table

	//Columns
	StudentID postgres.ColumnString
	Mark      postgres.ColumnInteger
	Semester  postgres.ColumnInteger

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type MarksTable struct {
	marksTable

	EXCLUDED marksTable
}

// AS creates new MarksTable with assigned alias
func (a MarksTable) AS(alias string) *MarksTable {
	return newMarksTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new MarksTable with assigned schema name
func (a MarksTable) FromSchema(schemaName string) *MarksTable {
	return newMarksTable(schemaName, a.TableName(), a.Alias())
}

func newMarksTable(schemaName, tableName, alias string) *MarksTable {
	return &MarksTable{
		marksTable: newMarksTableImpl(schemaName, tableName, alias),
		EXCLUDED:   newMarksTableImpl("", "excluded", ""),
	}
}

func newMarksTableImpl(schemaName, tableName, alias string) marksTable {
	var (
		StudentIDColumn = postgres.StringColumn("student_id")
		MarkColumn      = postgres.IntegerColumn("mark")
		SemesterColumn  = postgres.IntegerColumn("semester")
		allColumns      = postgres.ColumnList{StudentIDColumn, MarkColumn, SemesterColumn}
		mutableColumns  = postgres.ColumnList{StudentIDColumn, MarkColumn, SemesterColumn}
	)

	return marksTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		StudentID: StudentIDColumn,
		Mark:      MarkColumn,
		Semester:  SemesterColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}