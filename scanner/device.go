package main

import "time"

type Device struct {
	IP           string
	Hostname     string
	Online       bool
	ResponseTime time.Duration
}
