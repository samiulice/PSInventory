package repository

import (
	"PSInventory/internal/models"
)

type DatabaseRepo interface {
	//Head Accounts
	AddHeadAccount(ha models.HeadAccount) (int, error)
	GetAvailableHeadAccounts() ([]*models.HeadAccount, error)
	// HR Management
	AddEmployee(employee models.Employee) (int, error)
	GetEmployeeByID(id int) (models.Employee, error)
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
	ReturnProductUnitsToSupplier(JobID string, transactionDate string, ProductUnitsID []int, TotalUnits int, TotalPrices int) (int, error)
	GetProductList() ([]*models.Product, error)
	GetProductByID(id int) (models.Product, error)
	GetActiveProducts() ([]*models.Product, error)
	GetAvailableProductsByCategoryID(cat_id int) ([]*models.Product, error)
	GetAvailableProductsDetails() ([]*models.Product, error)

	GetInstockProductListByPurchaseIDAndProductID(purchaseID, productID int) (*models.Product, error)
	GetSoldProductListBySalesIDAndProductID(SalesID, productID int) (*models.Product, error)
	GetProductItemsListByProductID(productID int) (*models.Product, error)
	GetInStockItemDetailsBySerialNumber(serialNumber string) (*models.Product, error)
	GetSoldItemDetailsBySerialNumber(serialNumber string) (*models.Product, error)
	GetItemDetailsBySerialNumber(serialNumber string) (*models.Product, error)

	UpdateProductQuantityByProductID(quantity, productID int) error
	UpdateProductItemStatusByProductUnitsID(productUnitsID, status int) error

	AddProductSerialNumbers(purchase *models.Purchase) error

	AddToPurchaseHistory(purchase *models.Purchase) (int, error)

	//purchase
	GetPurchaseHistoryByMemoNo(memo_no string) ([]*models.Purchase, error)
	//sales
	GetSalesHistoryByMemoNo(memo_no string) ([]*models.Sale, error)
	SaleProducts(sale *models.Sale) error
	GetSalesHistoryByID(id int) (models.Sale, error)
	//Memo
	GetMemoListBySupplierID(supplierID int) ([]*models.Purchase, error)
	GetMemoListByCustomerID(customerID int) ([]*models.Sale, error)
	RestockProduct(purchase *models.Purchase) error

	//Helper functions
	CountRows(tableName string) (int, error)
}
