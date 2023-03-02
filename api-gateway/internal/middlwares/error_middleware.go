package middlwares

import (
	"fmt"
	"net/http"

	"google.golang.org/grpc/status"
)

func CheckErrorMiddlware(prev func(w http.ResponseWriter, r *http.Request) error) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := prev(w, r)
		if err != nil {
			st, ok := status.FromError(err)
			if !ok {
				fmt.Println("123")
				fmt.Println(err)
			}
			fmt.Println(st.Code(), st.Message())
		}
	})
}
