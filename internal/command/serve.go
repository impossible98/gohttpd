package command

import (
	// import built-in packages
	"fmt"
	"net/http"
)

func Serve(dir string) {
	http.Handle("/", http.FileServer(http.Dir(dir)))
	fmt.Println("Serving " + dir + " on http://localhost:80")
	http.ListenAndServe(":80", nil)
}
