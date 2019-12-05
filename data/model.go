package data

type Task struct {
	Id         int64
	RunnerHost string
	RunnerType string
	Cmd        string
}

type Worker struct {
	Id       int64
	Host     string
	IsActive bool
}
