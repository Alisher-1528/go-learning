package main

import "fmt"

func main() {
	N := 4
	A := 1
	B := 7
	count := 0
	for i := 0; i < B; i++ {
		if i != A && i != B {
			count++
			fmt.Println(" ", i)
		}

	}
	fmt.Print(count)
	for i := 0; i < N; i++ {
		fmt.Println("\n Я готовлюсь к собеседованию в Go!")

	} // твоё решение
	a := 1
	b := 6
	for i := a; i < b; i++ {
		fmt.Printf("%d^2 = %d\n", i, i*i)
	}

	m := 5             // например, 5 кг
	pricePerKg := 10.5 // цена за 1 кг

	// Стоимость = цена за кг * количество кг
	totalCost := pricePerKg * float64(m) // float64(N) - превращаем N в дробное число

	fmt.Printf("Задача For4: %d кг конфет стоят %.2f руб.\n\n", m, totalCost)
	// %.2f - выводит число с 2 знаками после запятой
	sena := 5.5
	for i := 0.1; i < 1.0; i += 0.1 {
		res := sena * i

		fmt.Printf("  %.1f кг = %.2f руб.\n", i, res)
	}
}

// type Student struct {
// 	ID    int
// 	Name  string
// 	Age   int
// 	Score float64
// }

// var Students []Student

// func (s *Student) ShowInfo() {

// 	fmt.Printf("ID: %d | Имя: %s | Возраст: %d | Балл: %f", s.ID, s.Name, s.Age, s.Score)
// }
// func addStudent() {
// 	var name string
// 	var age int
// 	var score float64
// 	fmt.Println("Ведите имя")
// 	fmt.Scan(&name)
// 	fmt.Println("Ведите возвраст")
// 	fmt.Scan(&age)
// 	fmt.Println("Ведите бал")
// 	fmt.Scan(&score)
// 	newStudents := Student{
// 		Name:  name,
// 		ID:    len(Students) + 1,
// 		Age:   age,
// 		Score: score,
// 	}
// 	Students = append(Students, newStudents)
// }
// func showAllStudents() {
// 	if len(Students) == 0 {
// 		fmt.Println("Нет студентов")
// 	}
// 	fmt.Println("Список всех студентов \n")
// 	for _, stud := range Students {
// 		fmt.Printf("ID: %d \nИмя: %s \nВозраст: %d \nБалл: %f", stud.ID, stud.Name, stud.Age, stud.Score)
// 	}
// }
// func findStudent() {
// 	var name string
// 	fmt.Println("\nВедите имя для поиска")
// 	fmt.Scan(&name)
// 	for _, find := range Students {
// 		if name == find.Name {
// 			fmt.Printf("Студент по имени %s найдено", name)
// 			fmt.Printf("\n ID: %d \nИмя: %s \nВозраст: %d \nБалл: %f", find.ID, find.Name, find.Age, find.Score)
// 		} else {
// 			fmt.Printf("Студент по имени %s не найдено", name)
// 		}

// 	}
// }
// func deleteStudent() {

// 	if len(Students) == 0 {
// 		fmt.Println("Нет студентов")
// 	}
// 	showAllStudents()

// 	var id int
// 	fmt.Print("Введите ID студента для удаления: ")
// 	fmt.Scan(&id)

// 	// Создаём новый срез
// 	var newStudents []Student

// 	// Копируем всех, кроме удаляемого
// 	for i := 0; i < len(Students); i++ {
// 		if Students[i].ID != id {
// 			newStudents = append(newStudents, Students[i])
// 		}
// 	}

// 	// Проверяем, нашли ли студента
// 	if len(newStudents) == len(Students) {
// 		fmt.Println("❌ Студент не найден")
// 		return
// 	}

// 	// Заменяем срез
// 	Students = newStudents
// 	fmt.Println("🗑️ Студент удалён")
// }

// func main() {
// 	fmt.Println("🎓 СИСТЕМА УПРАВЛЕНИЯ СТУДЕНТАМИ")

