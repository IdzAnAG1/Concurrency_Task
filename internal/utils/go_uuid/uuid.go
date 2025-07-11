package go_uuid

import "github.com/google/uuid"

func Uid() string {
	return uuid.NewString()
}
