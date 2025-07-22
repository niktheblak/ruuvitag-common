package psql

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/spf13/cast"
)

var passwordRegexp = regexp.MustCompile(`password=\S+\s`)

// RemovePassword removes password from a psqlInfo style string
func RemovePassword(connString string) string {
	return passwordRegexp.ReplaceAllString(connString, "password=[redacted] ")
}

func CreateConnString(cfg map[string]any) (connString string, err error) {
	var (
		host     = cast.ToString(cfg["host"])
		port     = cast.ToInt(cfg["port"])
		username = cast.ToString(cfg["username"])
		password = cast.ToString(cfg["password"])
		database = cast.ToString(cfg["database"])
		table    = cast.ToString(cfg["table"])
		sslmode  = cast.ToString(cfg["sslmode"])
		sslcert  = cast.ToString(cfg["sslcert"])
		sslkey   = cast.ToString(cfg["sslkey"])
	)
	if host == "" {
		err = fmt.Errorf("PostgreSQL host must be specified")
		return
	}
	if database == "" {
		err = fmt.Errorf("PostgreSQL database name must be specified")
		return
	}
	if table == "" {
		err = fmt.Errorf("PostgreSQL table name must be specified")
		return
	}
	if sslmode == "" {
		sslmode = "disable"
	}
	builder := new(strings.Builder)
	builder.WriteString("host=")
	builder.WriteString(host)
	builder.WriteString(" ")
	builder.WriteString("port=")
	builder.WriteString(strconv.Itoa(port))
	builder.WriteString(" ")
	if username != "" {
		builder.WriteString("user=")
		builder.WriteString(username)
		builder.WriteString(" ")
	}
	if password != "" {
		builder.WriteString("password=")
		builder.WriteString(password)
		builder.WriteString(" ")
	}
	builder.WriteString("dbname=")
	builder.WriteString(database)
	builder.WriteString(" ")
	builder.WriteString("sslmode=")
	builder.WriteString(sslmode)
	if sslcert != "" && sslkey != "" {
		builder.WriteString(" ")
		builder.WriteString("sslcert=")
		builder.WriteString(sslcert)
		builder.WriteString(" ")
		builder.WriteString("sslkey=")
		builder.WriteString(sslkey)
	}
	connString = builder.String()
	return
}
