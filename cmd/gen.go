/*
Copyright © 2021 Ramin Zhong <zhongyuanjia.uni@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	_ "embed"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/iancoleman/strcase"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
	"strings"
	"text/template"
)

type Field struct {
	ParamName     string
	ParamListName string
	SQLName       string
	StructName    string
	Type          string
	Tag           string
	ZeroValue     string
}

type Model struct {
	Pkg         string
	Name        string
	SQLName     string
	Placeholder string
	Fields      []*Field
}

var prefix = ""
var appendPrefix = false
var pkg = ""

//go:embed repo.tmpl
var modelTemplate string

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "generate model repository",
	Run: func(cmd *cobra.Command, args []string) {
		tables := getTables()
		log.Printf("start generate %+v\n", tables)
		escapeChar := "`"

		tmpl := template.Must(template.New("model").Funcs(template.FuncMap{
			"toSnake": func(s string) string {
				return strcase.ToSnake(s)
			},
			"escape": func(s string) string {
				return fmt.Sprintf("%s%s%s", escapeChar, s, escapeChar)
			},
			"toLowerCamel": func(s string) string {
				return strcase.ToLowerCamel(s)
			},
			"stringsJoin": func(elems []string, sep string) string {
				return strings.Join(elems, sep)
			},
			"newSlice": func(args ...interface{}) interface{} {
				return args
			},
		}).Parse(modelTemplate))

		for _, table := range tables {
			name := table
			if !appendPrefix && prefix != "" && strings.HasPrefix(name, prefix) {
				name = name[len(prefix):]
			}
			if err := os.MkdirAll(pkg, 0755); err != nil {
				panic(err)
			}
			file, err := os.OpenFile(fmt.Sprintf("%s/%s.gen.go", pkg, name), os.O_CREATE|os.O_WRONLY|os.O_TRUNC,
				os.ModePerm)
			if err != nil {
				panic(err)
			}
			name = strcase.ToCamel(name)
			if err := tmpl.Execute(file, &Model{
				Pkg:         pkg,
				Name:        name,
				SQLName:     table,
				Fields:      getFieldsFromTable(table),
				Placeholder: "?",
			}); err != nil {
				panic(err)
			}
		}

		if err := exec.Command("go", "fmt", "./...").Run(); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(genCmd)
	genCmd.Flags().StringVarP(&prefix, "prefix", "p", "", "生成指定前缀的表")
	genCmd.Flags().BoolVar(&appendPrefix, "append-prefix", false, "生成的表名是否忽略前缀")
	genCmd.Flags().StringVar(&pkg, "package", "repo", "包名")
}

func getTables() []string {
	rows, err := db.Queryx("SHOW TABLES")
	if err != nil {
		panic(err)
	}
	tables := make([]string, 0)
	for rows.Next() {
		var str string
		if err := rows.Scan(&str); err != nil {
			panic(err)
		}
		tables = append(tables, str)
	}

	// 选择指定的表生成
	if prefix == "" {
		prompt := &survey.MultiSelect{
			Message:  "请选择要生成的表",
			Options:  tables,
			PageSize: 10,
		}

		tables = make([]string, 0)
		_ = survey.AskOne(prompt, &tables, survey.WithPageSize(10))
		return tables
	}

	result := make([]string, 0)
	for _, table := range tables {
		if !strings.HasPrefix(table, prefix) {
			continue
		}

		result = append(result, table)
	}

	return result
}

func getFieldsFromTable(table string) []*Field {
	rows, err := db.Queryx(fmt.Sprintf("SHOW FULL COLUMNS FROM `%s`", table))
	if err != nil {
		panic(err)
	}

	type ColumnInfo struct {
		Field      string  `db:"Field"`
		Type       string  `db:"Type"`
		Collation  *string `db:"Collation"`
		Null       string  `db:"Null"`
		Key        string  `db:"Key"`
		Default    *string `db:"Default"`
		Extra      string  `db:"Extra"`
		Privileges string  `db:"Privileges"`
		Comment    string  `db:"Comment"`
	}

	typeMap := map[string]string{
		"float":     "float32",
		"double":    "float64",
		"decimal":   "float64",
		"date":      "time.Time",
		"time":      "time.Time",
		"year":      "time.Time",
		"datetime":  "time.Time",
		"timestamp": "time.Time",
		"json":      "string",
	}

	keywordMap := map[string]bool{
		"break":       true,
		"case":        true,
		"chan":        true,
		"const":       true,
		"continue":    true,
		"default":     true,
		"defer":       true,
		"else":        true,
		"fallthrough": true,
		"for":         true,
		"func":        true,
		"go":          true,
		"goto":        true,
		"if":          true,
		"import":      true,
		"interface":   true,
		"map":         true,
		"package":     true,
		"range":       true,
		"return":      true,
		"select":      true,
		"struct":      true,
		"switch":      true,
		"type":        true,
		"var":         true,
	}
	result := make([]*Field, 0)
	for rows.Next() {
		info := ColumnInfo{}
		if err := rows.StructScan(&info); err != nil {
			panic(err)
		}
		typ := strings.ToLower(info.Type)
		if bracketIdx := strings.Index(typ, "("); bracketIdx != -1 {
			typ = typ[0:bracketIdx]
		}

		zeroValue := ""
		if strings.HasSuffix(typ, "char") || strings.HasSuffix(typ, "blob") || strings.HasSuffix(typ, "text") {
			typ = "string"
			zeroValue = `""`
		} else if strings.HasSuffix(typ, "int") || typ == "integer" {
			typ = "int"
			zeroValue = "0"
		} else {
			typ = typeMap[typ]
		}

		if strings.HasPrefix(typ, "int") || strings.HasPrefix(typ, "float") {
			zeroValue = "0"
		} else if typ == "string" {
			zeroValue = `""`
		} else if typ == "time.Time" {
			zeroValue = "time.Time{}"
		}

		name := info.Field
		if isKeyword := keywordMap[name]; isKeyword {
			name = "k_" + name
		}
		name = strcase.ToLowerCamel(name)

		result = append(result, &Field{
			ParamName:     name,
			ParamListName: strcase.ToLowerCamel(info.Field + "_list"),
			SQLName:       fmt.Sprintf("`%s`", info.Field),
			StructName:    strcase.ToCamel(info.Field),
			Type:          typ,
			Tag:           fmt.Sprintf(`db:"%s"`, info.Field),
			ZeroValue:     zeroValue,
		})
	}

	return result
}
