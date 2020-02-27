package main

import (
	htf "go.momenta.works/activity_cadence/tasks/http_task_func"
)

func iw_manu_increment_merge(input []byte) ([]byte, error) {
	return htf.HttpProcess(input)
}
