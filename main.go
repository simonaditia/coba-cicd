package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

package main

import (
	"net/http"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/ping", ping)
	router.GET("/hello", hello)
	router.GET("/ilt6", ilt6)
	router.GET("/halo", halo)

	// router.Run(":8080")
	
	// Jalankan server HTTP di latar belakang
	go func() {
		if err := http.ListenAndServe(":8080", router); err != nil {
			log.Fatal("Error starting HTTP server:", err)
		}
	}()

	// Konfigurasi server HTTPS
	certFile := "cert.pem"
	keyFile := "key.pem"

	// Jalankan server HTTPS di latar belakang
	go func() {
		if err := http.ListenAndServeTLS(":8443", certFile, keyFile, router); err != nil {
			log.Fatal("Error starting HTTPS server:", err)
		}
	}()

	// Jalankan reverse proxy untuk mengarahkan lalu lintas HTTP ke HTTPS
	if err := http.ListenAndServe(":80", http.HandlerFunc(redirectHTTP)); err != nil {
		log.Fatal("Error starting reverse proxy:", err)
	}
}

func ping(c *gin.Context) {
	c.Data(
		http.StatusOK,
		"application/json",
		[]byte(`{"message": "pong"}`),
	)
}

func hello(c *gin.Context) {
	c.Data(
		http.StatusOK,
		"application/json",
		[]byte(`{"message": "hello ilt 6!"}`),
	)
}

func ilt6(c *gin.Context) {
	c.Data(
		http.StatusOK,
		"application/json",
		[]byte(`{"message": "this is ilt6!"}`),
	)
}

func halo(c *gin.Context) {
	c.Data(
		http.StatusOK,
		"application/json",
		[]byte(`{"message": "halo halo!"}`),
	)
}

func redirectHTTP(w http.ResponseWriter, req *http.Request) {
	target := "https://" + req.Host + req.URL.Path
	http.Redirect(w, req, target, http.StatusMovedPermanently)
}



/*
package main

import (
	"net/http"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/ping", ping)
	router.GET("/hello", hello)
	router.GET("/ilt6", ilt6)
	router.GET("/halo", halo)

	// router.Run(":8080")
	
	// Jalankan server HTTP di latar belakang
	go func() {
		if err := http.ListenAndServe(":8080", router); err != nil {
			log.Fatal("Error starting HTTP server:", err)
		}
	}()

	// Konfigurasi server HTTPS
	certFile := "cert.pem"
	keyFile := "key.pem"

	// Jalankan server HTTPS di latar belakang
	go func() {
		if err := http.ListenAndServeTLS(":8443", certFile, keyFile, router); err != nil {
			log.Fatal("Error starting HTTPS server:", err)
		}
	}()

	// Jalankan reverse proxy untuk mengarahkan lalu lintas HTTP ke HTTPS
	if err := http.ListenAndServe(":80", http.HandlerFunc(redirectHTTP)); err != nil {
		log.Fatal("Error starting reverse proxy:", err)
	}
}

func ping(c *gin.Context) {
	c.Data(
		http.StatusOK,
		"application/json",
		[]byte(`{"message": "pong"}`),
	)
}

func hello(c *gin.Context) {
	c.Data(
		http.StatusOK,
		"application/json",
		[]byte(`{"message": "hello ilt 6!"}`),
	)
}

func ilt6(c *gin.Context) {
	c.Data(
		http.StatusOK,
		"application/json",
		[]byte(`{"message": "this is ilt6!"}`),
	)
}

func halo(c *gin.Context) {
	c.Data(
		http.StatusOK,
		"application/json",
		[]byte(`{"message": "halo halo!"}`),
	)
}

func redirectHTTP(w http.ResponseWriter, req *http.Request) {
	target := "https://" + req.Host + req.URL.Path
	http.Redirect(w, req, target, http.StatusMovedPermanently)
}
*/