package main

import (
    "log"
    "net/http"
    "golang.org/x/net/webdav"
)

func main() {

    checkForUser()
	checkFolder()

    handler := &webdav.Handler{
        Prefix: "/",
        FileSystem: webdav.Dir("./data"),
        LockSystem: webdav.NewMemLS(),
    }

    authMiddleware := func(h http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            username, password, ok := r.BasicAuth()
            if !ok || !auth(username, password) {
                w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
                http.Error(w, "Unauthorized", http.StatusUnauthorized)
                return
            }
            h.ServeHTTP(w, r)
        })
    }

    addr := ip()
    log.Println("Server is running!")
    // Opening and loging in in your browser may cause a flood of requests!
    log.Println("http://" + addr + ":8080")
    http.ListenAndServe(":8080", authMiddleware(handler))
}