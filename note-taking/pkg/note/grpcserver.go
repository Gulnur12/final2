package note

import (
	"context"
	"fmt"
	"strconv"

	"gulnur.net/pkg/api"
)

type NoteServer struct {
	noteStore  map[int32]*api.Note
	lastNoteID int32
}

func NewNoteServer() *NoteServer {
	return &NoteServer{
		noteStore:  make(map[int32]*api.Note),
		lastNoteID: 0,
	}
}

func (s *NoteServer) Add(ctx context.Context, req *api.AddRequest) (*api.AddResponse, error) {
	noteID := s.generateNoteID()
	note := &api.Note{
		Id:          noteID,
		Title:       req.Title,
		Description: req.Description,
	}
	s.noteStore[noteID] = note

	return &api.AddResponse{NoteId: noteID}, nil
}

func (s *NoteServer) generateNoteID() int32 {
	s.lastNoteID++
	return s.lastNoteID
}

func (s *NoteServer) Get(ctx context.Context, req *api.GetRequest) (*api.GetResponse, error) {
	noteID, err := strconv.Atoi(req.NoteID)
	if err != nil {
		return nil, fmt.Errorf("invalid note ID")
	}

	note, exists := s.noteStore[int32(noteID)]
	if !exists {
		return nil, fmt.Errorf("note not found")
	}

	return &api.GetResponse{Note: note}, nil
}

func (s *NoteServer) Update(ctx context.Context, req *api.UpdateRequest) (*api.UpdateResponse, error) {
	noteID, err := strconv.Atoi(req.NoteID)
	if err != nil {
		return nil, fmt.Errorf("invalid note ID")
	}
	note, exists := s.noteStore[int32(noteID)]
	if !exists {
		return nil, fmt.Errorf("note not found")
	}

	note.Title = req.NewTitle
	note.Description = req.NewBody

	return &api.UpdateResponse{Note: note}, nil
}

func (s *NoteServer) Delete(ctx context.Context, req *api.DeleteRequest) (*api.DeleteResponse, error) {
	noteID, err := strconv.Atoi(req.NoteID)
	if err != nil {
		return nil, fmt.Errorf("invalid note ID")
	}

	_, exists := s.noteStore[int32(noteID)]
	if !exists {
		return nil, fmt.Errorf("note not found")
	}

	delete(s.noteStore, int32(noteID))

	return &api.DeleteResponse{Success: true}, nil
}
