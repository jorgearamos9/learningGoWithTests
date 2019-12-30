package main

import (
	"encoding/json"
	"os"
	"fmt"
)

type FileSystemPlayerStore struct {
	database *json.Encoder // Needed to be able to use Seek to go back to the first byte read
	league   League
}

func initialisePlayerDBFile(file *os.File) error {
    file.Seek(0, 0)

    info, err := file.Stat()

    if err != nil {
        return fmt.Errorf("problem getting file info from file %s, %v", file.Name(), err)
    }

    if info.Size()==0 {
        file.Write([]byte("[]"))
        file.Seek(0, 0)
    }

    return nil
}

func NewFileSystemPlayerStore(file *os.File) (*FileSystemPlayerStore, error) {
	
	err := initialisePlayerDBFile(file)

    if err != nil {
        return nil, fmt.Errorf("problem initialising player db file, %v", err)
    }
	
	league, err := NewLeague(file)
	if err != nil {
        return nil, fmt.Errorf("NewFileSystemPlayerStore: problem loading player store from file %s, %v", file.Name(), err)
    }
	return &FileSystemPlayerStore{
		database: json.NewEncoder(&tape{file}),
		league:   league,
	}, nil
}

func (f *FileSystemPlayerStore) GetLeague() League {
	return f.league
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	player := f.league.Find(name)

	if player != nil {
		return player.Wins
	}

	return 0
}

func (f *FileSystemPlayerStore) RecordWin(name string) {
	player := f.league.Find(name)

	if player != nil {
		player.Wins++
	} else {
		f.league = append(f.league, Player{name, 1})
	}
	// f.database.Seek(0, 0)
	f.database.Encode(f.league)
}
