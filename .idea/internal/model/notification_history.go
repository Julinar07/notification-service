package model

type NotificationHistory struct {
	BaseModel
	ID               int64   `db:"id" json:"id"`
	CustomerId       string  `db:"customer_id" json:"customerId"`
	Body		   	 string  `db:"body" json:"body"`
	Title		   	 string  `db:"title" json:"title"`
	AdditionalInfo 	 string  `db:"additional_info" json:"additionalInfo"`
	Type		 	 string  `db:"type" json:"type"`
}
