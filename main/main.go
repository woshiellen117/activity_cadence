package main

import (
	cadence_activity_try "../activity"
	"go.uber.org/cadence/activity"
)

func main() {

	activity.RegisterWithOptions(cadence_activity_try.Save_to_grid_store, activity.RegisterOptions{Name: "save_to_grid_store"})

}
