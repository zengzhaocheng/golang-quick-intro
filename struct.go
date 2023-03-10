package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	println("========定义结构体==============")
	println("========创建新结构体==============")
	fmt.Println(Books{"Go", "www.go.com", "go tutorial", 12345})
	println("===========用kv格式创建===========")
	fmt.Println(Books{title: "Go", author: "www.go.com", subject: "go tutorial", book_id: 12345})
	println("==========忽略字段为0或空============")
	fmt.Println(Books{title: "Go", author: "www.runoob.com"})

	println("========访问结构体成员==============")

	var Book1, Book2 Books
	Book1.title = "Go"
	Book1.author = "runoob"
	Book1.subject = "go tutorial"
	Book1.book_id = 12345

	Book2.title = "python"
	Book2.author = "zeng"
	Book2.subject = "python tutorial"
	Book2.book_id = 12346
	fmt.Printf("Book 1 title : %s\n", Book1.title)
	fmt.Printf("Book 1 author : %s\n", Book1.author)
	fmt.Printf("Book 1 subject : %s\n", Book1.subject)
	fmt.Printf("Book 1 book_id : %d\n", Book1.book_id)

	/* 打印 Book2 信息 */
	fmt.Printf("Book 2 title : %s\n", Book2.title)
	fmt.Printf("Book 2 author : %s\n", Book2.author)
	fmt.Printf("Book 2 subject : %s\n", Book2.subject)
	fmt.Printf("Book 2 book_id : %d\n", Book2.book_id)

	println("========结构体作为函数参数==============")
	/* 打印 Book1 信息 */
	printBook(Book1)

	/* 打印 Book2 信息 */
	printBook(Book2)
	println("========结构体指针==============")

	/* 打印 Book1 信息 */
	printBook1(&Book1)

	/* 打印 Book2 信息 */
	printBook1(&Book2)
}

func printBook(book Books) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
}

func printBook1(book *Books) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
}
