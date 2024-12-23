package repository

import (
	"PSInventory/internal/models"
)

type DatabaseRepo interface {
	GetDashBoardData()(interface{} , error)
	//Head Accounts
	AddHeadAccount(ha models.HeadAccount) (int, error)
	GetAvailableHeadAccounts() ([]*models.HeadAccount, error)
	GetAvailableHeadAccountsByType(accountType string) ([]*models.HeadAccount, error)
	UpdateHeadAccountBalance(int, int) error

	//Administrative panel
	AddNewStakeHolder(models.StakeHolder)(int, error)
	// HR Management
	AddEmployee(employee models.Employee) (int, error)
	GetEmployeeListPaginated(accountType string, pageSize, currentPageIndex int) ([]*models.Employee, int, error)

	//MIS--customer
	AddCustomer(customer models.Customer) (int, error)
	GetCustomerByID(id int) (models.Customer, error)
	GetActiveCustomersIDAndName() ([]*models.Customer, error)
	GetCreditCustomersDetails() ([]*models.Customer, error)
	GetDebitCustomersDetails() ([]*models.Customer, error)
	GetCustomerListPaginated(accountType string, pageSize, currentPageIndex int) ([]*models.Customer, int, error)
	//MIS--supplier
	AddSupplier(supplier models.Supplier) (int, error)
	GetActiveSuppliersIDAndName() ([]*models.Supplier, error)
	GetCreditSuppliersDetails() ([]*models.Supplier, error)
	GetDebitSuppliersDetails() ([]*models.Supplier, error)
	GetSupplierListPaginated(accountType string, pageSize, currentPageIndex int) ([]*models.Supplier, int, error)

	//Inventory

	//product
	AddBrand(b models.Brand) (int, error)
	GetBrandList() ([]*models.Brand, error)
	GetActiveBrands() ([]*models.Brand, error)
	AddCategory(c models.Category) (int, error)
	GetCategoryList() ([]*models.Category, error)
	GetActiveCategoryList() ([]*models.Category, error)
	AddProduct(i models.Product) (int, error)
	ReturnProductUnitsToSupplier(PurchaseHistory models.Purchase, JobID string, transactionDate string, ProductUnitsID []int, TotalUnits int, TotalPrices int) (int, error)
	GetProductList() ([]*models.Product, error)
	GetProductByID(id int) (models.Product, error)
	GetActiveProducts() ([]*models.Product, error)
	GetAvailableProductsByCategoryID(cat_id int) ([]*models.Product, error)
	GetAllProductsByCategoryID(cat_id int) ([]*models.Product, error)
	GetAvailableProductsDetails() ([]*models.Product, error)
	GetExpenseList()([]*models.ExpenseType, error)
	GeActiveStakeHolderList() ([]*models.StakeHolder, error)

	GetInStockProductListByPurchaseIDAndProductID(purchaseID, productID int) (*models.Product, error)
	GetSoldProductListBySalesIDAndProductID(SalesID, productID int) (*models.Product, error)
	GetInStockProductItemsListByProductID(productID int) (*models.Product, error)
	GetInStockItemDetailsBySerialNumber(serialNumber string) (*models.Product, error)
	GetSoldItemDetailsBySerialNumber(serialNumber string) (*models.Product, error)
	GetItemDetailsBySerialNumber(serialNumber string) (*models.Product, error)
	GetProductItemsDetailsBySalesHistoryID(id int) ([]*models.Product, error)
	UpdateProductItemStatusByProductUnitsID(productUnitsID, status int) error

	AddProductSerialNumbers(purchase *models.Purchase) error

	AddToPurchaseHistory(purchase *models.Purchase) (int, error)

	//purchase
	GetPurchaseHistoryByMemoNo(memo_no string) ([]*models.Purchase, error)
	//sales
	GetSalesHistoryByMemoNo(memo_no string) ([]*models.Sale, error)
	SaleProductsToCustomer(sale *models.SalesInvoice) error
	GetSalesHistoryByID(id int) (models.Sale, error)
	//sale return
	SaleReturnDB(SalesHistory *models.Sale, SelectedItemsID []int, SaleReturnDate string, ReturnItemsCount int, ReturnAmount int, MemoNo string) error
	//warranty
	AddNewWarrantyClaim(memoPrefix string, serialID int, serialNumber, contactNumber, reportedProblem, receivedBy, warrantyHistoryIds string) (int, error)
	GetWarrantyList(SearchType string) ([]*models.Warranty, error)
	// GetWarrantyDetailsByID(id int) ([]*models.Warranty, error)
	CheckoutWarrantyProduct(warrantyHistoryID, productSerialID int, arrivalDate, newSerialNumber, comment string) error
	DeliveryWarrantyProduct(warrantyHistoryID, productSerialID int, deliveredBy string) error
	//Memo
	GetMemoListBySupplierID(supplierID int) ([]*models.Purchase, error)
	GetMemoListByCustomerID(customerID int) ([]*models.Sale, error)
	RestockProduct(purchase *models.PurchasePayload) error

	//accounts
	//Receive & Collection
	CompleteReceiveCollectionTransactions(summary []*models.Reception) error
	//Payment
	CompletePaymentTransactions(summary []*models.Payment) error
	//Amount Transfer
	CompleteAmountTransferTransactions(summary []*models.AmountTransfer) error
	//Amount Payable
	CompleteAmountPayableTransactions(summary []*models.AmountPayable) error
	//Amount Receivable
	CompleteAmountReceivableTransactions(summary []*models.AmountReceivable) error
	//Expenses
	CompleteExpensesTransactions(summary []*models.Expense) error
	//Cash-Adjustment
	CompleteFundAcquisitionProcess(summary []*models.FundAcquisition) error
	//Inventory Reports
	GetAllEmployeesList() ([]*models.Employee, error)
	GetAllSuppliersList() ([]*models.Supplier, error)
	GetAllCustomersList() ([]*models.Customer, error)
	GetCategoryListReport() ([]*models.Category, error)
	GetBrandListReport() ([]*models.Brand, error)
	GetProductListReport() ([]*models.Product, error)
	GetLowStockProductReport() ([]*models.Product, error)
	GetServiceListReport() ([]*models.Service, error)
	GetPurchaseHistoryReport() ([]*models.Purchase, error)
	GetSalesHistoryReport() ([]*models.Sale, error)

	//Accounts report
	GetCustomerDueHistoryReport() ([]*models.Customer, error)
	GetSupplierDueHistoryReport() ([]*models.Supplier, error)
	GetTransactionsHistoryReport() ([]*models.Transaction, error)
	GetCashBankStatement() ([]*models.Transaction, error)
	GetLedgerBookDetails(account_type string, account_id int) ([]*models.Transaction, error)
	GetExpensesHistoryReport() ([]*models.Transaction, error)
	GetIncomeStatementData(startDate,endDate string) (models.IncomeStatement, error)
	GetTopSheetReport() ([]*models.TopSheet, error)
	GetBalanceSheetReport() (models.BalanceSheet, error)
	GetTrialBalanceReport() (models.TrialBalance, error)
	//Helper functions
	LastIndex(tableName string) (int64, error)
	GetCompanyProfile() (models.CompanyProfile, error)

	ImportDatabase()error
	ExportDatabase()error
}
