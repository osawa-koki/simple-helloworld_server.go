package main

import (
	"fmt"
	"net/http"
)

func main() {

	// ハンドラファンクション -> デフォルトmultiplexerに登録。
	// 第一引数 -> パターン
	// 第二引数 -> 「レスポンスライター」「リクエスト」を引数として受け取る関数
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World 👍👍👍")
	})
	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "I am Greeting you 💓💓💓")
	})

	// リクエストメソッドによって処理を分岐
	http.HandleFunc("/methods", func(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodGet {
			w.Write([]byte("I am a GET man."))
			return
		}
		if r.Method == http.MethodPost {
			w.Write([]byte("I am POST man."))
			return
		}
		if r.Method == http.MethodPut {
			w.Write([]byte("I am PUT man"))
			return
		}
		if r.Method == http.MethodDelete {
			w.Write([]byte("I am DELETE man"))
			return
		}
	})

	http.HandleFunc("/header", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "header -> %s", r.Header)
	})

	http.HandleFunc("/query", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "query -> %s", r.URL.RawQuery)
	})

	http.HandleFunc("/body", func(w http.ResponseWriter, r *http.Request) {
		len := r.ContentLength
		body := make([]byte, len)
		r.Body.Read(body)
		fmt.Fprintln(w, string(body))
	})

	http.HandleFunc("/form", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		fmt.Fprintln(w, r.Form)
	})

	http.HandleFunc("/postform", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		fmt.Fprintln(w, r.PostForm)
	})

	// 静的ファイル配信
	http.Handle("/contents/", http.FileServer(http.Dir("./")))

	// ディレクトリ名とURLパスを変える場合
	http.Handle("/wwwroot/", http.StripPrefix("/wwwroot/", http.FileServer(http.Dir("./contents"))))

	// 第一引数 -> リッスンするアドレス
	// 第二引数 -> 使用するmultiplexer | デフォルトで用意されているものを使用するため、nilを設定
	http.ListenAndServe("0.0.0.0:80", nil)
}
