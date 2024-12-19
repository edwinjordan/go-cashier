package config

type Message struct {
	SuccessGetData       string
	SuccessCreateData    string
	SuccessUpdateData    string
	SuccessDeleteData    string
	GetDataByIdNotFound  string
	SuccessLogin         string
	ErrorLogin           string
	ErrorCreateDataExist string
	SuccessImportData    string
	FormatFileError      string
}

func LoadMessage() *Message {
	return &Message{
		SuccessGetData:       "Berhasil mendapatkan data",
		SuccessCreateData:    "Berhasil menambah data",
		SuccessUpdateData:    "Berhasil mengubah data",
		SuccessDeleteData:    "Berhasil menghapus data",
		GetDataByIdNotFound:  "Data tidak ditemukan",
		SuccessLogin:         "Berhasil masuk",
		ErrorLogin:           "Username atau password tidak sesuai.",
		ErrorCreateDataExist: "Data tidak bisa ditambahkan karena sudah ada.",
		FormatFileError:      "Format file yang anda masukkan salah.",
		SuccessImportData:    "Data berhasil diimport, silahkan menunggu notitikasi selanjutnya hingga process selesai.",
	}
}
