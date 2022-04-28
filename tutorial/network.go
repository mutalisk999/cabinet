package tutorial

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/mutalisk999/cabinet/utils"
	"io/ioutil"
	"math"
	"net"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func networkScreen(_ fyne.Window) fyne.CanvasObject {
	return container.NewMax()
}

func networkGetIPScreen(_ fyne.Window) fyne.CanvasObject {
	externalIPEntry := widget.NewMultiLineEntry()
	externalIPEntry.Wrapping = fyne.TextWrapWord
	externalIPEntry.SetPlaceHolder("This entry is for information of external ip")
	deviceIfIPEntry := widget.NewMultiLineEntry()
	deviceIfIPEntry.Wrapping = fyne.TextWrapWord
	deviceIfIPEntry.SetPlaceHolder("This entry is for information of device local ip")

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
						externalIPEntry.SetText(err.Error())
						return
					}
					if resp.StatusCode != 200 {
						externalIPEntry.SetText(resp.Status)
						return
					}
					defer resp.Body.Close()
					body, err := ioutil.ReadAll(resp.Body)
					if err != nil {
						externalIPEntry.SetText(err.Error())
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
						externalIPEntry.SetText(err.Error())
						return
					}
					externalIPEntry.SetText(fmt.Sprintf(
						"ip: %s        city: %s        region: %s        country: %s\n"+
							"loc: %s        timezone: %s\n"+
							"org: %s",
						ipInfo.Ip, ipInfo.City, ipInfo.Region, ipInfo.Country, ipInfo.Loc, ipInfo.Timezone, ipInfo.Org))
				}()
			})),
			externalIPEntry, nil, nil),
		widget.NewSeparator(),
		container.NewBorder(container.NewBorder(nil, nil,
			widget.NewLabelWithStyle("Get Device Interface IP:",
				fyne.TextAlignLeading, fyne.TextStyle{Bold: true, Italic: true}),
			widget.NewButton("get", func() {
				ipsInfoText := ""
				interfaceAddrs, err := net.InterfaceAddrs()
				if err != nil {
					deviceIfIPEntry.SetText(err.Error())
				}
				for _, addr := range interfaceAddrs {
					ipNet, isValidIpNet := addr.(*net.IPNet)
					if isValidIpNet && ipNet.IP.To4() != nil {
						ipInfoText := fmt.Sprintf("ip: %s      mask: %s      network: %s\n",
							ipNet.IP.String(), net.IP(ipNet.Mask).String(), ipNet.IP.Mask(ipNet.Mask).String())
						ipsInfoText = strings.Join([]string{ipsInfoText, ipInfoText}, "")
					}
				}
				deviceIfIPEntry.SetText(ipsInfoText)
			})),
			deviceIfIPEntry, nil, nil),
	)
}

