package dto

type CollectionState string

const (
	CollectionStateInit    CollectionState = "init"
	CollectionStateSucceed CollectionState = "succeed"
	CollectionStateRemoved CollectionState = "removed"
)

type CollectionListReq struct {
	State    CollectionState `form:"state" binding:"required,oneof=init succeed removed"`
	Keyword  string          `form:"keyword" binding:"required"`
	PageSize int             `form:"pageSize" binding:"required"`
	PageNum  int             `form:"pageNum" binding:"required"`
}
