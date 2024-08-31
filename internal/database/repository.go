package repository

import (
	"PSInventory/internal/models"
)

type DatabaseRepo interface {
	//Head Accounts
	AddHeadAccount(ha models.HeadAccount) (int, error)
	GetAvailableHeadAccounts()([]*models.HeadAccount, error)
	// HR Management
	AddEmployee(employee models.Employee) (int, error)
	GetEmployeeByID(id int) (models.Employee, error)
	GetEmployeeListPaginated(accountType string, pageSize, currentPageIndex int) ([]*models.Employee, int, error)

	//MIS
	AddCustomer(customer models.Customer) (int, error)
	GetCustomerListPaginated(accountType string, pageSize, currentPageIndex int) ([]*models.Customer, int, error)
	AddSupplier(supplier models.Supplier) (int, error)
	GetSuppliersIDAndName()([]*models.Supplier, error)
	GetSupplierListPaginated(accountType string, pageSize, currentPageIndex int) ([]*models.Supplier, int, error)
	

	//Inventory
	AddBrand(b models.Brand) (int, error)
	GetBrandList()([]*models.Brand, error)
	GetAvailableBrands()([]*models.Brand, error)
	AddCategory(c models.Category) (int, error)
	GetCategoryList()([]*models.Category, error)
	GetAvailableCategories()([]*models.Category, error)

	AddItem(i models.Item)(int, error)
	GetItemList()([]*models.Item, error)
	GetAvailableItems()([]*models.Item, error)
	GetAvailableItemsByCategoryID(cat_id int)([]*models.Item, error)
	GetAvailableItemsDetails()([]*models.Item, error)
	//Helper functions
	CountRows(tableName string) (int, error)
}
