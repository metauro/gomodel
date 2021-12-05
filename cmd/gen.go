package cmd

import (
	"bytes"
	_ "embed"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/iancoleman/strcase"
	"github.com/metauro/gomodel/internal/msql"
	"github.com/metauro/gomodel/internal/template"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Model struct {
	Pkg             string
	TitleName       string
	CamelName       string
	SnakeName       string
	SQLName         string
	Placeholder     string
	SoftDeleteField *Field
	Fields          []*Field
}

type Field struct {
	TitleName string
	CamelName string
	SnakeName string
	SQLName   string
	SQLType   string
	// GoType eg: string,null.String
	GoType string
	// GoRawType eg: string, int
	GoRawType string
	Tag       string
	ZeroValue string
	Nullable  bool
}

type GenOptions struct {
	DriveName        string
	Dsn              string
	Output           string
	SoftDeleteColumn string
	Tag              string
	TimeString       bool
	Tables           []string
}

var genOpts = &GenOptions{}

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate code from tables",
	Run: func(cmd *cobra.Command, args []string) {
		pkg := strcase.ToSnake(filepath.Base(genOpts.Output))
		db, err := msql.Open(genOpts.DriveName, genOpts.Dsn)
		if err != nil {
			panic(err)
		}

		if len(genOpts.Tables) == 0 {
			genOpts.Tables, err = getTables(db)
			if err != nil {
				panic(err)
			}
		}

		if len(genOpts.Tables) == 0 {
			log.Printf("no tables need to generate\n")
			return
		}

		log.Printf("start generate tables: %v\n", genOpts.Tables)

		wd, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		getPath := func(path string) string {
			return filepath.Join(wd, genOpts.Output, path)
		}

		templateMap := map[string]string{
			"constants": template.TemplateConstants,
			"db":        template.TemplateDB,
			"delete":    template.TemplateDelete,
			"event":     template.TemplateEvent,
			"insert":    template.TemplateInsert,
			"model":     template.TemplateModel,
			"order":     template.TemplateOrder,
			"select":    template.TemplateSelect,
			"update":    template.TemplateUpdate,
			"where":     template.TemplateWhere,
		}
		genOnceTemplateMap := map[string]string{
			"db":      template.TemplateDBAll,
			"runtime": template.TemplateRuntime,
			"hook":    template.TemplateHook,
		}

		models := make([]*Model, 0, len(genOpts.Tables))
		if err := os.MkdirAll(getPath(""), 0755); err != nil {
			panic(err)
		}

		for _, table := range genOpts.Tables {
			name := table
			name = strcase.ToCamel(name)
			fields, err := getFieldsFromTable(db, table)
			if err != nil {
				panic(err)
			}
			var softDeleteField *Field
			for _, field := range fields {
				if field.SnakeName == genOpts.SoftDeleteColumn {
					softDeleteField = field
				}
			}

			model := &Model{
				Pkg:             pkg,
				CamelName:       strcase.ToLowerCamel(name),
				TitleName:       strcase.ToCamel(name),
				SnakeName:       table,
				SQLName:         fmt.Sprintf("`%s`", table),
				Fields:          fields,
				Placeholder:     "?",
				SoftDeleteField: softDeleteField,
			}
			models = append(models, model)

			for name, content := range templateMap {
				fmt.Printf("gen %v template\n", name)
				err := template.WriteFile(
					template.TemplatePredefine+content,
					getPath(fmt.Sprintf("%s_%s.go", model.SnakeName, name)),
					model,
				)
				if err != nil {
					panic(err)
				}
			}
			log.Printf("generate %s success\n", table)
		}

		for name, content := range genOnceTemplateMap {
			fmt.Printf("gen %v template\n", name)
			err := template.WriteFile(
				content,
				getPath(fmt.Sprintf("%s.go", name)),
				models,
			)
			if err != nil {
				panic(err)
			}
		}

		log.Println("start format code")
		var errBuf bytes.Buffer
		fmtCmd := exec.Command("go", "fmt", "./...")
		fmtCmd.Stderr = &errBuf
		if err := fmtCmd.Run(); err != nil {
			panic(errBuf.String())
		}
		log.Println("format code success")
		log.Println("generate done")
	},
}

func init() {
	rootCmd.AddCommand(genCmd)
	genCmd.Flags().StringVarP(&genOpts.Output, "output", "o", "model", "Output folder")
	genCmd.Flags().StringVarP(
		&genOpts.SoftDeleteColumn,
		"soft-delete-column",
		"d",
		"",
		"Enable soft deletion by updated column timestamp, the column type should be one of DATE, DATETIME, TIMESTAMP",
	)
	genCmd.Flags().StringVarP(
		&genOpts.Tag,
		"tag",
		"t",
		`json:"{{.SnakeName}}" db:"{{.SnakeName}}"`,
		`Struct field tag template, Available Tags: 
TitleName - title case field name, eg: TestField
CamelName - camel case field name, eg: testField
SnakeName - snake case field name, eg: test_field 
SQLName   - escaped field name, eg: "test_field"  
GoType    - field golang type, string, int, time.Time ...
SQLType   - field sql type, varchar, char, int ....
Len 	  - field sql type len 
Unsigned  -  true if field type is unsigned numeric   
ZeroValue - field golang zero value
`,
	)
	genCmd.Flags().BoolVar(&genOpts.TimeString, "time-string", false, "Use string instead time.Time")
	genCmd.Flags().StringVar(&genOpts.Dsn, "dsn", "", "eg: root:root@(localhost:3306)/dev")
	genCmd.Flags().StringVar(&genOpts.DriveName, "drive-name", "mysql", "eg: mysql,postgres")
	genCmd.Flags().StringSliceVar(&genOpts.Tables, "table", nil, "eg: table_name")
}

