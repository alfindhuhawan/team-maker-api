package vo

import (
	"fmt"
	"strings"
)

type Address struct {
	Name                 Name             `json:"name,omitempty" bson:"name"`                                     //
	Phone                PhoneNumber      `json:"phone,omitempty" bson:"phone"`                                   //
	Address              string           `json:"address" bson:"address"`                                         // Jl Karapitan no 24
	DisplayedAddress     string           `json:"displayed_address" bson:"displayed_address"`                     // Jl Karapitan no 24
	District             string           `json:"district,omitempty" bson:"district"`                             // kecamatan
	EstimatedDate        string           `json:"estimated_date,omitempty" bson:"estimated_date"`                 // kecamatan
	EstimatedTimeMaximum int64            `json:"estimated_time_maximum,omitempty" bson:"estimated_time_maximum"` // kecamatan
	EstimatedTimeMinimum int64            `json:"estimated_time_minimum,omitempty" bson:"estimated_time_minimum"` // kecamatan
	Latitude             string           `json:"latitude,omitempty" bson:"latitude"`
	Longitude            string           `json:"longitude,omitempty" bson:"longitude"`
	ZipCode              string           `json:"zip_code,omitempty" bson:"zip_code"` // 14232
	IsPaxelMember        bool             `json:"is_paxel_member" bson:"is_paxel_member"`
	Email                string           `json:"email" bson:"email"`
	Note                 string           `json:"note" bson:"note"`
	Province             string           `json:"province" bson:"province"`
	Regency              string           `json:"regency" bson:"regency"`
	Village              string           `json:"village" bson:"village"`
	BoxID                string           `json:"boxId" bson:"boxId"`
	IsWasteBank          bool             `json:"is_waste_bank" bson:"is_waste_bank"`
	PHCode               string           `json:"ph_code" bson:"ph_code"`
	PaxelCounter         *PaxelCounter    `json:"paxel_counter,omitempty" bson:"paxel_counter,omitempty"`
	WasteBankDetail      *WasteBankDetail `json:"waste_bank_detail,omitempty" bson:"waste_bank_detail,omitempty"`
}

func (n Address) Validate(s string) error {

	if v := strings.TrimSpace(n.Name.String()); v == "" {
		return fmt.Errorf("Nama pada %s harus diisi", s)
	}

	if v := strings.TrimSpace(n.Name.String()); len(v) < 3 {
		return fmt.Errorf("Nama pada %s minimal 3 karakter", s)
	}

	if v := strings.TrimSpace(n.DisplayedAddress); v == "" {
		return fmt.Errorf("Detail %s harus diisi", s)
	}

	if v := strings.TrimSpace(n.ZipCode); v == "" {
		return fmt.Errorf("Kode pos %s harus diisi", s)
	}

	// if v := strings.TrimSpace(n.Village.String()); v == "" {
	// 	return fmt.Errorf("Kelurahan %s harus diisi", s)
	// }

	if v := strings.TrimSpace(n.District); v == "" {
		return fmt.Errorf("Kecamatan %s harus diisi", s)
	}

	// if n.LatLng[0] == 0 {
	// 	return fmt.Errorf("Latitude %s harus diisi", s)
	// }

	// if n.LatLng[1] == 0 {
	// 	return fmt.Errorf("Longitude %s harus diisi", s)
	// }

	if v := strings.TrimSpace(string(n.Phone)); v == "" {
		return fmt.Errorf("Nomor telepon %s harus diisi", s)
	}

	// if v := strings.TrimSpace(n.City.String()); v == "" {
	// 	return fmt.Errorf("Kota %s harus diisi", s)
	// }

	return nil
}

type PaxelCounter struct {
	Code                 string  `json:"code" bson:"code"`
	Name                 string  `json:"name" bson:"name"`
	Address              string  `json:"address" bson:"address"`
	Longitude            float64 `json:"longitude" bson:"longitude"`
	Latitude             float64 `json:"latitude" bson:"latitude"`
	ZipCode              string  `json:"zipcode" bson:"zipcode"`
	Village              string  `json:"village" bson:"village"`
	District             string  `json:"district" bson:"district"`
	City                 string  `json:"city" bson:"city"`
	Province             string  `json:"province" bson:"province"`
	ContactPersonName    string  `json:"contact_person_name" bson:"contact_person_name"`
	ContactPersonNumber  string  `json:"contact_person_number" bson:"contact_person_number"`
	StartOperationalHour string  `json:"start_operational_hour" bson:"start_operational_hour"`
	EndOperationalHour   string  `json:"end_operational_hour" bson:"end_operational_hour"`
}

type WasteBankDetail struct {
	BankCode                  string  `json:"bank_code" bson:"bank_code"`
	BankName                  string  `json:"bank_name" bson:"bank_name"`
	DistanceFromPickupAddress float64 `json:"distance_from_pickup_address" bson:"distance_from_pickup_address"`
}
