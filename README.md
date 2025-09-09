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

## 層の分け方の例

1. 核方針（コアポリシー）、業務概念（ドメイン）、処理手順（ユースケース）、詳細実装
2. ドメイン（コアポリシー + Entities + Value）、処理手順（ユースケース）、詳細実装
3. ボブおじさん(Robert C. Martin)の図

### ありがちな例と用語

1. ドメイン層

- 役割: ビジネスルールそのものを表す層で、外部の技術や I/O には依存しない
- 構成要素:
  - エンティティ (Entity): 業務上の概念を表現するモデル
  - ドメインサービス (Domain Service): エンティティに属さないルールのまとめ
- 特徴:
  - フレームワークやデータベースに一切依存しない「純粋なロジック」
  - 最も内側で、他の層から再利用される中心

2. アプリケーション層 (Application Layer)

- 役割: ドメインをどう使うかの「手順書」を定める層
- 構成要素:
  - ユースケース (Use Case / Interactor):
    入力を受け取り、ドメインサービスやエンティティを呼び出して、
    結果を出力にまとめる具体的な処理
  - ポート (Port): ユースケースが外部から呼び出されるためのインタフェース
    - InputPort / OutputPort がある。
- 特徴:
  - ドメイン層に依存するが、外部 (HTTP, DB) には依存しない
  - システムに「どんな機能があるか」をユースケースとして定義する場

3. インターフェース層 / アダプタ層 (Interface / Adapter Layer)

- 役割: 外界（Web, CLI, DB, ファイルなど）とアプリケーション層をつなぐ「翻訳者」
- 構成要素:
  - コントローラ / ハンドラ: HTTP や CLI からのリクエストを受けて、ユースケースを呼び出す
  - パーサ: 外部表現をドメインオブジェクトに変換
  - バリデータ: 入力の整合性チェックを行い、ドメインに渡す前に不正データを弾く
  - プレゼンター: ユースケースの出力を HTTP レスポンスや JSON 形式に整形する
- 特徴:
  - 外部フォーマットの依存をここに閉じ込める
  - ここを差し替えれば、同じユースケースを REST 以外（gRPC, CLI）でも利用可能

4. インフラストラクチャ層 (Infrastructure Layer)

- 役割: 実際の技術要素（フレームワーク、データベース、ログ、クラウド環境）を扱う
- 構成要素:
  - フレームワーク起動: Gin サーバ、ルーティング初期化
  - DI: 依存解決。どの Decoder 実装を使うかなど
  - 永続化や外部サービス連携: DB や API 呼び出し
  - デプロイ設定: Dockerfile, ECS タスク定義など
- 特徴:
  - 最も外側で、変化しやすい部分
  - 「実行の場」を提供するが、ビジネスルールには立ち入らない

## 作る手順の例１(Go)

### 1. 設計

https://architecting.hateblo.jp/entry/2025/05/02/162516

- ユビキタス言語の定義
- 入力 → ユースケースで実施すること → 出力 を考える
- その他、どんな構成・機能・入出力か 考える

### 2. 各層で実装することを考える

#### 例: ある bash のコマンドを実行した結果を返す実装

- `各層の実装例.md` を参照

### 3. 層をフォルダにマッピングする

#### 例: ある bash のコマンドを実行した結果を返す実装

- Command: Go で実行しやすい形式に落とし込んだ文字列
- Result: 実行結果
- CommandType: 命令の種類

```
/app
├─ cmd/server/main.go           # エントリポイント（DI呼び出し＋HTTP起動）
├─ internal/
│  ├─ domain/                   # 1. ドメイン層（純粋ルール）
│  │  ├─ command.go
│  │  ├─ result.go
│  │  └─ decoder.go
│  ├─ usecase/                  # 2. アプリケーション層（手順 = ユースケース）
│  │  ├─ ports.go
│  │  └─ decode_interactor.go
│  ├─ adapter/                  # 3. インターフェース/アダプタ層
│  │  ├─ http/
│  │  │  ├─ handler.go
│  │  │  └─ router.go
│  │  ├─ parse/
│  │  │  ├─ parser.go
│  │  │  └─ whitespace_parser.go
│  │  ├─ validate/
│  │  │  ├─ validator.go
│  │  │  └─ command_validator.go
│  │  └─ presenter/
│  │     └─ response.go
│  └─ platform/                 # 4. インフラ層（技術詳細/起動/DI）
│     ├─ di.go
│     └─ config.go
└─ go.mod
```

### 4. 実運用のための“運用”フォルダ（任意）

```
api/          # openapi.yml など、外部契約
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
