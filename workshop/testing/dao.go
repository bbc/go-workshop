package testing

import (
	"bufio"
	"errors"
	"math/rand"
	"os"
	"strings"
	"time"
)

const path = "../fixtures/songs.txt"

// DO NOT CHANGE THIS CODE

// DAO ...
type DAO interface {
	GetSongs() ([][]string, error)
}

// SongDao ...
type SongDao struct {
	path string
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// GetSongs ...
func (dao *SongDao) GetSongs() ([][]string, error) {
	fh, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer fh.Close()

	if (rand.Intn(10))%2 != 0 {
		return nil, errors.New("failed to fetch from source")
	}

	time.Sleep(10000 * time.Millisecond)

	songs := [][]string{}
	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		data := scanner.Text()
		rec := strings.Split(data, "|")
		songs = append(songs, rec)
	}
	return songs, nil
}

// NewDao ..
func NewDao() *SongDao {
	dao := new(SongDao)
	dao.path = path
	return dao
}
