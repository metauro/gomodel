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
	"github.com/spf13/viper"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
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
	SQLType   string
	Type      string
	Tag       string
	ZeroValue string
	Nullable  bool
}

type GenConfig struct {
	DriveName        string
	Dsn              string
	Prefix           string
	AppendPrefix     bool
	SelectAllTable   bool
	Output           string
	SoftDeleteColumn string
	Tag              string
	TimeString       bool
}

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate code from tables",
	Run: func(cmd *cobra.Command, args []string) {
		c := &GenConfig{
			DriveName:        viper.GetString("drive-name"),
			Dsn:              viper.GetString("dsn"),
			Prefix:           viper.GetString("prefix"),
			AppendPrefix:     viper.GetBool("append-prefix"),
			SelectAllTable:   viper.GetBool("all"),
			Output:           viper.GetString("output"),
			SoftDeleteColumn: viper.GetString("soft-delete-column"),
			Tag:              viper.GetString("tag"),
			TimeString:       viper.GetBool("time-string"),
		}
		pkg := strcase.ToSnake(filepath.Base(c.Output))

		db, err := msql.Open(c.DriveName, c.Dsn)
		if err != nil {
			panic(err)
		}

		tables, err := getTables(db, c)
		if err != nil {
			panic(err)
		}

		if len(tables) == 0 {
			log.Printf("no tables need to generate\n")
			return
		}

		log.Printf("start generate tables: %v\n", tables)

		wd, err := os.Getwd()
		if err != nil {
			panic(err)
		}

		var templateMap = map[string]string{
			"constants": template.TemplateConstants,
			"db":        template.TemplateDB,
			"delete":    template.TemplateDelete,
			"insert":    template.TemplateInsert,
			"model":     template.TemplateModel,
			"order":     template.TemplateOrder,
			"runtime":   template.TemplateRuntime,
			"select":    template.TemplateSelect,
			"update":    template.TemplateUpdate,
			"where":     template.TemplateWhere,
		}

		models := make([]*Model, 0, len(tables))

		for _, table := range tables {
			name := table
			if !c.AppendPrefix && c.Prefix != "" && strings.HasPrefix(name, c.Prefix) {
				name = name[len(c.Prefix):]
			}
			if err := os.MkdirAll(c.Output, 0755); err != nil {
				panic(err)
			}

			name = strcase.ToCamel(name)
			fields, err := getFieldsFromTable(db, table, c)
			if err != nil {
				panic(err)
			}
			softDeleteColumn := ""
			for _, field := range fields {
				if field.SnakeName == c.SoftDeleteColumn {
					softDeleteColumn = field.SQLName
				}
			}

			model := &Model{
				Pkg:              pkg,
				CamelName:        strcase.ToLowerCamel(name),
				TitleName:        strcase.ToCamel(name),
				SnakeName:        table,
				SQLName:          fmt.Sprintf("`%s`", table),
				Fields:           fields,
				Placeholder:      "?",
				SoftDeleteColumn: softDeleteColumn,
			}
			models = append(models, model)

			for name, content := range templateMap {
				fmt.Printf("gen %v template\n", name)
				file, err := os.OpenFile(
					filepath.Join(wd, c.Output, fmt.Sprintf("%s_%s.go", model.SnakeName, name)),
					os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm,
				)
				if err != nil {
					panic(err)
				}
				err = template.Execute(template.TemplateVariable+content, file, model)
				if err != nil {
					panic(err)
				}
			}
			log.Printf("generate %s success\n", table)
		}

		file, err := os.OpenFile(
			filepath.Join(wd, c.Output, "db.go"),
			os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm,
		)
		err = template.Execute(template.TemplateDBAll, file, models)
		if err != nil {
			panic(err)
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
	genCmd.Flags().StringP("prefix", "p", "", "Select the table that contains the prefix")
	genCmd.Flags().Bool("append-prefix", false, "Append prefix to struct name (default false)")
	genCmd.Flags().BoolP("all", "a", false, "Select all table to generate")
	genCmd.Flags().StringP("output", "o", "model", "Output folder")
	genCmd.Flags().StringP("soft-delete-column", "d", "", "Enable soft deletion by updated column timestamp, "+
		"the column type should be one of DATE, DATETIME, TIMESTAMP")
	genCmd.Flags().StringP("tag", "t", `json:"{{.SnakeName}}" db:"{{.SnakeName}}"`,
		`Struct field tag template, Available Tags: 
	TitleName - title case field name, eg: TestField
	CamelName - camel case field name, eg: testField
	SnakeName - snake case field name, eg: test_field 
	SQLName   - escaped field name, eg: "test_field"  
	Type      - field golang type, string, int, time.Time etc.
	SQLType   - field sql type, varchar(11), char(11), int(11) etc.
	SQLTypeNoArgs - field sql type, no 
	Unsigned -  true if field type is unsigned numeric   
	ZeroValue - field golang zero value
`)
	genCmd.Flags().Bool("time-string", false, "Use string instead time.Time")
	genCmd.Flags().String("dsn", "", "eg: root:root@(localhost:3306)/test?parseTime=true")
	genCmd.Flags().String("drive-name", "mysql", "eg: mysql,postgres")
	cobra.CheckErr(viper.BindPFlags(genCmd.Flags()))
}

func getTables(db *msql.DB, c *GenConfig) ([]string, error) {
	tables, err := db.GetTables()
	if err != nil {
		return nil, err
	}

	// 生成全部表
	if c.SelectAllTable {
		return tables, nil
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
		return tables, nil
	}

	// 按前缀筛选
	result := make([]string, 0)
	for _, table := range tables {
		if !strings.HasPrefix(table, c.Prefix) {
			continue
		}

		result = append(result, table)
	}

	return result, nil
}

func getFieldsFromTable(db *msql.DB, table string, c *GenConfig) ([]*Field, error) {
	typeMap := map[string]string{
		// numeric type
		"tinyint":   "int8",
		"smallint":  "int16",
		"mediumint": "int32",
		"int":       "int32",
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

	// TODO: null 和 zeroValue  使用 https://github.com/guregu/null 处理
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
		// sql 类型,如 "tinyint(1)"
		sqlType := strings.ToLower(column.Type)
		// sql 类型,去掉括号内容的,如 "tinyint"
		sqlTypeNoArgs := sqlType
		typ := typeMap[sqlType]
		zeroValue := ""
		unsigned := strings.Contains(sqlType, "unsigned")

		if bracketIdx := strings.Index(sqlType, "("); bracketIdx != -1 {
			sqlTypeNoArgs = sqlType[0:bracketIdx]
		}

		if spaceIdx := strings.Index(sqlTypeNoArgs, " "); spaceIdx != -1 {
			sqlTypeNoArgs = sqlTypeNoArgs[0:spaceIdx]
		}

		if typ == "" {
			if strings.HasPrefix(sqlType, "tinyint(1)") {
				typ = "bool"
			} else {
				typ = typeMap[sqlTypeNoArgs]
			}

			if typ == "" {
				typ = "interface{}"
			}
		}

		fmt.Printf("sqltype: %v, %v, typ: %v \n", sqlType, sqlTypeNoArgs, typ)

		// 处理无符号类型
		if strings.HasPrefix(typ, "int") && unsigned {
			typ = "u" + typ
		}

		if typ == "time.Time" && c.TimeString {
			typ = "string"
		}

		if typ == "string" {
			zeroValue = `""`
		} else if typ == "bool" {
			zeroValue = "false"
		} else if typ == "time.Time" {
			zeroValue = "time.Time{}"
		} else if strings.HasPrefix(typ, "int") ||
			strings.HasPrefix(typ, "uint") ||
			strings.HasPrefix(typ, "float") {
			zeroValue = "0"
		} else {
			zeroValue = "nil"
		}

		name := column.Field
		// sql 字段名与 go 关键字冲突时,添加前缀
		if isKeyword := keywordMap[name]; isKeyword {
			name = "k_" + name
		}

		field := &Field{
			TitleName: strcase.ToCamel(column.Field),
			CamelName: strcase.ToLowerCamel(name),
			SnakeName: column.Field,
			SQLName:   fmt.Sprintf("`%s`", column.Field),
			Type:      typ,
			SQLType:   sqlType,
			ZeroValue: zeroValue,
			Nullable:  column.Null == "YES",
		}
		if c.Tag != "" {
			var tag bytes.Buffer
			err := template.Execute(c.Tag, &tag, field)
			if err != nil {
				return nil, err
			}
			field.Tag = tag.String()
		}
		result = append(result, field)
	}

	return result, nil
}
