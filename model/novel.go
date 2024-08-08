package model

type Novel struct {
	Id          string `json:"id"`
	Judul       string `json:"judul"`
	Penerbit    string `json:"penerbit"`
	TahunTerbit string `json:"tahunTerbit"`
	Penulis     string `json:"penulis"`
}
