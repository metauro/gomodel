GoModel 是一个生成模型与CRUD相关代码的工具,目前仅支持 MySQL

# 安装

```bash
$ go install github.com/metauro/gomodel/cmd/gomodel@latest
```

# 创建表

```mysql
CREATE TABLE IF NOT EXISTS `test`
(
    `id`        INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    `create_at` DATETIME    NOT NULL DEFAULT NOW(),
    `update_at` DATETIME    NOT NULL DEFAULT NOW() ON UPDATE NOW(),
    `key`       VARCHAR(64) NOT NULL
);
```

# 生成代码

执行以下命令,选择要生成的表后回车

```bash
$ gomodel gen -dsn "root:root@(localhost:3306)/database" -table test
```

# 使用

生成完毕后即可使用

```go
package main

import (
	"github.com/jmoiron/sqlx"
	"mod/gomodel"
)

func main() {
	sqlDB, _ := sqlx.Open("mysql", "root:root@(localhost:3306)/database")
	db := gomodel.NewDB(sqlDB)
	ctx := context.Background
	
	// 插入数据
	db.Test.Insert().Values(&gomodel.Test{}).Exec(ctx)
	
	// 批量插入数据
	db.Test.Insert().Values(&gomodel.Test{}, &gomodel.Test{}).Exec(ctx)
	
	// 获取单条数据
	db.Test.Select().Get(ctx)
	
	// 获取多条数据
	db.Test.Select().List(ctx)
	
	// 更新数据
	db.Test.Update().SetKey("test_key").Exec(ctx)
	
	// 删除 id=1 的数据
	db.Test.Delete().Where(func(b *gomodel.TestWhereBuilder) {
		b.WhereIdEQ(1)
  }).Exec(ctx)
}
```
