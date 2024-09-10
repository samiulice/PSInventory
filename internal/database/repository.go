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
	GetProductList() ([]*models.Product, error)
	GetActiveProducts() ([]*models.Product, error)
	GetAvailableProductsByCategoryID(cat_id int) ([]*models.Product, error)
	GetAvailableProductsDetails() ([]*models.Product, error)
	UpdateProductQuantity(quantity, productID int) error

	AddProductSerialNumbers(purchase *models.Purchase) error

	AddToPurchaseHistory(purchase *models.Purchase) (int, error)
	GetMemoListWithPurchaseID(supplierID int)([]*models.Purchase, error)
	RestockProduct(purchase *models.Purchase)(error)
	//Helper functions
	CountRows(tableName string) (int, error)
}
