package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/hashicorp/consul/api"
)

func main() {

	//Configure COnsul Client
	config := api.DefaultConfig()
	client, err := api.NewClient(config)
	if err != nil{
		fmt.Println("err")
		log.Fatal("Error during Client Creation")
	}

	registration := &api.AgentServiceRegistration{
		ID: "aps-login",
		Name: "aps-login",
		Port: 8080,
	}


	//Register Service on Consul
	err = client.Agent().ServiceRegister(registration)
	if err != nil{
		fmt.Println(err)
		log.Fatal("Error during Client Registration")
	}
	fmt.Println("Login Service registered on Consul")

	//After the instantiation, we can run up ours services http requests
	
	//Keep aplication running until it is Unregistred
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
    <-sigs

	//Unregister service at Consul
	err = client.Agent().ServiceDeregister("aps-login")
	if err != nil{
		fmt.Println(err)
		log.Fatal("Error during service Unregistration")
	}
	fmt.Println("Service unregistred from Consul")


}