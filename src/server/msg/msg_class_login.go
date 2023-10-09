package msg

type Login struct {
	UserId   string
	Password string
}
type LoginBack struct {
	LogResault   bool
	LoginResault string
}
type Register struct {
	UserId   string
	Password string
	UserName string
}
type RegisterBack struct {
	RegisterRes bool
	RegisterMsg string
}
