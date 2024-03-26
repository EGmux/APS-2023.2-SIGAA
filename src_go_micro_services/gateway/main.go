package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/api"
)

func main() {
	// Configuração do cliente Consul
	config := api.DefaultConfig()
	config.Address = os.Getenv("CONSUL_HOST")+":"+os.Getenv("CONSUL_PORT")
	consulClient, err := api.NewClient(config)
	if err != nil {
		log.Fatal("Erro ao criar cliente Consul:", err)
	}

	// Criação de um objeto Gin
	router := gin.Default()

	// Rota para obter um serviço registrado no Consul
	router.GET("/service/:name", func(c *gin.Context) {
		studentUser := c.Query("studentUser")
		serviceName := c.Param("name")
		fmt.Print(serviceName)
		fmt.Println("Student Received: ", studentUser)
		// Consulta o serviço no Consul
		entries, _, err := consulClient.Health().Service(serviceName, "", true, nil)
		if err != nil {
			c.JSON(500, gin.H{"error": "Erro ao consultar serviço no Consul",
								"err": err})
			return
		}

		if len(entries) > 0 {
			service := entries[0].Service
			if service.ID == "scholarship"{
				fmt.Println(service.Address, service.Port, service.ID)
				query := "http://"+service.Address+":"+strconv.Itoa(service.Port)+"/"+service.ID+"?studentUser="+studentUser
				fmt.Println(query)
				c.Redirect(http.StatusMovedPermanently, query)
				return
			}
			/*c.JSON(200, gin.H{
				"service_name": service.Service,
				"address":      service.Address,
				"port":         service.Port,
			}*/
			fmt.Println(service.Address,service.Port,service.ID)
			c.Redirect(http.StatusMovedPermanently, "http://"+service.Address+":"+strconv.Itoa(service.Port)+"/"+service.ID)
			return
		}

		c.JSON(404, gin.H{"error": "Serviço não encontrado no Consul"})
	})

	registerServiceWithConsul()
	// Executa o servidor Gin na porta 8080
	err = router.Run(":8079")
	if err != nil {
		log.Fatal("Erro ao iniciar o servidor Gin:", err)
	}
}

func registerServiceWithConsul(){
	config := api.DefaultConfig()
	config.Address = os.Getenv("CONSUL_HOST")+":"+os.Getenv("CONSUL_PORT")
	fmt.Println(config.Address)

	consulClient, err := api.NewClient(config)
	if err != nil {
		log.Fatal("Error creating Consul client: ", err)
	}

	// Criar um novo serviço
	registration := new(api.AgentServiceRegistration)
	registration.ID = "gateway"
	registration.Name = "gateway"
	registration.Port = 8079 // Porta do serviço

	// Registrar o serviço
	err = consulClient.Agent().ServiceRegister(registration)
	if err != nil {
		log.Fatal("Error registering service with Consul: ", err)
	}

	fmt.Println("Service registered with Consul")


}