package main

import (
	"github.com/mhale/smtpd"
	"github.com/rs/zerolog/log"
	"net"
)

func main() {
	//TODO add opentelemetry
	mailSrv := smtpd.Server{
		Addr:         ":2525",
		Appname:      "webischia-go-smtp",
		AuthRequired: false,
		Handler:      handleMail,
		HandlerRcpt:  handleRcpt,
		Hostname:     "webischia",
	}
	if err := mailSrv.ListenAndServe(); err != nil {
		log.Fatal().Err(err).Msg("Uh oh")
	}

}

func handleMail(remoteAddr net.Addr, from string, to []string, data []byte) {
	log.Info().Msg(string(data))
}

func handleRcpt(remoteAddr net.Addr, from string, to string) bool {
	//TODO Send Metrics
	log.Info().Msg("Rcpt Accepted ! From : " + from + " To: " + to)
	return true
}
