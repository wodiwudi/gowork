package main

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//优雅关闭服务器
func main() {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		time.Sleep(time.Second * 10)
		c.String(http.StatusOK, "200")
	})
	server := &http.Server{
		Addr:    "8085",
		Handler: r,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen : %s\n", err.Error())
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("shutdown server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("sever shutdown :", err.Error())
	}

	log.Println("server exiting")
}
