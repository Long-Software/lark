package main

import (
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/Long-Software/lark/pkg/log"
)

type DomainDetail struct {
	domain      string
	hasMX       bool
	hasSPF      bool
	sprRecord   string
	hasDMARC    bool
	dmarcRecord string
}

func main() {
	lg := log.Logger{
		// HasLogFile: true,
		// LogFilepath: "output.log",
		HasTimestamp: true,
		HasFilepath: true,
		HasMethod: true,
	}
	domain := "news.cnn.com"
	details, err := checkDomain(domain)
	if err != nil {
		lg.NewLog(log.FATAL, err.Error())
		os.Exit(0)
	}

	fmt.Println(*details)
}

func checkDomain(domain string) (*DomainDetail, error) {
	var details DomainDetail
	// checking for mx records
	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		return nil, err
	}
	if len(mxRecords) > 0 {
		details.hasMX = true
	}

	// checking for txt
	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		return nil, err
	}
	for _, rec := range txtRecords {
		if strings.HasPrefix(rec, "v=spf1") {
			details.hasSPF = true
			details.sprRecord = rec
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		return nil, err
	}
	for _, rec := range dmarcRecords {
		if strings.HasPrefix(rec, "v=DMARC1") {
			details.hasDMARC = true
			details.dmarcRecord = rec
			break
		}
	}

	return &details, err
}
