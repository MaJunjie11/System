package user_moudle

type UserInfo struct {
	Uid       int64
	Name      string
	Sex       int32
	Phone     string
	Age       int32
	Passworld string
	UserType  UserType
}

type UserType int32

const (
	UserType_Unknow  = 0
	UserType_Student = 1
	UserType_Teacher = 2
	UserType_Manager = 3
)
