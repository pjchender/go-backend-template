# Jubo Space Backend

## Development

### 建立 database

參考 [database 建立說明](/docs/database.md)

### 啟動伺服器

```bash
$ go run main.go
```

或透過 [cosmtrek/air](https://github.com/cosmtrek/air) 進行開發：

```bash
$ go get -u github.com/cosmtrek/air
$ air
```

如果是第一次安裝 Golang，記得要在 `.bashrc` 或 `.zshrc` 中加入將 `$GOPATH/bin` 加入 `PATH` 中：

```
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```