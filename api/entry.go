package api

import (
	"net/http"

	db "github.com/dasotd/gobank/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createEntryRequest struct {
	Amount    int64 `json:"amount" binding:"required"`
	AccountID int64 `json:"accountId" binding:"required"`
}

func (server *Server) createEntry(ctx *gin.Context){
	var req createEntryRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.CreateEntryParams{
		Amount: req.Amount,
		AccountID: req.AccountID,
	}

	entry, err := server.bank.CreateEntry(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusCreated, entry)
}

type getEntryRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}
func (server *Server) getEntry(ctx *gin.Context) {
	var req getEntryRequest
	if err := ctx.BindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))

	}


	entry, err := server.bank.GetEntry(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusCreated, entry)

}


type listEntryRequest struct {
	AccountID int64 `uri:"accountId" binding:"required,min=1"`
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func(server *Server) listEntry(ctx *gin.Context){
	var req listEntryRequest
	if err:= ctx.BindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))

	}

	args:= db.ListEntriesParams{
		AccountID: req.AccountID,
		Limit: req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	
	}

	entries, err := server.bank.ListEntries(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusCreated, entries)

}