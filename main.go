package main

func main() {
	server := NewServer(":3000")

	server.Hanlde("GET", "/", HandleRoot)
	server.Hanlde("POST", "/create", PostRequest)
	server.Hanlde("POST", "/user", UserPostRequest)

	server.Hanlde("POST", "/api", server.AddMidleware(HandleHome, CheckAuth(), Logging()))

	err := server.Listen()
	if err != nil {
		panic(err)
	}
}
