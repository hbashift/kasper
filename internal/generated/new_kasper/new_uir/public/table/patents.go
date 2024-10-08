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

var Patents = newPatentsTable("public", "patents", "")

type patentsTable struct {
	postgres.Table

	//Columns
	PatentID         postgres.ColumnString
	WorksID          postgres.ColumnString
	Name             postgres.ColumnString
	RegistrationDate postgres.ColumnTimestampz
	Type             postgres.ColumnString
	AddInfo          postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type PatentsTable struct {
	patentsTable

	EXCLUDED patentsTable
}

// AS creates new PatentsTable with assigned alias
func (a PatentsTable) AS(alias string) *PatentsTable {
	return newPatentsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new PatentsTable with assigned schema name
func (a PatentsTable) FromSchema(schemaName string) *PatentsTable {
	return newPatentsTable(schemaName, a.TableName(), a.Alias())
}

func newPatentsTable(schemaName, tableName, alias string) *PatentsTable {
	return &PatentsTable{
		patentsTable: newPatentsTableImpl(schemaName, tableName, alias),
		EXCLUDED:     newPatentsTableImpl("", "excluded", ""),
	}
}

func newPatentsTableImpl(schemaName, tableName, alias string) patentsTable {
	var (
		PatentIDColumn         = postgres.StringColumn("patent_id")
		WorksIDColumn          = postgres.StringColumn("works_id")
		NameColumn             = postgres.StringColumn("name")
		RegistrationDateColumn = postgres.TimestampzColumn("registration_date")
		TypeColumn             = postgres.StringColumn("type")
		AddInfoColumn          = postgres.StringColumn("add_info")
		allColumns             = postgres.ColumnList{PatentIDColumn, WorksIDColumn, NameColumn, RegistrationDateColumn, TypeColumn, AddInfoColumn}
		mutableColumns         = postgres.ColumnList{WorksIDColumn, NameColumn, RegistrationDateColumn, TypeColumn, AddInfoColumn}
	)

	return patentsTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		PatentID:         PatentIDColumn,
		WorksID:          WorksIDColumn,
		Name:             NameColumn,
		RegistrationDate: RegistrationDateColumn,
		Type:             TypeColumn,
		AddInfo:          AddInfoColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
