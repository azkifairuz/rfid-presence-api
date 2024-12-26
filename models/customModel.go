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
    MahasiswaNim    string `json:"nim"`
    MahasiswaClass    string `json:"kelas_name"`
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

type CardResponse struct {
    Uid string `json:"uid"`
    MahasiswaName    string `json:"mahasiswa_name"`
    MahasiswaClass    string `json:"kelas"`
    MahasiswaNim    string `json:"nim"`
    MahasiswaId    int `json:"mahasiswa_id"`
    KelasId    int `json:"kelas_id"`

    
}