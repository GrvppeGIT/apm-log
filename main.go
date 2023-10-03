package main

import (
	"github.com/GrvppeGIT/apm-log/show/utils"

	"github.com/GrvppeGIT/apm-log/show/models"

	"log"

	"github.com/GrvppeGIT/apm-log/show/services"
)

var logger models.Log
var show services.Show
var dt utils.DateTime

func main() {
	// apm.Main("test", "0.0.0", "America/Sao_Paulo")

	log.SetFlags(0)

	opt := models.Options{Name: "Test Name", Version: "0.0.0", Tz: "America/Sao_Paulo"}

	dt.SetDateTime(opt.Tz, "2006-01-02T15:04:05.000")

	// set service name and version
	logger.SetService(opt.Name, opt.Version)

	show.Initialize(logger, dt)

	show.SetContext("Main")
	show.Log("teste")

	// apm.Main.Log("oi")
}