// 	for {
// 		fmt.Println("\n=== МЕНЮ ===")
// 		fmt.Println("1. Показать всех студентов")
// 		fmt.Println("2. Добавить студента")
// 		fmt.Println("3. Найти студента")
// 		fmt.Println("4. Удалить студента")
// 		fmt.Println("5. Увеличить балл")
// 		fmt.Println("6. Сохранить и выйти")
// 		fmt.Print("Выберите действие: ")

// 		var choice int
// 		fmt.Scan(&choice)

// 		switch choice {
// 		case 1:
// 			showAllStudents()
// 		case 2:
// 			addStudent()
// 		case 3:
// 			findStudent()
// 		case 4:
// 			deleteStudent()

// 			fmt.Println("👋 До свидания!")
// 			return
// 		default:
// 			fmt.Println("⚠️ Неверный выбор")
// 		}
// 	}
// }

// type cat struct {
// 	Name string
// }

// func (c cat) Say() string {
// 	return "MIYAV"
// }

// type Speaker interface {
// 	Say() string
// }

// func MakeSound(s Speaker) {
// 	fmt.Println(s.Say())
// }
// func main() {
// 	cate1 := cat{Name: "Sara"}
// 	MakeSound(cate1)
// }

// type Mashinka struct {
// 	Power int
// 	Name  string
// 	age   int
// }

// var Mashins []Mashinka

// func NewMashik(name string) {
// 	var power, agemashinki int

// 	fmt.Println("Ведите Мошности машинки ")
// 	fmt.Scan(&power)
// 	fmt.Println("Ведите год машинки ")
// 	fmt.Scan(&agemashinki)
// 	newMashin := Mashinka{
// 		Power: power,
// 		Name:  name,
// 		age:   agemashinki,
// 	}
// 	Mashins = append(Mashins, newMashin)
// 	fmt.Printf("Новая машинка %s с мощностию %d и год его выпуска %d", newMashin.Name, newMashin.Power, newMashin.age)

// }
// func shovMashink() {
// 	fmt.Println("\n Все машинки на скалде ")
// 	if len(Mashins) == 0 {
// 		fmt.Println("Склад пустой ")

// 	}
// 	for _, ShovMashin := range Mashins {
// 		fmt.Printf("На скаладе машинка %s  с мощностию %d и год его выпуска %d ", ShovMashin.Name, ShovMashin.Power, ShovMashin.age)
// 	}
// }
// func (n *Mashinka) UpPower() {
// 	n.Power = n.Power + 10
// }
// func main() {

// 	NewMashik("Bosh")
// 	shovMashink()
// 	Mashins[0].UpPower()
// 	fmt.Println("\n Стало:", Mashins[0].Power)
// }

// Структура Task (задача)
// type Task struct {
// 	ID    int    // номер задачи
// 	Title string // название
// 	Done  bool   // выполнена или нет
// }

// Глобальный срез для хранения всех задач
// var tasks []Task

// Функция добавления задачи
// func addTask() {
// 	var title string
// 	fmt.Print("Введите название задачи: ")
// 	fmt.Scan(&title)

// 	Создаём новую задачу
// 	newTask := Task{
// 		ID:    len(tasks) + 1, // ID = количество задач + 1
// 		Title: title,
// 		Done:  false,
// 	}

// 	Добавляем в срез
// 	tasks = append(tasks, newTask)
// 	fmt.Printf("✅ Задача добавлена (ID: %d)\n", newTask.ID)
// }

// Функция показа всех задач
// func listTasks() {
// 	Проверка на пустоту
// 	if len(tasks) == 0 {
// 		fmt.Println("Нет задач")
// 		return
// 	}

// 	fmt.Println("\n📋 Список задач:")
// 	for _, task := range tasks {
// 		Определяем статус: ✅ или ❌
// 		status := "❌"
// 		if task.Done {
// 			status = "✅"
// 		}
// 		fmt.Printf("%d. %s %s\n", task.ID, status, task.Title)
// 	}
// }

// Функция отметки задачи как выполненной (с указателем!)
// func completeTask() {
// 	var id int
// 	fmt.Print("Введите ID задачи для выполнения: ")
// 	fmt.Scan(&id)

