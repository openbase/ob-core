package main

import (
	"flag"
	"fmt"
	"os"

	ob "github.com/openbase/ob-core"
	obsrv "github.com/openbase/ob-core/server"
)

func main() {
	//	command-line flags
	dirPath := flag.String("hive", "", fmt.Sprintf("%s hive directory path to use.\nIf omitted, defaults to either current directory, or the path stored in\nthe $%s environment variable: '%s'.\n", ob.OB_TITLE, ob.ENV_OBHIVE, os.Getenv(ob.ENV_OBHIVE)))
	addr := flag.String("addr", ":23456", "TCP address to serve HTTP requests.\nSpecify ':http' for default HTTP port or ':https' for default HTTPS port\n")
	tlsCertFile := flag.String("tls_cert", "", "File name containing a certificate for HTTPS-serving via TLS.\nIf the certificate is signed by a certificate authority, tls_cert should be\nthe concatenation of the server's certificate followed by the CA's certificate.\n")
	tlsKeyFile := flag.String("tls_key", "", "File name containing a matching private key for TLS serving.\nFor HTTPS/TLS serving, BOTH tls_cert AND tls_key are required.")
	logToFile := flag.Bool("log_file", false, "If false, logs to 'standard output' (console),\nif true, logs to a new log file at [hive]/log/[date-time].log\n")
	silent := flag.Bool("silent", false, "If true, nothing is ever written to standard-output.\n")
	flag.Parse()
	//	run until the Halting Problem is solved
	if err := obsrv.Main(*dirPath, *addr, *tlsCertFile, *tlsKeyFile, *logToFile, *silent); err != nil {
		panic(err)
	}
}
