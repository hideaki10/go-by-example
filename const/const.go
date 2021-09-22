package main

import "fmt"

type ServiceType string

const (
	ServiceTypeClusterIP ServiceType = "ClusterIP"

	ServiceNodePortService ServiceType = "NodePort"
)

func main() {

	fmt.Printf(" daa  %T", ServiceTypeClusterIP)
	fmt.Printf(" daa  %T", ServiceNodePortService)


	if value , esxit := range map :
}
