package models

type DosenModel struct {
    DosenID      uint   `json:"dosen_id"`
    DosenName    string `json:"dosen_name"`
    ProdiID      uint   `json:"prodi_id"`
    ProdiName    string `json:"prodi_name"`
    FakultasID      uint   `json:"fakultas_id"`
    FakultasName string `json:"fakultas_name"`


}
type MhsModel struct {
    MahasiswaID      uint   `json:"mahasiswa_id"`
    MahasiswaName    string `json:"mahasiswa_name"`
    ProdiID      uint   `json:"prodi_id"`
    ProdiName    string `json:"prodi_name"`
    FakultasID      uint   `json:"fakultas_id"`
    FakultasName string `json:"fakultas_name"`
}

type JadwalModel struct {
    JadwalId uint  `json:"jadwalId"`
    Matkul string  `json:"matkul"`
    Kelas string    `json:"kelas"`
    Dosen string    `json:"dosen"`
    Prodi string    `json:"prodi"`
    Fakultas string    `json:"fakultas"`
}