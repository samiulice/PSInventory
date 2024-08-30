package repository

import (
	"PSInventory/internal/models"
)

type DatabaseRepo interface {
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
	AddCategory(c models.Category) (int, error)
	GetCategoryList()([]*models.Category, error)
	//Helper functions
	CountRows(tableName string) (int, error)
}
