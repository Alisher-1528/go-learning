package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	ID       int
	Title    string
	Completed bool
}

var tasks []Task
var nextID = 1

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n=== Менеджер задач ===")
		fmt.Println("1. Добавить задачу")
		fmt.Println("2. Показать все задачи")
		fmt.Println("3. Отметить задачу как выполненную")
		fmt.Println("4. Выход")
		fmt.Print("Выберите действие: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			addTask(reader)
		case "2":
			showTasks()
		case "3":
			markComplete(reader)
		case "4":
			fmt.Println("До свидания!")
			return
		default:
			fmt.Println("Неверный выбор, попробуйте снова.")
		}
	}
}

func addTask(reader *bufio.Reader) {
	fmt.Print("Введите название задачи: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	if title == "" {
		fmt.Println("Название задачи не может быть пустым.")
		return
	}

	task := Task{
		ID:        nextID,
		Title:     title,
		Completed: false,
	}
	tasks = append(tasks, task)
	nextID++

	fmt.Printf("Задача '%s' добавлена с ID %d\n", title, task.ID)
}

func showTasks() {
	if len(tasks) == 0 {
		fmt.Println("Список задач пуст.")
		return
	}

	fmt.Println("\n--- Список задач ---")
	for _, task := range tasks {
		status := "[ ]"
		if task.Completed {
			status = "[✓]"
		}
		fmt.Printf("%d. %s %s\n", task.ID, status, task.Title)
	}
}

func markComplete(reader *bufio.Reader) {
	showTasks()

	if len(tasks) == 0 {
		return
	}

	fmt.Print("Введите ID задачи для завершения: ")
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)

	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Неверный ID.")
		return
	}

	found := false
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Completed = true
			fmt.Printf("Задача '%s' отмечена как выполненная!\n", tasks[i].Title)
			found = true
			break
		}
	}

	if !found {
		fmt.Println("Задача с таким ID не найдена.")
	}
}
