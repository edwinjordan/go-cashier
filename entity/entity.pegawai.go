package entity

import "time"

type Pegawai struct {
	PegawaiId          string    `json:"pegawai_id"`
	PegawaiCode        string    `json:"pegawai_code"`
	PegawaiName        string    `json:"pegawai_name"`
	PegawaiGender      string    `json:"pegawai_gender"`
	PegawaiPhonenumber string    `json:"pegawai_phonenumber"`
	PegawaiCreateAt    time.Time `json:"-"`
	PegawaiUpdateAt    time.Time `json:"-"`
	PegawaiDeleteAt    time.Time `json:"-"`
}
