package common

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

/*
To embed a struct into another struct and save it to gorm
you must be implement 2 method:
- Scanner
- Valuer
convert json value to buy to store db
-> convert data from byte to json to return client
*/
type Image struct {
	Id        int    `json:"id" gorm:"column:id;"`
	Url       string `json:"url" gorm:"column:url;"`
	Width     int    `json:"width" gorm:"column:width;"`
	Height    int    `json:"height" gorm:"column:height"`
	CloudName string `json:"cloud_name,omitempty" gorm:"-"`
	Extension string `json:"extension,omitempty" gorm:"-"`
}

func (Image) TableName() string {
	return "images"
}

func (j *Image) Scan(value interface{}) error {

	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value: ", value))
	}

	var img Image
	if err := json.Unmarshal(bytes, &img); err != nil {
		return err
	}

	*j = img
	return nil
}

func (j *Image) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	return json.Marshal(j)
}

type Images []Image

func (i *Images) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value: ", value))
	}

	var imgs Images
	if err := json.Unmarshal(bytes, &imgs); err != nil {
		return err
	}

	*i = imgs

	return nil
}

func (i *Images) Value() (driver.Value, error) {
	if i == nil {
		return nil, nil
	}
	return json.Marshal(i)
}
