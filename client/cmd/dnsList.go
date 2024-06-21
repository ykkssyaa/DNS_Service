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

// dnsListCmd represents the dnsList command
var dnsListCmd = &cobra.Command{
	Use:   "dnsList",
	Short: "Get a dns list of a machine",
	Long:  `Get a dns list of a machine`,
	Run: func(cmd *cobra.Command, args []string) {

		conn, err := grpc.NewClient(consts.AddrGRPC+viper.GetString("ports.grpc"), grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		c := gen.NewDnsServiceClient(conn)

		res, err := c.ListDnsServers(context.Background(), &gen.Empty{})
		if err != nil {
			log.Fatalf("could not list DNS servers: %v", err)
		}

		for _, address := range res.Addresses {
			fmt.Println(address)
		}
	},
}

func init() {
	rootCmd.AddCommand(dnsListCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dnsListCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dnsListCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
