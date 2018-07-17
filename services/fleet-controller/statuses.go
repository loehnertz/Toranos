package main

import "github.com/loehnertz/Toranos/common"

const Error = "error"
const Reserved = "reserved"
const Driving = "driving"
const Done = "done"
const Canceled = "canceled"

func getStatusKeyByName(status string) int {
	return common.GetConfigIntByPath(conf, "service-settings", "fleet-controller", "statuses", status)
}
