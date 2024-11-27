package model

import "time"

type BaseModel struct {
	CreatedDate      time.Time `db:"created_date" json:"createdDate" default:"current_timestamp"`
	LastModifiedDate time.Time `db:"last_modified_date" json:"lastModifiedDate,omitempty"`
	CreatedBy        string    `db:"created_by" json:"createdBy"`
	LastModifiedBy   string    `db:"last_modified_by" json:"lastModifiedBy,omitempty"`
}
