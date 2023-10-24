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

var AuthorizationToken = newAuthorizationTokenTable("public", "authorization_token", "")

type authorizationTokenTable struct {
	postgres.Table

	// Columns
	TokenID        postgres.ColumnString
	ClientID       postgres.ColumnString
	TokenStatus    postgres.ColumnString
	TokenNumber    postgres.ColumnString
	CreatedAt      postgres.ColumnTimestampz
	UpdatedAt      postgres.ColumnTimestampz
	ExpirationDate postgres.ColumnDate

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type AuthorizationTokenTable struct {
	authorizationTokenTable

	EXCLUDED authorizationTokenTable
}

// AS creates new AuthorizationTokenTable with assigned alias
func (a AuthorizationTokenTable) AS(alias string) *AuthorizationTokenTable {
	return newAuthorizationTokenTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new AuthorizationTokenTable with assigned schema name
func (a AuthorizationTokenTable) FromSchema(schemaName string) *AuthorizationTokenTable {
	return newAuthorizationTokenTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new AuthorizationTokenTable with assigned table prefix
func (a AuthorizationTokenTable) WithPrefix(prefix string) *AuthorizationTokenTable {
	return newAuthorizationTokenTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new AuthorizationTokenTable with assigned table suffix
func (a AuthorizationTokenTable) WithSuffix(suffix string) *AuthorizationTokenTable {
	return newAuthorizationTokenTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newAuthorizationTokenTable(schemaName, tableName, alias string) *AuthorizationTokenTable {
	return &AuthorizationTokenTable{
		authorizationTokenTable: newAuthorizationTokenTableImpl(schemaName, tableName, alias),
		EXCLUDED:                newAuthorizationTokenTableImpl("", "excluded", ""),
	}
}

func newAuthorizationTokenTableImpl(schemaName, tableName, alias string) authorizationTokenTable {
	var (
		TokenIDColumn        = postgres.StringColumn("token_id")
		ClientIDColumn       = postgres.StringColumn("client_id")
		TokenStatusColumn    = postgres.StringColumn("token_status")
		TokenNumberColumn    = postgres.StringColumn("token_number")
		CreatedAtColumn      = postgres.TimestampzColumn("created_at")
		UpdatedAtColumn      = postgres.TimestampzColumn("updated_at")
		ExpirationDateColumn = postgres.DateColumn("expiration_date")
		allColumns           = postgres.ColumnList{TokenIDColumn, ClientIDColumn, TokenStatusColumn, TokenNumberColumn, CreatedAtColumn, UpdatedAtColumn, ExpirationDateColumn}
		mutableColumns       = postgres.ColumnList{ClientIDColumn, TokenStatusColumn, TokenNumberColumn, CreatedAtColumn, UpdatedAtColumn, ExpirationDateColumn}
	)

	return authorizationTokenTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		TokenID:        TokenIDColumn,
		ClientID:       ClientIDColumn,
		TokenStatus:    TokenStatusColumn,
		TokenNumber:    TokenNumberColumn,
		CreatedAt:      CreatedAtColumn,
		UpdatedAt:      UpdatedAtColumn,
		ExpirationDate: ExpirationDateColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
