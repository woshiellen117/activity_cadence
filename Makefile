default:
	go run tasks/get_skeleton/*.go
	go run tasks/iw_auto_clue_process/*.go
	go run tasks/iw_delete_by_roadcenter/*.go
	go run tasks/iw_fm_auto_topology/*.go
	go run tasks/iw_fm_auto_yantu/*.go
	go run tasks/iw_fm_manu_delete_line/*.go
	go run tasks/iw_fm_manu_lane/*.go
	go run tasks/iw_manu_increment_merge/*.go
	go run tasks/iw_mapping/*.go
	go run tasks/iw_mearth_3dtile_compiler/*.go
	go run tasks/iw_mesh/*.go
	go run tasks/iw_pole_auto_check_dup/*.go
	go run tasks/iw_pole_manu_check_dup/*.go
	go run tasks/iw_pose_write_back/*.go
	go run tasks/save_to_grid_store/*.go
	go run tasks/update_status/*.go
	go run tasks/wait_complete/*.go
