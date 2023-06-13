package main

import (
	"context"
	"flag"
	"log"

	"google.golang.org/grpc"
	"gulnur.net/pkg/api"
)

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		log.Fatal("not enough arguments")
	}

	action := flag.Arg(0)
	con, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	c := api.NewNoteClient(con)

	switch action {
	case "add":
		if flag.NArg() < 3 {
			log.Fatal("not enough arguments for add")
		}
		title := flag.Arg(1)
		description := flag.Arg(2)
		res, err := c.Add(context.Background(), &api.AddRequest{
			Title:       title,
			Description: description,
		})
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Note ID:", res.GetNoteId())

	case "get":
		if flag.NArg() < 2 {
			log.Fatal("not enough arguments for get")
		}
		noteID := flag.Arg(1)
		res, err := c.Get(context.Background(), &api.GetRequest{
			NoteID: noteID,
		})
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Note:", res.GetNote())

	case "update":
		if flag.NArg() < 4 {
			log.Fatal("not enough arguments for update")
		}
		noteID := flag.Arg(1)
		title := flag.Arg(2)
		description := flag.Arg(3)
		res, err := c.Update(context.Background(), &api.UpdateRequest{
			NoteID:   noteID,
			NewTitle: title,
			NewBody:  description,
		})
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Update successful:", res.GetNote())

	case "delete":
		if flag.NArg() < 2 {
			log.Fatal("not enough arguments for delete")
		}
		noteID := flag.Arg(1)
		res, err := c.Delete(context.Background(), &api.DeleteRequest{
			NoteID: noteID,
		})
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Delete successful:", res.GetSuccess())

	default:
		log.Fatal("unknown action:", action)
	}
}
