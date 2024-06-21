/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"github.com/ykkssyaa/DNS_Service/client/consts"
	"github.com/ykkssyaa/DNS_Service/server/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"

	"github.com/spf13/cobra"
)

// removeDnsCmd represents the removeDns command
var removeDnsCmd = &cobra.Command{
	Use:   "removeDns [address]",
	Short: "remove dns address from dns list",
	Long:  `remove dns address from dns list`,
	Run: func(cmd *cobra.Command, args []string) {

		conn, err := grpc.NewClient(consts.AddrGRPC, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		c := gen.NewDnsServiceClient(conn)

		var address string
		if len(args) > 0 {
			address = args[0]
		} else {
			log.Fatalln("there are not dns address in args")
		}

		_, err = c.RemoveDnsServer(context.Background(), &gen.DNS{Address: address})
		if err != nil {
			log.Fatalf("could not list DNS servers: %v", err)
		}

		fmt.Println("DNS server removed from dns list")
	},
}

func init() {
	rootCmd.AddCommand(removeDnsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// removeDnsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// removeDnsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
