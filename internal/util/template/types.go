package template

type Data struct {
	User    User
	Account Account
}

type User struct {
	Id    string
	Name  string
	Email string
}

type Account struct {
	Id   string
	Name string
}
