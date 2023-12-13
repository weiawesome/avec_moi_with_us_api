package redis

import (
	"context"
	"github.com/google/uuid"
	"time"
)

const hashMapName = "mail_id_hash"

func (r *Repository) GetHashValue(mail string) (string, error) {
	if mail == "system" {
		return mail, nil
	}
	ctx := context.Background()
	exists, err := r.client.Exists(ctx, hashMapName).Result()
	if err != nil {
		return "", err
	}

	if exists == 1 {
		value, err := r.client.HGet(ctx, hashMapName, mail).Result()
		if value != "" && err == nil {
			return value, nil
		} else {
			id := uuid.New().String()
			r.client.HSet(ctx, hashMapName, mail, id)
			return id, nil
		}
	} else if exists == 0 {
		r.client.HSet(ctx, hashMapName, "system", "system")
		expirationTime := time.Now().Truncate(24 * time.Hour).Add(24 * time.Hour)
		err := r.client.ExpireAt(ctx, hashMapName, expirationTime).Err()
		if err != nil {
			return "", err
		}
		id := uuid.New().String()
		err = r.client.HSet(ctx, hashMapName, mail, id).Err()
		return id, err
	} else {
		return "", err
	}
}
