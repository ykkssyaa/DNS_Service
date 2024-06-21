/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"github.com/ykkssyaa/DNS_Service/client/consts"
	"github.com/ykkssyaa/DNS_Service/server/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"

	"github.com/spf13/cobra"
)

// addDnsCmd represents the addDns command
var addDnsCmd = &cobra.Command{
	Use:   "addDns [address]",
	Short: "add dns address to dns list",
	Long:  `add dns address to dns list`,
	Run: func(cmd *cobra.Command, args []string) {

		conn, err := grpc.NewClient(consts.AddrGRPC+viper.GetString("ports.grpc"),
			grpc.WithTransportCredentials(insecure.NewCredentials()))
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

		_, err = c.AddDnsServer(context.Background(), &gen.DNS{Address: address})
		if err != nil {
			log.Fatalf("could not list DNS servers: %v", err)
		}

		fmt.Println("DNS server added to dns list")
	},
}

func init() {
	rootCmd.AddCommand(addDnsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addDnsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addDnsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