// 	Ищем задачу по ID
// 	for i := 0; i < len(tasks); i++ {
// 		if tasks[i].ID == id {
// 			Берём УКАЗАТЕЛЬ на задачу в срезе
// 			taskPtr := &tasks[i]
// 			taskPtr.Done = true // меняем оригинал
// 			fmt.Printf("✅ Задача \"%s\" выполнена!\n", taskPtr.Title)
// 			return
// 		}
// 	}

// 	fmt.Println("❌ Задача не найдена")
// }

// Функция удаления задачи
// func deleteTask() {
// 	var id int
// 	fmt.Print("Введите ID задачи для удаления: ")
// 	fmt.Scan(&id)

// 	Ищем индекс задачи
// 	for i := 0; i < len(tasks); i++ {
// 		if tasks[i].ID == id {
// 			Удаляем элемент: берём элементы до i и после i
// 			tasks = append(tasks[:i], tasks[i+1:]...)
// 			fmt.Printf("🗑️ Задача (ID: %d) удалена\n", id)
// 			return
// 		}
// 	}

// 	fmt.Println("❌ Задача не найдена")
// }

// Функция редактирования задачи (с указателем!)
// func editTask() {
// 	var id int
// 	fmt.Print("Введите ID задачи для редактирования: ")
// 	fmt.Scan(&id)

// 	Ищем задачу по ID
// 	for i := 0; i < len(tasks); i++ {
// 		if tasks[i].ID == id {
// 			Берём УКАЗАТЕЛЬ на задачу
// 			taskPtr := &tasks[i]

// 			fmt.Printf("Текущее название: %s\n", taskPtr.Title)
// 			fmt.Print("Введите новое название: ")

// 			var newTitle string
// 			fmt.Scan(&newTitle)

// 			Если пользователь ввёл не пустую строку — меняем
// 			if newTitle != "" {
// 				taskPtr.Title = newTitle
// 				fmt.Printf("✅ Задача обновлена: \"%s\"\n", taskPtr.Title)
// 			} else {
// 				fmt.Println("Название не изменено")
// 			}
// 			return
// 		}
// 	}

// 	fmt.Println("❌ Задача не найдена")
// }

// Главная функция
// func main() {
// 	fmt.Println("📌 Добро пожаловать в Todo-лист!")

// 	Добавляем пример задачи для теста
// 	tasks = append(tasks, Task{
// 		ID:    1,
// 		Title: "Пример задачи",
// 		Done:  false,
// 	})

// 	for {
// 		Меню
// 		fmt.Println("\n=== СПИСОК ЗАДАЧ ===")
// 		fmt.Println("1. Показать все задачи")
// 		fmt.Println("2. Добавить задачу")
// 		fmt.Println("3. Выполнить задачу")
// 		fmt.Println("4. Редактировать задачу")
// 		fmt.Println("5. Удалить задачу")
// 		fmt.Println("6. Выйти")
// 		fmt.Print("Выберите действие (1-6): ")

// 		var choice int
// 		fmt.Scan(&choice)

// 		switch choice {
// 		case 1:
// 			listTasks()
// 		case 2:
// 			addTask()
// 		case 3:
// 			completeTask()
// 		case 4:
// 			editTask()
// 		case 5:
// 			deleteTask()
// 		case 6:
// 			fmt.Println("👋 До свидания!")
// 			return
// 		default:
// 			fmt.Println("⚠️ Неверный выбор, попробуйте снова")
// 		}
// 	}
// }

// type book struct {
// 	Title  string
// 	Author string
// 	Year   int
// 	Rating int
// }

// var books []book

// func main() {
// 	books = append(books, book{
// 		Title:  "Мастер и Маргарита",
// 		Author: "Булгаков",
// 		Year:   1967,
// 		Rating: 5,
// 	})
// 	books = append(books, book{
// 		Title:  "Преступление и наказание",
// 		Author: "Достоевский",
// 		Year:   1866,
// 		Rating: 5,
// 	})
// 	for {
// 		fmt.Println("\n=== КАТАЛОГ КНИГ ===")
// 		fmt.Println("1. Показать все книги")
// 		fmt.Println("2. Добавить книгу")
// 		fmt.Println("3. Найти по автору")
// 		fmt.Println("4. Книги с рейтингом 5")
// 		fmt.Println("5. Выйти")
// 		fmt.Print("Выберите действие: ")

