package entity

type Major struct {
	MajorId       string `json:"major_id"`
	MajorName     string `json:"major_name"`
	MajorDeleteAt string `json:"-"`
}
