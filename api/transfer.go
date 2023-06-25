package api

import (
	// "database/sql"
	// "errors"
	// "fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	db "github.com/dasotd/gobank/db/sqlc"
	// "github.com/dasotd/gobank/token"
)

type transferRequest struct {
	FromAccountID int64  `json:"from_account_id" binding:"required,min=1"`
	ToAccountID   int64  `json:"to_account_id" binding:"required,min=1"`
	Amount        int64  `json:"amount" binding:"required,gt=0"`
	Currency      string `json:"currency" binding:"required"`
}

func (server *Server) createTransfer(ctx *gin.Context) {
	var req transferRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// fromAccount, valid := server.validAccount(ctx, req.FromAccountID, req.Currency)
	// if !valid {
	// 	return
	// }

	// authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	// if fromAccount.Owner != authPayload.Username {
	// 	err := errors.New("from account doesn't belong to the authenticated user")
	// 	ctx.JSON(http.StatusUnauthorized, errorResponse(err))
	// 	return
	// }

	// _, valid = server.validAccount(ctx, req.ToAccountID, req.Currency)
	// if !valid {
	// 	return
	// }

	arg := db.TransferTxParams{
		FromAccountID: req.FromAccountID,
		ToAccountID:   req.ToAccountID,
		Amount:        req.Amount,
	}

	result, err := server.bank.TransferTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, result)
}