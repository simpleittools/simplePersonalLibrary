package main

import (
	"fmt"
	"github.com/simpleittools/simplepersonallibrary/pkg/handlers"
	"net/http"
)

// TODO: move port to environmental variable
const port = ":3000"

func main() {

	// TODO: move this to a routes handler
	http.HandleFunc(
		"/", handlers.Home,
	)
	fmt.Println(fmt.Sprintf("The server is running at http://localhost%s", port))
	_ = http.ListenAndServe(port, nil)

}
