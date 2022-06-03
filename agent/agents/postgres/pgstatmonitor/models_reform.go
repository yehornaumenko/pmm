// Code generated by gopkg.in/reform.v1. DO NOT EDIT.

package pgstatmonitor

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type pgStatDatabaseViewType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("pg_catalog").
func (v *pgStatDatabaseViewType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("pg_stat_database").
func (v *pgStatDatabaseViewType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *pgStatDatabaseViewType) Columns() []string {
	return []string{
		"datid",
		"datname",
	}
}

// NewStruct makes a new struct for that view or table.
func (v *pgStatDatabaseViewType) NewStruct() reform.Struct {
	return new(pgStatDatabase)
}

// pgStatDatabaseView represents pg_stat_database view or table in SQL database.
var pgStatDatabaseView = &pgStatDatabaseViewType{
	s: parse.StructInfo{
		Type:      "pgStatDatabase",
		SQLSchema: "pg_catalog",
		SQLName:   "pg_stat_database",
		Fields: []parse.FieldInfo{
			{Name: "DatID", Type: "int64", Column: "datid"},
			{Name: "DatName", Type: "*string", Column: "datname"},
		},
		PKFieldIndex: -1,
	},
	z: new(pgStatDatabase).Values(),
}

// String returns a string representation of this struct or record.
func (s pgStatDatabase) String() string {
	res := make([]string, 2)
	res[0] = "DatID: " + reform.Inspect(s.DatID, true)
	res[1] = "DatName: " + reform.Inspect(s.DatName, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *pgStatDatabase) Values() []interface{} {
	return []interface{}{
		s.DatID,
		s.DatName,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *pgStatDatabase) Pointers() []interface{} {
	return []interface{}{
		&s.DatID,
		&s.DatName,
	}
}

// View returns View object for that struct.
func (s *pgStatDatabase) View() reform.View {
	return pgStatDatabaseView
}

// check interfaces
var (
	_ reform.View   = pgStatDatabaseView
	_ reform.Struct = (*pgStatDatabase)(nil)
	_ fmt.Stringer  = (*pgStatDatabase)(nil)
)

type pgUserViewType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("pg_catalog").
func (v *pgUserViewType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("pg_user").
func (v *pgUserViewType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *pgUserViewType) Columns() []string {
	return []string{
		"usesysid",
		"usename",
	}
}

// NewStruct makes a new struct for that view or table.
func (v *pgUserViewType) NewStruct() reform.Struct {
	return new(pgUser)
}

// pgUserView represents pg_user view or table in SQL database.
var pgUserView = &pgUserViewType{
	s: parse.StructInfo{
		Type:      "pgUser",
		SQLSchema: "pg_catalog",
		SQLName:   "pg_user",
		Fields: []parse.FieldInfo{
			{Name: "UserID", Type: "int64", Column: "usesysid"},
			{Name: "UserName", Type: "*string", Column: "usename"},
		},
		PKFieldIndex: -1,
	},
	z: new(pgUser).Values(),
}

// String returns a string representation of this struct or record.
func (s pgUser) String() string {
	res := make([]string, 2)
	res[0] = "UserID: " + reform.Inspect(s.UserID, true)
	res[1] = "UserName: " + reform.Inspect(s.UserName, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *pgUser) Values() []interface{} {
	return []interface{}{
		s.UserID,
		s.UserName,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *pgUser) Pointers() []interface{} {
	return []interface{}{
		&s.UserID,
		&s.UserName,
	}
}

// View returns View object for that struct.
func (s *pgUser) View() reform.View {
	return pgUserView
}

// check interfaces
var (
	_ reform.View   = pgUserView
	_ reform.Struct = (*pgUser)(nil)
	_ fmt.Stringer  = (*pgUser)(nil)
)

type pgStatMonitorSettingsViewType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *pgStatMonitorSettingsViewType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("pg_stat_monitor_settings").
func (v *pgStatMonitorSettingsViewType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *pgStatMonitorSettingsViewType) Columns() []string {
	return []string{
		"name",
		"value",
	}
}

// NewStruct makes a new struct for that view or table.
func (v *pgStatMonitorSettingsViewType) NewStruct() reform.Struct {
	return new(pgStatMonitorSettings)
}

// pgStatMonitorSettingsView represents pg_stat_monitor_settings view or table in SQL database.
var pgStatMonitorSettingsView = &pgStatMonitorSettingsViewType{
	s: parse.StructInfo{
		Type:    "pgStatMonitorSettings",
		SQLName: "pg_stat_monitor_settings",
		Fields: []parse.FieldInfo{
			{Name: "Name", Type: "string", Column: "name"},
			{Name: "Value", Type: "int64", Column: "value"},
		},
		PKFieldIndex: -1,
	},
	z: new(pgStatMonitorSettings).Values(),
}

// String returns a string representation of this struct or record.
func (s pgStatMonitorSettings) String() string {
	res := make([]string, 2)
	res[0] = "Name: " + reform.Inspect(s.Name, true)
	res[1] = "Value: " + reform.Inspect(s.Value, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *pgStatMonitorSettings) Values() []interface{} {
	return []interface{}{
		s.Name,
		s.Value,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *pgStatMonitorSettings) Pointers() []interface{} {
	return []interface{}{
		&s.Name,
		&s.Value,
	}
}

// View returns View object for that struct.
func (s *pgStatMonitorSettings) View() reform.View {
	return pgStatMonitorSettingsView
}

// check interfaces
var (
	_ reform.View   = pgStatMonitorSettingsView
	_ reform.Struct = (*pgStatMonitorSettings)(nil)
	_ fmt.Stringer  = (*pgStatMonitorSettings)(nil)
)

type pgStatMonitorSettingsTextValueViewType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *pgStatMonitorSettingsTextValueViewType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("pg_stat_monitor_settings").
func (v *pgStatMonitorSettingsTextValueViewType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *pgStatMonitorSettingsTextValueViewType) Columns() []string {
	return []string{
		"name",
		"value",
	}
}

// NewStruct makes a new struct for that view or table.
func (v *pgStatMonitorSettingsTextValueViewType) NewStruct() reform.Struct {
	return new(pgStatMonitorSettingsTextValue)
}

// pgStatMonitorSettingsTextValueView represents pg_stat_monitor_settings view or table in SQL database.
var pgStatMonitorSettingsTextValueView = &pgStatMonitorSettingsTextValueViewType{
	s: parse.StructInfo{
		Type:    "pgStatMonitorSettingsTextValue",
		SQLName: "pg_stat_monitor_settings",
		Fields: []parse.FieldInfo{
			{Name: "Name", Type: "string", Column: "name"},
			{Name: "Value", Type: "string", Column: "value"},
		},
		PKFieldIndex: -1,
	},
	z: new(pgStatMonitorSettingsTextValue).Values(),
}

// String returns a string representation of this struct or record.
func (s pgStatMonitorSettingsTextValue) String() string {
	res := make([]string, 2)
	res[0] = "Name: " + reform.Inspect(s.Name, true)
	res[1] = "Value: " + reform.Inspect(s.Value, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *pgStatMonitorSettingsTextValue) Values() []interface{} {
	return []interface{}{
		s.Name,
		s.Value,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *pgStatMonitorSettingsTextValue) Pointers() []interface{} {
	return []interface{}{
		&s.Name,
		&s.Value,
	}
}

// View returns View object for that struct.
func (s *pgStatMonitorSettingsTextValue) View() reform.View {
	return pgStatMonitorSettingsTextValueView
}

// check interfaces
var (
	_ reform.View   = pgStatMonitorSettingsTextValueView
	_ reform.Struct = (*pgStatMonitorSettingsTextValue)(nil)
	_ fmt.Stringer  = (*pgStatMonitorSettingsTextValue)(nil)
)

type pgStatMonitorErrorsViewType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *pgStatMonitorErrorsViewType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("pg_stat_monitor_errors").
func (v *pgStatMonitorErrorsViewType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *pgStatMonitorErrorsViewType) Columns() []string {
	return []string{
		"severity",
		"message",
		"msgtime",
		"calls",
	}
}

// NewStruct makes a new struct for that view or table.
func (v *pgStatMonitorErrorsViewType) NewStruct() reform.Struct {
	return new(pgStatMonitorErrors)
}

// pgStatMonitorErrorsView represents pg_stat_monitor_errors view or table in SQL database.
var pgStatMonitorErrorsView = &pgStatMonitorErrorsViewType{
	s: parse.StructInfo{
		Type:    "pgStatMonitorErrors",
		SQLName: "pg_stat_monitor_errors",
		Fields: []parse.FieldInfo{
			{Name: "Severity", Type: "string", Column: "severity"},
			{Name: "Message", Type: "string", Column: "message"},
			{Name: "MessageTime", Type: "string", Column: "msgtime"},
			{Name: "Calls", Type: "int64", Column: "calls"},
		},
		PKFieldIndex: -1,
	},
	z: new(pgStatMonitorErrors).Values(),
}

// String returns a string representation of this struct or record.
func (s pgStatMonitorErrors) String() string {
	res := make([]string, 4)
	res[0] = "Severity: " + reform.Inspect(s.Severity, true)
	res[1] = "Message: " + reform.Inspect(s.Message, true)
	res[2] = "MessageTime: " + reform.Inspect(s.MessageTime, true)
	res[3] = "Calls: " + reform.Inspect(s.Calls, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *pgStatMonitorErrors) Values() []interface{} {
	return []interface{}{
		s.Severity,
		s.Message,
		s.MessageTime,
		s.Calls,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *pgStatMonitorErrors) Pointers() []interface{} {
	return []interface{}{
		&s.Severity,
		&s.Message,
		&s.MessageTime,
		&s.Calls,
	}
}

// View returns View object for that struct.
func (s *pgStatMonitorErrors) View() reform.View {
	return pgStatMonitorErrorsView
}

// check interfaces
var (
	_ reform.View   = pgStatMonitorErrorsView
	_ reform.Struct = (*pgStatMonitorErrors)(nil)
	_ fmt.Stringer  = (*pgStatMonitorErrors)(nil)
)

func init() {
	parse.AssertUpToDate(&pgStatDatabaseView.s, new(pgStatDatabase))
	parse.AssertUpToDate(&pgUserView.s, new(pgUser))
	parse.AssertUpToDate(&pgStatMonitorSettingsView.s, new(pgStatMonitorSettings))
	parse.AssertUpToDate(&pgStatMonitorSettingsTextValueView.s, new(pgStatMonitorSettingsTextValue))
	parse.AssertUpToDate(&pgStatMonitorErrorsView.s, new(pgStatMonitorErrors))
}
