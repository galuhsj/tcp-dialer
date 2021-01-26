package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"
	"workspace/practice/tcp-dialer/model"
	"workspace/practice/tcp-dialer/utils"
)

func main() {
	fmt.Println("tcp-dialer is running...")

	config := utils.GetConfig()

	utils.InitializeLogSetting(&config.Log)

	firstLoop := true
	for {
		// Loop delay control
		if !firstLoop {
			time.Sleep(
				time.Duration(
					config.DialerTriggerIntervalInMinute,
				) * time.Minute,
			)
		}

		var tcps []model.TCP

		// read tcpList file
		contents, err := ioutil.ReadFile(fmt.Sprintf("docs/%s", config.FileName))
		if err != nil {
			log.Fatal(err)
		}

		contentsArr := strings.Split(string(contents), "\n")
		for _, content := range contentsArr {
			indexOf := strings.Index(content, ":")
			if indexOf != -1 {
				// get ip
				ip := content[:indexOf]

				// get ports
				ports := content[indexOf+1:]

				tcps = append(tcps, model.TCP{
					IP:    ip,
					Ports: ports,
				})
			}
		}

		if len(tcps) > 0 {
			errorMsg := ""
			for _, tcp := range tcps {
				portsArr := strings.Split(tcp.Ports, config.PortSeparator)

				results := utils.TCPGather(tcp.IP, portsArr, config.DialTimeoutInMinute)
				for _, port := range portsArr {
					log.Println(fmt.Sprintf("%s:%s => %s", tcp.IP, port, results[port]))

					if results[port] == "failed" {
						errorMsg = errorMsg + "\n" + fmt.Sprintf("%s:%s => %s", tcp.IP, port, results[port])
					}
				}
			}

			// send email when dialing error
			if errorMsg != "" {
				utils.SendEmail(config.EmailReceiver, "TCP Dialer", errorMsg, "TCP Dialer", config.SMTP)
			}
		}

		firstLoop = false
	}
}
