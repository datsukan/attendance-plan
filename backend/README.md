# attendance plan backend

受講計画ツールのバックエンドのアプリケーション実装。

## Usage

### GOPRIVATE

`go get` や `go mod tidy` を実行する場合は以下の通り環境変数を設定する。

```sh
export GOPRIVATE=github.com/datsukan
// or
export GOPRIVATE=github.com/datsukan/attendance-plan
```

### SAM

```sh
sam validate
sam build
sam local start-api --env-vars env.json --docker-network backend_default
```