func networkIPMaskScreen(_ fyne.Window) fyne.CanvasObject {
	// ip + mask bit -> mask / network ip/ broadcast ip / first valid ip / last valid ip
	// mask bit -> mask / ip count / valid ip count
	// ip count needed -> mask bit / mask / valid ip count
	ipv4Entry := utils.NewIPv4Entry()
	maskBitEntry := utils.NewMaskBitEntry()
	maskBitEntry2 := utils.NewMaskBitEntry()
	numberEntry := utils.NewNumberEntry()

	ipMaskBitCalcResultEntry := widget.NewMultiLineEntry()
	ipMaskBitCalcResultEntry.Wrapping = fyne.TextWrapWord
	ipMaskBitCalcResultEntry.SetPlaceHolder("This entry is for calculated result of ip/mask bit")
	maskBitCalcResultEntry := widget.NewMultiLineEntry()
	maskBitCalcResultEntry.Wrapping = fyne.TextWrapWord
	maskBitCalcResultEntry.SetPlaceHolder("This entry is for calculated result of mask bit")
	ipCountNeedCalcResultEntry := widget.NewMultiLineEntry()
	ipCountNeedCalcResultEntry.Wrapping = fyne.TextWrapWord
	ipCountNeedCalcResultEntry.SetPlaceHolder("This entry is for calculated result of ip count you needed")

	return container.NewVBox(
		widget.NewSeparator(),
		widget.NewLabelWithStyle("IP/MaskBit Calculator:",
			fyne.TextAlignLeading, fyne.TextStyle{Bold: true, Italic: true}),
		container.NewBorder(nil, nil,
			container.NewHBox(ipv4Entry, widget.NewLabel("/"), maskBitEntry, widget.NewLabel("[0-30]")),
			widget.NewButton("calculate", func() {
				if ipv4Entry.Validate() != nil || maskBitEntry.Validate() != nil {
					return
				}

				cidrStr := fmt.Sprintf("%s/%s", ipv4Entry.Text, maskBitEntry.Text)
				_, ipNet, err := net.ParseCIDR(cidrStr)
				if err != nil {
					ipMaskBitCalcResultEntry.SetText(err.Error())
					return
				}
				if ipNet == nil {
					return
				}

				mask, _ := strconv.Atoi(maskBitEntry.Text)
				ipNetCount := int(math.Pow(2, float64(32-mask)))

				ipNetBytes := ipNet.IP
				ipNetUint32 := binary.BigEndian.Uint32(ipNetBytes)
				ipBroadCastUint32 := ipNetUint32 + uint32(ipNetCount) - 1
				ipNetFirstUint32 := ipNetUint32 + 1
				ipNetLastUint32 := ipBroadCastUint32 - 1

				ipBroadCastBytes := make([]byte, 4)
				ipNetFirstBytes := make([]byte, 4)
				ipNetLastBytes := make([]byte, 4)
				binary.BigEndian.PutUint32(ipBroadCastBytes, ipBroadCastUint32)
				binary.BigEndian.PutUint32(ipNetFirstBytes, ipNetFirstUint32)
				binary.BigEndian.PutUint32(ipNetLastBytes, ipNetLastUint32)

				calcResultText := fmt.Sprintf("ip: %s      mask: %s      network: %s      broadcast: %s\nfirst valid ip: %s      last valid ip: %s\nvalid ip count: %d",
					ipv4Entry.Text, net.IP(ipNet.Mask).String(), ipNet.IP.String(),
					net.IPv4(ipBroadCastBytes[0], ipBroadCastBytes[1], ipBroadCastBytes[2], ipBroadCastBytes[3]).String(),
					net.IPv4(ipNetFirstBytes[0], ipNetFirstBytes[1], ipNetFirstBytes[2], ipNetFirstBytes[3]).String(),
					net.IPv4(ipNetLastBytes[0], ipNetLastBytes[1], ipNetLastBytes[2], ipNetLastBytes[3]).String(),
					ipNetCount-2)
				ipMaskBitCalcResultEntry.SetText(calcResultText)
			}),
		),
		ipMaskBitCalcResultEntry,
		widget.NewSeparator(),
		widget.NewLabelWithStyle("MaskBit Calculator:",
			fyne.TextAlignLeading, fyne.TextStyle{Bold: true, Italic: true}),
		container.NewBorder(nil, nil,
			container.NewHBox(maskBitEntry2, widget.NewLabel("[0-30]")),
			widget.NewButton("calculate", func() {
				if maskBitEntry2.Validate() != nil {
					return
				}

				mask, _ := strconv.Atoi(maskBitEntry2.Text)
				ipNetCount := int(math.Pow(2, float64(32-mask)))
				maskLong := 0xffffffff >> (32 - mask) << (32 - mask)
				ipV4Mask := net.IPv4Mask(byte(maskLong>>24), byte((maskLong&0x00ff0000)>>16), byte((maskLong&0x0000ff00)>>8), byte(maskLong&0x000000ff))

				calcResultText2 := fmt.Sprintf("mask: %s\nvalid ip count: %d",
					net.IP(ipV4Mask).String(), ipNetCount-2)
				maskBitCalcResultEntry.SetText(calcResultText2)
			}),
		),
		maskBitCalcResultEntry,
		widget.NewSeparator(),
		widget.NewLabelWithStyle("How many ips are needed:",
			fyne.TextAlignLeading, fyne.TextStyle{Bold: true, Italic: true}),
		container.NewBorder(nil, nil,
			numberEntry,
			widget.NewButton("calculate", func() {
				if numberEntry.Validate() != nil {
					return
				}

				ipNetCountFunc := func(val int) int {
					s := 1
					for {
						if s >= val {
							return s
						} else {
							s = s << 1
						}
					}
				}
				number, _ := strconv.Atoi(numberEntry.Text)
				if number <= 0 || number > 100000000 {
					numberEntry.SetText("number should be set in section 1-100000000")
					return
				}
				// add network ip & broadcast ip
				number = number + 2
				ipNetCount := ipNetCountFunc(number)
				maskLong := uint32(-ipNetCount)
				ipV4Mask := net.IPv4Mask(byte(maskLong>>24), byte((maskLong&0x00ff0000)>>16), byte((maskLong&0x0000ff00)>>8), byte(maskLong&0x000000ff))

				calcResultText3 := fmt.Sprintf("mask: %s\nvalid ip count: %d",
					net.IP(ipV4Mask).String(), ipNetCount-2)
				ipCountNeedCalcResultEntry.SetText(calcResultText3)
			}),
		),
		ipCountNeedCalcResultEntry,
	)
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

	// default value
	portEntry.Text = "30000"

	startServerButton := widget.NewButton("run", func() {
		if directorySetLabel.Text == "" || ipPortSetLabel.Text == "" {
			dialog.ShowError(errors.New("you need to set all of these: Dir/IP/Port"), win)
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
				folderOpen := dialog.NewFolderOpen(func(list fyne.ListableURI, err error) {
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
				folderOpen.Resize(fyne.Size{Width: 800, Height: 600})
				folderOpen.Show()
			}),
		),
		widget.NewSeparator(),
		widget.NewLabelWithStyle("Set IP/Port (port allowed section: [30000, 59999]) and Start Web Server:",
			fyne.TextAlignLeading, fyne.TextStyle{Bold: true, Italic: true}),
		container.NewHBox(ipSelector, widget.NewLabel(":"), portEntry),

		container.NewBorder(nil, nil, ipPortSetLabel, startServerButton),
	)
}
