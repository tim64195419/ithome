package ithome

import (
	log "github.com/sirupsen/logrus"
	"github.com/tim64195419/ithome/ironman"
)

func PrintItHome() {
	log.Info("hi home")
	ironman.PrintIronMan()
}
