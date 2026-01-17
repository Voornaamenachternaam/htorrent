package v1

type Info struct {
	Name         string `json:"name"`
	InfoHash     string `json:"infohash"`
	Description  string `json:"description"`
	Files        []File `json:"files"`
	CreationDate int64  `json:"creationDate"`
}

type File struct {
	Path   string `json:"path"`
	Length int64  `json:"length"`
}

type TorrentMetrics struct {
	Magnet   string        `json:"magnet"`
	InfoHash string        `json:"infohash"`
	Files    []FileMetrics `json:"files"`
	Peers    int           `json:"peers"`
}

type FileMetrics struct {
	Path      string `json:"path"`
	Length    int64  `json:"length"`
	Completed int64  `json:"completed"`
}
