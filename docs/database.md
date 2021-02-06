# Database for DEV

## 使用 homebrew

安裝 postgreSQL

```bash
$ brew install postgresql
$ brew services start postgresql
```

## 建立使用者

```bash
$ psql postgres  # 登入 DB
```

透過 SQL 建立名為 `postgres` 的 superuser（預設的 superuser 會是系統的使用者）：

```postgresql
-- create superuser of "postgres"
CREATE ROLE postgres LOGIN SUPERUSER CREATEDB CREATEROLE REPLICATION BYPASSRLS;
GRANT
    ALL
    ON ALL TABLES IN SCHEMA "public" TO postgres;
```

## 建立 Database

```bash
$ createdb jubo_space -O postgres -E utf8
```

## Migration

使用 GORM 的 autoMigrate，當有新建的 models 時，需於 `internal/database/database.go` 的 `d.DB.AutoMigrate()`
新增對應的 model 名稱

## 刪除 Database 中的所有資料

```bash
$ make db-drop
```

當有新建的 models 時，需於 `internal/database/database.go` 的 `d.DB.Migrator().DropTable()` 中新增對應的 model 名稱
