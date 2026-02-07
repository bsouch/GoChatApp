package main

import "net/http"

type ReadinessPayload struct {
	Result string
}

func handlerReadiness(writer http.ResponseWriter, request *http.Request) {
	jsonResponse(writer, 200, ReadinessPayload{Result: "Success"})
}
