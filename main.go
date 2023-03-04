package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// httpcode represents data about a http code.
type httpcode struct {
	Code    float64 `json:"code"`
	Message string  `json:"message"`
}

// codes slice to seed http code data.
var codes = []httpcode{
	{Code: 100, Message: "Continue"},
	{Code: 101, Message: "Switching Protocols"},
	{Code: 102, Message: "Processing"},
	{Code: 103, Message: "Early Hints"},
	{Code: 200, Message: "OK"},
	{Code: 201, Message: "Created"},
	{Code: 202, Message: "Accepted"},
	{Code: 203, Message: "Non-Authoritative Information"},
	{Code: 204, Message: "No Content"},
	{Code: 206, Message: "Partial Content"},
	{Code: 207, Message: "Multi-Status"},
	{Code: 300, Message: "Multiple Choices"},
	{Code: 301, Message: "Moved Permanently"},
	{Code: 302, Message: "Found"},
	{Code: 303, Message: "See Other"},
	{Code: 304, Message: "Not Modified"},
	{Code: 305, Message: "Use Proxy"},
	{Code: 307, Message: "Temporary Redirect"},
	{Code: 308, Message: "Permanent Redirect"},
	{Code: 400, Message: "Bad Request"},
	{Code: 401, Message: "Unauthorized"},
	{Code: 402, Message: "Payment Required"},
	{Code: 403, Message: "Forbidden"},
	{Code: 404, Message: "Not Found"},
	{Code: 405, Message: "Method Not Allowed"},
	{Code: 406, Message: "Not Acceptable"},
	{Code: 407, Message: "Proxy Authentication Required"},
	{Code: 408, Message: "Request Timeout"},
	{Code: 409, Message: "Conflict"},
	{Code: 410, Message: "Gone"},
	{Code: 411, Message: "Length Required"},
	{Code: 412, Message: "Precondition Failed"},
	{Code: 413, Message: "Payload Too Large"},
	{Code: 414, Message: "Request-URI Too Long"},
	{Code: 415, Message: "Unsupported Media Type"},
	{Code: 416, Message: "Request Range Not Satisfiable"},
	{Code: 417, Message: "Expectation Failed"},
	{Code: 418, Message: "I’m a teapot"},
	{Code: 420, Message: "Enhance Your Calm"},
	{Code: 421, Message: "Misdirected Request"},
	{Code: 422, Message: "Unprocessable Entity"},
	{Code: 423, Message: "Locked"},
	{Code: 424, Message: "Failed Dependency"},
	{Code: 425, Message: "Too Early"},
	{Code: 426, Message: "Upgrade Required"},
	{Code: 429, Message: "Too Many Requests"},
	{Code: 431, Message: "Request Header Fields Too Large"},
	{Code: 444, Message: "No Response"},
	{Code: 450, Message: "Blocked by Windows Parental Controls"},
	{Code: 451, Message: "Unavailable For Legal Reasons"},
	{Code: 497, Message: "HTTP Request Sent to HTTPS Port"},
	{Code: 498, Message: "Token expired/invalid"},
	{Code: 499, Message: "Client Closed Request"},
	{Code: 500, Message: "Internal Server Error"},
	{Code: 501, Message: "Not Implemented"},
	{Code: 502, Message: "Bad Gateway"},
	{Code: 503, Message: "Service Unavailable"},
	{Code: 504, Message: "Gateway Timeout"},
	{Code: 506, Message: "Variant Also Negotiates"},
	{Code: 507, Message: "Insufficient Storage"},
	{Code: 508, Message: "Loop Detected"},
	{Code: 509, Message: "Bandwidth Limit Exceeded"},
	{Code: 510, Message: "Not Extended"},
	{Code: 511, Message: "Network Authentication Required"},
	{Code: 521, Message: "Web Server Is Down"},
	{Code: 522, Message: "Connection Timed Out"},
	{Code: 523, Message: "Origin Is Unreachable"},
	{Code: 525, Message: "SSL Handshake Failed"},
	{Code: 599, Message: "Network Connect Timeout Error"},
}

func main() {
	app := gin.Default()

	app.Static("/static", "./static")
	app.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Status": "Healthy",
		})
	})

	app.LoadHTMLGlob("templates/*")

	app.GET("/codes", getCodes)
	app.GET("/:code", index)
	// listen and serve on 0.0.0.0:8080
	app.Run()
}

// getCodes responds with the list of all codes in JSON format.
func getCodes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, codes)
}

func index(c *gin.Context) {
	startTime := time.Now()

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"Code":     c.Param("code"),
		"LoadTime": time.Since(startTime),
		// "Image":     "static/images/" + c.Param("code") + ".jpg",
		"Animation": "static/" + c.Param("code") + ".json",
	})
}
