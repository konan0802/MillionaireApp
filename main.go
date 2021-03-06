package main

import (
	"log"
	"net/http"
	"os"

	handler "MillionaireApp/handler/rest"
	"MillionaireApp/infra"
	"MillionaireApp/usecase"

	"github.com/julienschmidt/httprouter"
)

func main() {
	// 依存関係を注入
	ss := infra.NewSpreadsheets(os.Getenv("SPREADSHEETID"), os.Getenv("CREDENTIALFILEPATH"))
	budgetInfra := infra.NewBudgetInfraSS(ss)
	budgetUseCase := usecase.NewBudgetUseCase(budgetInfra)
	budgetHandler := handler.NewBudgetHandler(budgetUseCase)
	receiptInfra := infra.NewReceiptInfraSS(ss)
	receiptUseCase := usecase.NewReceiptUseCase(receiptInfra)
	receiptHandler := handler.NewReceiptHandler(receiptUseCase)

	// ルーティングの設定
	router := httprouter.New()
	router.GET("/receipt/monthlyreceipts", receiptHandler.GetMonthlyReceipts)
	router.POST("/receipt/create", receiptHandler.Create)
	router.GET("/budget/monthlybudget", budgetHandler.GetMonthlybudget)

	// サーバ起動
	port := os.Getenv("PORT")
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal(err)
	}
}
