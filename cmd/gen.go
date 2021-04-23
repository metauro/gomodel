package cmd

import (
	"bytes"
	_ "embed"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/iancoleman/strcase"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/exec"
	"strings"
	"text/template"
)

type Model struct {
	Pkg              string
	TitleName        string
	CamelName        string
	SnakeName        string
	SQLName          string
	Placeholder      string
	SoftDeleteColumn string
	Fields           []*Field
}

type Field struct {
	TitleName string
	CamelName string
	SnakeName string
	SQLName   string
	Type      string
	Tag       string
	ZeroValue string
}

type GenConfig struct {
	Prefix           string
	AppendPrefix     bool
	SelectAllTable   bool
	Pkg              string
	Output           string
	SoftDeleteColumn string
	Tag              string
}

//go:embed model.tmpl
var modelTemplate string

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate code from tables",
	Run: func(cmd *cobra.Command, args []string) {
		c := &GenConfig{
			Prefix:           viper.GetString("prefix"),
			AppendPrefix:     viper.GetBool("append-prefix"),
			SelectAllTable:   viper.GetBool("all"),
			Pkg:              viper.GetString("pkg"),
			Output:           viper.GetString("output"),
			SoftDeleteColumn: viper.GetString("soft-delete-column"),
			Tag:              viper.GetString("tag"),
		}
		tables := getTables(c)

		if len(tables) == 0 {
			log.Printf("no tables need to generate")
			return
		}

		log.Printf("start generate %+v\n", tables)
		escapeChar := "`"

		tmpl := template.Must(template.New("").Funcs(template.FuncMap{
			"snakeCase": func(s string) string {
				return strcase.ToSnake(s)
			},
			"escape": func(s string) string {
				return fmt.Sprintf("%s%s%s", escapeChar, s, escapeChar)
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
		}).Parse(modelTemplate))

		for _, table := range tables {
			name := table
			if !c.AppendPrefix && c.Prefix != "" && strings.HasPrefix(name, c.Prefix) {
				name = name[len(c.Prefix):]
			}
			if err := os.MkdirAll(c.Output, 0755); err != nil {
				panic(err)
			}
			file, err := os.OpenFile(fmt.Sprintf("%s/%s.gen.go", c.Pkg, name), os.O_CREATE|os.O_WRONLY|os.O_TRUNC,
				os.ModePerm)
			if err != nil {
				panic(err)
			}
			name = strcase.ToCamel(name)
			fields := getFieldsFromTable(table, c)
			softDeleteColumn := ""

			for _, field := range fields {
				if field.SnakeName == c.SoftDeleteColumn {
					softDeleteColumn = field.SQLName
				}
			}

			if err := tmpl.Execute(file, &Model{
				Pkg:              c.Pkg,
				CamelName:        strcase.ToLowerCamel(name),
				TitleName:        strcase.ToCamel(name),
				SnakeName:        table,
				SQLName:          fmt.Sprintf("`%s`", table),
				Fields:           fields,
				Placeholder:      "?",
				SoftDeleteColumn: softDeleteColumn,
			}); err != nil {
				panic(err)
			}
			log.Printf("generate %s success\n", table)
		}

		log.Println("start format code")
		if err := exec.Command("go", "fmt", "./...").Run(); err != nil {
			panic(err)
		}
		log.Println("format code success")
		log.Println("generate done")
	},
}

func init() {
	rootCmd.AddCommand(genCmd)
	genCmd.Flags().StringP("prefix", "p", "", "Select the table that contains the prefix")
	genCmd.Flags().Bool("append-prefix", false, "Append prefix to struct name (default false)")
	genCmd.Flags().String("pkg", "model", "Generate package name")
	genCmd.Flags().BoolP("all", "a", false, "Select all table to generate")
	genCmd.Flags().StringP("output", "o", "model", "Output folder")
	genCmd.Flags().StringP("soft-delete-column", "d", "", "Enable soft deletion by updated column timestamp, "+
		"the column type should be one of DATE, DATETIME, TIMESTAMP")
	genCmd.Flags().StringP("tag", "t", `json:"{{.SnakeName}}" db:"{{.SnakeName}}"`, "Struct field tag template")
	cobra.CheckErr(viper.BindPFlags(genCmd.Flags()))
}

func getTables(c *GenConfig) []string {
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

	// 生成全部表
	if c.SelectAllTable {
		return tables
	}

	// 如果没有指定前缀,则用户手动选择要生成的表
	if c.Prefix == "" {
		prompt := &survey.MultiSelect{
			Message:  "please select what tables you need to generate",
			Options:  tables,
			PageSize: 10,
		}

		tables = make([]string, 0)
		_ = survey.AskOne(prompt, &tables, survey.WithPageSize(10))
		return tables
	}

	// 按前缀筛选
	result := make([]string, 0)
	for _, table := range tables {
		if !strings.HasPrefix(table, c.Prefix) {
			continue
		}

		result = append(result, table)
	}

	return result
}

func getFieldsFromTable(table string, c *GenConfig) []*Field {
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

	tmpl := template.Must(template.New("").Parse(c.Tag))
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

		field := &Field{
			TitleName: strcase.ToCamel(info.Field),
			CamelName: strcase.ToLowerCamel(name),
			SnakeName: info.Field,
			SQLName:   fmt.Sprintf("`%s`", info.Field),
			Type:      typ,
			ZeroValue: zeroValue,
		}
		if c.Tag != "" {
			var tag bytes.Buffer
			cobra.CheckErr(tmpl.Execute(&tag, field))
			field.Tag = tag.String()
		}
		result = append(result, field)
	}

	return result
}
