/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/ykkssyaa/DNS_Service/client/cmd"
	"github.com/ykkssyaa/DNS_Service/client/consts"
	"github.com/ykkssyaa/DNS_Service/server/pkg/config"
	"log"
)

func main() {

	err := config.InitConfig(consts.ConfigFilePath)

	if err != nil {
		log.Fatal(err)
	}

	cmd.Execute()
}
