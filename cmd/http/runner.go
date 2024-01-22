package main

import (
	http2 "gotest/internal/controller/http"
	"gotest/internal/core/config"
	"gotest/internal/core/server/http"
	"gotest/internal/core/service"
	infraConf "gotest/internal/infra/config"
	"gotest/internal/infra/repository"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

func main() {
	instance := gin.New()
	instance.Use(gin.Recovery())

	db, err := repository.NewDB(
		infraConf.DBConfig{
			Driver:                  "mysql",
			Url:                     "backnumber:11111111@tcp(127.0.0.1:3306)/react?charset=utf8mb4&parseTime=true&loc=UTC&tls=false&readTimeout=3s&writeTimeout=3s&timeout=3s&clientFoundRows=true",
			ConnMaxLifetimeInMinute: 3,
			MaxOpenConns:            10,
			MaxIdleConns:            1,
		},
	)

	if err != nil {
		log.Fatalf("failed database err=%s\n", err.Error())
	}

	personRepo := repository.NewPersonRepository(db)
	personService := service.NewPersonService(personRepo)
	personController := http2.NewUserController(instance, personService)

	todoRepo := repository.NewTodoRepository(db)
	todoService := service.NewTodoService(todoRepo)
	todoController := http2.NewTodoController(instance, todoService)

	personController.InitRouter()
	todoController.InitRouter()

	httpServer := http.NewHttpServer(
		instance,
		config.HttpServerConfig{
			Port: 8000,
		},
	)

	httpServer.Start()

	defer func(httpServer http.HttpServer) {
		err := httpServer.Close()

		if err != nil {
			log.Printf("failed to close server %v", err)
		}
	}(httpServer)

	log.Println("Listening signals...")
	c := make(chan os.Signal, 1)
	signal.Notify(
		c,
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGTERM,
	)

	<-c
	log.Println("graceful shutdown")

}
