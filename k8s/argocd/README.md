# ArgoCD Setup

## ローカル環境での起動方法

1. ArgoCD をインストール:

```bash
kubectl apply -f install.yaml
```

2. アプリケーションの定義を適用:

```bash
# GitHubリポジトリのURLを更新後に実行
kubectl apply -f applications.yaml
```

3. ArgoCD の UI にアクセス:

- URL: http://localhost:31181

## 初期設定

1. 初期管理者パスワードの取得:

```bash
kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d
```

2. ログイン情報:

- ユーザー名: admin
- パスワード: 上記コマンドで取得したパスワード

## 注意事項

- `applications.yaml`内の GitHub リポジトリ URL を、実際のリポジトリ URL に更新してください。
- 初回ログイン後は、セキュリティのためパスワードを変更することを推奨します。
- ArgoCD は自動的に Git リポジトリの変更を検知し、Kubernetes クラスターに反映します。
