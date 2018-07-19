package testing

import (
	"errors"
	"strings"
)

// Song ...
type Song struct {
	ID     string
	Name   string
	Artist string
	Genre  string
}

// BestSongClient ...
type BestSongClient struct {
	songs []string
	dao   DAO
}

func toSong(rec []string) *Song {
	song := new(Song)
	song.ID = rec[0]
	song.Name = rec[1]
	song.Artist = rec[2]
	song.Genre = rec[3]
	return song
}

// GetSong ...
func (client *BestSongClient) GetSong(name string) (*Song, error) {
	recs, err := client.dao.GetSongs()
	if err != nil {
		return nil, errors.New("failed to fetch songs from source")
	}
	for _, rec := range recs {
		if strings.ToLower(rec[1]) == strings.ToLower(name) {
			return toSong(rec), nil
		}
	}
	return new(Song), nil
}

// GetAllSongs ...
func (client *BestSongClient) GetAllSongs() ([]*Song, error) {
	recs, err := client.dao.GetSongs()
	if err != nil {
		return nil, errors.New("failed to fetch songs from source")
	}

	var allSongs []*Song
	for _, rec := range recs {
		allSongs = append(allSongs, toSong(rec))
	}
	return allSongs, nil
}

// NewClient ...
func NewClient(dao DAO) *BestSongClient {
	client := new(BestSongClient)
	if dao == nil {
		client.dao = NewDao()
	} else {
		client.dao = dao
	}
	return client
}
