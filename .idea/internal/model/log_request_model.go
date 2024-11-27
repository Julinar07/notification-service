package model

type LogRequest struct {
	BaseModel
	ID               int64   `db:"id" json:"id"`
	FcmId			 string  `db:"fcm_id" json:"fcmId"`
	CustomerId       string  `db:"customer_id" json:"customerId"`
	DeviceId         string  `db:"device_id" json:"deviceId"`
	OsType         	 string  `db:"os_type" json:"osType"`
	Status         	 string  `db:"status" json:"status"`
	ErrorMessage   	 string  `db:"error_message" json:"errorMessage"`
	Body		   	 string  `db:"body" json:"body"`
	Title		   	 string  `db:"title" json:"title"`
	AdditionalInfo 	 string  `db:"additional_info" json:"additionalInfo"`
	Type		 	 string  `db:"type" json:"type"`
}
