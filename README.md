# Init
1. create file docs/tcpList.txt with format ip:ports
    ex:
        10.10.10.10:3000,8000
        20.20.20.20:3000
1. create file config/config.json
    {
        "DialerTriggerIntervalInMinute": 30,
        "Log": {
            "Filename": "tcp-dialer.log",
            "MaxSize": 10,
            "MaxBackups": 5,
            "MaxAge": 120,
            "Compress": true
        },
        "FileName": "tcpList.txt",
        "PortSeparator": ",",
        "DialTimeoutInMinute": 3,
        "EmailReceiver": "",
        "SMTP": {
            "Host": "",
            "Port": 0,
            "Sender": "",
            "Password": ""
        }
    }
2. fill the SMTP config in config/config.json for error dial email notif

# Run locally
1. go run main.go

# Build & run for non docker image
1. ./build.sh
2. tar -xzf build.gz
3. ./tcp-dialer &

# Build & run for docker image
1. docker build --rm -f "Dockerfile" -t tcp-dialer:latest .
2. docker run -d -it tcp-dialer:latest

# Log
1. tcp-dialer.log
2. email (error only)
