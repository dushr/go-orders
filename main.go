package orders

import "net/http"

func init() {
	router := NewRouter()
	http.Handle("/", router)
}
