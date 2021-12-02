CREATE TABLE IF NOT EXISTS `gomodel`(
    `tinyint`    TINYINT NOT NULL,
    `smallint`   SMALLINT,
    `mediumint`  MEDIUMINT,
    `int`        INT,
    `bigint`     BIGINT,
    `float`      FLOAT,
    `double`     DOUBLE,
    `decimal`    DECIMAL(10, 2),

    `utinyint`   TINYINT UNSIGNED,
    `usmallint`  SMALLINT UNSIGNED,
    `umediumint` MEDIUMINT UNSIGNED,
    `uint`       INT UNSIGNED,
    `ubigint`    BIGINT UNSIGNED,
    `ufloat`     FLOAT UNSIGNED,
    `udouble`    DOUBLE UNSIGNED,
    `udecimal`   DECIMAL(10, 2) UNSIGNED,

    `date`       DATE,
    `datetime`   DATETIME,
    `timestamp`  TIMESTAMP,
    `time`       TIME,
    `year`       YEAR,

    `char`       CHAR(11),
    `varchar`    VARCHAR(11),
    `binary`     BINARY,
    `varbinary`  VARBINARY(10),
    `tinyblob`   TINYBLOB,
    `tinytext`   TINYTEXT,
    `blob`       BLOB,
    `text`       TEXT,
    `mediumblob` MEDIUMBLOB,
    `mediumtext` MEDIUMTEXT,
    `longblob`   LONGBLOB,
    `longtext`   LONGTEXT,
    `enum`       ENUM ('v1', 'v2'),
    `set`        SET ('v1', 'v2'),
    `json`       JSON,

    `tinybool`   TINYINT(1),
    `bool`       BOOL
);
