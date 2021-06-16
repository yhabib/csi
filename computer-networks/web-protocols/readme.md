# Proxy Cache Server

How to run it?

* Start the server with `python3 server.py`

* Start the proxu with `go run http_proxy.go`

* Run a client in localhost:8080

## Questions

* Server creates socket, `serverSocket` to listen for new connections. It is where the 3-way handshake happense, after this it creates a new dedicated socket in a new PORT for each connection, `connectionSocket`. How does the client know about this change of Socket?

* hop-by-hop headers, so proxies consume them. Are they indicated for the Client-Proxy connection? In the exercise, should the Proxy implement the keep-alive connection w/ the server?

* .py server in 1.0, is this the root for the issues w/ the payload not being there sometimes?

* Connection pooling vs load balancing? A pool for each server?
