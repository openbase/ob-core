package obsrv

import (
	"os"

	ugo "github.com/metaleap/go-util"
	ustr "github.com/metaleap/go-util/str"

	ob "github.com/openbase/ob-core"
)

//	Called by func main() in cmd/ob-server/main.go package.
//	In theory, you could 'abuse' this by writing your own custom server
//	executable instead of using ob-server, but this isn't really intended.
//
//	(Do note, this function does run 'forever' and thus defers all cleanups.)
func Main(hiveDir, httpAddr, tlsCertFile, tlsKeyFile string, logToFile, silent bool) (err error) {
	//	pre-init
	if len(hiveDir) == 0 {
		hiveDir, _ = os.Getwd()
		hiveDir = ob.Hive.GuessDir(hiveDir)
	}
	ob.Opt.Server = true

	//	init
	logger := ob.NewLogger(ugo.Ifw(silent, nil, os.Stdout))
	if err = ob.Init(hiveDir, logger); err != nil {
		return
	}
	defer ob.Dispose()

	//	create logger file?
	if logToFile {
		var (
			logFile     *os.File
			logFilePath string
		)
		if logFilePath, logFile, err = ob.Hive.CreateLogFile(); err != nil {
			return err
		} else {
			defer logFile.Close()
			logger.Infof("LOG @ %s", logFilePath)
			logger = ob.NewLogger(logFile)
			ob.Opt.Log = logger
		}
	}

	//	all systems go!
	proto := ugo.Ifs(len(tlsCertFile) > 0 && len(tlsKeyFile) > 0, "https", "http")
	logger.Infof("LIVE @ %s://%s%s", proto, ugo.HostName(), ustr.StripSuffix(httpAddr, ":"+proto))
	Init()
	return ListenAndServe(httpAddr, tlsCertFile, tlsKeyFile)
}
