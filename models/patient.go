package models

type Patient struct {
	ID            int64
	Uhid          string `binding:"required"`
	Barcode       string `binding:"required"`
	Name          string `binding:"required"`
	LabourId      string `binding:"required"`
	Age           int64  `binding:"required"`
	Gender        string `binding:"required"`
	Mobile        string `binding:"required"`
	District      string `binding:"required"`
	Taluk         string `binding:"required"`
	Camp          string `binding:"required"`
	LabTestStatus int
	ReportUrl     string
}