// 		var choice int
// 		fmt.Scan(&choice)

// 		switch choice {
// 		case 1:
// 			showAllBooks()
// 		case 2:
// 			addBook()
// 		case 3:
// 			findByAuthor()
// 		case 4:
// 			topRated()
// 		case 5:
// 			fmt.Println("До свидания!")
// 			return
// 		default:
// 			fmt.Println("Неверный выбор!")
// 		}
// 	}
// }
// func showAllBooks() {
// 	for i, book := range books {
// 		fmt.Printf("%d. %s (%s, %d г.) - Рейтинг: %d/5\n",
// 			i+1, book.Title, book.Author, book.Year, book.Rating)

// 	}
// }
// func addBook() {
// 	var title, author string
// 	var year, rating int
// 	fmt.Print("Введите название: ")
// 	fmt.Scan(&title)
// 	fmt.Print("Введите автора: ")
// 	fmt.Scan(&author)
// 	fmt.Print("Введите год издания: ")
// 	fmt.Scan(&year)
// 	fmt.Print("Введите рейтинг (1-5): ")
// 	fmt.Scan(&rating)
// 	if rating < 1 || rating > 5 {
// 		fmt.Println("❌ Рейтинг должен быть от 1 до 5!")
// 		return
// 	}
// 	newbook := book{
// 		Title:  title,
// 		Author: author,
// 		Year:   year,
// 		Rating: rating,
// 	}
// 	books = append(books, newbook)
// 	fmt.Printf("✅ Книга \"%s\" добавлена!\n", title)

// }
// func findByAuthor() {
// 	var author string
// 	fmt.Print("Введите автора: ")
// 	fmt.Scan(&author)
// 	flags := false
// 	fmt.Printf("\n📖 Книги автора %s:\n", author)
// 	for _, book := range books {
// 		if author == book.Author {
// 			fmt.Printf("  - %s (%d г.) - Рейтинг: %d/5\n", book.Author, book.Year, book.Rating)
// 			flags = true

// 		}
// 	}
// 	if !flags {
// 		fmt.Printf("❌ Книги автора %s не найдены\n", author)
// 	}

// }
// func topRated() {
// 	falgs := false
// 	for _, book := range books {
// 		if book.Rating == 5 {
// 			falgs = true
// 			fmt.Printf(book.Author, book.Title, book.Year, book.Rating)
// 		}
// 	}
// 	if !falgs {
// 		fmt.Println("Нет книг с таким рейтингом")
// 	}

// }

// var phnebooks = make(map[string]string)

// func main() {
// 	phnebooks["alisher"] = "+7 -983-348-52-42"
// 	phnebooks["Tah"] = "+7 -983-348-52-42"
// 	for {
// 		var chois int
// 		fmt.Println("\n=== МЕНЮ ===")
// 		fmt.Println("1. Показать все контакты")
// 		fmt.Println("2. Добавить контакт")
// 		fmt.Println("3. Найти контакт")
// 		fmt.Println("4. Удалить контакт")
// 		fmt.Println("5. Показать количество контактов")
// 		fmt.Println("6. Выйти")
// 		fmt.Print("Выберите действие (1-5): ")
// 		fmt.Scan(&chois)
// 		switch chois {
// 		case 1:
// 			showcontakts()
// 		case 2:
// 			addcontact()
// 		case 3:
// 			findcontact()
// 		case 4:
// 			deletcontact()
// 		case 5:
// 			fmt.Println("Контактов:", getcontact())
// 		case 6:
// 			fmt.Println("👋 До свидания!")
// 			return // Выход из программыs
// 		}
// 	}
// }
// func showcontakts() {
// 	if len(phnebooks) == 0 {
// 		fmt.Println("Телефонная книга пуста")
// 		return
// 	}
// 	fmt.Println("\n📖 Все контакты:")
// 	for name, number := range phnebooks {
// 		fmt.Printf("  %s: %s\n", name, number)
// 	}
// }
// func addcontact() {
// 	var name, number string
// 	fmt.Println("Ведите имя контакта ")
// 	fmt.Scan(&name)
// 	fmt.Println("Ведите номер контакта ")
// 	fmt.Scan(&number)
// 	phnebooks[name] = number
// 	fmt.Printf("Контакт %s добавлен ", name)

