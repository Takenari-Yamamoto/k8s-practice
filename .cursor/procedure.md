# Go アプリケーション開発から Kubernetes デプロイまでの手順

## 1. 開発環境のセットアップ

1. Go 言語のインストール確認
2. Docker のインストール確認
3. Minikube のインストール確認
4. kubectl のインストール確認
5. Node.js と npm のインストール確認

## 2. バックエンド: Go アプリケーションの開発

1. プロジェクトディレクトリの作成
2. net/http を使用した簡単な Web サーバーの実装
3. CORS 設定の追加（フロントエンドからのアクセスを許可）
4. ローカルでの動作確認

## 3. フロントエンド: React アプリケーションの開発

1. Create React App でプロジェクト作成
   ```bash
   npx create-react-app frontend --template typescript
   ```
2. 必要なパッケージのインストール
   ```bash
   npm install axios @types/axios
   ```
3. API 呼び出しの実装
4. ローカルでの動作確認

## 4. Docker コンテナ化

1. バックエンド
   - Dockerfile の作成
   - イメージのビルド
2. フロントエンド
   - Dockerfile の作成（nginx ベース）
   - イメージのビルド
3. Minikube の docker 環境の有効化
   ```bash
   eval $(minikube docker-env)
   ```
4. 両イメージの Minikube 環境でのビルド

## 5. Kubernetes 環境の準備

1. Minikube クラスタの起動
2. クラスタの状態確認

## 6. Kubernetes へのデプロイ

1. バックエンド
   - デプロイメント用の YAML ファイル作成
   - Service (NodePort) の YAML ファイル作成
2. フロントエンド
   - デプロイメント用の YAML ファイル作成
   - Service (NodePort) の YAML ファイル作成
3. Ingress の設定（オプション）
   - Ingress Controller の有効化
   - Ingress リソースの作成
4. kubectl を使用したデプロイ
5. サービスの公開とアクセス確認

## 7. 追加設定（オプション）

1. Helm のセットアップと利用
2. ArgoCD のセットアップと GitOps の実践
3. 環境変数の管理（ConfigMap/Secret）
4. マイクロサービスの監視設定

## 注意事項

- Minikube での docker-env 設定は必須
- アプリケーションへのアクセスは NodePort または Ingress を使用
- コンテナのビルドは Minikube 環境内で実行すること
- フロントエンドからバックエンドへの通信は Service 名で行う
- 本番環境では適切な CORS 設定が必要

## ディレクトリ構成

```
.
├── golang-app/          # バックエンド
│   ├── cmd/
│   ├── Dockerfile
│   └── go.mod
├── frontend/           # フロントエンド
│   ├── src/
│   ├── package.json
│   └── Dockerfile
└── k8s/               # Kubernetes マニフェスト
    ├── backend/
    │   ├── deployment.yaml
    │   └── service.yaml
    ├── frontend/
    │   ├── deployment.yaml
    │   └── service.yaml
    └── ingress.yaml
```
