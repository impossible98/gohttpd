package command

import (
	// import built-in packages
	"os"
	// import third-party packages
	"github.com/joho/godotenv"
	// import local packages
	"github.com/impossible98/gohttpd/internal/server"
)

func Serve(dir string) {
	err := godotenv.Load()
	if err != nil {
		port := "80"
		server.InitServer(dir, port)
	}
	port := os.Getenv("PORT")
	server.InitServer(dir, port)
}
