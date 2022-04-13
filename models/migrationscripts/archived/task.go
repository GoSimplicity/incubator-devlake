package archived

import (
	"time"

	"gorm.io/datatypes"
)

type Task struct {
	Model
	Plugin        string `gorm:"index"`
	Options       datatypes.JSON
	Status        string
	Message       string
	Progress      float32
	FailedSubTask string
	PipelineId    uint64 `gorm:"index"`
	PipelineRow   int
	PipelineCol   int
	BeganAt       *time.Time
	FinishedAt    *time.Time `gorm:"index"`
	SpentSeconds  int
}

func (Task) TableName() string {
	return "_devlake_tasks"
}
