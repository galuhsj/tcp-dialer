package model

import "time"

// Config ...
type Config struct {
	DialerTriggerIntervalInMinute int
	Log                           Log
	FileName                      string
	PortSeparator                 string
	DialTimeoutInMinute           time.Duration
	EmailReceiver                 string
	SMTP                          SMTP
}

// Log ...
type Log struct {
	Filename   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
}

// SMTP ...
type SMTP struct {
	Host     string
	Port     int
	Sender   string
	Password string
}
