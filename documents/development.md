# :wrench: ブログの開発をする
## :rocket: 環境構築手順
### ブログを開発する

[投稿者用の手順](/documents/posting.md)と、[Makefile](/Makefile) を参照してください。

### 投稿内容検証用 Go のテストを実行する

```
go test ./scripts/posts_test.go
```

### その他の Go のテストを実行する

```
go test ./scripts/test_utils/...
```
