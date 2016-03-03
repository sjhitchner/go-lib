package main

import (
	"flag"
	"fmt"
	libio "github.com/sjhitchner/go-lib/io"
	"os"
	"strings"
	"text/template"
	"time"
)

var (
	packageName string
	dateName    string
	dateFormat  string
)

func init() {
	flag.StringVar(&packageName, "package", "", "Name of package")
	flag.StringVar(&dateName, "name", "", "Name of date type")
	flag.StringVar(&dateFormat, "format", "", "Format of date type")
}

const HEADER = `package {{.Package}}

import (
	"encoding/json"
	"fmt"
	"time"
)
`
const TEMPLATE = `
const {{.NameUpper}}_FORMAT = "{{.Format}}"

type {{.Name}} struct {
	time.Time
}

func (t *{{.Name}}) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Date should be a string, got %s", data)
	}

	d, err := time.Parse({{.NameUpper}}_FORMAT, s)
	if err != nil {
		return fmt.Errorf("Date should be in format: %s, got %v", {{.NameUpper}}_FORMAT, s)
	}
	t.Time = d
	return nil
}

func (t {{.Name}}) MarshalJSON() ([]byte, error) {
	s := t.Format({{.NameUpper}}_FORMAT)
	return []byte(s), nil
}

`

type HeaderType struct {
	Package string
}

type DateType struct {
	Name      string
	NameUpper string
	Format    string
}

func main() {
	flag.Parse()

	if packageName == "" {
		Error("Name of package must be set")
	}

	if dateName == "" {
		Error("Name of date must be set")
	}

	if dateFormat == "" {
		Error("Format of date must be set")
	}

	if _, err := time.Parse(dateFormat, dateFormat); err != nil {
		Error("Date format provided is invalid: " + dateFormat)
	}

	header, err := template.New("header").Parse(HEADER)
	if err != nil {
		Error(err.Error())
	}

	body, err := template.New("body").Parse(TEMPLATE)
	if err != nil {
		Error(err.Error())
	}

	fileName := "jsondate_generated.go"
	f, isNew, err := libio.OpenForAppendOrNew(fileName)
	if err != nil {
		Error(err.Error())
	}
	defer f.Close()

	if isNew {
		if err := header.Execute(f, HeaderType{packageName}); err != nil {
			Error(err.Error())
		}
	}

	dateType := DateType{
		Name:      dateName,
		NameUpper: strings.ToUpper(dateName),
		Format:    dateFormat,
	}

	if err := body.Execute(f, dateType); err != nil {
		Error(err.Error())
	}
}

func Error(err string) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(-1)
}
