package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	ping "github.com/NaddiNadja/peer-to-peer/grpc"
	"google.golang.org/grpc"
)

func main() {

	// If the file doesn't exist, create it or append to the file
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)

	arg1, _ := strconv.ParseInt(os.Args[1], 10, 32)
	ownPort := int32(arg1) + 5000

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if ownPort == 5000 {
		log.Println("\n \n \n Start of log file for given communication of clients:")
	}

	p := &peer{
		id:            ownPort,
		amountOfPings: make(map[int32]int32),
		clients:       make(map[int32]ping.PingClient),
		ctx:           ctx,
		wanted:        false,
		held:          false,
		timesAccessed: 0,
	}

	// Create listener tcp on port ownPort
	list, err := net.Listen("tcp", fmt.Sprintf(":%v", ownPort))
	if err != nil {
		log.Fatalf("Failed to listen on port: %v", err)
	}
	grpcServer := grpc.NewServer()
	ping.RegisterPingServer(grpcServer, p)

	go func() {
		if err := grpcServer.Serve(list); err != nil {
			log.Fatalf("failed to server %v", err)
		}
	}()

	for i := 0; i < 3; i++ {
		port := int32(5000) + int32(i)

		if port == ownPort {
			continue
		}

		var conn *grpc.ClientConn
		fmt.Printf("Trying to dial: %v\n", port)
		conn, err := grpc.Dial(fmt.Sprintf(":%v", port), grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("Could not connect: %s", err)
		}
		defer conn.Close()
		c := ping.NewPingClient(conn)
		p.clients[port] = c
	}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		log.Printf("%v is trying to get permission to send its message", p.id)
		p.askPermission(scanner.Text())
	}
}

type peer struct {
	ping.UnimplementedPingServer
	id            int32
	amountOfPings map[int32]int32
	clients       map[int32]ping.PingClient
	ctx           context.Context
	wanted        bool
	held          bool
	timesAccessed int32
}

func (p *peer) ReturnInfo(ctx context.Context, req *ping.Request) (*ping.ReturnInfoReply, error) {

	rep := &ping.ReturnInfoReply{Id: p.id, TimesAccessed: p.timesAccessed, Wanted: p.wanted, Held: p.held}
	return rep, nil
}

func (p *peer) askPermission(message string) {
	p.wanted = true
	permission := false

	request := &ping.Request{Id: p.id}
	for id, client := range p.clients {
		log.Printf("%v is asking permission from id %v\n", p.id, id)
		returnInfoReply, err := client.ReturnInfo(p.ctx, request)
		if err != nil {
			fmt.Println("something went wrong")
		}
		if returnInfoReply.Held {
			permission = false
			log.Printf("%v Did not get permission from id %v\n", p.id, id)
			break
		} else if returnInfoReply.Wanted {
			if returnInfoReply.TimesAccessed < p.timesAccessed {
				permission = false
				log.Printf("%v Did not get permission from id %v\n", p.id, id)
				break
			} else if returnInfoReply.TimesAccessed > p.timesAccessed {
				permission = true
				log.Printf("%v Got permission from id %v\n", p.id, id)
			} else if returnInfoReply.TimesAccessed == p.timesAccessed {
				if returnInfoReply.Id < p.id {
					permission = false
					log.Printf("%v Did not get permission from id %v\n", p.id, id)
					break
				} else if returnInfoReply.Id > p.id {
					permission = true
					log.Printf("%v Got permission from id %v\n", p.id, id)
				}
			}
		} else if !returnInfoReply.Wanted {
			permission = true
		}
	}

	if permission {
		p.held = true
		log.Printf("%v Got permission from all, trying to write", p.id)
		p.criticalWriting(message)
	} else if !permission {
		log.Printf("%v Did not get permission from all, retrying", p.id)
		time.Sleep(1 * time.Second)
		p.askPermission(message)
		return
	}

	p.wanted = false
	p.held = false
	p.timesAccessed++
}

func (p *peer) criticalWriting(message string) {
	messageParts := strings.Split(message, " ")
	log.Printf("%v is trying to write: "+message, p.id)
	for part := range messageParts {
		time.Sleep(1 * time.Second)
		log.Print(messageParts[part] + " ")
	}

}
