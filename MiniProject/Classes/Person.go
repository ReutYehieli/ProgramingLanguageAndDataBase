package Classes

type Person struct {
	id                          string
	name                        string
	email                       string
	favoriteProgrammingLanguage string
	tasks                       []Task
}

//func CreatePerson(id_curr, name_curr, email_curr, favoriteProgrammingLanguage_curr) Person {
//id :=uuid.New()
////tasks =  initilize the array task.
//return Person{id: id_curr, name: name_curr, email: email_curr, favoriteProgrammingLanguage: favoriteProgrammingLanguage_curr, tasks}

//}

func addTask(task Task) bool {
	return false
}

func (p Person) listTasks() []Task {
	return p.tasks
}
