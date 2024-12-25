package models

type DosenModel struct {
    DosenID      uint   `json:"dosen_id"`
    DosenName    string `json:"dosen_name"`
    ProdiName    string `json:"prodi_name"`
    FakultasName string `json:"fakultas_name"`
}