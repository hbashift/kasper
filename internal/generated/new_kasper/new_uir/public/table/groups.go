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

var Groups = newGroupsTable("public", "groups", "")

type groupsTable struct {
	postgres.Table

	//Columns
	GroupID   postgres.ColumnInteger
	GroupName postgres.ColumnString
	Archived  postgres.ColumnBool

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type GroupsTable struct {
	groupsTable

	EXCLUDED groupsTable
}

// AS creates new GroupsTable with assigned alias
func (a GroupsTable) AS(alias string) *GroupsTable {
	return newGroupsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new GroupsTable with assigned schema name
func (a GroupsTable) FromSchema(schemaName string) *GroupsTable {
	return newGroupsTable(schemaName, a.TableName(), a.Alias())
}

func newGroupsTable(schemaName, tableName, alias string) *GroupsTable {
	return &GroupsTable{
		groupsTable: newGroupsTableImpl(schemaName, tableName, alias),
		EXCLUDED:    newGroupsTableImpl("", "excluded", ""),
	}
}

func newGroupsTableImpl(schemaName, tableName, alias string) groupsTable {
	var (
		GroupIDColumn   = postgres.IntegerColumn("group_id")
		GroupNameColumn = postgres.StringColumn("group_name")
		ArchivedColumn  = postgres.BoolColumn("archived")
		allColumns      = postgres.ColumnList{GroupIDColumn, GroupNameColumn, ArchivedColumn}
		mutableColumns  = postgres.ColumnList{GroupNameColumn, ArchivedColumn}
	)

	return groupsTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		GroupID:   GroupIDColumn,
		GroupName: GroupNameColumn,
		Archived:  ArchivedColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
