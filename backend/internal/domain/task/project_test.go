package task

import (
	"errors"
	"testing"
	"time"
)

func TestNewProject(t *testing.T) {
	projectType := mustProjectType(t, "work")

	project, err := NewProject("project-1", "user-1", projectType, "Build task domain")
	if err != nil {
		t.Fatalf("NewProject() error = %v", err)
	}

	if project.ID != "project-1" || project.UserID != "user-1" || project.Type != projectType {
		t.Errorf("project = %+v, want IDs and type to be set", project)
	}
	if project.Goal != "" || project.Description != "" || project.Progress != 0 {
		t.Errorf("project = %+v, want only required values to be set", project)
	}
}

func TestNewProjectWithDetails(t *testing.T) {
	projectType := mustProjectType(t, "work")
	startAt := time.Date(2026, 7, 18, 9, 0, 0, 0, time.UTC)
	endAt := startAt.Add(24 * time.Hour)

	project, err := NewProjectWithDetails("project-1", "user-1", projectType, "Build task domain", "Ship it", "Entity layer", 10, startAt, endAt)
	if err != nil {
		t.Fatalf("NewProjectWithDetails() error = %v", err)
	}

	if project.Goal != "Ship it" || project.Description != "Entity layer" || project.StartAt != startAt || project.EndAt != endAt {
		t.Errorf("project = %+v, want details to be set", project)
	}
}

func TestNewProjectValidation(t *testing.T) {
	projectType := mustProjectType(t, "work")

	tests := []struct {
		name        string
		id          ProjectID
		userID      UserID
		projectType ProjectType
		title       string
		wantErr     error
	}{
		{name: "empty ID", userID: "user-1", projectType: projectType, title: "Project", wantErr: ErrProjectIDEmpty},
		{name: "empty user ID", id: "project-1", projectType: projectType, title: "Project", wantErr: ErrProjectUserIDEmpty},
		{name: "invalid type", id: "project-1", userID: "user-1", projectType: ProjectType{Value: "fitness"}, title: "Project", wantErr: ErrProjectTypeInvalid},
		{name: "blank title", id: "project-1", userID: "user-1", projectType: projectType, title: " ", wantErr: ErrProjectTitleEmpty},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewProject(tt.id, tt.userID, tt.projectType, tt.title)
			assertTaskDomainErrorIs(t, err, tt.wantErr)
			if got.ID != tt.id || got.UserID != tt.userID || got.Title != tt.title {
				t.Errorf("project = %+v, want input values to be preserved", got)
			}
		})
	}
}

func TestNewProjectWithDetailsValidation(t *testing.T) {
	projectType := mustProjectType(t, "work")
	startAt := time.Date(2026, 7, 18, 9, 0, 0, 0, time.UTC)

	tests := []struct {
		name     string
		progress int
		startAt  time.Time
		endAt    time.Time
		wantErr  error
	}{
		{name: "negative progress", progress: -1, startAt: startAt, endAt: startAt.Add(time.Hour), wantErr: ErrProjectProgressInvalid},
		{name: "progress too high", progress: 101, startAt: startAt, endAt: startAt.Add(time.Hour), wantErr: ErrProjectProgressInvalid},
		{name: "end before start", progress: 0, startAt: startAt, endAt: startAt, wantErr: ErrProjectEndAtMustBeAfterStartAt},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewProjectWithDetails("project-1", "user-1", projectType, "Project", "", "", tt.progress, tt.startAt, tt.endAt)
			assertTaskDomainErrorIs(t, err, tt.wantErr)
			if got.Progress != tt.progress || got.StartAt != tt.startAt || got.EndAt != tt.endAt {
				t.Errorf("project = %+v, want input details to be preserved", got)
			}
		})
	}
}

func mustProjectType(t *testing.T, value string) ProjectType {
	t.Helper()
	projectType, err := NewProjectType(value)
	if err != nil {
		t.Fatalf("NewProjectType() error = %v", err)
	}
	return projectType
}

func assertTaskDomainErrorIs(t *testing.T, got, want error) {
	t.Helper()
	if !errors.Is(got, want) {
		t.Errorf("error = %v, want %v", got, want)
	}
}
