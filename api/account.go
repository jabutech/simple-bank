package api

import (
	"database/sql"
	db "github/jabutech/simplebank/db/sqlc"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Request create account struct
type createAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,oneof=USD EUR"`
}

// Function for handler createAccount
func (server *Server) createAccount(ctx *gin.Context) {
	var req createAccountRequest
	// Get request from client and check if error
	if err := ctx.ShouldBindJSON(&req); err != nil {
		// return response code 400 (bad request), and error
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// If data is valid and no error
	// Create new object argument
	arg := db.CreateAccountParams{
		Owner:    req.Owner,
		Currency: req.Currency,
		Balance:  0,
	}
	// Create new account, with send object argument
	account, err := server.store.CreateAccount(ctx, arg)
	// If error
	if err != nil {
		// return response code 500 (internal server error), and error
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// If no error, and create account success,
	// response code 200 (sucess), and new account data
	ctx.JSON(http.StatusOK, account)
}

// Get uri request for uri id
type getAccountRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

// Function handler getAccount by id
func (server *Server) getAccount(ctx *gin.Context) {
	var req getAccountRequest

	// Get uri id from client
	if err := ctx.ShouldBindUri(&req); err != nil {
		// return response code 400 (bad request), and error
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// If no error, get account by id
	account, err := server.store.GetAccount(ctx, req.ID)
	// If error
	if err != nil {
		// If error same with sql.ErrNpRows / data not found
		if err == sql.ErrNoRows {
			// response code 400 (not found), and error
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		// If no, response code 500 (internal server error), and error
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// If no error, and data is available,
	// Response code 200 (success), and send data account
	ctx.JSON(http.StatusOK, account)
}

// Get query request
type listAccountRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

// Function handler getAccount by id
func (server *Server) listAccount(ctx *gin.Context) {
	var req listAccountRequest

	// Get form query from client
	if err := ctx.ShouldBindQuery(&req); err != nil {
		// return response code 400 (bad request), and error
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Create object argument
	arg := db.ListAccountsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID) * req.PageSize,
	}

	// If no error, get list accounts
	account, err := server.store.ListAccounts(ctx, arg)
	// If error
	if err != nil {
		// If no, response code 500 (internal server error), and error
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// If no error, and data is available,
	// Response code 200 (success), and send data account
	ctx.JSON(http.StatusOK, account)
}
