package main

import (
	htf "go.momenta.works/activity_cadence/tasks/http_task_func"
)

func get_skeleton(input []byte) ([]byte, error) {
	return htf.HttpProcess(input)
}
