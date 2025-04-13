# Go アプリケーション開発から Kubernetes デプロイまでの手順

## 1. 開発環境のセットアップ

1. Go 言語のインストール確認
2. Docker のインストール確認
3. Minikube のインストール確認
4. kubectl のインストール確認

## 2. Go アプリケーションの開発

1. プロジェクトディレクトリの作成
2. net/http を使用した簡単な Web サーバーの実装
3. ローカルでの動作確認

## 3. Docker コンテナ化

1. Dockerfile の作成
2. Minikube の docker 環境の有効化
   ```bash
   eval $(minikube docker-env)
   ```
3. Docker イメージのビルド

## 4. Kubernetes 環境の準備

1. Minikube クラスタの起動
2. クラスタの状態確認

## 5. Kubernetes へのデプロイ

1. デプロイメント用の YAML ファイル作成
   - アプリケーションのデプロイメント定義
   - NodePort タイプのサービス定義
2. kubectl を使用したデプロイ
3. サービスの公開とアクセス確認

## 6. 追加設定（オプション）

1. Helm のセットアップと利用
2. ArgoCD のセットアップと GitOps の実践

## 注意事項

- Minikube での docker-env 設定は必須
- アプリケーションへのアクセスは NodePort を使用
- コンテナのビルドは Minikube 環境内で実行すること
