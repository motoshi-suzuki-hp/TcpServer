package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Println("===クライアント起動===")

	conn, err := net.Dial("tcp", "localhost:80")
	if err != nil {
		fmt.Println("エラー: ", err)
		return
	}
	defer conn.Close()

	fmt.Println("===サーバーに接続済===")

	fileName := "./client_send.txt"
	request, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("エラー: ", err)
		return
	}

	_, err = conn.Write(request)
	if err != nil {
		fmt.Println("エラー: ", err)
		return
	}

	response := make([]byte, 4096)
	_, err = conn.Read(response)
	if err != nil {
		fmt.Println("エラー: ", err)
		return
	}

	fileName = "./client_recv.txt"
	err = os.WriteFile(fileName, response, 0644)
	if err != nil {
		fmt.Println("エラー: ", err)
		return
	}

	fmt.Println("===クライアント終了===")
}
