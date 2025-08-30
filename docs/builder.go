type User struct {
	Name  string
	Email string
	Age   int
}

type UserBuilder struct {
 user User
}
func (ub *UserBuilder) SetName(name string) *UserBuilder {
 ub.user.Name = name
 return ub
}
func (ub *UserBuilder) SetEmail(email string) *UserBuilder {
 ub.user.Email = email
 return ub
}
func (ub *UserBuilder) SetAge(age int) *UserBuilder {
 ub.user.Age = age
 return ub
}
func (ub *UserBuilder) Build() User {
 return ub.user
}
func main() {
 user := (&UserBuilder{}).SetName("Alice").SetEmail("alice@mail.com").SetAge(30).Build()
 fmt.Println(user)
}