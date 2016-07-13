// +build linux darwin freebsd

package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/xtaci/kcp-go"
)

func init() {
	go sig_handler()
}

func sig_handler() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGUSR1)

	for {
		switch <-ch {
		case syscall.SIGUSR1:
			log.Printf("KCP SNMP:%+v", kcp.DefaultSnmp.Copy())
		}
	}
}