func getTables(db *msql.DB) ([]string, error) {
	tables, err := db.GetTables()
	if err != nil {
		return nil, err
	}

	prompt := &survey.MultiSelect{
		Message:  "please select what tables you need to generate",
		Options:  tables,
		PageSize: 10,
	}

	tables = make([]string, 0)
	_ = survey.AskOne(prompt, &tables, survey.WithPageSize(10))
	return tables, nil
}

func getFieldsFromTable(db *msql.DB, table string) ([]*Field, error) {
	typeMap := map[string]string{
		// numeric type
		"tinyint":   "int8",
		"smallint":  "int16",
		"mediumint": "int32",
		"int":       "int",
		"bigint":    "int64",
		"float":     "float32",
		"double":    "float64",
		"decimal":   "float64",
		// date type
		"date":      "time.Time",
		"datetime":  "time.Time",
		"timestamp": "time.Time",
		"time":      "time.Time",
		"year":      "time.Time",
		// string type
		"char":       "string",
		"varchar":    "string",
		"binary":     "string",
		"varbinary":  "string",
		"tinyblob":   "string",
		"tinytext":   "string",
		"blob":       "string",
		"text":       "string",
		"mediumblob": "string",
		"mediumtext": "string",
		"longblob":   "string",
		"longtext":   "string",
		"enum":       "string",
		"set":        "string",
		"json":       "string",
	}

	nullableTypeMap := map[string]string{
		// numeric type
		"tinyint":   "null.Int8",
		"smallint":  "null.Int16",
		"mediumint": "null.Int32",
		"int":       "null.Int",
		"bigint":    "null.Int64",
		"float":     "null.Float32",
		"double":    "null.Float64",
		"decimal":   "null.Float64",
		// date type
		"date":      "null.Time",
		"datetime":  "null.Time",
		"timestamp": "null.Time",
		"time":      "null.Time",
		"year":      "null.Time",
		// string type
		"char":       "null.String",
		"varchar":    "null.String",
		"binary":     "null.String",
		"varbinary":  "null.String",
		"tinyblob":   "null.String",
		"tinytext":   "null.String",
		"blob":       "null.String",
		"text":       "null.String",
		"mediumblob": "null.String",
		"mediumtext": "null.String",
		"longblob":   "null.String",
		"longtext":   "null.String",
		"enum":       "null.String",
		"set":        "null.String",
		"json":       "null.String",
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

	columns, err := db.GetColumns(table)
	if err != nil {
		return nil, err
	}

	result := make([]*Field, 0)
	for _, column := range columns {
		var goType string
		var goRawType string
		var zeroValue string

		if column.Type == "tinyint" && column.Len == 1 {
			goType = "bool"
		} else {
			goType = typeMap[column.Type]
		}

		// 处理无符号类型
		if column.Unsigned && !strings.HasPrefix(goType, "float") {
			goType = "u" + goType
		}

		if goType == "time.Time" && genOpts.TimeString {
			goType = "string"
		}

		if goType == "string" {
			zeroValue = `""`
		} else if goType == "bool" {
			zeroValue = "false"
		} else if goType == "time.Time" {
			zeroValue = "time.Time{}"
		} else if strings.HasPrefix(goType, "int") ||
			strings.HasPrefix(goType, "uint") ||
			strings.HasPrefix(goType, "float") {
			zeroValue = "0"
		}
		goRawType = goType

		if column.Nullable {
			if column.Type == "tinyint" && column.Len == 1 {
				goType = "null.Bool"
			} else {
				goType = nullableTypeMap[column.Type]
			}

			if column.Unsigned && !strings.HasPrefix(goType, "null.Float") {
				goType = "null.U" + strings.ToLower(goType[5:])
			}

			if goType == "null.Time" && genOpts.TimeString {
				goType = "null.String"
			}
		}

		fieldName := column.Field
		// sql 字段名与 go 关键字冲突时,添加前缀
		if isKeyword := keywordMap[fieldName]; isKeyword {
			fieldName = "k_" + fieldName
		}

		field := &Field{
			TitleName: strcase.ToCamel(column.Field),
			CamelName: strcase.ToLowerCamel(fieldName),
			SnakeName: column.Field,
			SQLName:   fmt.Sprintf("`%s`", column.Field),
			GoType:    goType,
			GoRawType: goRawType,
			SQLType:   column.Type,
			ZeroValue: zeroValue,
			Nullable:  column.Nullable,
		}
		if genOpts.Tag != "" {
			var tag bytes.Buffer
			err := template.Execute(genOpts.Tag, &tag, field)
			if err != nil {
				return nil, err
			}
			field.Tag = tag.String()
		}
		result = append(result, field)
	}

	return result, nil
}
