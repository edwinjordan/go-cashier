package helpers

func ContentTypeBase64(mime string) string {
	return "data:" + mime + ";base64,"
}

func ExtToMimeType(ext string) string {
	switch ext {
	case "xlsx":
		return "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"

	case "xls":
		return "application/vnd.ms-excel"

	case "zip":
		return "application/zip"

	case "docx":
		return "application/vnd.openxmlformats-officedocument.wordprocessingml.document"

	case "doc":
		return "application/msword"

	case "jpg":
		return "image/jpeg"

	case "png":
		return "image/png"

	case "pdf":
		return "application/pdf"

	case "ppt":
		return "application/vnd.ms-powerpoint"

	case "pptx":
		return "application/vnd.openxmlformats-officedocument.presentationml.presentation"

	case "rar":
		return "application/vnd.rar"

	case "apk":
		return "application/vnd.android.package-archive"
	default:
		return ""
	}
}

func MimeTypeToExt(mm string) string {
	switch mm {
	case "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet":
		return "xlsx"
	case "application/vnd.ms-excel":
		return "xls"
	case "application/zip":
		return "zip"
	case "application/vnd.openxmlformats-officedocument.wordprocessingml.document":
		return "docx"
	case "application/msword":
		return "doc"
	case "image/jpeg":
		return "jpg"
	case "image/png":
		return "png"
	case "application/pdf":
		return "pdf"
	case "application/vnd.ms-powerpoint":
		return "ppt"
	case "application/vnd.openxmlformats-officedocument.presentationml.presentation":
		return "pptx"
	case "application/vnd.rar":
		return "rar"
	case "application/vnd.android.package-archive":
		return "apk"
	default:
		return ""

	}
}
