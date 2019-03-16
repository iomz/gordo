package model

import (
	"database/sql"
	"fmt"
	"strings"
)

type Resource struct {
	Ct int
	Path,
	Rt,
	If,
	Anchor,
	Title,
	Rel sql.NullString
}

// Resources is a LinkFormatter
type Resources []Resource

func (rs Resources) LinkFormat() string {
	var lf string
	for i := range rs {
		lf += toLinkFormat(rs[i])
		if i != len(rs)-1 {
			lf += ","
		}
	}
	return lf
}

func NewResource() *Resource {
	return &Resource{Ct: -1}
}

func isNonEmpty(s sql.NullString) bool {
	return s.Valid && s.String != ""
}

func toLinkFormat(r Resource) string {
	lf := fmt.Sprintf("<%s>", r.Path.String)

	if r.Ct >= 0 {
		lf += fmt.Sprintf(";ct=%d", r.Ct)
	}

	if isNonEmpty(r.Title) {
		lf += fmt.Sprintf(";title=\"%s\"", r.Title.String)
	}

	if isNonEmpty(r.Rel) {
		lf += fmt.Sprintf(";rel=\"%s\"", r.Rel.String)
	}

	if isNonEmpty(r.Rt) {
		lf += fmt.Sprintf(";rt=\"%s\"", r.Rt.String)
	}

	if isNonEmpty(r.If) {
		lf += fmt.Sprintf(";if=\"%s\"", r.If.String)
	}

	if isNonEmpty(r.Anchor) {
		lf += fmt.Sprintf(";anchor=\"%s\"", r.Anchor.String)
	}

	return lf
}

// only allow queries with the following criteria
var safeLookupKeys = map[string]string{
	"rt": "string",
	"ct": "int",
	"if": "string",
}

func buildQueryStmt(qry []string) string {
	stmt := `SELECT path, ct, rt, if, anchor, title, rel FROM resource`

	var cond []string
	for i := range qry {
		kv := strings.Split(qry[i], "=")
		if len(kv) == 2 {
			_, ok := safeLookupKeys[kv[0]]
			if ok {
				// TODO(tho) type based formatting
				cond = append(cond, fmt.Sprintf("%s = '%s'", kv[0], kv[1]))
			}
		}
	}

	if len(cond) > 0 {
		stmt += fmt.Sprintf(" WHERE ")
		for i := range cond {
			stmt += cond[i]
			if i != len(cond)-1 {
				stmt += " AND "
			}
		}
	}

	return stmt
}

// LookupRes queries the DB and returns the result set as an array of resources
func (m *Model) ResourceLookup(qry []string) (Resources, error) {
	stmt := buildQueryStmt(qry)

	fmt.Println(qry)
	fmt.Println(stmt)

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rs Resources
	for rows.Next() {
		r := NewResource()
		err = rows.Scan(
			&r.Path,
			&r.Ct,
			&r.Rt,
			&r.If,
			&r.Anchor,
			&r.Title,
			&r.Rel,
		)
		if err != nil {
			return nil, err
		}
		rs = append(rs, *r)
	}

	return rs, nil
}
