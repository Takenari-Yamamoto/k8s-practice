# マイクロサービスアプリケーションの開発からデプロイまでの手順

## 1. 開発環境のセットアップ

必要なツールのインストールと確認：

- Go 言語
- Docker Desktop
- kubectl
- Node.js と npm

## 2. バックエンド開発（Go）

1. プロジェクト作成

   ```bash
   mkdir -p golang-app/cmd/server
   cd golang-app
   go mod init golang-app
   ```

2. Web サーバーの実装

   - net/http パッケージを使用
   - `/books` エンドポイントの実装
   - CORS 設定の追加

3. ローカル動作確認
   ```bash
   go run cmd/server/main.go
   curl http://localhost:8080/books
   ```

## 3. フロントエンド開発（React）

1. プロジェクト作成

   ```bash
   npx create-react-app react-app --template typescript
   cd react-app
   ```

2. 必要なパッケージのインストール

   ```bash
   npm install axios @types/axios
   ```

3. コンポーネントの実装

   - 型定義の作成
   - API クライアントの実装
   - 本一覧表示の実装

4. ローカル動作確認
   ```bash
   npm start
   # http://localhost:3000 で確認
   ```

## 4. コンテナ化

1. バックエンド

   ```dockerfile
   # golang-app/Dockerfile
   FROM golang:1.24-alpine as builder
   WORKDIR /app
   COPY . .
   RUN go build -o server ./cmd/server/main.go

   FROM alpine
   COPY --from=builder /app/server /app/
   CMD ["/app/server"]
   ```

2. フロントエンド

   ```dockerfile
   # react-app/Dockerfile
   FROM node:18-alpine as builder
   WORKDIR /app
   COPY . .
   RUN npm install && npm run build

   FROM nginx:alpine
   COPY --from=builder /app/build /usr/share/nginx/html
   COPY nginx.conf /etc/nginx/conf.d/default.conf
   ```

## 5. Kubernetes マニフェストの作成

1. ディレクトリ構造

   ```
   k8s/
   ├── backend/
   │   ├── deployment.yaml
   │   └── service.yaml
   └── frontend/
       ├── deployment.yaml
       └── service.yaml
   ```

2. バックエンドの設定

   - Deployment: 2 レプリカ
   - Service: ClusterIP タイプ

3. フロントエンドの設定
   - Deployment: 2 レプリカ
   - Service: NodePort タイプ（開発環境用）

## 6. ローカル環境でのデプロイ

1. Docker イメージのビルド

   ```bash
   eval $(minikube docker-env)
   docker build -t golang-app:v1 ./golang-app
   docker build -t react-app:v1 ./react-app
   ```

2. Kubernetes へのデプロイ

   ```bash
   kubectl apply -f k8s/backend/
   kubectl apply -f k8s/frontend/
   ```

3. アクセス確認
   ```bash
   minikube service react-app
   ```

## 7. ArgoCD によるデプロイ自動化

1. ArgoCD のインストール

   ```bash
   kubectl create namespace argocd
   kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
   ```

2. ArgoCD の設定

   ```bash
   # サービスタイプの変更（開発環境用）
   kubectl patch svc argocd-server -n argocd -p '{"spec": {"type": "NodePort"}}'

   # 初期パスワードの取得
   kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d

   # UI へのアクセス
   minikube service argocd-server -n argocd
   ```

3. アプリケーション定義の作成

   ```yaml
   # k8s/argocd/application.yaml
   apiVersion: argoproj.io/v1alpha1
   kind: Application
   metadata:
     name: k8s-practice
     namespace: argocd
   spec:
     project: default
     source:
       repoURL: https://github.com/your-username/k8s-practice.git
       targetRevision: HEAD
       path: k8s
     destination:
       server: https://kubernetes.default.svc
       namespace: default
     syncPolicy:
       automated:
         prune: true
         selfHeal: true
   ```

4. アプリケーションのデプロイ
   ```bash
   kubectl apply -f k8s/argocd/application.yaml
   ```

## 8. 動作確認とモニタリング

1. Pod の状態確認

   ```bash
   kubectl get pods
   kubectl get services
   ```

2. ログの確認

   ```bash
   kubectl logs deploy/golang-app
   kubectl logs deploy/react-app
   ```

3. ArgoCD での同期状態の確認
   - UI でアプリケーションの状態を確認
   - 自動同期の動作確認
   - リソースツリーの確認

## 注意事項

1. 開発環境での注意点

   - Minikube 環境では `imagePullPolicy: Never` を使用
   - フロントエンドからバックエンドへは Service 名で通信
   - 環境変数は ConfigMap で管理

2. Git 管理の注意点
   - マニフェストファイルの変更は ArgoCD により自動反映
   - イメージタグの更新は明示的に行う
   - 機密情報は Secret として管理

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
