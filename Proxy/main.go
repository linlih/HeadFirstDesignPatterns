package main

import (
	"Proxy/nginx"
	"fmt"
)

func main() {
	nginxServer := nginx.NewNginxServer()
	appStatusURL := "/app/status"
	createuserURL := "/create/user"

	httpCode, body := nginxServer.HandleRequest(appStatusURL, "GET")
	fmt.Printf("\nURL: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.HandleRequest(appStatusURL, "GET")
	fmt.Printf("\nURL: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.HandleRequest(appStatusURL, "GET")
	fmt.Printf("\nURL: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.HandleRequest(createuserURL, "POST")
	fmt.Printf("\nURL: %s\nHttpCode: %d\nBody: %s\n", createuserURL, httpCode, body)

	httpCode, body = nginxServer.HandleRequest(createuserURL, "GET")
	fmt.Printf("\nURL: %s\nHttpCode: %d\nBody: %s\n", createuserURL, httpCode, body)
}
