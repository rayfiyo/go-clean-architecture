# go-clean-architecture

- An example of clean architecture in Go and notes
- Go 言語におけるクリーンアーキテクチャの一例とメモ

# クリーンアーキテクチャの骨子

- 目的に対して揮発性の低い核（不変方針）に、外側（可変詳細）が依存する構造則
  - 目的: ビジネスルールやシステムが解決すべき課題についての目的
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

## 作る手順の例

1. イベントを並べて業務機能（サブドメイン）ごとに箱を作る
2. その中で「ユースケース → モデル → コアポリシー」へと安定度順に粒度を細かくする
   - コアポリシー: 計算・制約だけを持つ最も安定した層
   - モデル: コアポリシーを包む
   - ユースケース: 外部との橋渡し
3. 外界との I/O はポート (interface) に閉じ込め、実装は最外殻のアダプタへ
4. 依存は常に不安定 → 安定 の一方向になるよう interface を配置
5. コアだけで動くテストが通るかで構造を検証し、方針に依存し詳細は差し替えられる状態を確認

- https://chatgpt.com/share/685b4add-5104-8006-8f71-310c38da8477
