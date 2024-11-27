package repo

import (
	"context"
	_commonDB "notification-service/.idea/internal/common"
	model "notification-service/.idea/internal/model"
)

type UserBindingIFace interface {
	AddUserBinding(ctx context.Context, t *_commonDB.T, detail *model.UserBinding) (int64, error)
}

type UserBindingRepo struct {
}

func (r *UserBindingRepo) AddUserBinding(ctx context.Context, t *_commonDB.T, detail *model.UserBinding) (int64, error) {
	query := `
		INSERT INTO app.user_binding
		(fcm_id,
		 customer_id, 
		 device_id, 
		 os_type,
		 created_date, 
		 created_by)
		VALUES ($1, $2, $3, $4, CURRENT_TIMESTAMP, $5)
		RETURNING id;
`
	args := []interface{}{
		detail.FcmId,
		detail.CustomerId,
		detail.DeviceId,
		detail.OsType,
		detail.CreatedBy,
	}

	var insertId int64
	var err error
	if t.Tx != nil {
		err = t.Tx.QueryRow(ctx, query, args...).Scan(&insertId)
	} else {
		err = t.Db.QueryRow(ctx, query, args...).Scan(&insertId)
	}
	if err != nil {
		return 0, err
	}

	return insertId, nil
}