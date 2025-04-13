# ArgoCD Setup Guide

## 前提条件

- minikube
- kubectl
- argocd CLI (オプション)

## セットアップ手順

### 1. minikube の起動

```bash
minikube start
```

### 2. ArgoCD のインストール

```bash
# ArgoCDの名前空間を作成
kubectl create namespace argocd

# ArgoCDをインストール
kubectl apply -f install.yaml
```

### 3. サーバー設定の適用

```bash
# ArgoCDサーバーの設定を適用
kubectl apply -f server.yaml
```

### 4. アクセス方法

#### ポートフォワーディングを使用する場合（推奨）

```bash
# HTTPSでアクセスする場合（8080ポート）
kubectl port-forward svc/argocd-server -n argocd 8080:443

# HTTPでアクセスする場合（9000ポート）
kubectl port-forward svc/argocd-server -n argocd 9000:80
```

その後、以下のいずれかの URL でアクセス：

- HTTPS: https://localhost:8080 （自己署名証明書の警告が表示されます）
- HTTP: http://localhost:9000

※ HTTPS アクセスの場合、自己署名証明書の警告が表示されますが、開発環境では「詳細設定」から「安全でないサイトにアクセスする」を選択して進めることができます。

#### NodePort を使用する場合

```bash
# NodePortとminikubeのIPアドレスを確認
kubectl get svc argocd-server -n argocd
minikube ip

# 出力例:
# NAME            TYPE       CLUSTER-IP     EXTERNAL-IP   PORT(S)                      AGE
# argocd-server   NodePort   10.x.x.x      <none>        80:31181/TCP,443:30692/TCP   1m
#
# minikube ip の出力例:
# 192.168.49.2
```

その後、以下のいずれかの URL でアクセス：

- HTTP: http://[minikube-ip]:[HTTP NodePort]
  例: http://192.168.49.2:31181
- HTTPS: https://[minikube-ip]:[HTTPS NodePort]
  例: https://192.168.49.2:30692

### 5. ログイン情報

- ユーザー名: `admin`
- パスワード: 以下のコマンドで取得

```bash
kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d
```

## アプリケーションのデプロイ

### CLI を使用する場合

```bash
# CLIでログイン
argocd login localhost:8080

# アプリケーションの作成
argocd app create [APP_NAME] \
  --repo [GIT_REPO_URL] \
  --path [MANIFEST_PATH] \
  --dest-server https://kubernetes.default.svc \
  --dest-namespace default

# アプリケーションの同期
argocd app sync [APP_NAME]
```

### UI を使用する場合

1. ArgoCD の UI にアクセス
2. 「NEW APP」ボタンをクリック
3. アプリケーション情報を入力
   - Application Name: アプリケーション名
   - Project: default
   - Repository URL: Git リポジトリの URL
   - Path: マニフェストファイルのパス
   - Destination: https://kubernetes.default.svc
   - Namespace: デプロイ先の名前空間

## トラブルシューティング

### 証明書の警告が表示される場合

開発環境では自己署名証明書を使用しているため、ブラウザで警告が表示されます。
「詳細設定」から「安全でないサイトにアクセスする」を選択してください。

### イメージのプル失敗

ローカルで作成したイメージを使用する場合は、以下のコマンドで minikube にイメージをロードする必要があります：

```bash
minikube image load [IMAGE_NAME]:[TAG]
```

### Pod が起動しない場合

以下のコマンドで Pod の状態とログを確認できます：

```bash
# Podの状態確認
kubectl get pods -n argocd

# Podのログ確認
kubectl logs -n argocd [POD_NAME]
```
