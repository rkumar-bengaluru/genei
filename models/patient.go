package models

import (
	"errors"

	"example.com/rest-api/db"
)

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

func (p *Patient) Save() error {
	query := `
	INSERT INTO patients(uhid, barcode, name, labour_id, age,
	                     gender, mobile, district, taluk, camp) 
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(p.Uhid, p.Barcode, p.Name, p.LabourId, p.Age,
		p.Gender, p.Mobile, p.District, p.Taluk, p.Camp)
	if err != nil {
		return err
	}
	_, err = result.LastInsertId()
	return err
}

func (p *Patient) GetPatientByUhid(uhid string) (*Patient, error) {
	query := `SELECT uhid, barcode, name, labour_id, age,
	                 gender, mobile, district, taluk, camp FROM patients WHERE uhid = ?`
	row := db.DB.QueryRow(query, uhid)

	err := row.Scan(&p.Uhid, &p.Barcode, &p.Name, &p.LabourId, &p.Age,
		&p.Gender, &p.Mobile, &p.District, &p.Taluk, &p.Camp)

	if err != nil {
		return nil, errors.New("Credentials invalid")
	}

	return p, nil
}
