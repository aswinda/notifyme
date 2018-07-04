package interfaces

type IUserService interface {
	GetUserDetail(userId int) (string, error)
}