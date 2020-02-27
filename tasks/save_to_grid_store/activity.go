package main

import htf "go.momenta.works/activity_cadence/tasks/http_task_func"

func save_to_grid_store(input []byte) ([]byte, error) {
	return htf.HttpProcess(input)
}
