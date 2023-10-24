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

var Supervisor = newSupervisorTable("public", "supervisor", "")

type supervisorTable struct {
	postgres.Table

	// Columns
	UserID       postgres.ColumnString
	FullName     postgres.ColumnString
	SupervisorID postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type SupervisorTable struct {
	supervisorTable

	EXCLUDED supervisorTable
}

// AS creates new SupervisorTable with assigned alias
func (a SupervisorTable) AS(alias string) *SupervisorTable {
	return newSupervisorTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new SupervisorTable with assigned schema name
func (a SupervisorTable) FromSchema(schemaName string) *SupervisorTable {
	return newSupervisorTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new SupervisorTable with assigned table prefix
func (a SupervisorTable) WithPrefix(prefix string) *SupervisorTable {
	return newSupervisorTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new SupervisorTable with assigned table suffix
func (a SupervisorTable) WithSuffix(suffix string) *SupervisorTable {
	return newSupervisorTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newSupervisorTable(schemaName, tableName, alias string) *SupervisorTable {
	return &SupervisorTable{
		supervisorTable: newSupervisorTableImpl(schemaName, tableName, alias),
		EXCLUDED:        newSupervisorTableImpl("", "excluded", ""),
	}
}

func newSupervisorTableImpl(schemaName, tableName, alias string) supervisorTable {
	var (
		UserIDColumn       = postgres.StringColumn("user_id")
		FullNameColumn     = postgres.StringColumn("full_name")
		SupervisorIDColumn = postgres.StringColumn("supervisor_id")
		allColumns         = postgres.ColumnList{UserIDColumn, FullNameColumn, SupervisorIDColumn}
		mutableColumns     = postgres.ColumnList{UserIDColumn, FullNameColumn}
	)

	return supervisorTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		UserID:       UserIDColumn,
		FullName:     FullNameColumn,
		SupervisorID: SupervisorIDColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
