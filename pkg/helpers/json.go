package helpers

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

func ReadFromRequestBody(r *http.Request, result interface{}) {
	contentType := r.Header.Get("Content-Type")
	if contentType == "application/json" {
		decodeJSON(r, result)
	} else if contentType == "application/x-www-form-urlencoded" {
		urlEncodedFormToJSON(r, result)
	} else {
		formDataToJSON(r, result)
	}
}

func WriteToResponseBody(w http.ResponseWriter, response interface{}) {
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(response)
	PanicIfError(err)
}

func formDataToJSON(r *http.Request, result interface{}) {
	err := r.ParseMultipartForm(2048)
	PanicIfError(err)
	data := map[string]interface{}{}
	for i, _ := range r.MultipartForm.File {
		file, fileHeader, err := r.FormFile(i)
		if err == nil {
			bytes, _ := ioutil.ReadAll(file)
			// contentType := ContentTypeBase64(http.DetectContentType(bytes))
			data[i+"_data"] = map[string]interface{}{
				"Base64":   toBase64([]byte(bytes)),
				"FileName": fileHeader.Filename,
				"Size":     fileHeader.Size,
				"Ext":      MimeTypeToExt(fileHeader.Header["Content-Type"][0]),
				// "Ext":      MimeTypeToExt(http.DetectContentType(bytes)),
			}

		}
	}
	dataType := GetStructDataType(&result, "snake")
	for i, _ := range r.MultipartForm.Value {

		val := r.PostFormValue(i)

		switch dataType[i] {
		case "int":
			res, _ := strconv.Atoi(val)
			data[i] = res
		case "float32":
			res, _ := strconv.ParseFloat(val, 32)
			data[i] = res
		case "float64":
			res, _ := strconv.ParseFloat(val, 64)
			data[i] = res
		default:
			data[i] = val
		}
	}

	ra, err := json.Marshal(data)
	PanicIfError(err)
	json.Unmarshal([]byte(ra), result)
}

func urlEncodedFormToJSON(r *http.Request, result interface{}) {
	err := r.ParseForm()
	PanicIfError(err)
	data := map[string]interface{}{}
	for i, _ := range r.Form {
		data[i] = r.PostFormValue(i)
	}
	ra, err := json.Marshal(data)
	PanicIfError(err)

	json.Unmarshal([]byte(ra), result)
}

func decodeJSON(r *http.Request, result interface{}) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(result)
	PanicIfError(err)
}

func toBase64(f []byte) string {
	return base64.StdEncoding.EncodeToString(f)
}
