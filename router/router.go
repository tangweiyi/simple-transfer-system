package router

import (
	"fmt"
	"strconv"
	"transfer-service/service"
	"transfer-service/view"

	"github.com/gin-gonic/gin"
)

type AccountRouter struct {
	accountService service.AccountService
}

func NewAccountHandler(accountService service.AccountService) AccountRouter {
	return AccountRouter{
		accountService: accountService,
	}
}

func (r *AccountRouter) Register(engine *gin.Engine) {
	engine.POST("/accounts", r.CreateAccount)
	engine.GET("/accounts/:account_id", r.QueryAccount)
	engine.POST("/transactions", r.MakeTransaction)
}

func (r *AccountRouter) CreateAccount(c *gin.Context) {
	req := view.CreateAccountRequest{}
	if e := c.BindJSON(&req); e != nil {
		c.JSON(200, view.NewInvalidRequestResp(e))
		return
	}
	err := r.accountService.CreateAccount(c, req.AccountID, req.InitialBalance)
	if err != nil {
		fmt.Printf("Error %+v", err)
		c.JSON(200, view.NewGeneralErrorResp(err))
		return
	}
	c.JSON(200, view.NewSuccessResp(nil))
}

func (r *AccountRouter) QueryAccount(c *gin.Context) {
	accountID, e := strconv.ParseInt(c.Param("account_id"), 10, 64)
	if e != nil {
		c.JSON(200, view.NewInvalidRequestResp(e))
		return
	}
	account, err := r.accountService.GetAccount(c, accountID)
	if err != nil {
		fmt.Printf("Error %+v", err)
		c.JSON(200, view.NewGeneralErrorResp(err))
		return
	}
	c.JSON(200, view.NewSuccessResp(view.BuildAccountQueryData(account)))
}

func (r *AccountRouter) MakeTransaction(c *gin.Context) {
	req := view.TransactionRequest{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(200, view.NewInvalidRequestResp(err))
		return
	}
	err := r.accountService.MakeTransfer(c, req.SourceAccountID, req.DestinationAccountID, req.Amount)
	if err != nil {
		fmt.Printf("Error %+v", err)
		c.JSON(200, view.NewGeneralErrorResp(err))
		return
	}
	c.JSON(200, view.NewSuccessResp(nil))
}
