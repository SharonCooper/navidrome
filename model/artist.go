package model

import "time"

type Artist struct {
	ID              string `json:"id"          orm:"column(id)"`
	Name            string `json:"name"`
	AlbumCount      int    `json:"albumCount"`
	SongCount       int    `json:"songCount"`
	FullText        string `json:"fullText"`
	SortArtistName  string `json:"sortArtistName"`
	OrderArtistName string `json:"orderArtistName"`

	// Annotations
	PlayCount int64     `json:"playCount"   orm:"-"`
	PlayDate  time.Time `json:"playDate"    orm:"-"`
	Rating    int       `json:"rating"      orm:"-"`
	Starred   bool      `json:"starred"     orm:"-"`
	StarredAt time.Time `json:"starredAt"   orm:"-"`
}

type Artists []Artist

type ArtistIndex struct {
	ID      string
	Artists Artists
}
type ArtistIndexes []ArtistIndex

type ArtistRepository interface {
	CountAll(options ...QueryOptions) (int64, error)
	Exists(id string) (bool, error)
	Put(m *Artist) error
	Get(id string) (*Artist, error)
	GetStarred(options ...QueryOptions) (Artists, error)
	Search(q string, offset int, size int) (Artists, error)
	Refresh(ids ...string) error
	GetIndex() (ArtistIndexes, error)
	AnnotatedRepository
}
