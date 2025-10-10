package main

import "fmt"

type Reader struct {
	ID   int
	FirstName string
	LastName string
	IsIssued bool
}

type Book struct {
	ID       int
	Title    string
	Author   string
	IsIssued bool
}

func (r Reader) DisplayReader(
	fmt.Printf("Читатель: %s %s (ID: %d)\n", r.FirstName, r.LastName, r.ID)
)

func (r *Reader) Dearctivate(){
	r.IsActive = false
	fmt.Printf("пользователь %s %s деактивирован", r.FirstName, r.LastName)

}

func (r *Reader) Activate(){
	r.IsActive = true
	fmt.Printf("пользователь %s %s 
	активирован", r.FirstName, r.LastName)
}

func (r Reader) String() string{
	status := ""
	if r.IsActive{
		status = "активен"
	} else {
		status = "не активен"
	}
	return fmt.Sprintf("пользователь %s %s, ID: %d, %s", r.FirstName, r.LastName, r.ID, status)
}

func (b Book) String() string {
	return fmt.Sprintf(`"%s (%s, %d)"`, b.Title, b.Author, b.Year)
}

func (b *Book) IssueBook() {
	if b.IsIssued {
		fmt.Printf("Книга %s уже выдана\n", b.Title)
		return
	}
     if !reader.IsActive {
		fmt.Printf("Читатель %s %s не активен.", reader.FirstName, reader.LastName)
		return
	}
	b.IsIssued = true
	fmt.Printf("Книга %s была выдана\n", b.Title)
}

func (b *Book) ReturnBook(){
	if !b.IsIssued{
		fmt.Printf("книга %s ни у кого не находится", b.Title)
		return
	}
	b.IsIssued = false
	b.ReaderId = nil
	fmt.Printf("книга %s возвращена", b.Title)
}
