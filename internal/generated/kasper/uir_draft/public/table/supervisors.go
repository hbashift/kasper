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

var Supervisors = newSupervisorsTable("public", "supervisors", "")

type supervisorsTable struct {
	postgres.Table

	// Columns
	ClientID     postgres.ColumnString
	FullName     postgres.ColumnString
	SupervisorID postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type SupervisorsTable struct {
	supervisorsTable

	EXCLUDED supervisorsTable
}

// AS creates new SupervisorsTable with assigned alias
func (a SupervisorsTable) AS(alias string) *SupervisorsTable {
	return newSupervisorsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new SupervisorsTable with assigned schema name
func (a SupervisorsTable) FromSchema(schemaName string) *SupervisorsTable {
	return newSupervisorsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new SupervisorsTable with assigned table prefix
func (a SupervisorsTable) WithPrefix(prefix string) *SupervisorsTable {
	return newSupervisorsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new SupervisorsTable with assigned table suffix
func (a SupervisorsTable) WithSuffix(suffix string) *SupervisorsTable {
	return newSupervisorsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newSupervisorsTable(schemaName, tableName, alias string) *SupervisorsTable {
	return &SupervisorsTable{
		supervisorsTable: newSupervisorsTableImpl(schemaName, tableName, alias),
		EXCLUDED:         newSupervisorsTableImpl("", "excluded", ""),
	}
}

func newSupervisorsTableImpl(schemaName, tableName, alias string) supervisorsTable {
	var (
		ClientIDColumn     = postgres.StringColumn("client_id")
		FullNameColumn     = postgres.StringColumn("full_name")
		SupervisorIDColumn = postgres.StringColumn("supervisor_id")
		allColumns         = postgres.ColumnList{ClientIDColumn, FullNameColumn, SupervisorIDColumn}
		mutableColumns     = postgres.ColumnList{ClientIDColumn, FullNameColumn}
	)

	return supervisorsTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ClientID:     ClientIDColumn,
		FullName:     FullNameColumn,
		SupervisorID: SupervisorIDColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
