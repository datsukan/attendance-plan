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

### DB

#### Up

```sh
make up
```

#### Down

```sh
make down
```

### Test

```sh
make test
```

### SAM

#### 形式チェック

```sh
sam validate
```

#### ローカルサーバーの起動

```sh
make dev
```

#### 本番デプロイ

```sh
make deploy
```

初回はデプロイ後にAPI Gateway のエンドポイントとカスタムドメインのマッピングとしてDNSにCNAMEレコードの追加が必要
