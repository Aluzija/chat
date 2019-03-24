package main

import (
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte(`
            <html>
                <head>
                    <meta charset="utf-8" />
                    <title>Komunikator</title>
                </head>
                <body>
                    Pogadajmy!
                </body>
            </html>
        `))
    })
    // uruchomienie serwera WWW
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal("ListenAndServe:", err)
    }
}
