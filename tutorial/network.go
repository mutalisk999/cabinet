package tutorial

import (
	"encoding/json"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
)

func networkScreen(_ fyne.Window) fyne.CanvasObject {
	return container.NewMax()
}

func networkGetIPScreen(_ fyne.Window) fyne.CanvasObject {
	externalIPEntryValidated := widget.NewMultiLineEntry()
	externalIPEntryValidated.Wrapping = fyne.TextWrapWord
	deviceIfIPEntryValidated := widget.NewMultiLineEntry()
	deviceIfIPEntryValidated.Wrapping = fyne.TextWrapWord

	return container.NewVBox(
		widget.NewSeparator(),
		container.NewBorder(container.NewBorder(nil, nil,
			widget.NewLabelWithStyle("Get External IP:",
				fyne.TextAlignLeading, fyne.TextStyle{Bold: true, Italic: true}),
			widget.NewButton("get", func() {
				go func() {
					// DefaultTransport will use ProxyFromEnvironment
					client := http.Client{
						//Transport: http.DefaultTransport,
						Transport: &http.Transport{},
					}
					resp, err := client.Get("https://ipinfo.io/")
					if err != nil {
						externalIPEntryValidated.SetText(err.Error())
						return
					}
					if resp.StatusCode != 200 {
						externalIPEntryValidated.SetText(resp.Status)
						return
					}
					defer resp.Body.Close()
					body, err := ioutil.ReadAll(resp.Body)
					if err != nil {
						externalIPEntryValidated.SetText(err.Error())
						return
					}

					type IpInfo struct {
						Ip       string `json:"ip"`
						City     string `json:"city"`
						Region   string `json:"region"`
						Country  string `json:"country"`
						Loc      string `json:"loc"`
						Org      string `json:"org"`
						Timezone string `json:"timezone"`
					}
					var ipInfo IpInfo
					err = json.Unmarshal(body, &ipInfo)
					if err != nil {
						externalIPEntryValidated.SetText(err.Error())
						return
					}
					externalIPEntryValidated.SetText(fmt.Sprintf(
						"ip: %s        city: %s        region: %s        country: %s\n"+
							"loc: %s        timezone: %s\n"+
							"org: %s",
						ipInfo.Ip, ipInfo.City, ipInfo.Region, ipInfo.Country, ipInfo.Loc, ipInfo.Timezone, ipInfo.Org))
				}()
			})),
			externalIPEntryValidated, nil, nil),
		widget.NewSeparator(),
		container.NewBorder(container.NewBorder(nil, nil,
			widget.NewLabelWithStyle("Get Device Interface IP:",
				fyne.TextAlignLeading, fyne.TextStyle{Bold: true, Italic: true}),
			widget.NewButton("get", func() {
				ipsInfoText := ""
				interfaceAddrs, err := net.InterfaceAddrs()
				if err != nil {
					deviceIfIPEntryValidated.SetText(err.Error())
				}
				for _, addr := range interfaceAddrs {
					ipNet, isValidIpNet := addr.(*net.IPNet)
					if isValidIpNet && ipNet.IP.To4() != nil {
						ipInfoText := fmt.Sprintf("ip: %s      mask: %s      network: %s\n",
							ipNet.IP.String(), ipNet.Mask.String(), ipNet.IP.Mask(ipNet.Mask).String())
						ipsInfoText = strings.Join([]string{ipsInfoText, ipInfoText}, "")
					}
				}
				deviceIfIPEntryValidated.SetText(ipsInfoText)
			})),
			deviceIfIPEntryValidated, nil, nil),
	)
}
