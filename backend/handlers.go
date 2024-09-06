package backend

import (
	"database/sql"
	"fmt"
	"net/http"

	"PSInventory/internal/models"
	"path"
	"strconv"
)

type JSONResponse struct {
	Error   bool        `json:"error,omitempty"`
	Message string      `json:"message,omitempty"`
	Result  interface{} `json:"result,omitempty"`
}

// .....................HR Management Panel Handlers......................
// GetEmployeeList return list of employees to the corresponded category in JSON format
func (app *application) GetEmployees(w http.ResponseWriter, r *http.Request) {
	accountType := path.Base(r.URL.Path)

	id, err := strconv.Atoi(accountType)
	if err == nil {
		employee, err := app.DB.GetEmployeeByID(id)
		if err != nil {
			app.errorLog.Println(err)
			app.badRequest(w, err)
			return
		}
		app.writeJSON(w, http.StatusOK, employee)
	} else {
		var payload struct {
			PageSize         int `json:"page_size,omitempty"`
			CurrentPageIndex int `json:"current_page_index,omitempty"`
		}
		err := app.readJSON(w, r, &payload)
		if err != nil {
			app.badRequest(w, err)
			return
		}
		employees, totalRecords, err := app.DB.GetEmployeeListPaginated(accountType, payload.PageSize, payload.CurrentPageIndex)

		if err != nil {
			app.badRequest(w, err)
			return
		}
		var Resp struct {
			Error            bool               `json:"error,omitempty"`
			Message          string             `json:"message,omitempty"`
			PageSize         int                `json:"page_size,omitempty"`
			CurrentPageIndex int                `json:"current_page_index,omitempty"`
			TotalRecords     int                `json:"total_records,omitempty"`
			Employees        []*models.Employee `json:"employees,omitempty"`
		}
		Resp.Error = false
		Resp.Message = "Data successfully fetched"
		Resp.PageSize = payload.PageSize
		Resp.CurrentPageIndex = payload.CurrentPageIndex
		Resp.TotalRecords = totalRecords
		Resp.Employees = employees
		app.writeJSON(w, http.StatusOK, Resp)
	}
}

// AddEmployee add new employee info to the database
func (app *application) AddEmployee(w http.ResponseWriter, r *http.Request) {
	var employee models.Employee
	var resp JSONResponse
	err := app.readJSON(w, r, &employee)
	if err != nil {
		app.badRequest(w, err)
		return
	}
	//Sanitize blank info
	if employee.AccountCode == "" {
		n, err := app.DB.CountRows("employees")
		if err == nil {
			employee.AccountCode = "em-" + fmt.Sprintf("%06d", n+1)
		}
	}
	employee.AccountStatus = true
	_, err = app.DB.AddEmployee(employee)

	if err != nil {
		resp = JSONResponse{
			Error: true,
			// Message: "Internal Server Error! Please try again or contact to developer",
			Message: err.Error(),
		}
		app.writeJSON(w, http.StatusInternalServerError, resp)
		return
	}
	resp = JSONResponse{
		Error:   false,
		Message: "Employee added successfully",
		Result:  employee,
	}
	app.writeJSON(w, http.StatusOK, resp)
}

// .....................MIS Handlers......................
// AddEmployee add new employee info to the database
func (app *application) AddCustomer(w http.ResponseWriter, r *http.Request) {
	var customer models.Customer
	var resp JSONResponse
	err := app.readJSON(w, r, &customer)
	if err != nil {
		app.badRequest(w, err)
		return
	}
	//Sanitize blank info
	if customer.AccountCode == "" {
		n, err := app.DB.CountRows("customers")
		if err == nil {
			customer.AccountCode = "em-" + fmt.Sprintf("%06d", n+1)
		}
	}
	customer.AccountStatus = true
	_, err = app.DB.AddCustomer(customer)

	if err != nil {
		resp = JSONResponse{
			Error: true,
			// Message: "Internal Server Error! Please try again or contact to developer",
			Message: err.Error(),
		}
		app.writeJSON(w, http.StatusInternalServerError, resp)
		return
	}
	resp = JSONResponse{
		Error:   false,
		Message: "Customer added successfully",
		Result:  customer,
	}
	app.writeJSON(w, http.StatusOK, resp)
}