// }
// func findcontact() {
// 	var name string
// 	fmt.Println("Ведите имя контакта ")
// 	fmt.Scan(&name)
// 	phone, exists := phnebooks[name]
// 	if exists {
// 		fmt.Printf("📞 %s: %s\n", name, phone)
// 	} else {
// 		fmt.Printf("❌ Абонент %s не найден\n", name)
// 	}

// }
// func deletcontact() {
// 	var name string
// 	fmt.Println("Ведите имя контакта для удаление ")
// 	fmt.Scan(&name)

// 	_, exists := phnebooks[name]
// 	if exists {
// 		delete(phnebooks, name)
// 		fmt.Printf("🗑️ Контакт %s удалён\n", name)
// 	} else {
// 		fmt.Printf("❌ Абонент %s не найден\n", name)
// 	}

// }
// func getcontact() int {
// 	return len(phnebooks)
// }

// number := map[string]string{
// 	"alisher": "+7 -983-348-52-42",
// 	"homid":   "+7 -983-348-52-42",
// 	"fotima":  "+7 -983-348-52-42",
// 	"Ganj":    "+7 -983-348-52-42",
// 	"Tah":     "+7 -983-348-52-42",
// }
// for {
// 	fmt.Println("\n=== МЕНЮ ===")
// 	fmt.Println("1. Показать все контакты")
// 	fmt.Println("2. Добавить контакт")
// 	fmt.Println("3. Найти контакт")
// 	fmt.Println("4. Удалить контакт")
// 	fmt.Println("5. Выйти")
// 	fmt.Print("Выберите действие (1-5): ")

// 	var choice int
// 	fmt.Scan(&choice)
// 	switch choice {
// 	case 1:
// 		fmt.Print("1. Показать все контакты\n")
// 		if len(number) == 0 {
// 			fmt.Print("Телефонная книга пуста")
// 		}
// 		for name, numbers := range number {
// 			fmt.Printf("%s: %s\n", name, numbers)
// 		}
// 	case 2:
// 		fmt.Println(" Добавить контакт")
// 		var name, numberss string
// 		fmt.Println("Ведите номер")
// 		fmt.Scan(&numberss)
// 		fmt.Println("Ведите имя")
// 		fmt.Scan(&name)
// 		number[name] = numberss
// 		fmt.Printf("✅ Контакт %s добавлен!\n", name)
// 	case 3:
// 		fmt.Println("Найти контакт")
// 		var name string
// 		fmt.Println("Ведите имя для поиска")
// 		fmt.Scan(&name)
// 		number, exits := number[name]
// 		if exits {
// 			fmt.Printf("📞 %s: %s\n", name, number)
// 		} else {
// 			fmt.Printf("❌ Абонент %s не найден\n", number)
// 		}
// 	}
//}
// var n []int
// numbers := []int{2, 2, 2, 2, 3, 3, 5, 4, 5, 5}
// sum := 0
// max := numbers[0]
// min := numbers[0]
// caunt := 0
// n = append(n, 2, 2, 4, 5, 6)
// for i := 0; i < len(numbers); i++ {
// 	fmt.Println(numbers[i])

// }
// for _, value := range numbers {

// 	sum += value
// 	if value > max {
// 		max = value
// 	}
// 	if value < min {
// 		min = value
// 	}
// 	if value%2 == 0 {
// 		caunt++
// 	}

// }
// fmt.Printf("Сумма: %d Макс значение: %d Мин значение: %d Четное: %d Нечетное: %d ", sum, max, min, caunt, len(numbers)-caunt)
// var n, m int
// var s string
// fmt.Println("Ведите число: ")
// fmt.Scan(&n)
// fmt.Println("Ведите число до скольки умножать: ")
// fmt.Scan(&m)
// fmt.Println("Ведите матиматическую операцию хотите провести(*,/,-,+): ")
// fmt.Scan(&s)
// for i := 0; i < m; i++ {
// 	switch s {
// 	case "*":
// 		fmt.Printf("%d x %d = %d\n", n, i, n*i)
// 	case "+":
// 		fmt.Printf("%d + %d = %d\n", n, i, n+i)
// 	case "-":
// 		fmt.Printf("%d - %d = %d\n", n, i, n-i)
// 	case "/":
// 		fmt.Printf("%d / %d = %d\n", n, i, n/i)
// 	default:
// 		fmt.Println("Неправильная операция ")

