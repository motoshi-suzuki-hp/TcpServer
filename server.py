import socket

class Server:
    def serve(self):
        print("===サーバー起動中===")

        try:
            serverSocket = socket.socket()
            serverSocket.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)

            # localhost8080番ポートで待ち受け
            serverSocket.bind(('localhost', 8080))
            serverSocket.listen(10)

            print("===サーバー起動済===")

            # コネクション確立
            (clientSocket, address) = serverSocket.accept()

            print(f"===クライアント接続済み:{address}===")

            # クライアントからのリクエスト受け取り
            request = clientSocket.recv(4096)

            with open('./py_server_recv.txt', 'wb') as f:
                f.write(request)
            
            with open('./server_send.txt', 'rb') as f:
                response = f.read()
            
            clientSocket.send(response)

            clientSocket.close()
            

        finally:
            print("===サーバー終了中===")

if __name__ == "__main__":
    server = Server()
    server.serve()
