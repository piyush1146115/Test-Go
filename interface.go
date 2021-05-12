package main

import "log"

type User struct {
	Name  string
	Email string
}

//We have declared a new type called Admin and embedded
//the User type within the struct declaration. This is not inheritance but composition
type Admin struct {
	User
	Level string
}

//func (u *User) Notify() error {
//	panic("implement me")
//}

func (u *User) Notify() error {
	log.Printf("User: Sending User Email To %s<%s>\n",
		u.Name,
		u.Email)

	return nil
}

func (a *Admin) Notify() error {
	log.Printf("Admin: Sending Admin Email To %s<%s>\n",
		a.Name,
		a.Email)

	return nil
}

type Notifier interface {
	Notify() error
}

func SendNotification(notify Notifier) error {
	return notify.Notify()
}

func main() {

	// Value of type User can be used to call the method
	// with a pointer receiver.
	bill := User{"Bill", "bill@email.com"}
	bill.Notify()

	// Pointer of type User can be used to call the method
	// with a pointer receiver.
	jill := &User{"Jill", "jill@email.com"}
	jill.Notify()


	user := &User{
		Name:  "janet jones",
		Email: "janet@email.com",
	}
	SendNotification(user)

	admin := &Admin{
		User: User{
			Name:  "john smith",
			Email: "john@email.com",
		},
		Level: "super",
	}
	admin.User.Notify() //Here we are accessing the field and method set of the inner type through the use of the type’s name
	admin.Notify() // these fields and methods are also promoted to the outer type
	SendNotification(admin)
}