// Copyright 2012-2016 Apcera Inc. All rights reserved.
// +build ignore

package main

import (
	"flag"
	"log"
	"runtime"

	"github.com/nats-io/go-nats"
)

// NOTE: Use tls scheme for TLS, e.g. nats-rply -s tls://demo.nats.io:4443 foo hello
func usage() {
	log.Fatalf("Usage: nats-rply [-s server][-t] <subject>\n")
}

func printMsg(m *nats.Msg, i int) {
	log.Printf("[#%d] Received on [%s]: '%s'\n", i, m.Subject, string(m.Data))
}

func checkMsg(m *nats.Msg) (string) {
	if len(string(m.Data)) > 10 {
		return "OK"
	}  else {
		return "NG"
	}
}
func main() {
	//var urls = flag.String("s", nats.DefaultURL, "The nats server URLs (separated by comma)")
	var showTime = flag.Bool("t", false, "Display timestamps")

	log.SetFlags(0)
	flag.Usage = usage
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		usage()
	}
	var servers = "nats://nats_a:4222, nats://nats_a:4223"

	nc, err := nats.Connect(servers)
	if err != nil {
		log.Fatalf("Can't connect: %v\n", err)
	}

	subj, reply, i := args[0], "", 0

	nc.Subscribe(subj, func(msg *nats.Msg) {
		i++
		printMsg(msg, i)
		reply = checkMsg(msg)
		nc.Publish(msg.Reply, []byte(reply))
	})
	nc.Flush()

	if err := nc.LastError(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Listening on [%s]\n", subj)
	if *showTime {
		log.SetFlags(log.LstdFlags)
	}

	runtime.Goexit()
}
