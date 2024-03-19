package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/hashicorp/consul/api"
)

func main() {
    // Inicialize o roteador do Gin
    router := gin.Default()

    // Defina a rota
    router.GET("/hello", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "Hello, World!",
        })
    })

    // Inicialize o cliente do Consul
    consulConfig := api.DefaultConfig()
    consulClient, err := api.NewClient(consulConfig)
    if err != nil {
        log.Fatal("Failed to create Consul client:", err)
    }

    // Registre o serviço no Consul
    registration := new(api.AgentServiceRegistration)
    registration.ID = "my-service-1"
    registration.Name = "my-service"
    registration.Address = "localhost"
    registration.Port = 8080
    registration.Check = &api.AgentServiceCheck{
        HTTP:     "http://localhost:8080/hello",
        Interval: "10s",
    }

    err = consulClient.Agent().ServiceRegister(registration)
    if err != nil {
        log.Fatal("Failed to register service with Consul:", err)
    }
    defer consulClient.Agent().ServiceDeregister(registration.ID)

    // Inicie o servidor HTTP
    err = router.Run(":8080")
    if err != nil {
        log.Fatal("Failed to start HTTP server:", err)
    }

    fmt.Println("Server running on port 8080")
}