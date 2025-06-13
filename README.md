# 📦 Project Structure Overview

このプロジェクトは Go (Gin) をベースにした REST API バックエンドです。各ディレクトリの役割は以下の通りです。

---

## 🗂️ ディレクトリ構成と説明(アーキテクチャ)

```
├── controller/ # HTTPリクエストを受け取り、serviceを呼び出してレスポンスを返す層
├── service/ # ビジネスロジック（ユースケース）を記述する層
├── middleware/ # 認証やログなど共通処理を定義する層
├── routes/ # エンドポイントとコントローラを紐づけるルーティング定義
├── models/ # ドメインモデル・構造体（User, Articleなど）を定義[マイグレーションファイル的なイメージ]
├── db/ # DB接続やマイグレーション処理を担当
├── go.mod # Go Modules定義ファイル
├── go.sum # Go Modules依存バージョンファイル
└── main.go # エントリーポイント（APIサーバー起動）
```

## 🔧 実例

エンドユーザーからリクエストが来る (/profile/:id)

↓

**middleware**でログインしてるかどうかの認証(ログイン済みの人のみ/profile にアクセスできる)

↓

**routing**で/profile/:id を controller に飛ばす

↓

**controller**で受け取ったら http の窓口だからこの場合は":id"の部分を加工して service を呼び出す

↓

**service**で実際に db にアクセスして JSON を返す
