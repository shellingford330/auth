package middleware

import "net/http"

func CommonHeader(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// CORS対応
		writer.Header().Add("Access-Control-Allow-Origin", "*")
		writer.Header().Add("Access-Control-Allow-Headers", "Content-Type,Accept,Origin,x-token")

		// 共通のレスポンスヘッダを設定
		writer.Header().Add("Content-Type", "application/json")
		nextFunc(writer, request)
	}
}
