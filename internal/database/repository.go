package repository

import (
	"PSInventory/internal/models"
	"time"
)

type DatabaseRepo interface {
	//Head Accounts
	AddHeadAccount(ha models.HeadAccount) (int, error)
	GetAvailableHeadAccounts() ([]*models.HeadAccount, error)
	// HR Management
	AddEmployee(employee models.Employee) (int, error)
	GetEmployeeByID(id int) (models.Employee, error)
	GetEmployeeListPaginated(accountType string, pageSize, currentPageIndex int) ([]*models.Employee, int, error)

	//MIS
	AddCustomer(customer models.Customer) (int, error)
	GetCustomerListPaginated(accountType string, pageSize, currentPageIndex int) ([]*models.Customer, int, error)
	AddSupplier(supplier models.Supplier) (int, error)
	GetSuppliersIDAndName() ([]*models.Supplier, error)
	GetSupplierListPaginated(accountType string, pageSize, currentPageIndex int) ([]*models.Supplier, int, error)

	//Inventory
	AddBrand(b models.Brand) (int, error)
	GetBrandList() ([]*models.Brand, error)
	GetActiveBrands() ([]*models.Brand, error)
	AddCategory(c models.Category) (int, error)
	GetCategoryList() ([]*models.Category, error)
	GetActiveCategoryList() ([]*models.Category, error)
	AddProduct(i models.Product) (int, error)
	ReturnProductUnitsToSupplier(JobID string, transactionDate time.Time, ProductUnitsID []int, TotalUnits int, TotalPrices int) (int, error)
	GetProductList() ([]*models.Product, error)
	GetProductByID(id int) (models.Product, error)
	GetActiveProducts() ([]*models.Product, error)
	GetAvailableProductsByCategoryID(cat_id int) ([]*models.Product, error)
	GetAvailableProductsDetails() ([]*models.Product, error)
	GetPurchaseHistoryByMemoNo(memo_no string) ([]*models.Purchase, error)
	GetProductListByPurchaseIDAndProductID(purchaseID, productID int) (*models.Product, error)

	UpdateProductQuantityByProductID(quantity, productID int) error
	UpdateProductItemStatusByProductUnitsID(productUnitsID, status int) error

	AddProductSerialNumbers(purchase *models.Purchase) error

	AddToPurchaseHistory(purchase *models.Purchase) (int, error)
	GetMemoListBySupplierID(supplierID int) ([]*models.Purchase, error)
	RestockProduct(purchase *models.Purchase) error

	//Helper functions
	CountRows(tableName string) (int, error)
}
