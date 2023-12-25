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

var TeachingLoad = newTeachingLoadTable("public", "teaching_load", "")

type teachingLoadTable struct {
	postgres.Table

	// Columns
	LoadID         postgres.ColumnString
	StudentID      postgres.ColumnString
	Semester       postgres.ColumnInteger
	Hours          postgres.ColumnInteger
	AdditionalLoad postgres.ColumnString
	LoadType       postgres.ColumnString
	MainTeacher    postgres.ColumnString
	GroupName      postgres.ColumnString
	SubjectName    postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type TeachingLoadTable struct {
	teachingLoadTable

	EXCLUDED teachingLoadTable
}

// AS creates new TeachingLoadTable with assigned alias
func (a TeachingLoadTable) AS(alias string) *TeachingLoadTable {
	return newTeachingLoadTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new TeachingLoadTable with assigned schema name
func (a TeachingLoadTable) FromSchema(schemaName string) *TeachingLoadTable {
	return newTeachingLoadTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new TeachingLoadTable with assigned table prefix
func (a TeachingLoadTable) WithPrefix(prefix string) *TeachingLoadTable {
	return newTeachingLoadTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new TeachingLoadTable with assigned table suffix
func (a TeachingLoadTable) WithSuffix(suffix string) *TeachingLoadTable {
	return newTeachingLoadTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newTeachingLoadTable(schemaName, tableName, alias string) *TeachingLoadTable {
	return &TeachingLoadTable{
		teachingLoadTable: newTeachingLoadTableImpl(schemaName, tableName, alias),
		EXCLUDED:          newTeachingLoadTableImpl("", "excluded", ""),
	}
}

func newTeachingLoadTableImpl(schemaName, tableName, alias string) teachingLoadTable {
	var (
		LoadIDColumn         = postgres.StringColumn("load_id")
		StudentIDColumn      = postgres.StringColumn("student_id")
		SemesterColumn       = postgres.IntegerColumn("semester")
		HoursColumn          = postgres.IntegerColumn("hours")
		AdditionalLoadColumn = postgres.StringColumn("additional_load")
		LoadTypeColumn       = postgres.StringColumn("load_type")
		MainTeacherColumn    = postgres.StringColumn("main_teacher")
		GroupNameColumn      = postgres.StringColumn("group_name")
		SubjectNameColumn    = postgres.StringColumn("subject_name")
		allColumns           = postgres.ColumnList{LoadIDColumn, StudentIDColumn, SemesterColumn, HoursColumn, AdditionalLoadColumn, LoadTypeColumn, MainTeacherColumn, GroupNameColumn, SubjectNameColumn}
		mutableColumns       = postgres.ColumnList{StudentIDColumn, SemesterColumn, HoursColumn, AdditionalLoadColumn, LoadTypeColumn, MainTeacherColumn, GroupNameColumn, SubjectNameColumn}
	)

	return teachingLoadTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		LoadID:         LoadIDColumn,
		StudentID:      StudentIDColumn,
		Semester:       SemesterColumn,
		Hours:          HoursColumn,
		AdditionalLoad: AdditionalLoadColumn,
		LoadType:       LoadTypeColumn,
		MainTeacher:    MainTeacherColumn,
		GroupName:      GroupNameColumn,
		SubjectName:    SubjectNameColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}