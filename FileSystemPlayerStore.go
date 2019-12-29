package main

import (
	"io"
)

type FileSystemPlayerStore struct {
	database io.ReadSeeker // Needed to be able to use Seek to go back to the first byte read
}

func (f *FileSystemPlayerStore) GetLeague() []Player {
	f.database.Seek(0, 0)
    league, _ := NewLeague(f.database)
    return league
}