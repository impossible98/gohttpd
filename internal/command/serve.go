package command

import (
	// import local packages
	"github.com/impossible98/gohttpd/internal/server"
	"github.com/impossible98/gohttpd/pkg/env"
)

func Serve(dir string) {
	port := env.GetEnv("PORT")
	server.InitServer(dir, port)
}
