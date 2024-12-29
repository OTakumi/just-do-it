package main

type Task struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (t *Task) Create() {
	println("Create Task")
}

func main() {
	println("Hello World")

	var sample_task = &Task{ID: "1", Title: "Sample", CreatedAt: "Today", UpdatedAt: "Today"}

	println(sample_task.ID, sample_task.Title)
}
