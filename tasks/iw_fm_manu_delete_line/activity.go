package main

import (
	htf "go.momenta.works/activity_cadence/tasks/http_task_func"
)

func iw_fm_manu_delete_line(input []byte) ([]byte, error) {
	return htf.HttpProcess(input)
}
