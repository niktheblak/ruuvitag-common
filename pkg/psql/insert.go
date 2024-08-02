package psql

import (
	"fmt"
	"strings"
	"text/template"

	"github.com/niktheblak/ruuvitag-common/pkg/columnmap"
	"github.com/niktheblak/ruuvitag-common/pkg/sensor"
)

func BuildInsertQuery(table string, columns map[string]string) (string, error) {
	var tmplBuilder strings.Builder
	tmplBuilder.WriteString("INSERT INTO ")
	tmplBuilder.WriteString(table)
	tmplBuilder.WriteString("(")
	var includedColumns []string
	for _, c := range sensor.DefaultColumns {
		_, ok := columns[c]
		if ok {
			includedColumns = append(includedColumns, fmt.Sprintf("{{.%s}}", c))
		}
	}
	tmplBuilder.WriteString(strings.Join(includedColumns, ","))
	tmplBuilder.WriteString(")")
	tmpl, err := template.New("insertQuery").Parse(tmplBuilder.String())
	if err != nil {
		return "", err
	}
	var queryBuilder strings.Builder
	if err := tmpl.Execute(&queryBuilder, columns); err != nil {
		return "", err
	}
	var placeholders []string
	queryBuilder.WriteString(" VALUES (")
	for i := 1; i < len(columns)+1; i++ {
		placeholders = append(placeholders, fmt.Sprintf("$%d", i))
	}
	queryBuilder.WriteString(strings.Join(placeholders, ","))
	queryBuilder.WriteString(")")
	return queryBuilder.String(), nil
}

func BuildQueryArguments(columns map[string]string, data sensor.Data) []any {
	var args []any
	columnmap.Collect(columns, data, func(_ string, v any) {
		args = append(args, v)
	})
	return args
}
