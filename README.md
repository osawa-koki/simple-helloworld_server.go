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
- /post-form (フォームデータとクエリ文字列をセット)
