package main

import (
	"github.com/XWS-BSEP-Tim-13/Dislinkt_CompanyService/startup"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_CompanyService/startup/config"
)

func main() {
	config := config.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
