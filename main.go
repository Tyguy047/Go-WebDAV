package main

import (
    "log"
    "net/http"
    "time"
    "golang.org/x/net/webdav"
)

func main() {

    checkForUser()
	checkFolder()

    handler := &webdav.Handler{
        Prefix: "/",
        FileSystem: webdav.Dir("./data"),
        LockSystem: webdav.NewMemLS(),
        Logger: func(r *http.Request, err error) {
            // Silent logger - suppress WebDAV internal logging for performance
        },
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

    // Custom HTTP server with optimized settings for large file transfers
    server := &http.Server{
        Addr:           ":8080",
        Handler:        authMiddleware(handler),
        ReadTimeout:    0,                 // No timeout for large file uploads
        WriteTimeout:   0,                 // No timeout for large file downloads
        IdleTimeout:    120 * time.Second, // Keep connections alive between requests
        MaxHeaderBytes: 1 << 20,           // 1 MB max header size
    }

    log.Fatal(server.ListenAndServe())
}