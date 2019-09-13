package model

// Song is the Json
type Song struct {
	ID      string `json:"song_id"`
	Name    string `json:"song_name"`
	Artist  string `json:"song_artist"`
	Link    string `json:"song_link"`
	Genre   string `json:song_genre`
	Album   string `json:song_album`
	Release string `json:song_releasedate`
}
