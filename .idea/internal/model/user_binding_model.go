package model

type UserBinding struct {
	BaseModel
	ID               int64   `db:"id" json:"id"`
	FcmId			 string  `db:"fcm_id" json:"fcmId"`
	CustomerId       string  `db:"customer_id" json:"customerId"`
	DeviceId         string  `db:"device_id" json:"deviceId"`
	OsType         	 string  `db:"os_type" json:"osType"`
}
