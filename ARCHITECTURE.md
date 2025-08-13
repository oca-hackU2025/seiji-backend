# seiji-backend アーキテクチャ & ログイン実装ドキュメント

## プロジェクト概要
seiji-backendは政治家に関するソーシャルアプリケーションのバックエンドAPIサービスです。Go言語でFirebase認証とPostgreSQLを使用して構築されています。

## システムアーキテクチャ

### 技術スタック
- **言語**: Go
- **フレームワーク**: Gin Web Framework
- **データベース**: PostgreSQL (GORM ORM使用)
- **認証**: Firebase Authentication + JWT
- **デプロイメント**: Render

### アーキテクチャ構成図
```
Client (Frontend)
       ↓ HTTP Request
┌─────────────────────────┐
│     Gin Router          │
│   (CORS Middleware)     │
└─────────┬───────────────┘
          ↓
┌─────────────────────────┐
│    Routes Layer         │
│  - Public Routes (/api) │
│  - Private Routes       │
│    (/api/private)       │
└─────────┬───────────────┘
          ↓
┌─────────────────────────┐
│   Middleware Layer      │
│ - Firebase Auth         │
│ - JWT Auth              │
└─────────┬───────────────┘
          ↓
┌─────────────────────────┐
│   Controller Layer      │
│ - auth_controller.go    │
│ - like_controller.go    │
│ - seijika_controller.go │
│ - test_controller.go    │
└─────────┬───────────────┘
          ↓
┌─────────────────────────┐
│    Service Layer        │
│ - Business Logic        │
│ - randomSeijika_service │
└─────────┬───────────────┘
          ↓
┌─────────────────────────┐
│     Models Layer        │
│ - GORM Data Models      │
│ - Database Entities     │
└─────────┬───────────────┘
          ↓
┌─────────────────────────┐
│   Database Layer        │
│    PostgreSQL           │
└─────────────────────────┘
```

### ディレクトリ構造
```
seiji-backend/
├── main.go                    # アプリケーションエントリーポイント
├── controller/                # HTTPリクエストハンドラー
│   ├── auth_controller.go     # 認証関連API
│   ├── like_controller.go     # いいね機能API
│   ├── seijika_controller.go  # 政治家関連API
│   └── test_controller.go     # テスト用API
├── db/                        # データベース関連
│   ├── connect.go             # DB接続設定
│   ├── initializer.go         # DB初期化
│   ├── migrate.go             # マイグレーション
│   └── seed.go                # シードデータ
├── lib/                       # 外部サービス統合
│   ├── firebase/              # Firebase関連
│   │   ├── init.go            # Firebase初期化
│   │   └── auth.go            # トークン検証
│   └── jwt/                   # JWT処理
│       └── jwt.go             # JWT生成・検証
├── middleware/                # ミドルウェア
│   └── firebase_auth.go       # 認証ミドルウェア
├── models/                    # データモデル
│   ├── user.go                # ユーザーモデル
│   ├── like.go                # いいねモデル
│   ├── seijika.go             # 政治家モデル
│   └── ...
├── routes/                    # ルーティング定義
│   └── routes.go              # ルート設定
├── service/                   # ビジネスロジック
│   └── randomSeijika_service.go
├── utils/                     # ユーティリティ
│   └── context.go             # コンテキスト処理
└── config/                    # 設定ファイル
    └── firebase-adminsdk.json # Firebase認証情報
```

## ログイン実装詳細

### 認証フロー
```
1. クライアント → Firebase Auth でログイン
2. Firebase ID Token取得
3. バックエンドに ID Token送信 (/api/auth/login)
4. Firebase Admin SDK でトークン検証
5. ユーザー情報をDBに保存/取得
6. JWT AccessToken生成・返却
7. 以降のAPI呼び出しはJWTで認証
```

### 実装コンポーネント

#### 1. ログインエンドポイント (`controller/auth_controller.go:18`)
```go
func Login(c *gin.Context) {
    // 1. Firebase ID Tokenを受け取り
    // 2. Firebase Admin SDKで検証
    // 3. ユーザー作成/取得
    // 4. JWT生成・返却
}
```

**処理フロー:**
1. `LoginRequest`でクライアントからID Tokenを受信
2. `firebase.VerifyIDToken()`でFirebase側で検証
3. `firebase_user_id`でDBからユーザー検索
4. 存在しない場合は新規ユーザー作成
5. `jwt.GenerateToken()`でJWTアクセストークン生成
6. レスポンスで`accessToken`を返却

#### 2. Firebase認証 (`lib/firebase/auth.go:11`)
```go
func VerifyIDToken(idToken string) (*auth.Token, error) {
    // Firebase Admin SDK でID Token検証
    // UIDとクレーム情報を返却
}
```

#### 3. JWT処理 (`lib/jwt/jwt.go`)
**JWT生成** (`GenerateToken:13`)
- ユーザーID、有効期限（24時間）をクレームに設定
- HMAC SHA256で署名
- 環境変数`JWT_SECRET`でシークレット管理

**JWT検証** (`VerifyTokenAndExtractUserID:52`)
- 署名検証
- 有効期限チェック
- ユーザーIDを抽出して返却

#### 4. 認証ミドルウェア (`middleware/firebase_auth.go`)

**Firebase認証ミドルウェア** (`FirebaseAuthMiddleware:13`)
- Firebase ID Token検証用
- 直接Firebase認証を使用する場合

**JWT認証ミドルウェア** (`JWTAuthMiddleware:35`)
- プライベートAPI用の認証
- Bearerトークン形式のJWTを検証
- 検証成功時は`userID`をコンテキストに設定

#### 5. ルート設定 (`routes/routes.go`)
```go
// パブリック（認証不要）
api.POST("/auth/login", controller.Login)

// プライベート（JWT認証必須）
private := r.Group("/api/private")
private.Use(middleware.JWTAuthMiddleware())
{
    private.GET("/politicians/generate", controller.GetRandomSeijika)
    private.POST("/likes", controller.CreateLike)
    private.GET("/likes", controller.GetLikedSeijikaList)
}
```

### データモデル (`models/user.go:5`)
```go
type User struct {
    ID             uint      // プライマリキー
    FirebaseUserID string    // Firebase UID
    CreatedAt      time.Time // 作成日時
    Likes          []Like    // リレーション（いいね履歴）
}
```

### セキュリティ設計
1. **二重認証システム**
   - 初回ログイン: Firebase ID Token
   - API呼び出し: JWT Access Token

2. **トークン管理**
   - JWT有効期限: 24時間
   - シークレットキー: 環境変数で管理

3. **CORS設定** (`main.go:34`)
   - 許可オリジン: `http://localhost:3000`
   - 許可メソッド: GET, POST, PUT, DELETE
   - 認証情報の送信を許可

### 環境変数
- `FIREBASE_CREDENTIALS`: Firebase認証ファイルパス
- `JWT_SECRET`: JWT署名用シークレット
- `DATABASE_URL`: PostgreSQL接続URL
- `GIN_MODE`: 実行モード (development/release)
- `PORT`: サーバーポート番号

### デプロイメント設定
- **開発環境**: Docker Compose使用
- **本番環境**: Render.com
- **設定ファイル**: `render.yaml`
- **環境分離**: `GIN_MODE`で開発/本番を切り替え

### エラーハンドリング
- Firebase認証エラー: 401 Unauthorized
- JWT検証エラー: 401 Unauthorized  
- DB操作エラー: 500 Internal Server Error
- リクエスト形式エラー: 400 Bad Request

この実装により、セキュアで拡張性の高い認証システムが構築されています。