package main

import htf "go.momenta.works/activity_cadence/tasks/http_task_func"

func iw_pole_manu_check_dup(input []byte) ([]byte, error) {
	return htf.HttpProcess(input)
}
