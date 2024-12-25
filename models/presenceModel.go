package models

import (
	"time"
	"gorm.io/gorm"
)

// Fakultas model
type Fakultas struct {
	gorm.Model
	ID   int    `gorm:"primaryKey"`
	Name string `gorm:"type:varchar(100);not null"`
}

// Prodi model
type Prodi struct {
	gorm.Model
	ID        int    `gorm:"primaryKey"`
	Name      string `gorm:"type:varchar(100);not null"`
	FakultasID int    `gorm:"not null"`
	Fakultas  Fakultas `gorm:"foreignKey:FakultasID"`
}

// Dosen model
type Dosen struct {
	gorm.Model
	ID       int    `gorm:"primaryKey"`
	NIP      string `gorm:"type:varchar(20);not null;unique"`
	Name     string `gorm:"type:varchar(100);not null"`
	FakultasID int    `gorm:"not null"`
	Fakultas  Fakultas `gorm:"foreignKey:FakultasID"`
	Accounts []Account `gorm:"foreignKey:UserID"`
	Jadwals  []Jadwal  `gorm:"foreignKey:DosenID"`
}

// Mahasiswa model
type Mahasiswa struct {
	gorm.Model

	ID        int    `gorm:"primaryKey"`
	NIM       string `gorm:"type:varchar(20);not null;unique"`
	Name      string `gorm:"type:varchar(100);not null"`
	ProdiID   int    `gorm:"not null"`
	Prodi     Prodi  `gorm:"foreignKey:ProdiID"`
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
	Hour    time.Time `gorm:"not null"`
	Days    string    `gorm:"type:varchar(50);not null"`
	Date	time.Time
	ProdiID int       `gorm:"not null"`
	Prodi   Prodi     `gorm:"foreignKey:ProdiID"`
	DosenID int       `gorm:"not null"`
	Dosen   Dosen     `gorm:"foreignKey:DosenID"`
	Presensi []Presensi `gorm:"foreignKey:JadwalID"`
}

// Account model
type Account struct {
	gorm.Model

	ID       int        `gorm:"primaryKey"`
	UserID   int        `gorm:"not null;unique"`
	Email    string
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
	Name      string    `gorm:"type:varchar(100);not null"`
	FakultasID int      `gorm:"not null"`
	Fakultas  Fakultas `gorm:"foreignKey:FakultasID"`
	Accounts []Account `gorm:"foreignKey:UserID"`
}

// Presensi model
type Presensi struct {
	gorm.Model

	ID        int      `gorm:"primaryKey"`
	MahasiswaID int    `gorm:"not null"`
	DosenID   int      `gorm:"not null"`
	JadwalID  int      `gorm:"not null"`
	Date      time.Time `gorm:"not null"`
	Mahasiswa Mahasiswa `gorm:"foreignKey:MahasiswaID"`
	Dosen     Dosen     `gorm:"foreignKey:DosenID"`
	Jadwal    Jadwal    `gorm:"foreignKey:JadwalID"`
}