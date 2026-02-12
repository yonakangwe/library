package systems

import (
	"library/package/client"
)

const (
	otherBaseURL = "http://x.x.x.x/others/api/v1"
)

var InstitutionClient *client.Client

func Init() {
	// log.Infoln("Initialising clients...")
	// cfg, _ := config.New()

	// librarrySystem := "librarry"
	// librarryKey, err := cfg.GetSystemPrivateKey(librarrySystem)
	// if util.IsError(err) {
	// 	panic("system is not well started")
	// }

	// InstitutionClient, err = client.New(otherBaseURL, librarryKey, librarrySystem)
	// if util.IsError(err) {
	// 	panic("institution client could not be initiated")
	// }
}
