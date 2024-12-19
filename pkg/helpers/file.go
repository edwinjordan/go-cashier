package helpers

import (
	"bytes"
	"encoding/base64"
	"errors"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"strings"

	"github.com/jolebo/e-canteen-cashier-api/config"
)

/*
	fileExt : png,jpg,pdf
	extAllowed : png/ png|jpg
*/

func FileUploadFormat(fileExt string, extAllowed string) (bool, error) {
	allExt := ""
	formatAllowed := strings.Split(extAllowed, "|")
	for i, v := range formatAllowed {
		allExt += "." + v
		if i != len(formatAllowed)-1 {
			allExt += ","

		}
		if fileExt == v {
			return true, nil
		}
	}
	return false, errors.New(config.LoadMessage().FormatFileError + " Hanya masukkan format " + allExt)
}

func SaveImageFromBase64(fName string, data string) error {
	idx := strings.Index(data, ";base64,")
	if idx < 0 {
		return errors.New("invalid files")
	}

	ImageType := data[11:idx]

	unbased, err := base64.StdEncoding.DecodeString(data[idx+8:])
	PanicIfError(err)
	r := bytes.NewReader(unbased)
	switch ImageType {
	case "png":
		im, err := png.Decode(r)
		PanicIfError(err)

		f, err := os.OpenFile("./uploaded_files/"+fName, os.O_WRONLY|os.O_CREATE, 0777)
		PanicIfError(err)

		png.Encode(f, im)
	case "jpeg":
		im, err := jpeg.Decode(r)
		PanicIfError(err)

		f, err := os.OpenFile("./uploaded_files/"+fName, os.O_WRONLY|os.O_CREATE, 0777)
		PanicIfError(err)

		jpeg.Encode(f, im, nil)
	case "gif":
		im, err := gif.Decode(r)
		PanicIfError(err)

		f, err := os.OpenFile("./uploaded_files/"+fName, os.O_WRONLY|os.O_CREATE, 0777)
		PanicIfError(err)

		gif.Encode(f, im, nil)
	}
	return nil
}

/* example
err := helpers.SaveFileFromBase64(file["FileName"].(string), file["Base64"].(string), "./uploaded_files/")
*/
func SaveFileFromBase64(fName string, data string, dir string) error {
	index := strings.Index(data, "base64,")
	dec, err := base64.StdEncoding.DecodeString(data[index+1:])
	PanicIfError(err)
	f, err := os.Create(dir + fName)
	PanicIfError(err)

	defer f.Close()
	if _, err := f.Write(dec); err != nil {
		PanicIfError(err)
	}
	f.Sync()
	return nil
}

func CreatePathIfNotExist(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		errr := os.MkdirAll(path, os.ModePerm)
		PanicIfError(errr)
	}
}

func RemoveFile(fName string, dir string) error {
	os.Remove(dir + fName)
	return nil
}
