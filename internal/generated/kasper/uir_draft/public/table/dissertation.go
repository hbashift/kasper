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

var Dissertation = newDissertationTable("public", "dissertation", "")

type dissertationTable struct {
	postgres.Table

	// Columns
	StudentID      postgres.ColumnString
	Status         postgres.ColumnString
	Title          postgres.ColumnString
	CreatedAt      postgres.ColumnDate
	UpdatedAt      postgres.ColumnDate
	DissertationID postgres.ColumnString
	Feedback       postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type DissertationTable struct {
	dissertationTable

	EXCLUDED dissertationTable
}

// AS creates new DissertationTable with assigned alias
func (a DissertationTable) AS(alias string) *DissertationTable {
	return newDissertationTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new DissertationTable with assigned schema name
func (a DissertationTable) FromSchema(schemaName string) *DissertationTable {
	return newDissertationTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new DissertationTable with assigned table prefix
func (a DissertationTable) WithPrefix(prefix string) *DissertationTable {
	return newDissertationTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new DissertationTable with assigned table suffix
func (a DissertationTable) WithSuffix(suffix string) *DissertationTable {
	return newDissertationTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newDissertationTable(schemaName, tableName, alias string) *DissertationTable {
	return &DissertationTable{
		dissertationTable: newDissertationTableImpl(schemaName, tableName, alias),
		EXCLUDED:          newDissertationTableImpl("", "excluded", ""),
	}
}

func newDissertationTableImpl(schemaName, tableName, alias string) dissertationTable {
	var (
		StudentIDColumn      = postgres.StringColumn("student_id")
		StatusColumn         = postgres.StringColumn("status")
		TitleColumn          = postgres.StringColumn("title")
		CreatedAtColumn      = postgres.DateColumn("created_at")
		UpdatedAtColumn      = postgres.DateColumn("updated_at")
		DissertationIDColumn = postgres.StringColumn("dissertation_id")
		FeedbackColumn       = postgres.StringColumn("feedback")
		allColumns           = postgres.ColumnList{StudentIDColumn, StatusColumn, TitleColumn, CreatedAtColumn, UpdatedAtColumn, DissertationIDColumn, FeedbackColumn}
		mutableColumns       = postgres.ColumnList{StudentIDColumn, StatusColumn, TitleColumn, CreatedAtColumn, UpdatedAtColumn, FeedbackColumn}
	)

	return dissertationTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		StudentID:      StudentIDColumn,
		Status:         StatusColumn,
		Title:          TitleColumn,
		CreatedAt:      CreatedAtColumn,
		UpdatedAt:      UpdatedAtColumn,
		DissertationID: DissertationIDColumn,
		Feedback:       FeedbackColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
