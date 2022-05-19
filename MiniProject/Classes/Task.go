package Classes

import "time"

type Task struct {
	id      string
	ownerId string
	status  Status
}

type Chore struct {
	task        Task
	description string
	size        Size
}

type HomeWork struct {
	task    Task
	course  string
	dueDate time.Time
	details string
}

func createChoreTask(id_curr string, ownerId_curr string, status_curr Status, description_curr string, size_curr Size) Chore {
	choretask := Task{id: id_curr, ownerId: ownerId_curr, status: status_curr}
	return Chore{task: choretask, description: description_curr, size: size_curr}
}

func createHomeWorkTask(id_curr string, ownerId_curr string, status_curr Status, course_curr string, dueDate_curr time.Time, details_curr string) HomeWork {
	homeWorkTask := Task{id: id_curr, ownerId: ownerId_curr, status: status_curr}
	return HomeWork{task: homeWorkTask, course: course_curr, dueDate: dueDate_curr, details: details_curr}
}

func (t Task) setOwner(newOwnerId string) {
	t.ownerId = newOwnerId
}

func (t Task) setStatus(newStatus Status) {
	t.status = newStatus
}
