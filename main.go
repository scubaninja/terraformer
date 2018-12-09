package main

import (
	"log"
	"os"

	"waze/terraformer/aws_terraforming"
	"waze/terraformer/gcp_terraforming"
)

func main() {
	provider := os.Args[1]
	service := os.Args[2]
	args := []string{}
	if len(os.Args) > 2 {
		args = os.Args[3:]
	}
	var err error
	switch provider {
	case "aws":
		err = aws_terraforming.Generate(service, args)
	case "google":
		err = gcp_terraforming.Generate(service, args)
	}
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
