package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/ykkssyaa/DNS_Service/client/consts"
	"github.com/ykkssyaa/DNS_Service/server/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

// hostnameCmd represents the hostname command
var hostnameCmd = &cobra.Command{
	Use:   "hostname",
	Short: "Get the hostname of a machine",
	Long:  `Get the hostname of a machine`,
	Run: func(cmd *cobra.Command, args []string) {

		conn, err := grpc.NewClient(consts.AddrGRPC, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		c := gen.NewDnsServiceClient(conn)

		res, err := c.GetHostname(context.Background(), &gen.Empty{})
		if err != nil {
			log.Fatalf("could not list DNS servers: %v", err)
		}

		fmt.Println(res.Name)
	},
}

func init() {
	rootCmd.AddCommand(hostnameCmd)
}
