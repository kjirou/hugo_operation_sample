# :wrench: ブログの開発をする
## :rocket: 環境構築手順

- [投稿者用の手順](/documents/posting.md) に含まれる Docker の設定を行ってください。
- Golang のインストールを行ってください。
  - バージョンは、[go.mod](https://github.com/kjirou/hugo_operation_sample/blob/main/go.mod) に記載のある値です。


## :memo: 開発ケース別の操作例
### ブログを開発する

以下で開発サーバを起動し、その後は Hugo と GitHub の作法に則り開発をしてください。

```
make hugo_server
```

最終的に本番に配信される成果物は、以下のコマンドにより生成されます。

```
make hugo_build
```

### 投稿内容検証用 Go のテストを実行する

```
go test ./scripts/posts_test.go
```

### その他の Go のテストを実行する

```
go test ./scripts/test_utils/...
```
