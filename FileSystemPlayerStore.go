package main

import (
	"io"
)

type FileSystemPlayerStore struct {
	database io.ReadWriteSeeker // Needed to be able to use Seek to go back to the first byte read
}

func (f *FileSystemPlayerStore) GetLeague() []Player {
	f.database.Seek(0, 0)
    league, _ := NewLeague(f.database)
    return league
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
    var wins int

    for _, player := range f.GetLeague() {
        if player.Name == name {
            wins = player.Wins
            break
        }
    }

    return wins
}

func (f *FileSystemPlayerStore) RecordWin(name string) {

}