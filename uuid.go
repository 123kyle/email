package email

import "github.com/google/uuid"

func UUID() string {
	return uuid.New().String()
}

func ValidUUIDString(sid string) bool {
	_, err := uuid.Parse(sid)
	return err == nil
}
