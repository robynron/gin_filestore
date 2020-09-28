package handler

import "net/http"

//拦截器
func HTTPInterceptor(h http.HandlerFunc) http.HandlerFunc{
	return func(writer http.ResponseWriter, request *http.Request) {
			request.ParseForm()

			username := request.Form.Get("username")
			token := request.Form.Get("token")

			if len(username) < 3 || !IsTokenValid(token){
				writer.WriteHeader(http.StatusForbidden)
				return
			}

			h(writer, request)
		}
}