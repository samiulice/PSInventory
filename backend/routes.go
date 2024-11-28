package backend

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	//Secure routes
	mux.Route("/api", func(mux chi.Router) {
		// mux.Use(app.AuthAdmin)
		mux.Post("/get-dash-board-data", app.FetchDashBoardData)
		mux.Post("/admin/add-stake-holder", app.AddNewStakeHolder)
		mux.Post("/company-profile", app.FetchCompanyProfile)

		//HR Management
		mux.Post("/hr/view-all-employees", app.GetAllEmployees)
		mux.Post("/hr/add-employee", app.AddEmployee)
		//MIS
		mux.Post("/mis/view-customer/{type}", app.GetCustomers)
		mux.Post("/mis/view-all-customers", app.GetAllCustomers)
		mux.Post("/mis/add-customer", app.AddCustomer)
		mux.Post("/mis/view-supplier/{type}", app.GetSuppliers)
		mux.Post("/mis/view-all-suppliers", app.GetAllSuppliers)
		mux.Post("/mis/get-supplier-id-name-list", app.GetActiveSuppliersIDAndName)
		mux.Post("/mis/get-customer-id-name-list", app.GetActiveCustomersIDAndName)
		mux.Post("/mis/add-supplier", app.AddSupplier)

		//Inventory
		mux.Post("/inventory/add-brand", app.AddBrand)
		mux.Post("/inventory/add-category", app.AddCategory)
		mux.Post("/inventory/add-product", app.AddProduct)
		//memo--products list
		mux.Post("/inventory/memo/get-purchase-product-list", app.FetchPurchaseMemoProductItems)

		mux.Post("/inventory/products/get-list-by-id", app.FetchProductItemsByProductID)
		mux.Post("/inventory/products/search-instock-products-by-serial", app.FetchInstockProductItemBySerialNumber) //search in-stock items
		mux.Post("/inventory/products/search-sold-products-by-serial", app.FetchSoldProductItemBySerialNumber)       //search sale/sold items
		mux.Post("/inventory/products/search-products-by-serial", app.FetchProductItemBySerialNumber)                //search all type items
		//memo--supplier
		mux.Post("/inventory/get-supplier-memo-list", app.GetMemoListBySupplierID)
		//memo--customer
		//sale-return page
		mux.Post("/inventory/memo/get-sales-product-list", app.FetchProductItemsBySalesHistory)
		mux.Post("/inventory/get-customer-memo-list", app.GetMemoListByCustomerID)
		mux.Post("/inventory/sale/return-products", app.ReturnProductsFromCustomer)
		//purchase
		mux.Post("/inventory/restock-product", app.RestockProduct)
		//purchase-return
		mux.Post("/inventory/return-product-to-supplier", app.ReturnProductsToSupplier)
		//sale
		mux.Post("/inventory/sale-products", app.SaleProductsToCustomer)
		//page details
		mux.Post("/inventory/purchase/get-page-details", app.GetPurchasePageDetails)
		mux.Post("/inventory/sale/get-page-details", app.GetSalePageDetails)
		//warranty

		mux.Post("/inventory/products/claim-warranty-by-serial-id", app.ClaimWarrantyBySerialID)
		mux.Post("/inventory/products/warranty/checkout", app.CheckoutWarrantyProduct)
		mux.Post("/inventory/products/warranty/delivery", app.DeliveryWarrantyProduct)
		mux.Post("/inventory/products/warranty/checkout/get-list", app.GetClaimWarrantyList)
		mux.Post("/inventory/products/warranty/get-history", app.GetClaimWarrantyList)

		//accounts
		//Receive & Collection
		mux.Post("/accounts/receive-collection/get-page-details", app.GetReceiveCollectionPageDetails)
		mux.Post("/accounts/receive-collection/complete-submission", app.CompleteReceiveCollectionProcess)

		//payment
		mux.Post("/accounts/payment/get-page-details", app.GetPaymentPageDetails)
		mux.Post("/accounts/payment/complete-submission", app.CompletePaymentProcess)

		//Amount Transfer
		mux.Post("/accounts/amount-transfer/get-page-details", app.GetAmountTransferPageDetails)
		mux.Post("/accounts/amount-transfer/complete-submission", app.CompleteAmountTransferProcess)
		//Amount Payable
		mux.Post("/accounts/amount-payable/get-page-details", app.GetAmountPayablePageDetails)
		mux.Post("/accounts/amount-payable/complete-submission", app.CompleteAmountPayableProcess)
		//Amount Receivable
		mux.Post("/accounts/amount-receivable/get-page-details", app.GetAmountReceivablePageDetails)
		mux.Post("/accounts/amount-receivable/complete-submission", app.CompleteAmountReceivableProcess)
		//Amount Receivable
		mux.Post("/accounts/expenses/get-page-details", app.GetExpensesPageDetails)
		mux.Post("/accounts/expenses/complete-submission", app.CompleteExpensesProcess)
		//Amount Receivable
		mux.Post("/accounts/adjustment/get-page-details", app.GetAdjustmentPageDetails)
		// mux.Post("/accounts/adjustment/complete-submission", app.CompleteAdjustmentProcess)

		mux.NotFound(app.PathNotFound)
		//.......................Inventory Reports.......................
		mux.Post("/reports/inventory/category-list", app.GetCategoryListReport)
		mux.Post("/reports/inventory/brand-list", app.GetBrandListReport)
		mux.Post("/reports/inventory/product-list", app.GetProductListReport)
		mux.Post("/reports/inventory/service-list", app.GetServiceListReport)
		mux.Post("/reports/inventory/stock-report", app.GetProductListReport)
		mux.Post("/reports/inventory/stock-alert-report", app.GetLowStockProductReport)
		mux.Post("/reports/inventory/purchase-history", app.GetPurchaseHistoryReport)
		mux.Post("/reports/inventory/sales-history", app.GetSalesHistoryReport)
		// Category List
		// Product List
		// Service List
		// Stock Report
		// Item History Report
		// Purchase History
		// Sales History
		// Service History

		//.......................Accounts Reports.......................
		mux.Post("/reports/accounts/transactions-report", app.GetTransactionsReport)
		mux.Post("/reports/accounts/cash-bank-statement", app.GetCashBankStatement)
		mux.Post("/reports/accounts/ledger-book-details", app.GetLedgerBookDetails)
		mux.Post("/reports/accounts/expenses-report", app.GetExpensesReport)
		mux.Post("/reports/accounts/income-statement", app.GetIncomeStatementReport)
		mux.Post("/reports/accounts/customer-due-report", app.GetCustomerDueReport)
		mux.Post("/reports/accounts/top-sheet-report", app.GetTopSheetReport)

	})
	return mux
}
