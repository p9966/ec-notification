package repository

type IRestfulConstraint interface {
	GetMsgCount(customerNumber string, configName string) (string, error)
}
