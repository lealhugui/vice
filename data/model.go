package data

type Task struct {
	Id            int64
	RunnerHost    string
	RequirementId int64
	Requirement   Attribute `gorm:"foreignkey:RequirementId"`
	Cmd           string
}

type Worker struct {
	Id       int64
	Host     string
	IsActive bool
}

type Attribute struct {
	Id      int64
	Name    string
	CallCmd string
}

type WorkerAttribute struct {
	Id              int64
	WorkerId        int64
	Worker          Worker `gorm:"foreignkey:WorkerId"`
	FunctionalityId int64
	Functionality   Attribute `gorm:"foreignkey:FunctionalityId"`
	IsActive        bool
}
