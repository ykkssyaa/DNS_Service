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

// setHostnameCmd represents the setHostname command
var setHostnameCmd = &cobra.Command{
	Use:   "setHostname",
	Short: "Set new hostname on a machine",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		conn, err := grpc.NewClient(consts.AddrGRPC, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		c := gen.NewDnsServiceClient(conn)

		res, err := c.ListDnsServers(context.Background(), &gen.Empty{})
		if err != nil {
			log.Fatalf("could not list DNS servers: %v", err)
		}

		for _, addr := range res.Addresses {
			fmt.Println(addr)
		}
	},
}

func init() {
	rootCmd.AddCommand(setHostnameCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setHostnameCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setHostnameCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