// GetCustomers returns list of customers to the corresponded category in JSON format
func (app *application) GetCustomers(w http.ResponseWriter, r *http.Request) {
	accountType := path.Base(r.URL.Path)

	_, err := strconv.Atoi(accountType)
	if err == nil {
		// employee, err := app.DB.GetEmployeeByID(id)
		// if err != nil {
		// 	app.errorLog.Println(err)
		// 	app.badRequest(w, err)
		// 	return
		// }
		// app.writeJSON(w, http.StatusOK, employee)
		// TODO:
	} else {
		var payload struct {
			PageSize         int `json:"page_size,omitempty"`
			CurrentPageIndex int `json:"current_page_index,omitempty"`
		}
		err := app.readJSON(w, r, &payload)
		if err != nil {
			app.badRequest(w, err)
			return
		}
		customers, totalRecords, err := app.DB.GetCustomerListPaginated(accountType, payload.PageSize, payload.CurrentPageIndex)

		if err != nil {
			app.badRequest(w, err)
			return
		}
		var Resp struct {
			Error            bool               `json:"error,omitempty"`
			Message          string             `json:"message,omitempty"`
			PageSize         int                `json:"page_size,omitempty"`
			CurrentPageIndex int                `json:"current_page_index,omitempty"`
			TotalRecords     int                `json:"total_records,omitempty"`
			Customers        []*models.Customer `json:"customers,omitempty"`
		}
		Resp.Error = false
		Resp.Message = "Data successfully fetched"
		Resp.PageSize = payload.PageSize
		Resp.CurrentPageIndex = payload.CurrentPageIndex
		Resp.TotalRecords = totalRecords
		Resp.Customers = customers
		app.writeJSON(w, http.StatusOK, Resp)
	}
}

// AddSupplier add new supplier info to the database
func (app *application) AddSupplier(w http.ResponseWriter, r *http.Request) {
	var supplier models.Supplier
	var resp JSONResponse
	err := app.readJSON(w, r, &supplier)
	if err != nil {
		app.badRequest(w, err)
		return
	}
	//Sanitize blank info
	if supplier.AccountCode == "" {
		n, err := app.DB.CountRows("suppliers")
		if err == nil {
			supplier.AccountCode = "sup-" + fmt.Sprintf("%06d", n+1)
		}
	}
	supplier.AccountStatus = true
	_, err = app.DB.AddSupplier(supplier)

	if err != nil {
		resp = JSONResponse{
			Error: true,
			// Message: "Internal Server Error! Please try again or contact to developer",
			Message: err.Error(),
		}
		app.writeJSON(w, http.StatusInternalServerError, resp)
		return
	}
	resp = JSONResponse{
		Error:   false,
		Message: "Supplier added successfully",
		Result:  supplier,
	}
	app.writeJSON(w, http.StatusOK, resp)
}

// GetSuppliers returns list of suppliers to the corresponded category in JSON format
func (app *application) GetSuppliers(w http.ResponseWriter, r *http.Request) {
	accountType := path.Base(r.URL.Path)

	_, err := strconv.Atoi(accountType)
	if err == nil {
		// employee, err := app.DB.GetEmployeeByID(id)
		// if err != nil {
		// 	app.errorLog.Println(err)
		// 	app.badRequest(w, err)
		// 	return
		// }
		// app.writeJSON(w, http.StatusOK, employee)
		// TODO
	} else {
		var payload struct {
			PageSize         int `json:"page_size,omitempty"`
			CurrentPageIndex int `json:"current_page_index,omitempty"`
		}
		err := app.readJSON(w, r, &payload)
		if err != nil {
			app.badRequest(w, err)
			return
		}
		suppliers, totalRecords, err := app.DB.GetSupplierListPaginated(accountType, payload.PageSize, payload.CurrentPageIndex)

		if err != nil {
			app.badRequest(w, err)
			return
		}
		var Resp struct {
			Error            bool               `json:"error,omitempty"`
			Message          string             `json:"message,omitempty"`
			PageSize         int                `json:"page_size,omitempty"`
			CurrentPageIndex int                `json:"current_page_index,omitempty"`
			TotalRecords     int                `json:"total_records,omitempty"`
			Suppliers        []*models.Supplier `json:"suppliers,omitempty"`
		}
		Resp.Error = false
		Resp.Message = "Data successfully fetched"
		Resp.PageSize = payload.PageSize
		Resp.CurrentPageIndex = payload.CurrentPageIndex
		Resp.TotalRecords = totalRecords
		Resp.Suppliers = suppliers
		app.writeJSON(w, http.StatusOK, Resp)
	}
}

// .....................Inventory Handlers......................

