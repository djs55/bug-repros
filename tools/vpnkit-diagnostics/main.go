package main

import (
	"os"
	"context"
	"io"
	"log"
	"time"

	"github.com/Microsoft/go-winio"

)

const socket =  `\\.\pipe\dockerVpnKitDiagnostics`

func main(){

	f, err := os.Create("vpnkit.tar")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	c, err := winio.DialPipeContext(ctx, socket)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	if err := c.SetDeadline(time.Now().Add(10 * time.Second)); err != nil {
		log.Fatal(err)
	}
	if _, err := io.Copy(f, c); err != nil {
		log.Fatal(err)
	}
}
