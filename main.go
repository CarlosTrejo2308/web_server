package main

func main() {
	// Set a server on port 3000
	server := NewServer(":3000")

	// Set endpoints
	server.Hanlde("GET", "/", HandleRoot)
	server.Hanlde("POST", "/create", PostRequest)
	server.Hanlde("POST", "/user", UserPostRequest)

	server.Hanlde("POST", "/api", server.AddMidleware(HandleHome, CheckAuth(), Logging()))

	// Start the server
	err := server.Listen()
	if err != nil {
		panic(err)
	}
}
