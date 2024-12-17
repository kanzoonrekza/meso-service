package db

type TaskWithSubtasks struct {
	Task
	Subtasks []Task `json:"subtasks,omitempty"`
}
