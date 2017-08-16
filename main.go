package main

import (
	"github.com/google/logger"
	"github.com/selvakn/libzt"
	"gopkg.in/alecthomas/kingpin.v2"
	"io"
	"os"
	"p2p-port-forward/client"
	"p2p-port-forward/utils"
	"p2p-port-forward/server"
)

var (
	network     = kingpin.Flag("network", "zerotier network id").Short('n').Default("8056c2e21c000001").String()
	forwardPort = kingpin.Flag("forward-port", "port to forward (in listen mode)").Short('f').Default("22").String()
	acceptPort  = kingpin.Flag("accept-port", "port to accept (in connect mode)").Short('a').Default("2222").String()
	useUDP      = kingpin.Flag("use-udp", "UDP instead of TCP (TCP default)").Short('u').Default("false").Bool()

	connectTo = kingpin.Flag("connect-to", "server (zerotier) ip to connect").Short('c').String()
)

func main() {
	logger.Init("p2p-port-forward", false, false, os.Stdout)

	kingpin.Version("1.0.1")
	kingpin.Parse()

	zt := libzt.Init(*network, "./zt")

	logger.Infof("ipv4 = %v \n", zt.GetIPv4Address().String())
	logger.Infof("ipv6 = %v \n", zt.GetIPv6Address().String())

	var closableConn io.Closer

	if len(*connectTo) == 0 {
		forwarderServer := server.New(zt, *forwardPort, utils.GetIPProto(*useUDP))
		closableConn = forwarderServer.Listen()
	} else {
		forwarderClient := client.New(zt, *connectTo, *acceptPort, utils.GetIPProto(*useUDP))
		closableConn = forwarderClient.ListenAndSync()
	}

	<-utils.SetupCleanUpOnInterrupt(func() {
		if closableConn != nil {
			closableConn.Close()
		}
	})

}
