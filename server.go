package main

import (
	"fmt"
	"bufio"
	"net"
	"os"
	"net/http"
	"net/http/httputil"
)


func main() {
	fmt.Println("===サーバー起動中===")


	// localhost8080番ポートで待ち受け
	listener, err := net.Listen("tcp", "localhost:8080")
	// 待ち受け終了の遅延実行
	defer listener.Close()
	// 待ち受け失敗したらエラー吐く
	if err != nil {
		fmt.Println("エラー: ", err)
		return
	}

	fmt.Println("===サーバー起動済===")

	// コネクション確立
	conn, err := listener.Accept()
	// 確立終了の遅延実行
	defer conn.Close()
	// 確立失敗したらエラー吐く
	if err != nil {
		fmt.Println("エラー: ", err)
		return
	}

	fmt.Printf("===クライアント接続済: %s===\n", conn.RemoteAddr())

	// クライアントからのリクエスト受け取り
	// buf := make([]byte, 4096)
	request, err := http.ReadRequest(bufio.NewReader(conn))
	if err != nil {
		panic(err)
	}
	dump, err := httputil.DumpRequest(request, true)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dump))

	// クライアントからのリクエストをファイルに出力
	// err = os.WriteFile("./go_server_recv.txt", request, 0644)	// "err = ~~~"になってるのは、文字通りエラーがerrに格納されるから。
	// 書き込み失敗したらエラー吐く
	// if err != nil {
	// 	fmt.Println("エラー: ", err)
	// }

	// レスポンスをファイルから取得
	response, err := os.ReadFile("./server_send.txt")
	// 取得失敗したらエラー吐く
	if err != nil {
		fmt.Println("エラー: ", err)
	}

	// レスポンスをクライアントに送る
	_, err = conn.Write(response)

	// 送るのに失敗したらエラー吐く
	if err != nil {
		fmt.Println("エラー: ", err)
	}
	
	fmt.Println("===サーバー終了中===")
}