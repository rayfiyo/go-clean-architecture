# go-clean-architecture

- An example of clean architecture in Go and notes
- Go 言語におけるクリーンアーキテクチャの一例とメモ

# クリーンアーキテクチャの骨子

- 課題に対して揮発性の低い核（不変方針）に、外側（可変詳細）が依存する構造則
  - 課題: ビジネスルールやシステムが解決すべき課題
  - 揮発性の低い: **時間経過で**変化しやすい度合いの低さ
    - 時間経過の不変性の高い、時間経過の安定性の高い、変更耐性の高い とも言い換えられる
      - 可変性は変更のしやすさ
    - ボブおじさん(Robert C. Martin)は『Clean Architecture』以前から
      “volatility-based decomposition” と呼んでいるらしい（**要出典**）
      - https://www.cool-man.org/software-architecture/
      - https://nadermedhatthoughts.medium.com/granularity-in-software-architecture-bec7c432d6d3

## 用語対応

- エンティティ(domain):
  業務ルールの中心（純粋な構造体や値オブジェクト、ドメインサービス）
- ユースケース(usecase / application):
  操作手順・入力/出力ポート（ビジネスフロー）
- インターフェースアダプタ(interface / adapters):
  DB/HTTP など外界との橋渡し（実装が集まる）
- フレームワーク&ドライバ(infrastructure / presentation):
  Web フレームワーク、DB クライアント等

## 層の分け方の例

1. 核方針（コアポリシー）、業務概念（ドメイン）、処理手順（ユースケース）、詳細実装
2. ドメイン（コアポリシー + Entities + Value）、処理手順（ユースケース）、詳細実装
3. ボブおじさん(Robert C. Martin)の図

## 作る手順の例１(Go)

### 1. 「単一プロダクト」の外枠を決める

- 外部公開バイナリの入口は `cmd/<appname>/main.go`
- ライブラリ化しないアプリなら、実装は `internal/` の下に寄せる
  - 外部から import 不能にできるため

```
hoge-app-repo/
├─ cmd/
│  └─ hoge-app/
│     └─ main.go
├─ internal/
└─ go.mod
```

## 2. 層をフォルダにマッピングする

- 最も迷いにくい無難案：

```
internal/
├─ domain/          # エンティティ、ドメインサービス、リポジトリIF
├─ usecase/         # 入力/出力ポート、ユースケース実装
├─ interface/       # コントローラ、プレゼンター、ゲートウェイIF実装の薄い層
├─ infrastructure/  # DB, 外部API, ロガーなど具体実装
└─ presentation/    # WebサーバやHTTPルーティング（Echo/Fiber等）
```

- 許容される依存方向の例
  - `presentation -> interface -> usecase -> domain`
  - `interface -> infrastructure`

## 3. 「境界の契約（インターフェース）」を先に置く

- `domain`: エンティティ（構造体）、リポジトリのインターフェース
- `usecase`: ユースケース入力/出力ポート（インターフェース）
- 実装は後で `infrastructure` / `presentation` に置く

### 例：ToDo アプリの最小雛形

- このレポジトリの実装のこと
- 次の順番で実装する（上から順に実装）

```
internal/
├── domain
│   ├── todo.go
│   └── todo_repository.go
├─── usecase
│   └── add_todo.go
├── interface
│   └── http
│       └── todo_handler.go
├── infrastructure
│   └── sqlite
│       └── todo_repository_sqlite.go
└── presentation
    └── httpserver
        └── router.go
cmd/
└── todoapp
    └── main.go # ここで、具体的な実装を組み立てる（依存性逆転）、内部は外部を知らない
```

- 大規模の場合、`internal/todo/domain/` のように、機能別にフォルダを切ると良い

### 4. テストの置き場

- ユースケースは、ブラックボックスでテスト
  - fake リポジトリを `usecase` 直下や `test/` に置く
- インフラは、結合テスト用に `test/integration/` を用意
- Go は、`*_test.go` を各パッケージに置くのが自然

### 5. 実運用のための“運用”フォルダ（任意）

```
build/        # Dockerfile, CIスクリプト
configs/      # 設定テンプレート（YAML/TOML等）
migrations/   # DBマイグレーション
scripts/      # 開発用スクリプト（fishでもOK）
```

## 作る手順の例２

1. イベントを並べて業務機能（サブドメイン）ごとに箱を作る
2. その中で「ユースケース → モデル → コアポリシー」へと安定度順に粒度を細かくする
   - コアポリシー: 計算・制約だけを持つ最も安定した層
   - モデル: コアポリシーを包む
   - ユースケース: 外部との橋渡し
3. 外界との I/O はポート (interface) に閉じ込め、実装は最外殻のアダプタへ
4. 依存は常に不安定 → 安定 の一方向になるよう interface を配置
5. コアだけで動くテストが通るかで構造を検証し、方針に依存し詳細は差し替えられる状態を確認

- https://chatgpt.com/share/685b4add-5104-8006-8f71-310c38da8477