// 	}

//}
//var name string
// var age int
// var rost float64
// var student string
// var sity string
// const BIRTH_YEAR = 1996

// fmt.Println("Ведите имя")
// fmt.Scan(&name)
// fmt.Println("Возраст ")
// fmt.Scan(&age)
// fmt.Println("Ваш рост ")
// fmt.Scan(&rost)
// fmt.Println("Вы студент  ")
// fmt.Scan(&student)

// fmt.Println("С какого города вы ")
// fmt.Scan(&sity)
// calculatedAge := 2026 - BIRTH_YEAR

// // Выводим результаты
// fmt.Println("\n--- Личная карточка ---")
// fmt.Println("Имя:", name)
// fmt.Println("Возраст:", age)
// if age > 0 && age <=12 {
// 	fmt.Println("Возрастная группа: Ребенок ")
// } else if age >=13 && age <= 17{
// 	fmt.Println("Возрастная группа: Подросток ")
// } else if age>=18 && age <=25{
// 	fmt.Println("Возрастная группа: Молодой взрослый ")
// }else if age>=26 && age <=60{
// 	fmt.Println("Возрастная группа: Взрослый ")
// } else {
// 	fmt.Println("Возрастная группа: Пенсионер ")
// }

// fmt.Println("Рост:", rost, "м")
// fmt.Println("Студент:", student)
// fmt.Println("Город:", sity)
// fmt.Println("Год рождения (из константы):", BIRTH_YEAR)
// fmt.Println("Возраст по году рождения:", calculatedAge)
// if calculatedAge == age {
// 	fmt.Println("✅ Возраст совпадает с годом рождения")
// } else {
// 	fmt.Println("⚠️ Возраст не совпадает с годом рождения")
// }

/*
	lens := 0
	fmt.Print("Ведите длину масива ")
	fmt.Scan(&lens)

	numbers := make([]int, lens)
	for i := 0; i < lens; i++ {
		fmt.Printf("Введите число %d: ", i+1)
		fmt.Scan(&numbers[i])
	}
	fmt.Println("Исходный срез  ", numbers)

	max := numbers[0]
	min := numbers[0]
	sums := 0
	cout := 0
	leg := len(numbers)
	fmt.Println(leg)

	for _, num := range numbers {
		if num%2 == 0 {
			cout++
		}
		sums += num
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}

	}
	fmt.Println("Максимальное число ", max)
	fmt.Println("Минимальное число ", min)
	fmt.Println("Количество четных ", cout)
	fmt.Println("Сумма ", sums)
*/
//}

/*m := []int{1, 2, 3, 4, 5}
	for _, r := range m {
		print(r)
	}
	rav := []struct {
		name  string
		agewe int
	}{
		{"alsiher", 2},
		{"alisher", 3},
	}
	fmt.Println(m, rav)
	var n [2]string
	n[0] = "ALisher"
	n[1] = "salom"
	fmt.Println(n[0], n[1])
	fmt.Println(n)
	type student struct {
		Name string
		age  int
		bal  int
	}
	Alisher := student{Name: "Alisher Mirzoev", age: 30, bal: 5}

	fmt.Println(Alisher)

	var a int
	i := 10
	p := &i
	*p = 2
	fmt.Println(i)

	P := 10
	var count, propeg int
	propeg = 11

	for propeg < 200 {
		propeg = propeg + P
		count++

	}
	fmt.Print(count, " за столько дней пробегажал ", propeg)
	fmt.Scan(&a)

	switch a {
	case 12, 1, 2:
		fmt.Println("Зима")
	case 3, 4, 5:
		fmt.Println("Весна")
	case 6, 7, 8:
		fmt.Println("Лето")
	case 9, 10, 11:
		fmt.Println("Осень")
	default:
		fmt.Println("Нет такое время года")

	}
	kartafan := 2.01
	for i := 0.1; i < 1; i = i + 0.1 {
		prise := kartafan * i
		fmt.Println(prise)
	}

}

/*http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	for _, p := range products {
		fmt.Fprintln(w, p.Name, p.Price, p.Strock)
	}
})
http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Главная страница")

})
http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
	for _, p := range products {
		fmt.Fprintln(w, p.Name, p.Price, p.Strock)
	}
})

fmt.Println("Server started on :8080")
http.ListenAndServe(":8080", nil)
*/
//}

