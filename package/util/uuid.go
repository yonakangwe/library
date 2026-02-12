package util

import (
	"crypto/md5"
	"strings"

	"github.com/google/uuid"
)

func GenerateUUID(data string) string {
	appID := "0e9fcc7f-e56d-414e-8e66-5c5a8d893d27"
	version := 4
	space := uuid.MustParse(appID)
	cleanString := strings.TrimRight(data, ".")
	cleanString = strings.ReplaceAll(cleanString, " ", "")
	cleanString = strings.ToLower(cleanString)
	id := uuid.NewHash(md5.New(), space, []byte(cleanString), version)
	return id.String()

}
