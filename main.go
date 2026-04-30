package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Note struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var Notes []Note
var nextID = 1

func addNote() {
	var title, content string
	fmt.Print("Введите заголовок: ")
	fmt.Scan(&title)
	fmt.Print("Введите содержание: ")
	fmt.Scan(&content)

	newNote := Note{
		ID:      nextID,
		Title:   title,
		Content: content,
	}
	Notes = append(Notes, newNote)
	fmt.Printf("✅ Заметка добавлена! ID: %d\n", nextID)
	nextID++
}

func showAllNotes() {
	if len(Notes) == 0 {
		fmt.Println("📝 ===== Нет заметок =====")
		return
	}
	fmt.Println("📝 ===== ВСЕ ЗАМЕТКИ =====")
	for _, note := range Notes {
		fmt.Printf("[ID: %d] %s - %s\n", note.ID, note.Title, note.Content)
	}
}

func editNote() {
	if len(Notes) == 0 {
		fmt.Println("📝 ===== Нет заметок =====")
		return
	}
	var id int
	fmt.Print("Введите ID: ")
	fmt.Scan(&id)

	for i := 0; i < len(Notes); i++ {
		if Notes[i].ID == id {
			fmt.Print("Новый заголовок: ")
			fmt.Scan(&Notes[i].Title)
			fmt.Print("Новое содержание: ")
			fmt.Scan(&Notes[i].Content)
			fmt.Println("✅ Обновлено")
			return
		}
	}
	fmt.Println("❌ Заметка с таким ID не найдена")
}

func deleteNote() {
	if len(Notes) == 0 {
		fmt.Println("📝 ===== Нет заметок =====")
		return
	}
	var id int
	fmt.Print("Введите ID для удаления: ")
	fmt.Scan(&id)

	for i := 0; i < len(Notes); i++ {
		if Notes[i].ID == id {
			Notes = append(Notes[:i], Notes[i+1:]...)
			fmt.Println("🗑️ Заметка удалена")
			return
		}
	}
	fmt.Println("❌ Заметка с таким ID не найдена")
}

func searchByTitle() {
	var title string
	fmt.Print("Введите заголовок для поиска: ")
	fmt.Scan(&title)

	found := false
	for _, note := range Notes {
		if note.Title == title {
			fmt.Printf("🔍 [ID:%d] %s - %s\n", note.ID, note.Title, note.Content)
			found = true
		}
	}
	if !found {
		fmt.Println("❌ Заметка не найдена")
	}
}

// HTTP сервер
func startHTTPServer() {
	http.HandleFunc("/notes", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method == "GET" {
			json.NewEncoder(w).Encode(Notes)
			return
		}

		if r.Method == "POST" {
			var newNote Note
			json.NewDecoder(r.Body).Decode(&newNote)
			newNote.ID = nextID
			nextID++
			Notes = append(Notes, newNote)
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(newNote)
			return
		}
	})

	http.HandleFunc("/notes/search", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		title := r.URL.Query().Get("title")

		var result []Note
		for _, note := range Notes {
			if note.Title == title {
				result = append(result, note)
			}
		}
		json.NewEncoder(w).Encode(result)
	})

	fmt.Println("🌐 HTTP сервер запущен на http://localhost:8080")
	fmt.Println("GET  /notes          - все заметки")
	fmt.Println("POST /notes          - создать заметку")
	fmt.Println("GET  /notes/search?title=... - поиск")
	http.ListenAndServe(":8080", nil)
}

func main() {
	// Запускаем HTTP сервер в фоне
	go startHTTPServer()

	// Консольное меню
	for {
		fmt.Println("\n---------------------------------")
		fmt.Println("1. Добавить 2. Показать 3. Редактировать 4. Удалить 5. Поиск 6. Выход")
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			addNote()
		case 2:
			showAllNotes()
		case 3:
			editNote()
		case 4:
			deleteNote()
		case 5:
			searchByTitle()
		case 6:
			fmt.Println("👋 До свидания!")
			return
		default:
			fmt.Println("⚠️ Неверный выбор")
		}
	}
}
