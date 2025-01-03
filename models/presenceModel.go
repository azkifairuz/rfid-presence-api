package models

import (
	"time"
	"gorm.io/gorm"
)

// Fakultas model
type Fakultas struct {
	gorm.Model
	ID   int    `gorm:"primaryKey"`
	Name string `gorm:"type:varchar(100);not null;unique"`
}

// Prodi model
type Prodi struct {
	gorm.Model
	ID        int    `gorm:"primaryKey"`
	Name      string `gorm:"type:varchar(100);not null;unique"`
	FakultasID int    `gorm:"not null"`
	Fakultas  Fakultas `gorm:"foreignKey:FakultasID"`
}

// Dosen model
type Dosen struct {
	gorm.Model
	ID       int    `gorm:"primaryKey"`
	Nip      string `gorm:"type:varchar(20);not null;unique"`
	Name     string `gorm:"type:varchar(100);not null"`
	ProdiID   int    `gorm:"not null"`
	Prodi     Prodi  `gorm:"foreignKey:ProdiID"`
}

// Mahasiswa model
type Mahasiswa struct {
	gorm.Model

	ID        int    `gorm:"primaryKey"`
	Nim       string `gorm:"type:varchar(20);not null;unique"`
	Name      string `gorm:"type:varchar(100);not null"`
	KelasID   int    `gorm:"not null"`
	Kelas     Kelas  `gorm:"foreignKey:KelasID"`
	Presensi  []Presensi `gorm:"foreignKey:MahasiswaID"`
}

// Kelas model
type Kelas struct {
	gorm.Model

	ID      int    `gorm:"primaryKey"`
	Name    string `gorm:"type:varchar(100);not null"`
	ProdiID int    `gorm:"not null"`
	Prodi   Prodi  `gorm:"foreignKey:ProdiID"`
}

// Jadwal model
type Jadwal struct {
	gorm.Model

	ID      int       `gorm:"primaryKey"`
	Hour    string `gorm:"not null" json:"hour"`
	Matkul  string 	  `gorm:"not null"`
	Days    string    `gorm:"type:varchar(50);not null"`
	KelasID int       `gorm:"not null"`
	Kelas   Kelas     `gorm:"foreignKey:KelasID"`
	DosenID int       `gorm:"not null"`
	Dosen   Dosen     `gorm:"foreignKey:DosenID"`
}

// Account model
type Account struct {
	gorm.Model

	ID       int        `gorm:"primaryKey"`
	UserID   int        `gorm:"not null"`
	Email    string `gorm:"not null"`
	Password    string  `gorm:"not null"`
	AccountType AccountType `gorm:"type:varchar(50);not null"`
}


// Enum untuk tipe account
type AccountType string

const (
	DosenType  AccountType = "dosen"
	StaffType  AccountType = "staff"
)

// Staf model
type Staf struct {
	gorm.Model

	ID        int       `gorm:"primaryKey"`
	Nip		string		`gorm:"type:varchar(20);not null;unique"`
	Name      string    `gorm:"type:varchar(100);not null"`
	FakultasID int      `gorm:"not null"`
	Fakultas  Fakultas `gorm:"foreignKey:FakultasID"`
}

// Presensi model
type Presensi struct {
	gorm.Model

	ID        int      `gorm:"primaryKey"`
	MahasiswaID int    `gorm:"not null"`
	KelasID   int      `gorm:"not null"`
	Date      time.Time `gorm:"not null"`
	Mahasiswa Mahasiswa `gorm:"foreignKey:MahasiswaID"`
	Kelas     Kelas     `gorm:"foreignKey:KelasID"`
}

type MhsCard struct {
	ID int `gorm:"primaryKey"`
	CardUid string `gorm:"not null"`
	MahasiswaID int    `gorm:"not null"`
	Mahasiswa Mahasiswa `gorm:"foreignKey:MahasiswaID"`

}
