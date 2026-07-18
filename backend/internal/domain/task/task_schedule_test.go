package task

import (
	"testing"
	"time"
)

func TestNewTaskSchedule(t *testing.T) {
	schedule, err := NewTaskSchedule("schedule-1", "task-1", "Focus block")
	if err != nil {
		t.Fatalf("NewTaskSchedule() error = %v", err)
	}

	if schedule.ID != "schedule-1" || schedule.TaskID != "task-1" || schedule.Title != "Focus block" {
		t.Errorf("task schedule = %+v, want required values to be set", schedule)
	}
	if !schedule.StartAt.IsZero() || !schedule.EndAt.IsZero() || schedule.Description != "" {
		t.Errorf("task schedule = %+v, want only required values to be set", schedule)
	}
}

func TestNewTaskScheduleWithDetails(t *testing.T) {
	startAt := time.Date(2026, 7, 18, 9, 0, 0, 0, time.UTC)
	endAt := startAt.Add(time.Hour)
	dueAt := startAt.Add(2 * time.Hour)

	schedule, err := NewTaskScheduleWithDetails("schedule-1", "task-1", "Focus block", "Deep work", "Home", startAt, endAt, dueAt)
	if err != nil {
		t.Fatalf("NewTaskScheduleWithDetails() error = %v", err)
	}

	if schedule.ID != "schedule-1" || schedule.TaskID != "task-1" || schedule.DueAt != dueAt {
		t.Errorf("task schedule = %+v, want IDs and due time to be set", schedule)
	}
}

func TestNewTaskScheduleValidation(t *testing.T) {
	tests := []struct {
		name    string
		id      TaskScheduleID
		taskID  TaskID
		title   string
		wantErr error
	}{
		{name: "empty ID", taskID: "task-1", title: "Schedule", wantErr: ErrTaskScheduleIDEmpty},
		{name: "empty task ID", id: "schedule-1", title: "Schedule", wantErr: ErrTaskScheduleTaskIDEmpty},
		{name: "blank title", id: "schedule-1", taskID: "task-1", title: " ", wantErr: ErrTaskScheduleTitleEmpty},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTaskSchedule(tt.id, tt.taskID, tt.title)
			assertTaskDomainErrorIs(t, err, tt.wantErr)
			if got.ID != tt.id || got.TaskID != tt.taskID || got.Title != tt.title {
				t.Errorf("task schedule = %+v, want input values to be preserved", got)
			}
		})
	}
}

func TestNewTaskScheduleWithDetailsValidation(t *testing.T) {
	startAt := time.Date(2026, 7, 18, 9, 0, 0, 0, time.UTC)

	got, err := NewTaskScheduleWithDetails("schedule-1", "task-1", "Schedule", "", "", startAt, startAt, time.Time{})
	assertTaskDomainErrorIs(t, err, ErrTaskScheduleEndAtMustBeAfterStartAt)
	if got.StartAt != startAt || got.EndAt != startAt {
		t.Errorf("task schedule = %+v, want input details to be preserved", got)
	}
}
