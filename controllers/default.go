package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/beego/i18n"
	"strings"
	"whois/services"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	domain := fmt.Sprintf("%v", this.GetString("domain", ""))

	if domain != "" {
		if !strings.HasSuffix(domain, ".vn") {
			result, err := services.GetGlobalDomain(domain)
			if err != nil {
				logs.Error(err)
			}

			if result.Registrar.CreatedDate == "" {
				this.Data["message"] = i18n.Tr(services.Lang, "domain_available")
				result.Registrar.DomainName = domain
			} else {
				this.Data["message"] = i18n.Tr(services.Lang, "domain_information")
				statuses := strings.Split(result.Registrar.DomainStatus, " ")
				status := statuses[0]
				result.Registrar.DomainStatus = this.getStatus(status)
			}
			this.Data["result"] = result
		} else {
			result, err := services.GetVietnamDomain(domain)
			if err != nil {
				logs.Error(err)
			}

			if result.Registrar.CreatedDate == "" {
				this.Data["message"] = i18n.Tr(services.Lang, "domain_available")
				result.Registrar.DomainName = domain
			} else {
				this.Data["message"] = i18n.Tr(services.Lang, "domain_information")
				statuses := strings.Split(result.Registrar.DomainStatus, " ")
				status := statuses[0]
				result.Registrar.DomainStatus = this.getStatus(status)
			}
			this.Data["result"] = result
		}

		this.Data["domain"] = domain
	}

	this.Data["Email"] = "kientranvan@gmail.com"
	this.TplName = "index.tpl"
}

func (this *MainController) getStatus(status string) string {
	return status
}
