package server

import (
	// import built-in packages
	"net/http"
	// import local packages
	"github.com/impossible98/gohttpd/pkg/server"
)

func InitServer(dir string, port string) {
	http.Handle("/", http.FileServer(http.Dir(dir)))
	server.RunTip(port)
	http.ListenAndServe(":"+port, nil)
}
