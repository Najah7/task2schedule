package task

import "testing"

func TestNewTodoItem(t *testing.T) {
	item, err := NewTodoItem("todo-item-1", "task-1", "Buy milk")
	if err != nil {
		t.Fatalf("NewTodoItem() error = %v", err)
	}

	if item.ID != "todo-item-1" || item.TaskID != "task-1" || item.Position != 0 {
		t.Errorf("todo item = %+v, want ID, task ID, and position to be set", item)
	}
	if item.Description != "" || item.Completed || item.Position != 0 {
		t.Errorf("todo item = %+v, want only required values to be set", item)
	}
}

func TestNewTodoItemWithDetails(t *testing.T) {
	item, err := NewTodoItemWithDetails("todo-item-1", "task-1", "Buy milk", "At the store", true, 2)
	if err != nil {
		t.Fatalf("NewTodoItemWithDetails() error = %v", err)
	}

	if item.Description != "At the store" || !item.Completed || item.Position != 2 {
		t.Errorf("todo item = %+v, want details to be set", item)
	}
}

func TestNewTodoItemValidation(t *testing.T) {
	tests := []struct {
		name     string
		id       TodoItemID
		taskID   TaskID
		title    string
		position int
		wantErr  error
	}{
		{name: "empty ID", taskID: "task-1", title: "Todo", position: 0, wantErr: ErrTodoItemIDEmpty},
		{name: "empty task ID", id: "todo-item-1", title: "Todo", position: 0, wantErr: ErrTodoItemTaskIDEmpty},
		{name: "blank title", id: "todo-item-1", taskID: "task-1", title: " ", position: 0, wantErr: ErrTodoItemTitleEmpty},
		{name: "negative position", id: "todo-item-1", taskID: "task-1", title: "Todo", position: -1, wantErr: ErrTodoItemPositionLess},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTodoItemWithDetails(tt.id, tt.taskID, tt.title, "", false, tt.position)
			assertTaskDomainErrorIs(t, err, tt.wantErr)
			if got.ID != tt.id || got.TaskID != tt.taskID || got.Title != tt.title {
				t.Errorf("todo item = %+v, want input values to be preserved", got)
			}
		})
	}
}
