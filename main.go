package main

import (
	"flag"
	"io"
	"os"
	"ztfw/client"
	"ztfw/logger"
	"ztfw/server"
	"ztfw/utils"

	"ztfw/libzt"
)

var (
	network = *flag.String("n", getEnv("ZTFW_NETWORK", "8056c2e21c000001"), "zerotier network id")
	forward = *flag.String("f", getEnv("ZTFW_FORWARD", "127.0.0.1:22"), "forward target in listen mode")
	listen  = *flag.String("l", getEnv("ZTFW_LISTEN", "0.0.0.0:2222"), "port to listen on local in connect mode")
	useUDP  = *flag.Bool("u", getEnv("ZTFW_UDP", "") != "", "udp instead of tcp")

	connect = *flag.String("c", getEnv("ZTFW_SERVER", ""), "zerotier server ip")
)

var log = logger.Logger

func main() {
	flag.Parse()

	zt := libzt.Init(network, "./zt-home")

	log.Infof("ipv4 = %v ", zt.GetIPv4Address().String())
	log.Infof("ipv6 = %v ", zt.GetIPv6Address().String())

	var closableConn io.Closer

	if len(connect) == 0 {
		forwarderServer := server.New(zt, forward, utils.GetIPProto(useUDP))
		closableConn = forwarderServer.Listen()
	} else {
		forwarderClient := client.New(zt, connect, listen, utils.GetIPProto(useUDP))
		closableConn = forwarderClient.ListenAndSync()
	}

	<-utils.SetupCleanUpOnInterrupt(func() {
		if closableConn != nil {
			closableConn.Close()
		}
	})

}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
