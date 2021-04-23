GoModel 是一个生成模型与CRUD相关代码的工具,目前仅支持 MySQL

# 安装

```bash
$ go install github.com/metauro/gomodel/cmd/gomodel@latest
```

# 配置

在`$HOME/.gomodel.toml`下配置以下内容

```toml
[mysql]
username = "root"
password = "root"
host = "localhost"
port = 3306
database = "db"
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
$ gomodel gen
```

# 使用

生成完毕后即可使用

```go
package main

import (
	"github.com/metauro/gomodel"
)

func main() {
	db, err := gomodel.Open("mysql", "dsn")
	if err != nil {
		panic(err)
	}
	repo := model.NewTestRepo(db)
	// 插入单条数据
	repo.Insert(&model.Test{
		Key: "a",
	}).MustExec()
	// 插入多条数据
	repo.Insert(&model.Test{
		Key: "a",
	}, &model.Test{
		Key: "b",
	}).MustExec()
	// 更新数据
	repo.Update().SetKey("update").WhereKeyEqual("a").MustExec()
	// 查询一条数据
	repo.Select().WhereKeyEqual("a").MustGet()
	// 查询多条数据
	repo.Select().MustSelect()
	// 删除数据
	repo.Delete().WhereKeyEqual("a").Limit(1).MustExec()
}
```