/*fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)

	products := []Product{
		{Name: "Apple", Price: 1100, Strock: 3},
		{Name: "Naushnik", Price: 1200, Strock: 5},
		{Name: "Mishka", Price: 8, Strock: 5},
		{Name: "Hamid", Price: 100, Strock: 241},
	}
	expensive := filterExpensiveProducts(products)

	expensiveStock := filterByStock(expensive, 5)
	for _, p := range expensiveStock {
		fmt.Println(p.Name, "цена", p.Price, "на скадле", p.Strock)

	}
	fmt.Println("Всего дорогих товаров:", len(expensive))

	for _, p := range expensive {
		fmt.Println(p.Name, "цена", p.Price, "на скадле", p.Strock)

	}
	fmt.Println("Всего дорогих товаров:", len(expensive))
}

/*


func isExpensive(value int) bool {
	return value >= 1000
}
func filterExpensiveProducts(product []Product) []Product {
	var result []Product
	for _, st := range product {
		if isExpensive(st.Price) {
			result = append(result, st)
		}

	}
	return result
}
func main() {
	products := []Product{
		{Name: "Apple", Price: 1000, Strock: 100},
		{Name: "Samsung", Price: 2000, Strock: 50},
		{Name: "Mause", Price: 4000, Strock: 40},
	}
	expensive := filterExpensiveProducts(products)
	for _, p := range expensive {
		fmt.Println(p.Name, "цена", p.Price, "на скадле", p.Strock)

	}
	fmt.Println("Всего дорогих товаров:", len(expensive))
}

/*func sadlExamen(value int) bool {
	return value >= 60

}

type User struct {
	Name  string
	Sorse int
	age   int
	City  string
}
type product struct {
	Name   string
	Price  int
	Strock int
}




func main() {
	count := 0
	student := make(map[string]User)

	student["Salimjon"] = User{Name: "Aliser", Sorse: 100, age: 23}
	student["Tahmin"] = User{Name: "Tahmin", Sorse: 66, age: 22, City: "Bohtar"}
	student["Lahmin"] = User{Name: "Lahmin", Sorse: 77, age: 22, City: "Bohtar"}
	for product, price := range student {
		if sadlExamen(price.Sorse) {
			fmt.Println(product, "Здал экзамен на ", price.Sorse, "балов", "Город ", price.City)
			count++
		}

	}
	fmt.Println(count, "студентов сдали экзамен  больше 60 балов ")

}

/*func isEven(n int) bool {
	return n%2 == 0
}

type User struct {
	Name strfmt
	age  int
}

func isAdult(user User) bool {
	return user.Name == "Alish"

}
func main() {
	users := make(map[string]User)
users["Ali"] = User{Name: "Alisher", age: 23}
	count := 0
	user1 := User{
		Name: "Alisher",
		age:  9,
	}
	user2 := User{
		Name: "Ali",
		age:  26,
	}
	user3 := User{
		Name: "Alis",
		age:  21,
	}
	user4 := User{
		Name: "Alish",
		age:  22,
	}
	users := []User{user1, user2, user3, user4}

	for _, us := range users {
		if isAdult(us) {
			fmt.Println(us.Name, "is adult")
			count++
		}
	}
	fmt.Println("количество", count)
	}
*/

/*count := 0
ages := []int{15, 22, 30, 17, 40}
for _, age := range ages {

	if isAdult(age) {
		fmt.Println("Взросыле", age)
		count++
	}
}
fmt.Println("количество", count)
*/
