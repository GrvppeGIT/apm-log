package logger

import (
	"log"

	"github.com/GrvppeGIT/apm-log/logger/utils"

	"github.com/GrvppeGIT/apm-log/logger/models"
)

type Log struct {
	Printer Printer
	Options models.Options
}

var Logger models.Log
var Dt utils.DateTime
var MainLog Log

func StartLog(nameApp string, version string, tz string) Log {

	MainLog.Printer = Printer{}
	MainLog.Options = models.Options{Name: nameApp, Version: version, Tz: tz}

	log.SetFlags(0)

	Dt.SetDateTime(MainLog.Options.Tz, "2006-01-02T15:04:05.000")

	// set service name and version
	Logger.SetService(MainLog.Options.Name, MainLog.Options.Version)

	MainLog.Printer.Initialize(Logger, Dt)
	MainLog.Printer.SetContext("Main")
	MainLog.Printer.Log("Module apm-log started...")

	return MainLog
}

// "Test Name", Version: "0.0.0", Tz: "America/Sao_Paulo"
