package tutorial

import (
	"encoding/json"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/mutalisk999/cabinet/utils"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"os/exec"
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

func networkIPMaskScreen(_ fyne.Window) fyne.CanvasObject {
	// ip + mask bit -> mask / network ip/ broadcast ip / first valid ip / last valid ip
	// mask bit -> mask / ip count / valid ip count
	// ip count needed -> mask bit / mask / valid ip count
	return container.NewMax()
}

func networkWebServerScreen(win fyne.Window) fyne.CanvasObject {
	directorySetLabel := widget.NewLabelWithStyle("", fyne.TextAlignCenter, fyne.TextStyle{})
	ipPortSetLabel := widget.NewLabelWithStyle("", fyne.TextAlignCenter, fyne.TextStyle{})
	portEntry := utils.NewPortEntry()

	// get all ip addresses
	allInterfaceIps := make([]string, 0)
	interfaceAddrs, err := net.InterfaceAddrs()
	if err != nil {
		dialog.ShowError(err, win)
	}
	for _, addr := range interfaceAddrs {
		ipNet, isValidIpNet := addr.(*net.IPNet)
		if isValidIpNet && ipNet.IP.To4() != nil {
			allInterfaceIps = append(allInterfaceIps, ipNet.IP.String())
		}
	}
	allInterfaceIps = append(allInterfaceIps, "0.0.0.0")

	ipSelector := widget.NewSelect(allInterfaceIps, func(strVal string) {
		err := portEntry.Validate()
		if err != nil {
			ipPortSetLabel.SetText("")
			return
		}
		ipPortSetLabel.SetText(fmt.Sprintf("http://%s:%s", strVal, portEntry.Text))
	})
	portEntry.OnChanged = func(strVal string) {
		err := portEntry.Validate()
		if err != nil {
			ipPortSetLabel.SetText("")
			return
		}
		if ipSelector.Selected == "" {
			ipPortSetLabel.SetText("")
			return
		}
		ipPortSetLabel.SetText(fmt.Sprintf("http://%s:%s", ipSelector.Selected, strVal))
	}

	startServerButton := widget.NewButton("run", func() {
		if directorySetLabel.Text == "" || ipPortSetLabel.Text == "" {
			return
		}

		_, err := exec.LookPath("file_server")
		if err != nil {
			dialog.ShowError(err, win)
			return
		}
		go func() {
			cmd := exec.Command("file_server",
				"-e", fmt.Sprintf("%s:%s", ipSelector.Selected, portEntry.Text),
				"-d", directorySetLabel.Text)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			err := cmd.Start()
			if err != nil {
				dialog.ShowError(err, win)
				return
			}

			fyne.CurrentApp().SendNotification(&fyne.Notification{
				Title:   "Info",
				Content: fmt.Sprintf("file_server start, pid: %d", cmd.Process.Pid),
			})
			stat, err := cmd.Process.Wait()
			if err != nil {
				dialog.ShowError(err, win)
				return
			} else {
				fyne.CurrentApp().SendNotification(&fyne.Notification{
					Title:   "Info",
					Content: fmt.Sprintf("file_server stopped, pid: %d, exit_code: %d", stat.Pid(), stat.ExitCode()),
				})
			}
		}()
	})

	return container.NewVBox(
		widget.NewSeparator(),
		widget.NewLabelWithStyle("Select File Directory:",
			fyne.TextAlignLeading, fyne.TextStyle{Bold: true, Italic: true}),
		container.NewBorder(nil, nil,
			directorySetLabel,
			widget.NewButton("open folder", func() {
				dialog.ShowFolderOpen(func(list fyne.ListableURI, err error) {
					if err != nil {
						dialog.ShowError(err, win)
						return
					}
					if list == nil {
						return
					}
					if len(list.String()) >= 7 && list.String()[0:7] == "file://" {
						directorySetLabel.SetText(list.String()[7:])
					} else {
						directorySetLabel.SetText(list.String())
					}
				}, win)
			}),
		),
		widget.NewSeparator(),
		widget.NewLabelWithStyle("Set IP/Port and Start Web Server:",
			fyne.TextAlignLeading, fyne.TextStyle{Bold: true, Italic: true}),
		container.NewHBox(ipSelector, widget.NewLabel(":"), portEntry),

		container.NewBorder(nil, nil, ipPortSetLabel, startServerButton),
	)
}
