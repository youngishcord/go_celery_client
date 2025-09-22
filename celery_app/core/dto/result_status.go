package dto

type Status string

const (
	SUCCESS Status = "SUCCESS"
	FAILURE Status = "FAILURE"
	RETRY   Status = "RETRY"
	STARTED Status = "STARTED"
	PENDING Status = "PENDING"
)
