package main

import (
	"fmt"
	"net/http"
)

func main() {
	// TODO: move port to environmental variable
	port := ":3000"

	// TODO: move this to a routes handler
	http.HandleFunc(
		"/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello From Simple Personal Library")
		},
	)
	fmt.Println(fmt.Sprintf("The server is running at http://localhost%s", port))
	_ = http.ListenAndServe(port, nil)

}
