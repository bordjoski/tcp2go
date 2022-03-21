package cmd

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"

	"github.com/spf13/cobra"
)

var port int
var host string

// proxyCmd represents the proxy command
var proxyCmd = &cobra.Command{
	Use:   "proxy",
	Short: "Simple tcp proxy",
	Long:  `Simple TCP proxy to bypass firewall rules`,
	Run:   runProxy,
}

func handle(src net.Conn) {
	dst, err := net.Dial("tcp", host)
	if err != nil {
		log.Fatalln("Unreachable host")
	}
	defer dst.Close()

	go func() {
		if _, err := io.Copy(dst, src); err != nil {
			log.Fatalln(err)
		}
	}()
	if _, err := io.Copy(src, dst); err != nil {
		log.Fatalln(err)
	}
}

func runProxy(cmd *cobra.Command, args []string) {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalln("Unable to bind to port", port)
		os.Exit(1)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatalln("Can not accept connection")
		}
		log.Println("Accepting connection...")
		go handle(conn)
	}
}

func init() {
	rootCmd.AddCommand(proxyCmd)
	proxyCmd.Flags().IntVarP(&port, "port", "p", 3001, "Listen on given port")
	proxyCmd.Flags().StringVarP(&host, "target", "t", "", "Host. Eg example.com:8080")
	proxyCmd.MarkFlagRequired("target")
}
