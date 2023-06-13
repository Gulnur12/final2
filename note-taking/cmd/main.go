package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"gulnur.net/pkg/api"
	"gulnur.net/pkg/note"
)

func main(){
	s := grpc.NewServer()
	srv := note.NewNoteServer()
	api.RegisterNoteServer(s, srv)

	l, err := net.Listen("tcp", ":8080")
	if err != nil{
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil{
		log.Fatal(err)
	}
}