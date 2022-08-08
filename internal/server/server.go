package server

import (
	// import built-in packages
	"fmt"
	"net/http"
	// import local packages
	"github.com/impossible98/gohttpd/global"
)

func InitServer(dir string, port string) {
	http.Handle("/", http.FileServer(http.Dir(dir)))
	fmt.Println()
	fmt.Println("  " + global.APP_NAME + " v" + global.VERSION + " server running at:")
	fmt.Println()
	fmt.Println("  > Local:   http://localhost:" + port)
	fmt.Println("  > Network: http://localhost:" + port)
	fmt.Println()
	http.ListenAndServe(":"+port, nil)
}
