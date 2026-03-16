package main

import "fmt"

func main() {
	var n []int
	numbers := []int{2, 2, 2, 2, 3, 3, 5, 4, 5, 5}
	sum := 0
	max := numbers[0]
	min := numbers[0]
	caunt := 0
	n = append(n, 2, 2, 4, 5, 6)
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])

	}
	for _, value := range numbers {

		sum += value
		if value > max {
			max = value
		}
		if value < min {
			min = value
		}
		if value%2 == 0 {
			caunt++
		}

	}
	fmt.Printf("Сумма: %d, Макс значение: %d Мин значение: %d Четное: %d Нечетное: %d ", sum, max, min, caunt, len(numbers)-caunt)
}

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