// AddBrand Handles category adding process
func (app *application) AddBrand(w http.ResponseWriter, r *http.Request) {

	var brand models.Brand
	err := app.readJSON(w, r, &brand)
	if err != nil {
		app.badRequest(w, err)
		return
	}

	id, err := app.DB.AddBrand(brand)
	if err != nil {
		app.badRequest(w, err)
		return
	}
	brand.ID = id
	var resp = JSONResponse{
		Error:   false,
		Message: "Brand Added Succesfully",
		Result:  brand,
	}
	app.writeJSON(w, http.StatusOK, resp)
}

// AddCategory Handles category adding process
func (app *application) AddCategory(w http.ResponseWriter, r *http.Request) {

	var category models.Category
	err := app.readJSON(w, r, &category)
	if err != nil {
		app.badRequest(w, err)
		return
	}

	id, err := app.DB.AddCategory(category)
	if err != nil {
		app.badRequest(w, err)
		return
	}
	category.ID = id
	var resp = JSONResponse{
		Error:   false,
		Message: "Category Added Succesfully",
		Result:  category,
	}
	app.writeJSON(w, http.StatusOK, resp)
}

// AddProduct insert new product type to database
func (app *application) AddProduct(w http.ResponseWriter, r *http.Request) {

	var product models.Product
	err := app.readJSON(w, r, &product)
	if err != nil {
		app.badRequest(w, err)
		return
	}
	//sanitize blank variables
	code, err := app.DB.CountRows("products")
	if err != nil {
		app.badRequest(w, err)
		return
	}
	product.ProductCode = fmt.Sprintf("i-%06d", code)
	product.ProductStatus = true

	id, err := app.DB.AddProduct(product)
	if err != nil {
		app.badRequest(w, err)
		return
	}
	product.ID = id
	var resp = JSONResponse{
		Error:   false,
		Message: "Product Added Succesfully",
		Result:  product,
	}
	app.writeJSON(w, http.StatusOK, resp)
}

// RestockProduct add new purchased product to database
func (app *application) RestockProduct(w http.ResponseWriter, r *http.Request) {
	var purchase_details models.Purchase
	err := app.readJSON(w, r, &purchase_details)
	if err != nil {
		app.badRequest(w, err)
		return
	}
	err = app.DB.RestockProduct(&purchase_details)
	if err != nil {
		app.badRequest(w, err)
		return
	}
	var resp = JSONResponse{
		Error:   false,
		Message: "restock product Succesfully",
		Result:  purchase_details,
	}
	app.writeJSON(w, http.StatusOK, resp)
}

// GetPageDetails scrape data from the database for init purchase page
func (app *application) GetPageDetails(w http.ResponseWriter, r *http.Request) {
	//supplier
	//category
	//product
	//account
	var resp struct {
		Error      bool               `json:"error,omitempty"`
		Message    string             `json:"message,omitempty"`
		Suppliers  []*models.Supplier `json:"suppliers,omitempty"`
		Categories []*models.Category `json:"categories,omitempty"`
		Brands     []*models.Brand    `json:"brands,omitempty"`
		Products   []*models.Product  `json:"products,omitempty"`

		HeadAccounts []*models.HeadAccount `json:"head_accounts,omitempty"`
	}

	//retrive suppliers from the database
	suppliers, err := app.DB.GetSuppliersIDAndName()
	if err == sql.ErrNoRows {
		resp.Message += "||No Supplier Available||"
	} else if err != nil {
		app.badRequest(w, err) //send error response
		return
	}
	//retrive categories from the database
	categories, err := app.DB.GetActiveCategoryList()
	if err == sql.ErrNoRows {
		resp.Message += "||No Category Available||"
	} else if err != nil {
		app.badRequest(w, err) //send error response
		return
	}
	//retrive brands from the database
	brands, err := app.DB.GetActiveBrands()
	if err == sql.ErrNoRows {
		resp.Message += "||No Brand Available||"
	} else if err != nil {
		app.badRequest(w, err) //send error response
		return
	}
	//retrive products from the database
	products, err := app.DB.GetActiveProducts()
	if err == sql.ErrNoRows {
		resp.Message += "||No Product Available||"
	} else if err != nil {
		app.badRequest(w, err) //send error response
		return
	}
	//retrive accounts from the database
	headAccounts, err := app.DB.GetAvailableHeadAccounts()
	if err == sql.ErrNoRows {
		resp.Message += "||No Accounts Available||"
	} else if err != nil {
		app.badRequest(w, err) //send error response
		return
	}

	//TODO: Retrive accounts and send to frontend
	resp.Error = false
	resp.Message += "||data Succesfully fetched||"
	resp.Suppliers = suppliers
	resp.Categories = categories
	resp.Brands = brands
	resp.Products = products
	resp.HeadAccounts = headAccounts

	app.writeJSON(w, http.StatusOK, resp)
}
