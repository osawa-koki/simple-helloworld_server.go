# simple-helloworld_server.go

Go言語でAPIサーバを構築する際のテンプレート。  

## 実行方法

```shell
# Dockerfileのビルドの実行
docker build -t simple-go-server .
docker run -p 80:80 -it --rm --name my-simple-go-server simple-go-server

# 一行で書くなら、、、
docker build -t simple-go-server . && docker run -p 80:80 -it --rm --name my-simple-go-server simple-go-server
```

## 動作確認

以下のパスに対してアクセスしてみる。  

- /
- /greet
- /query?ilove=tako
- /methods (GET/POST/PUT/DELETE)
- /contents/tako.png
- /wwwroot/tako.png
- /body (ボディ部にデータをセット)
- /form (フォームデータとクエリ文字列をセット)
- /postform (フォームデータとクエリ文字列をセット)

## デプロイ設定(Render.com)

| キー | バリュー |
| ---- | ---- |
| Name | simple-go-restapi-server |
| Region | Oregon(US West) |
| Branch | main |
| Root Directory |  |
| Environment | Docker |
| Dockerfile Path | ./Dockerfile |
| Docker Build Context Directory |  |
| Docker Command |  |
