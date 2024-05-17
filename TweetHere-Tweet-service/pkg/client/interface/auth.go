package interfaces


type AuthClient interface{
	DoesUserExist(id int64)(bool,error)
	FindUserName(id int64)(string,error)
}