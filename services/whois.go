package services

import (
	"fmt"
	"github.com/likexian/whois-parser-go"
	"github.com/undiabler/golang-whois"
	"io/ioutil"
	"net/http"
	"strings"
)

const Lang string = "vi-VN"
const URL_DOMAIN string = "http://www.whois.net.vn/whois.php?domain=%s"
const URL_WHOIS string = "http://www.whois.net.vn/whois.php?domain=%s&act=getwhois"

func GetGlobalDomain(domain string) (whois_parser.WhoisInfo, error) {
	raw, err := whois.GetWhois(domain)
	if err != nil {
		return whois_parser.WhoisInfo{}, err
	}

	result, err := whois_parser.Parse(raw)
	if err != nil {
		return whois_parser.WhoisInfo{}, err
	}

	return result, nil
}

func GetVietnamDomain(domain string) (whois_parser.WhoisInfo, error) {
	response, err := http.Get(fmt.Sprintf(URL_DOMAIN, domain))
	if err != nil {
		return whois_parser.WhoisInfo{}, err
	}

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return whois_parser.WhoisInfo{}, err
	}

	response.Body.Close()

	whoisInfo := whois_parser.WhoisInfo{}
	whoisInfo.Registrar.DomainName = domain
	if string(content) != "1" {
		return whoisInfo, nil
	}

	response, err = http.Get(fmt.Sprintf(URL_WHOIS, domain))
	if err != nil {
		return whois_parser.WhoisInfo{}, err
	}
	defer response.Body.Close()

	content, err = ioutil.ReadAll(response.Body)

	result := string(content)
	rows := strings.Split(result, "<br/>")

	for _, row := range rows {
		row := strings.TrimSpace(row)
		if strings.Contains(row, ":") {
			if strings.HasPrefix(row, "Domain") {
				domains := strings.Split(row, ":")
				whoisInfo.Registrar.DomainName = domains[1]
			}

			if strings.HasPrefix(row, "Status") {
				statuses := strings.Split(row, ":")
				whoisInfo.Registrar.DomainStatus = statuses[1]
			}

			if strings.HasPrefix(row, "Issue Date") {
				issuses := strings.Split(row, ":")
				whoisInfo.Registrar.CreatedDate = issuses[1]
			}

			if strings.HasPrefix(row, "Expired Date") {
				expireses := strings.Split(row, ":")
				whoisInfo.Registrar.ExpirationDate = expireses[1]
			}

			if strings.HasPrefix(row, "Registrar Name") {
				registers := strings.Split(row, ":")
				whoisInfo.Registrar.RegistrarName = registers[1]
			}

			if strings.HasPrefix(row, "Owner Name") {
				owners := strings.Split(row, ":")
				whoisInfo.Registrant.Name = owners[1]
			}
		}
	}

	return whoisInfo, nil
}
