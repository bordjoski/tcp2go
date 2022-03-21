/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"

	"github.com/spf13/cobra"
)

// connectCmd represents the connect command
var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Client.",
	Long:  `Execute arbitary commands on remote server.`,
	Run:   clientConnect,
}

type Reader struct{}
type Writer struct{}

func (fooReader *Reader) Read(b []byte) (int, error) {
	return os.Stdin.Read(b)
}

// Write writes data to Stdout.
func (fooWriter *Writer) Write(b []byte) (int, error) {
	fmt.Print("\n")
	return os.Stdout.Write(b)
}

func clientConnect(cmd *cobra.Command, args []string) {
	var (
		reader Reader
		writer Writer
	)

	connection, err := net.Dial("tcp", host)

	if err != nil {
		fmt.Println("Can not connect to host")
	}

	fmt.Println("Connected to", host)

	go func() {
		if _, err := io.Copy(connection, &reader); err != nil {
			log.Fatalln("Unable to read/write data")
		}
	}()
	if _, err := io.Copy(&writer, connection); err != nil {
		log.Fatalln("Unable to read/write data")
	}
}

func init() {
	rootCmd.AddCommand(connectCmd)
	connectCmd.Flags().StringVarP(&host, "target", "t", "", "Host. Eg example.com:8080")
}
