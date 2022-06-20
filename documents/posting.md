# :writing_hand: ブログ記事を投稿する
## :tipping_hand_woman: 全体の流れ

投稿から公開までの流れは、大まかには以下のようになります。

1. 投稿者が、記事を Markdown ファイルとして作成する。
2. 投稿者が、本リポジトリに対して、Pull-Request でその Markdown ファイルを提出する。
3. 運営者が、記事をレビューする。運営者から修正の依頼があれば、投稿者は対応する。
4. 運営者が、記事を公開する。

## :rocket: 執筆環境を構築する
### PC へインストールが必要なソフトウェア

[Docker](https://www.docker.com/) のインストールをしてください。  
実際に使うソフトウェアは [Hugo](https://gohugo.io/) というブログエンジンですが、それを Docker 経由で動かします。

### アプリケーションのインストール

```
git clone git@github.com:kjirou/hugo_operation_sample.git
```

## :black_nib: 記事を執筆する
### 記事の雛形を作る

```
cd /path/to/hugo_operation_sample
make hugo_new
```

この後に入力する記事のファイル名が、公開 URL の Slug になります。

### プレビュー環境を起動する

```
cd /path/to/hugo_operation_sample
make hugo_server
```

起動したら、 http://localhost:1313/ をブラウザで開いてください。

### Front Matter を設定する

まずは、以下の Front Matter 部分の設定をしてください。

```
---
title: "Hello World"
date: 2022-01-01 09:00:00
authors:
  - 投稿者名
---
```

- `title`
  - いわゆるタイトルです。公開時に `<title>` や `<h1>` などへ展開されます。
  - 文字数の上限は特にありません。
- `date`
  - 公開日時として表示されます。
  - 提出時は、初期値のままで良いです。
  - 投稿者からの要求がなければ、最終的な値は運営者により設定されます。
- `authors`
  - 投稿者名として表示されます。実名でなくても構いません。
  - 匿名を希望するときは、`チームで決めたデフォルトの名前` と設定してください。
  - 1 つのみ設定できます。

### 本文を書く

- Markdown の記法は、基本的には [GitHub Flavored Markdown](https://github.github.com/gfm/) に準拠しています。
- 画像を置くときの構文は、`![{任意でalt属性}](/external/posts/{記事の投稿年}/{記事のファイル名}/{自由な画像名}.{拡張子})` です。
- 外部コンテンツを埋め込むときの構文は、`{{< コンテンツに相当するShortcodesの名前 引数1 引数2 ... >}}` です。
  - 対応しているコンテンツの種類は、[Hugo のビルトインに含まれているもの](https://gohugo.io/content-management/shortcodes/) と [独自に対応しているもの](/layouts/shortcodes) です。
  - 引数については、Shortcode の元になっている各コンテンツを埋め込むための HTML コードを見て、推測してください。
    - なお、独自対応に含まれている Speaker Deck は、普通に URL に含まれている識別子っぽい部分ではなく、Embed をするためのコードに含まれている `data-id` を使います。
- レビューの指摘を減らすために、本文を書く前に [内容についてのガイドライン](/documents/content-guidelines.md) を一読願います。

## :mailbox_with_mail: 記事を提出する

- Pull-Request の各属性値を以下のように設定してください。
  - タイトルへ `title` と同じ文章を設定する。
  - Assignees へ自分を設定する。
  - Open を設定する。
    - Draft で Pull-Request を作成することも可能です。そのときは、運営者は未完成として認識します。
- CI で、記事の内容についての定型的な検証を行っています。
  - もし、`validate-posts` の Workflow が失敗していたら、CI の出力を確認して修正してください。
  - その他の Workflow が失敗していたら、不慮のエラーの可能性があるため、開発者までご一報ください。
- 提出後は、Reviewers へ登録されている運営者たちによるレビューが行われます。
  - 一般的な Pull-Request の作法に準拠して、指摘された点を解決してください。
  - 指摘された点を全て解決するか、特に指摘がないときは誰かひとりから Approve をもらえば完了です。
- 基本的に、rebase や force-push は禁止です。運営者側のレビューが行いにくくなるためです。実行するときは、事前に運営者へ確認してください。

## :soon: その後

- 運営者によって記事の公開作業が行われます。公開されたときに、チャットで周知されます。
- もちろん、公開後の修正も可能です。そのときも、一般的な Pull-Request の作法に準拠して修正を行ってください。

投稿いただき、ありがとうございました！
