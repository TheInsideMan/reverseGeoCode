package log

import (
	"github.com/op/go-logging"
	"os"
)

var Log = logging.MustGetLogger("example")

var format = logging.MustStringFormatter(
	"%{color}%{time:15:04:05.000} %{shortfile} ▶ %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}",
)

func init() {

	backend2 := logging.NewLogBackend(os.Stderr, "", 0)
	backend2Formatter := logging.NewBackendFormatter(backend2, format)
	logging.SetBackend(backend2Formatter)

	// Log.Info("info")
	// Log.Notice("notice")
	// Log.Warning("warning")
	// Log.Error("err")
	// Log.Critical("crit")

}

