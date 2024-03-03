import socket

class Client:
    def request(self):
        print("===クライアント起動===")

        try:
            clientSocket = socket.socket()
            clientSocket.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)

            print("===サーバーに接続中===")
            clientSocket.connect(('localhost',80))
            print("===サーバーに接続済===")

            with open('./client_send.txt', 'rb') as f:
                request = f.read()

            clientSocket.send(request)

            response = clientSocket.recv(4096)

            with open('./client_recv.txt', 'wb') as f:
                f.write(response)
            
            clientSocket.close()
            

        finally:
            print("===クライアント終了中===")

if __name__ == '__main__':
    client = Client()
    client.request()