package hardinfo

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"
	"strings"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

/*
***************************************************

	功能:读取WIFI连接的接口信息
	输入:
		wifiname string:WIFI连接的名称
	输出:
		map[string]string:接口信息
		error:错误信息
	说明:
	时间:2024年2月22日
	编辑:wang_jp

***************************************************
*/
func WlanInfo(wifiname string) (map[string]string, error) {
	cmd := exec.Command("netsh", "wlan", "show", "interfaces", fmt.Sprintf("name=%s", wifiname)) //列出所有任务列表
	var stdout bytes.Buffer
	cmd.Stdout = &stdout // 标准输出
	err := cmd.Run()
	if err != nil {
		return nil, err
	}
	outstr := string(gbkToUtf8(stdout.Bytes()))
	rst := decodeWlanInfo(outstr)
	return rst, nil
}

/*
***************************************************

	功能:解析读取到的WIFI连接的接口信息
	输入:
		cmdoutstr string:WIFI接口信息的查询输出字符串
	输出:
		map[string]string:接口信息
	说明:
	时间:2024年2月22日
	编辑:wang_jp

***************************************************
*/
func decodeWlanInfo(cmdoutstr string) map[string]string {
	wlaninfo := make(map[string]string)
	outstrs := strings.Split(cmdoutstr, "\n")
	for _, s := range outstrs {
		s = strings.TrimSpace(s)
		if len(s) > 0 {
			kvs := strings.Split(s, ": ")
			if len(kvs) >= 2 {
				val := strings.TrimSpace(kvs[1])
				switch strings.TrimSpace(kvs[0]) {
				case "名称", "Name":
					wlaninfo["Name"] = val
				case "说明", "Description":
					wlaninfo["Description"] = val
				case "GUID":
					wlaninfo["GUID"] = val
				case "物理地址", "Physical address":
					wlaninfo["Physical address"] = val
				case "状态", "State":
					wlaninfo["State"] = val
				case "SSID":
					wlaninfo["SSID"] = val
				case "BSSID":
					wlaninfo["BSSID"] = val
				case "网络类型", "Network type":
					wlaninfo["Network type"] = val
				case "无线电类型", "Radio type":
					wlaninfo["Network type"] = val
				case "身份验证", "Authentication":
					wlaninfo["Authentication"] = val
				case "密码", "Cipher":
					wlaninfo["Cipher"] = val
				case "连接模式", "Connection mode":
					wlaninfo["Connection mode"] = val
				case "频带", "Frequency":
					wlaninfo["Frequency"] = val
				case "通道", "Channel":
					wlaninfo["Channel"] = val
				case "接收速率(Mbps)", "Receive rate (Mbps)":
					wlaninfo["Receive rate"] = val
				case "传输速率 (Mbps)", "Transmit rate (Mbps)":
					wlaninfo["Transmit rate"] = val
				case "信号", "Signal":
					wlaninfo["Signal"] = val
				}
			}
		}
	}

	return wlaninfo
}

/*
***************************************************

	功能:字符串转码,将bgk码转换为utf8码
	输入:[s] 原始字符串
	输出:转码后的字符串
	说明:
	时间:2020年6月11日
	编辑:wang_jp

***************************************************
*/
func gbkToUtf8(s []byte) []byte {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := io.ReadAll(reader)
	if e != nil {
		return nil
	}
	var bstr []byte
	for _, c := range d {
		if c > 0 {
			bstr = append(bstr, c)
		}
	}
	return bstr
}
