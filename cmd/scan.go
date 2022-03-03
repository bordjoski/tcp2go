package cmd

import (
	"fmt"
	"math/rand"
	"net"
	"sort"
	"time"

	"github.com/spf13/cobra"
	"honnef.co/go/netdb"
)

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan for opet ports",
	Long:  `Scan for open ports for given target. It will also show service version if known.`,
	Run:   scan,
}

func worker(target string, ports, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("%s:%d", target, p)
		var dialer net.Dialer
		dialer.Timeout = time.Second * 3
		conn, err := dialer.Dial("tcp", address)
		if err != nil {
			results <- 0
			continue
		}

		conn.Close()
		time.Sleep(time.Duration(50) + time.Millisecond*time.Duration(rand.Int31n(250)))
		results <- p
	}
}

func scan(cmd *cobra.Command, args []string) {
	start := time.Now()
	fmt.Println("Scan started at", start)
	ports := make(chan int, 250)
	results := make(chan int)
	var openports []int

	for i := 0; i < cap(ports); i++ {
		go worker(args[0], ports, results)
	}

	go func() {
		for i := 1; i <= 65535; i++ {
			ports <- i
		}
	}()

	for i := 1; i < 65535; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}

	sort.Ints(openports)

	var proto *netdb.Protoent
	var serv *netdb.Servent

	for _, port := range openports {
		proto = netdb.GetProtoByName("tcp")
		serv = netdb.GetServByPort(port, proto)
		fmt.Println("---------------------------------------")
		if serv != nil {
			fmt.Println("Service name : ", serv.Name)
			fmt.Println("Service is using port number : ", serv.Port)
		} else {
			fmt.Println("Service name unknown")
			fmt.Println("Service is using port number : ", port)
		}
	}
	close(ports)
	close(results)
	duration := time.Since(start)
	fmt.Println("---------------------------------------")
	fmt.Println("Scan completed in", duration)
}

func init() {
	rootCmd.AddCommand(scanCmd)
}
