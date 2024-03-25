package domain

type Header struct {
	PrevHash string `json:"prevHash"`
	Time     int64  `json:"time"`
}
