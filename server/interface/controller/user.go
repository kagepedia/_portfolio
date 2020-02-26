package controller

import (
	"fmt"
	"net/http"
)

// =============================
//    TEST
// =============================
func helloHandle(w http.ResponseWriter, r *http.Request) {
	h := `
        <html>
            <head>
                <title>Hello</title>
            </head>
            <body>
                Hello
            </body>
        </html>
        `
	fmt.Fprint(w, h)
}

func goodbyeHandle(w http.ResponseWriter, r *http.Request) {
	h := `
        <html>
            <head>
                <title>goodbye</title>
            </head>
            <body>
                goodbye
            </body>
        </html>
        `
	fmt.Fprint(w, h)
}

func landingHandle(w http.ResponseWriter, r *http.Request) {
	h := `
        <html>
            <head>
                <title>Landing</title>
            </head>
            <body>
                Landing
            </body>
        </html>
        `
	fmt.Fprint(w, h)
}
