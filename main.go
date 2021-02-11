package main

func main() {
	server := newServer(":3000")
	server.Handle("GET","/", HandleRoot)
	server.Handle("POST", "/create", POSTRequest)
	server.Handle("POST", "/user", UserPOSTRequest)
	server.Handle("POST","/api", server.AddMiddleware(HandleHome, CheckAuth(), Loging()))
	_ = server.Listen()
}
