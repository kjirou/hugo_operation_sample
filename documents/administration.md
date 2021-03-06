# :computer: ブログを運営する
## :tipping_hand_woman: 全体の流れ

前提知識として、[投稿者用のドキュメント](/documents/posting.md) を一読願います。

本ドキュメントへは、それに加えて必要になる、運営者用の情報のみを記載します。

## :eyes: 記事をレビューする
### レビューを開始する

Open 状態の Pull-Request が、レビューを開始できる対象です。

### コミュニケーションの方法

本システム独自のルールはないので、一般的な GitHub の操作・慣習に従ってください。

### CI が失敗しているときの対応

CI が失敗しているときは、その Pull-Request はマージしないでください。

「Checks」のタブを開き、 `Validate Posts` の Workflow が失敗していたときは、おそらくは記事データに不備があります。  
そのときは、まずは投稿者へ記事の見直しを要求してください。

上記ではないところが失敗していたら、開発チャンネルまで報告をお願いします。

## :+1: 記事を Approve する

Pull-Request の内容に問題がなくなったら、投稿者へそれを明示的に伝えるために、PR の状態を Approve へ変更してください。

## :shipit: 記事を公開する
### 必ず `date` 属性値を設定する

公開日時を意味する `date` 属性値は、実際に公開作業を行う運営者が設定してください。  
ただし、常に UTC として扱われるため注意が必要です。以下、具体的に `date` の影響を解説します。

- 日付へと加工され、記事の個別または一覧画面で出力されます。
  - このとき、タイムゾーンの影響は無視できます。例えば、`2021-02-09 00:00:00` と `2021-02-09 23:59:59` は両方とも、`2021年2月9日` という出力になります。
  - `date` が UTC な一方で、Hugo の出力処理も UTC で両者が一致するためです。結果的に無視できます。
- 記事の整列順を決める際に参照されます。
  - トップ画面などの記事を一覧する箇所では、`date` の降順に記事が整列します。
  - このときも、相対的な比較しかされないため、タイムゾーンの影響は無視できます。
- 公開されるかの判定で参照されます。
  - 未来日時が設定されている記事は公開されません。
  - このときは、タイムゾーンの考慮が必要です。例えば、JST の時刻を見ているなら、それより `-09:00` 以前の時刻を `date` へ設定しないと、記事が公開されません。

### 本番環境へデプロイする

ここへ本番環境へデプロイする時の操作を記述すると良さそうです。  
どこかへ `make hugo_build` の結果を配信することを想定しています。
