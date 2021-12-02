package template

import (
	_ "embed"
	"fmt"
	"github.com/iancoleman/strcase"
	"io"
	"strings"
	"text/template"
)

var (
	//go:embed constants.tmpl
	TemplateConstants string
	//go:embed db.tmpl
	TemplateDB string
	//go:embed db_all.tmpl
	TemplateDBAll string
	//go:embed delete.tmpl
	TemplateDelete string
	//go:embed insert.tmpl
	TemplateInsert string
	//go:embed model.tmpl
	TemplateModel string
	//go:embed order.tmpl
	TemplateOrder string
	//go:embed runtime.tmpl
	TemplateRuntime string
	//go:embed select.tmpl
	TemplateSelect string
	//go:embed update.tmpl
	TemplateUpdate string
	//go:embed variable.tmpl
	TemplateVariable string
	//go:embed where.tmpl
	TemplateWhere string
)

var funcMap = template.FuncMap{
	"snakeCase": func(s string) string {
		return strcase.ToSnake(s)
	},
	"escape": func(s string) string {
		return fmt.Sprintf("`%s`", s)
	},
	"camelCase": func(s string) string {
		return strcase.ToLowerCamel(s)
	},
	"titleCase": func(s string) string {
		return strcase.ToCamel(s)
	},
	"stringsJoin": func(elems []string, sep string) string {
		return strings.Join(elems, sep)
	},
	"newSlice": func(args ...interface{}) interface{} {
		return args
	},
}

func Execute(text string, wr io.Writer, data interface{}) error {
	tmpl, err := template.New("").Funcs(funcMap).Parse(text)
	if err != nil {
		return err
	}
	return tmpl.Execute(wr, data)
}
