package dtos

type GetMsgCountInputDto struct {
	CustomerNumber string `uri:"customerNumber" binding:"required"`
}
