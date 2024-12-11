package backend

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"
	"time"

	"PSInventory/internal/models"
	"path"
	"strconv"
)

type JSONResponse struct {
	Error   bool        `json:"error,omitempty"`
	Message string      `json:"message,omitempty"`
	Result  interface{} `json:"result,omitempty"`
}

func (app *application) PathNotFound(w http.ResponseWriter, r *http.Request) {
	// Serve the custom 404 HTML page
	http.ServeFile(w, r, "frontend/src/404.html")
}
func (app *application) FetchDashBoardData(w http.ResponseWriter, r *http.Request) {
	data, err := app.DB.GetDashBoardData()
	if err != nil {
		app.badRequest(w, fmt.Errorf("ERROR:FetchDashBoardData: Unable to get data %w", err))
		return
	}

	var resp struct {
		Error   bool        `json:"error,omitempty"`
		Message string      `json:"message,omitempty"`
		Result  interface{} `json:"result,omitempty"`
	}

	resp.Error = false
	resp.Message = "Data fetched successfully"
	resp.Result = data

	app.writeJSON(w, http.StatusOK, resp)
}
func (app *application) FetchCompanyProfile(w http.ResponseWriter, r *http.Request) {

	cp, err := app.DB.GetCompanyProfile()
	if err != nil {
		app.badRequest(w, fmt.Errorf("ERROR: FetchCompanyProfile => Unable to retrieve Company Details:  %w", err))
		return
	}
	var resp struct {
		Error          bool                  `json:"error"`
		Message        string                `json:"message"`
		CompanyProfile models.CompanyProfile `json:"company_profile"`
	}
	resp.Message = "Company data fetched successfully"
	resp.CompanyProfile = cp

	app.writeJSON(w, http.StatusOK, resp)
}

// .....................Administrative Panel Handlers......................
func (app *application) AddNewStakeHolder(w http.ResponseWriter, r *http.Request) {
	var stk models.StakeHolder

	err := app.readJSON(w, r, &stk)

	if err != nil {
		app.badRequest(w, fmt.Errorf("ERROR: AddNewStakeHolder: Unable to read JSON => %w", err))
		return
	}

	//insert to company_stakeholders
	n, err := app.DB.LastIndex("company_stakeholders")
	if err != nil {
		app.badRequest(w, err)
		return
	}
	pre, err := app.GenerateRandomAlphanumericCode(4)
	if err != nil {
		app.badRequest(w, err)
		return
	}
	if stk.AccountType == "Owner" {
		stk.AccountCode = "SH-" + pre + strconv.FormatInt(n, 10)
	} else if stk.AccountType == "Investor" {
		stk.AccountCode = "IN-" + pre + strconv.FormatInt(n, 10)
	}
	stk.AccountStatus = true

	_, err = app.DB.AddNewStakeHolder(stk)
	if err != nil {
		app.badRequest(w, fmt.Errorf("ERROR: AddNewStakeHolder => %w", err))
		return
	}
	var resp struct {
		Error              bool               `json:"error"`
		Message            string             `json:"message"`
		CompanyStakeholder models.StakeHolder `json:"company_stakeholder"`
	}
	resp.Message = fmt.Sprintf("%s-%s Added Successfully", stk.AccountType, stk.AccountName)
	resp.CompanyStakeholder = stk

	app.writeJSON(w, http.StatusOK, resp)
}

// .....................HR Management Panel Handlers......................

// GetAllEmployees returns list of all employees info from the database
func (app *application) GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	var Resp struct {
		Error     bool               `json:"error,omitempty"`
		Message   string             `json:"message,omitempty"`
		Employees []*models.Employee `json:"employees,omitempty"`
	}
	employees, err := app.DB.GetAllEmployeesList()
	if err == sql.ErrNoRows {
		Resp.Error = false
		Resp.Message = "No Data available"
		app.writeJSON(w, http.StatusOK, Resp)
		return
	}
	if err != nil {
		app.badRequest(w, err)
		return
	}
	Resp.Error = false
	Resp.Message = "Data successfully fetched"
	Resp.Employees = employees
	app.writeJSON(w, http.StatusOK, Resp)
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
	n, err := app.DB.LastIndex("employees")
	if err != nil {
		app.badRequest(w, err)
		return
	}
	pre, err := app.GenerateRandomAlphanumericCode(4)
	if err != nil {
		app.badRequest(w, err)
		return
	}
	employee.AccountCode = "E-" + pre + strconv.FormatInt(n, 10)
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
	n, err := app.DB.LastIndex("customers")
	if err != nil {
		app.badRequest(w, err)
		return
	}
	pre, err := app.GenerateRandomAlphanumericCode(4)
	if err != nil {
		app.badRequest(w, err)
		return
	}
	customer.AccountCode = "C-" + pre + strconv.FormatInt(n, 10)
	customer.AccountStatus = true
	id, err := app.DB.AddCustomer(customer)

	if err != nil {
		resp = JSONResponse{
			Error: true,
			// Message: "Internal Server Error! Please try again or contact to developer",
			Message: err.Error(),
		}
		app.writeJSON(w, http.StatusInternalServerError, resp)
		return
	}
	customer.ID = id
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

// GetAllCustomers returns list of all customers info from the database
func (app *application) GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	var Resp struct {
		Error     bool               `json:"error,omitempty"`
		Message   string             `json:"message,omitempty"`
		Customers []*models.Customer `json:"customers,omitempty"`
	}
	customers, err := app.DB.GetAllCustomersList()
	if err == sql.ErrNoRows {
		Resp.Error = false
		Resp.Message = "No Data available"
		app.writeJSON(w, http.StatusOK, Resp)
		return
	}
	if err != nil {
		app.badRequest(w, err)
		return
	}
	Resp.Error = false
	Resp.Message = "Data successfully fetched"
	Resp.Customers = customers
	app.writeJSON(w, http.StatusOK, Resp)
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
	n, err := app.DB.LastIndex("suppliers")
	if err != nil {
		app.badRequest(w, err)
		return
	}
	pre, err := app.GenerateRandomAlphanumericCode(4)
	if err != nil {
		app.badRequest(w, err)
		return
	}
	supplier.AccountCode = "S-" + pre + strconv.FormatInt(n, 10)
	supplier.AccountStatus = true
	id, err := app.DB.AddSupplier(supplier)

	if err != nil {
		resp = JSONResponse{
			Error: true,
			// Message: "Internal Server Error! Please try again or contact to developer",
			Message: err.Error(),
		}
		app.writeJSON(w, http.StatusInternalServerError, resp)
		return
	}
	supplier.ID = id
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

// GetSuppliers returns list of suppliers to the corresponded category in JSON format
func (app *application) GetAllSuppliers(w http.ResponseWriter, r *http.Request) {

	var Resp struct {
		Error     bool               `json:"error,omitempty"`
		Message   string             `json:"message,omitempty"`
		Suppliers []*models.Supplier `json:"suppliers,omitempty"`
	}
	suppliers, err := app.DB.GetAllSuppliersList()
	if err == sql.ErrNoRows {
		Resp.Error = false
		Resp.Message = "No Data available"
		app.writeJSON(w, http.StatusOK, Resp)
		return
	}
	if err != nil {
		app.badRequest(w, err)
		return
	}
	Resp.Error = false
	Resp.Message = "Data successfully fetched"
	Resp.Suppliers = suppliers
	app.writeJSON(w, http.StatusOK, Resp)
}

// GetActiveSuppliersIDAndName returns a list of supplier's id and name
func (app *application) GetActiveSuppliersIDAndName(w http.ResponseWriter, r *http.Request) {
	//supplier
	var resp struct {
		Error     bool               `json:"error,omitempty"`
		Message   string             `json:"message,omitempty"`
		Suppliers []*models.Supplier `json:"suppliers,omitempty"`
	}

	//retrieve suppliers from the database
	suppliers, err := app.DB.GetActiveSuppliersIDAndName()
	if err == sql.ErrNoRows {
		resp.Message += "||No Supplier Available||"
	} else if err != nil {
		app.badRequest(w, err) //send error response
		return
	}

	//TODO: Retrive accounts and send to frontend
	resp.Error = false
	resp.Message += "data Succesfully fetched"
	resp.Suppliers = suppliers
	app.writeJSON(w, http.StatusOK, resp)
}

// GetActiveCustomersIDAndName retruns a list of customers's id and name
func (app *application) GetActiveCustomersIDAndName(w http.ResponseWriter, r *http.Request) {
	//supplier
	var resp struct {
		Error     bool               `json:"error,omitempty"`
		Message   string             `json:"message,omitempty"`
		Customers []*models.Customer `json:"customers,omitempty"`
	}

	//retrieve suppliers from the database
	customers, err := app.DB.GetActiveCustomersIDAndName()
	if err == sql.ErrNoRows {
		resp.Message += "||No Customer Available||"
	} else if err != nil {
		app.badRequest(w, err) //send error response
		return
	}

	//TODO: Retrive accounts and send to frontend
	resp.Error = false
	resp.Message += "data Succesfully fetched"
	resp.Customers = customers
	app.writeJSON(w, http.StatusOK, resp)
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
	code, err := app.DB.LastIndex("products")
	if err != nil {
		app.badRequest(w, err)
		return
	}
	product.ProductCode = fmt.Sprintf("i-%06d", code+1)
	product.ProductStatus = true

	id, err := app.DB.AddProduct(product)
	if err != nil {
		app.badRequest(w, err)
		return
	}
	product, err = app.DB.GetProductByID(id)
	if err != nil {
		app.badRequest(w, err)
		return
	}
	var resp = JSONResponse{
		Error:   false,
		Message: "Product Added Successfully",
		Result:  product,
	}
	app.writeJSON(w, http.StatusOK, resp)
}

// FetchPurchaseMemoProductItems retrieve purchased products list of a memo
func (app *application) FetchPurchaseMemoProductItems(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		MemoNo string `json:"memo_no"`
	}
	err := app.readJSON(w, r, &payload) //read json body
	if err != nil {
		app.badRequest(w, err)
		return
	}
	//get purchase history associated with memo_no
	purchaseHistory, err := app.DB.GetPurchaseHistoryByMemoNo(payload.MemoNo)
	if err != nil {
		app.badRequest(w, fmt.Errorf("ErrorPurchaseHistory:: %v", err))
	}
	//Get product ids for this memo with associated purchase_id for the given memo
	//get detailed Product info for these ids
	//retrieve all product-serial of each product_id && purchase_id
	var products []*models.Product
	for _, v := range purchaseHistory {
		product, err := app.DB.GetInStockProductListByPurchaseIDAndProductID(v.ID, v.Product.ID)
		if err != nil {
			app.badRequest(w, fmt.Errorf("ErrorProducts:: %v", err))
		}
		products = append(products, product)
	}

	if err != nil {
		app.badRequest(w, err)
		return
	}

	var resp struct {
		Error           bool               `json:"error,omitempty"`
		Message         string             `json:"message,omitempty"`
		Product         []*models.Product  `json:"product,omitempty"`
		PurchaseHistory []*models.Purchase `json:"purchase_history,omitempty"`
	}
	resp.Error = false
	resp.Message = "Data fetched successfully"
	resp.Product = products
	resp.PurchaseHistory = purchaseHistory
	app.writeJSON(w, http.StatusOK, resp)
}

// FetchProductItemsBySalesHistory retrieve purchased products list of a memo
func (app *application) FetchProductItemsBySalesHistory(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		ID int `json:"sales_history_id"`
	}
	err := app.readJSON(w, r, &payload) //read json body
	if err != nil {
		app.badRequest(w, err)
		return
	}
	//get purchase history associated with memo_no
	salesHistory, err := app.DB.GetSalesHistoryByID(payload.ID)
	if err != nil {
		app.badRequest(w, fmt.Errorf("ErrorSalesHistory:: %v", err))
		return
	}

	//get product details by sales id
	products, err := app.DB.GetProductItemsDetailsBySalesHistoryID(salesHistory.ID)

	if err != nil {
		app.badRequest(w, fmt.Errorf("DBERROR:=>GetProductItemsDetailsBySalesHistoryID: %w", err))
		return
	}
	var resp struct {
		Error        bool              `json:"error"`
		Message      string            `json:"message"`
		Products     []*models.Product `json:"products"`
		SalesHistory models.Sale       `json:"sales_history"`
	}
	resp.Error = false
	resp.Message = "Data fetched successfully"
	resp.Products = products
	resp.SalesHistory = salesHistory
	app.writeJSON(w, http.StatusOK, resp)
}

// FetchProductItemsByProductID retrieve in stock product items
func (app *application) FetchProductItemsByProductID(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		ProductID int `json:"product_id"`
	}
	err := app.readJSON(w, r, &payload) //read json body
	if err != nil {
		app.badRequest(w, err)
		return
	}
	//get purchase history associated with memo_no
	productItems, err := app.DB.GetInStockProductItemsListByProductID(payload.ProductID)
	if err != nil {
		app.badRequest(w, fmt.Errorf("ErrorPurchaseHistory:: %v", err))
		return
	}

	var resp struct {
		Error        bool            `json:"error,omitempty"`
		Message      string          `json:"message,omitempty"`
		ProductItems *models.Product `json:"product_items,omitempty"`
	}
	resp.Error = false
	resp.Message = "Data fetched successfully"
	resp.ProductItems = productItems
	app.writeJSON(w, http.StatusOK, resp)
}

// FetchProductItemBySerialNumber retrieve product item by serial number
func (app *application) FetchInstockProductItemBySerialNumber(w http.ResponseWriter, r *http.Request) {

	//Define struct for JSON object
	var payload struct {
		SerialNumber string `json:"product_serial_number"`
	}
	var resp struct {
		Error        bool            `json:"error,omitempty"`
		Message      string          `json:"message,omitempty"`
		ProductItems *models.Product `json:"product_items,omitempty"`
	}
	//Read JSON body
	err := app.readJSON(w, r, &payload)
	if err != nil {
		resp.Error = true
		resp.Message = "cannot read JSON: " + err.Error()
		app.writeJSON(w, http.StatusOK, resp)
		return
	}

	//Get product item details for the given serial number(e.g. id,serial_number, product_id,purchase_history_id,warranty, max_retail_price, purchase_rate)
	productItem, err := app.DB.GetInStockItemDetailsBySerialNumber(payload.SerialNumber)
	if err != nil {
		resp.Error = true
		resp.Message = "Error retrieving data from database: " + err.Error()
		app.writeJSON(w, http.StatusOK, resp)
		return
	}
	//Get product details for the product id

	resp.Error = false
	resp.Message = "Data fetched successfully"
	resp.ProductItems = productItem
	app.writeJSON(w, http.StatusOK, resp)
}

// FetchSoldProductItemBySerialNumber retrieve sold product item by serial number
func (app *application) FetchSoldProductItemBySerialNumber(w http.ResponseWriter, r *http.Request) {

	//Define struct for JSON object
	var payload struct {
		SerialNumber string `json:"product_serial_number"`
	}
	var resp struct {
		Error        bool             `json:"error,omitempty"`
		Message      string           `json:"message,omitempty"`
		ProductItem  *models.Product  `json:"product_item,omitempty"`
		SalesHistory *models.Sale     `json:"sales_history,omitempty"`
		Customer     *models.Customer `json:"customer,omitempty"`
	}
	//Read JSON body
	err := app.readJSON(w, r, &payload)
	if err != nil {
		resp.Error = true
		resp.Message = "cannot read JSON: " + err.Error()
		app.writeJSON(w, http.StatusOK, resp)
		return
	}

	var productItem *models.Product
	var salesHistory models.Sale
	var customer models.Customer

	//Get product item details for the given serial number(e.g. id,serial_number, product_id,purchase_history_id,warranty, max_retail_price, purchase_rate)
	productItem, err = app.DB.GetItemDetailsBySerialNumber(payload.SerialNumber)
	if err == sql.ErrNoRows {
		resp.Message = "Product item not found"
		app.writeJSON(w, http.StatusOK, resp)
		return
	} else if err != nil {
		app.badRequest(w, err)
		return
	}

	//Get sales history for the fetched product item
	salesHistory, err = app.DB.GetSalesHistoryByID(productItem.ProductMetadata[0].SalesHistoryID)
	if err == sql.ErrNoRows {
		resp.Message = "sales history not found for the item"
		app.writeJSON(w, http.StatusOK, resp)
		return
	} else if err != nil {
		app.badRequest(w, err)
		return
	}
	//Get customer details for the fetched product item
	customer, err = app.DB.GetCustomerByID(salesHistory.CustomerID)
	if err == sql.ErrNoRows {
		resp.Message = "sales history not found for the item"
		app.writeJSON(w, http.StatusOK, resp)
		return
	} else if err != nil {
		app.badRequest(w, err)
		return
	}

	resp.Error = false
	resp.Message = "Data fetched successfully"
	resp.ProductItem = productItem
	resp.SalesHistory = &salesHistory
	resp.Customer = &customer
	app.writeJSON(w, http.StatusOK, resp)
}

// FetchProductItemBySerialNumber retrieve product item by serial number
func (app *application) FetchProductItemBySerialNumber(w http.ResponseWriter, r *http.Request) {

	//Define struct for JSON object
	var payload struct {
		SerialNumber string `json:"product_serial_number"`
	}
	var resp struct {
		Error        bool            `json:"error,omitempty"`
		Message      string          `json:"message,omitempty"`
		ProductItems *models.Product `json:"product_items,omitempty"`
	}
	//Read JSON body
	err := app.readJSON(w, r, &payload)
	if err != nil {
		resp.Error = true
		resp.Message = "cannot read JSON: " + err.Error()
		app.writeJSON(w, http.StatusOK, resp)
		return
	}
	//Get product item details for the given serial number(e.g. id,serial_number, product_id,purchase_history_id,warranty, max_retail_price, purchase_rate)
	productItem, err := app.DB.GetItemDetailsBySerialNumber(payload.SerialNumber)
	if err != nil {
		resp.Error = true
		resp.Message = "Error retrieving data from database: " + err.Error()
		app.writeJSON(w, http.StatusOK, resp)
		return
	}
	//Get product details for the product id
	resp.Error = false
	resp.Message = "Data fetched successfully"
	resp.ProductItems = productItem
	app.writeJSON(w, http.StatusOK, resp)
}

// GetMemoListBySupplierID return memo list against supplierID in JSON format
func (app *application) GetMemoListBySupplierID(w http.ResponseWriter, r *http.Request) {
	var supplier models.Supplier
	err := app.readJSON(w, r, &supplier)
	if err != nil {
		app.badRequest(w, err) //send error response
		return
	}
	//Memo
	var resp struct {
		Error    bool               `json:"error,omitempty"`
		Message  string             `json:"message,omitempty"`
		Purchase []*models.Purchase `json:"purchase,omitempty"`
	}

	//retrieve suppliers from the database
	purchase, err := app.DB.GetMemoListBySupplierID(supplier.ID)
	if err == sql.ErrNoRows {
		resp.Message += "||No Memo Available For this selected supplier||"
	} else if err != nil {
		app.badRequest(w, err) //send error response
		return
	}

	//TODO: Retrieve accounts and send to frontend
	resp.Error = false
	resp.Message = "data Successfully fetched"
	resp.Purchase = purchase
	app.writeJSON(w, http.StatusOK, resp)
}

// GetMemoListByCustomerID return memo list against customer id in JSON format
func (app *application) GetMemoListByCustomerID(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		CustomerID int `json:"customer_id"`
	}
	err := app.readJSON(w, r, &payload)
	if err != nil {
		app.badRequest(w, err) //send error response
		return
	}
	//Memo
	var resp struct {
		Error   bool           `json:"error,omitempty"`
		Message string         `json:"message,omitempty"`
		Sale    []*models.Sale `json:"sales,omitempty"`
	}

	//retrieve memo list for given supplier id from sales_history table
	sale, err := app.DB.GetMemoListByCustomerID(payload.CustomerID)
	if err == sql.ErrNoRows {
		resp.Message = "No Memo Available"
		app.writeJSON(w, http.StatusOK, resp) //send error response
		return
	} else if err != nil {
		app.badRequest(w, err) //send error response
		return
	}

	//TODO: Retrieve accounts and send to frontend
	resp.Error = false
	resp.Message = "data Succesfully fetched"
	resp.Sale = sale
	app.writeJSON(w, http.StatusOK, resp)
}

// ReturnProductsToSupplier read json and update product items state from purchase to purchase-return
func (app *application) ReturnProductsToSupplier(w http.ResponseWriter, r *http.Request) {
	var ReturnProductsInfo struct {
		JobID           string          `json:"job_id"`
		PurchaseHistory models.Purchase `json:"purchase_history"`
		ReturnedDate    string          `json:"returned_date"`
		SupplierID      int             `json:"supplier_id"`
		ProductUnitsID  []int           `json:"product_units_id"`
		TotalUnits      int             `json:"total_units"`
		TotalPrices     int             `json:"total_prices"`
	}
	err := app.readJSON(w, r, &ReturnProductsInfo)

	if err != nil {
		app.badRequest(w, err)
		return
	}
	//Update product_serial_numbers
	id, err := app.DB.ReturnProductUnitsToSupplier(ReturnProductsInfo.PurchaseHistory, ReturnProductsInfo.JobID, ReturnProductsInfo.ReturnedDate, ReturnProductsInfo.ProductUnitsID, ReturnProductsInfo.TotalUnits, ReturnProductsInfo.TotalPrices)
	var resp struct {
		Error   bool   `json:"error,omitempty"`
		Message string `json:"message,omitempty"`
		ID      int    `json:"id"`
	}
	if err != nil {
		app.badRequest(w, err)
		return
	}
	resp.Error = false
	resp.Message = "Products Returned to supplier succesfully"
	resp.ID = id

	app.writeJSON(w, http.StatusOK, resp)
}

// RestockProduct add new purchased product to database
func (app *application) RestockProduct(w http.ResponseWriter, r *http.Request) {
	var purchase_details models.PurchasePayload
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
		Message: "restock product Successfully",
		Result:  purchase_details,
	}
	app.writeJSON(w, http.StatusOK, resp)
}

// SaleProductsToCustomer sale products, update product items state from in stock to sold in the database
func (app *application) SaleProductsToCustomer(w http.ResponseWriter, r *http.Request) {
	var err error
	var salesInfo models.SalesInvoice

	err = app.readJSON(w, r, &salesInfo)
	if err != nil {
		app.badRequest(w, fmt.Errorf("ERROR: SaleProductsToCustomer:Unable to read JSON => %w", err))
		return
	}
	err = app.DB.SaleProductsToCustomer(&salesInfo)
	if err != nil {
		app.badRequest(w, fmt.Errorf("ERROR: SaleProductsToCustomer => %w", err))
		return
	}
	cp, err := app.DB.GetCompanyProfile()
	if err != nil {
		app.badRequest(w, fmt.Errorf("ERROR: SaleProductsToCustomer => %w", err))
		return
	}
	var resp struct {
		Error            bool                   `json:"error"`
		Message          string                 `json:"message"`
		SalesInvoiceData *models.SalesInvoice   `json:"sales_invoice_data"`
		CompanyProfile   *models.CompanyProfile `json:"company_profile"`
	}

	resp.Error = false
	resp.Message = "product checked out successfully"
	resp.SalesInvoiceData = &salesInfo
	resp.CompanyProfile = &cp
	app.writeJSON(w, http.StatusOK, resp)
}

// ReturnProductsFromCustomer updates database for returned products
func (app *application) ReturnProductsFromCustomer(w http.ResponseWriter, r *http.Request) {
	var SaleReturnPayload struct {
		Customer         *models.Customer  `json:"customer_info"`
		Products         []*models.Product `json:"product_Items"`
		SalesHistory     *models.Sale      `json:"sales_history"`
		SelectedItemsID  []int             `json:"selected_items"`
		SaleReturnDate   string            `json:"sale_return_date"`
		ReturnItemsCount int               `json:"return_items_number"`
		ReturnAmount     int               `json:"return_items_amount"`
		MemoNo           string            `json:"memo_no"`
	}
	// customer_info: customers[parseInt(document.getElementById("customer").value)],
	// product_Items: products,
	// sales_history: salesHistory,
	// selected_items: returnedID,
	// sale_return_date: document.getElementById("single_cal4").value,
	// return_items_number: total_returned_items,
	// return_items_amount: total_returned_values,
	// memo_no: document.getElementById("memo").value

	err := app.readJSON(w, r, &SaleReturnPayload)
	if err != nil {
		app.badRequest(w, fmt.Errorf("ERROR:=>ReturnProductsFromCustomer: Unable to read JSON : %w", err))
		return
	}
	err = app.DB.SaleReturnDB(
		SaleReturnPayload.SalesHistory,
		SaleReturnPayload.SelectedItemsID,
		SaleReturnPayload.SaleReturnDate,
		SaleReturnPayload.ReturnItemsCount,
		SaleReturnPayload.ReturnAmount,
		SaleReturnPayload.MemoNo,
	)
	if err != nil {
		app.badRequest(w, err)
		return
	}
	var resp = JSONResponse{
		Error:   false,
		Message: "Successful",
		Result:  SaleReturnPayload,
	}
	app.writeJSON(w, http.StatusOK, resp)
}

// GetPurchasePageDetails scrape data from the database to initialize purchase page
func (app *application) GetPurchasePageDetails(w http.ResponseWriter, r *http.Request) {
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

	//retrieve suppliers from the database
	suppliers, err := app.DB.GetActiveSuppliersIDAndName()
	if err == sql.ErrNoRows {
		resp.Message += "||No Supplier Available||"
	} else if err != nil {
		app.badRequest(w, err) //send error response
		return
	}
	//retrieve categories from the database
	categories, err := app.DB.GetActiveCategoryList()
	if err == sql.ErrNoRows {
		resp.Message += "||No Category Available||"
	} else if err != nil {
		app.badRequest(w, err) //send error response
		return
	}
	//retrieve brands from the database
	brands, err := app.DB.GetActiveBrands()
	if err == sql.ErrNoRows {
		resp.Message += "||No Brand Available||"
	} else if err != nil {
		app.badRequest(w, err) //send error response
		return
	}
	//retrieve products from the database
	products, err := app.DB.GetActiveProducts()
	if err == sql.ErrNoRows {
		resp.Message += "||No Product Available||"
	} else if err != nil {
		app.badRequest(w, err) //send error response
		return
	}
	//retrieve accounts from the database
	headAccounts, err := app.DB.GetAvailableHeadAccountsByType("CASH & BANK ACCOUNTS")
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

// GetSalePageDetails scrape data from the database to initialize sale page
func (app *application) GetSalePageDetails(w http.ResponseWriter, r *http.Request) {
	var resp struct {
		Error        bool                  `json:"error,omitempty"`
		Message      string                `json:"message,omitempty"`
		Customers    []*models.Customer    `json:"customers,omitempty"`
		Categories   []*models.Category    `json:"categories,omitempty"`
		Brands       []*models.Brand       `json:"brands,omitempty"`
		Products     []*models.Product     `json:"products,omitempty"`
		HeadAccounts []*models.HeadAccount `json:"head_accounts,omitempty"`
	}

	//retrieve suppliers from the database
	customers, err := app.DB.GetActiveCustomersIDAndName()
	if err == sql.ErrNoRows {
		resp.Message += "||No Supplier Available||"
	} else if err != nil {
		app.badRequest(w, err) //send error response
		return
	}
	//retrieve categories from the database
	categories, err := app.DB.GetActiveCategoryList()
	if err == sql.ErrNoRows {
		resp.Message += "||No Category Available||"
	} else if err != nil {
		app.badRequest(w, err) //send error response
		return
	}
	//retrieve brands from the database
	brands, err := app.DB.GetActiveBrands()
	if err == sql.ErrNoRows {
		resp.Message += "||No Brand Available||"
	} else if err != nil {
		app.badRequest(w, err) //send error response
		return
	}
	//retrieve products from the database
	products, err := app.DB.GetActiveProducts()
	if err == sql.ErrNoRows {
		resp.Message += "||No Product Available||"
	} else if err != nil {
		app.badRequest(w, err) //send error response
		return
	}
	//retrieve accounts from the database
	headAccounts, err := app.DB.GetAvailableHeadAccountsByType("CASH & BANK ACCOUNTS")
	if err == sql.ErrNoRows {
		resp.Message += "||No Accounts Available||"
	} else if err != nil {
		app.badRequest(w, err) //send error response
		return
	}

	//TODO: Retrieve accounts and send to frontend
	resp.Error = false
	resp.Message += "||data Succesfully fetched||"
	resp.Customers = customers
	resp.Categories = categories
	resp.Brands = brands
	resp.Products = products
	resp.HeadAccounts = headAccounts

	app.writeJSON(w, http.StatusOK, resp)
}

// GetReceiveCollectionPageDetails scrape data from the database to initialize receive-collection page
func (app *application) GetReceiveCollectionPageDetails(w http.ResponseWriter, r *http.Request) {

	var resp struct {
		Error        bool                  `json:"error,omitempty"`
		Message      string                `json:"message,omitempty"`
		Customers    []*models.Customer    `json:"customers,omitempty"`
		HeadAccounts []*models.HeadAccount `json:"head_accounts,omitempty"`
	}

	//retrieve suppliers from the database
	customers, err := app.DB.GetCreditCustomersDetails()
	if err == sql.ErrNoRows {
		resp.Message += "||No Supplier Available||"
	} else if err != nil {
		app.badRequest(w, err) //send error response
		return
	}

	//retrieve accounts from the database
	headAccounts, err := app.DB.GetAvailableHeadAccountsByType("CASH & BANK ACCOUNTS")
	if err == sql.ErrNoRows {
		resp.Message += "||No Accounts Available||"
	} else if err != nil {
		app.badRequest(w, err) //send error response
		return
	}

	//TODO: Retrieve accounts and send to frontend
	resp.Error = false
	resp.Message += "||data Succesfully fetched||"
	resp.Customers = customers
	resp.HeadAccounts = headAccounts

	app.writeJSON(w, http.StatusOK, resp)
}

// CompleteReceiveCollection handles for completing reception process of that day
func (app *application) CompleteReceiveCollectionProcess(w http.ResponseWriter, r *http.Request) {
	var Summary []*models.Reception
	var resp struct {
		Error   bool   `json:"error,omitempty"`
		Message string `json:"message,omitempty"`
	}

	err := app.readJSON(w, r, &Summary)
	if err != nil {
		resp.Error = true
		resp.Message = "Unable to read JSON: " + err.Error()
		app.writeJSON(w, http.StatusAccepted, resp)
		return
	}

	//update Cash-Bank account
	//set current_balance += received_amount
	// err := app.DB.UpdateHeadAccountBalance(payload.Summary.DestinationAccount.ID, payload.Summary.ReceivedAmount)
	//update customer account
	//set due_mount -= received_amount
	//TODO::::::::::::::::::::::::::::::::::::::::::
	err = app.DB.CompleteReceiveCollectionTransactions(Summary)
	if err != nil {
		resp.Error = true
		resp.Message = "Unable to complete receive and collection process: " + err.Error()
		app.writeJSON(w, http.StatusAccepted, resp)
		return
	}
	//complete the process
	resp.Error = false
	resp.Message = "Amount received successfully"
	app.writeJSON(w, http.StatusOK, resp)
}

// GetPaymentPageDetails scrape data from the database to initialize payment page
func (app *application) GetPaymentPageDetails(w http.ResponseWriter, r *http.Request) {

	var resp struct {
		Error        bool                  `json:"error,omitempty"`
		Message      string                `json:"message,omitempty"`
		Suppliers    []*models.Supplier    `json:"suppliers,omitempty"`
		HeadAccounts []*models.HeadAccount `json:"head_accounts,omitempty"`
	}

	//retrieve suppliers from the database
	suppliers, err := app.DB.GetCreditSuppliersDetails()
	if err == sql.ErrNoRows {
		resp.Message += "||No Supplier Available||"
	} else if err != nil {
		app.badRequest(w, err) //send error response
		return
	}

	//retrieve accounts from the database
	headAccounts, err := app.DB.GetAvailableHeadAccountsByType("CASH & BANK ACCOUNTS")
	if err == sql.ErrNoRows {
		resp.Message += "||No Accounts Available||"
	} else if err != nil {
		app.badRequest(w, err) //send error response
		return
	}

	//TODO: Retrieve accounts and send to frontend
	resp.Error = false
	resp.Message += "||data Successfully fetched||"
	resp.Suppliers = suppliers
	resp.HeadAccounts = headAccounts

	app.writeJSON(w, http.StatusOK, resp)
}

// CompletePaymentProcess handles for completing payment process of that day
func (app *application) CompletePaymentProcess(w http.ResponseWriter, r *http.Request) {
	var Summary []*models.Payment
	var resp struct {
		Error   bool   `json:"error,omitempty"`
		Message string `json:"message,omitempty"`
	}

	err := app.readJSON(w, r, &Summary)
	if err != nil {
		resp.Error = true
		resp.Message = "Unable to read JSON: " + err.Error()
		app.writeJSON(w, http.StatusAccepted, resp)
		return
	}
	err = app.DB.CompletePaymentTransactions(Summary)
	if err != nil {
		resp.Error = true
		resp.Message = "Unable to complete payment transaction process: " + err.Error()
		app.writeJSON(w, http.StatusAccepted, resp)
		return
	}
	//complete the process
	resp.Error = false
	resp.Message = "Amount paid successfully"
	app.writeJSON(w, http.StatusOK, resp)
}

// GetAmountTransferPageDetails handle amount transfer between to accounts
func (app *application) GetAmountTransferPageDetails(w http.ResponseWriter, r *http.Request) {
	var resp struct {
		Error    bool                  `json:"error,omitempty"`
		Message  string                `json:"message,omitempty"`
		Accounts []*models.HeadAccount `json:"head_accounts"`
	}

	accounts, err := app.DB.GetAvailableHeadAccountsByType("CASH & BANK ACCOUNTS")
	if err != nil {
		app.badRequest(w, fmt.Errorf("ERROR: GetAmountTransferPageDetails => %w", err))
		return
	}
	resp.Message = "Data fetched successfully"
	resp.Accounts = accounts
	app.writeJSON(w, http.StatusOK, resp)
}

// CompleteAmountTransferProcess handles for completing amount transfer process of that day
func (app *application) CompleteAmountTransferProcess(w http.ResponseWriter, r *http.Request) {
	var Summary []*models.AmountTransfer

	var resp struct {
		Error   bool   `json:"error,omitempty"`
		Message string `json:"message,omitempty"`
	}

	err := app.readJSON(w, r, &Summary)
	if err != nil {
		resp.Error = true
		resp.Message = "Unable to read JSON: " + err.Error()
		app.writeJSON(w, http.StatusAccepted, resp)
		return
	}
	err = app.DB.CompleteAmountTransferTransactions(Summary)
	if err != nil {
		resp.Error = true
		resp.Message = "Unable to complete amount transfer process: " + err.Error()
		app.writeJSON(w, http.StatusAccepted, resp)
		return
	}
	//complete the process
	resp.Error = false
	resp.Message = "Amount transferred successfully"
	app.writeJSON(w, http.StatusOK, resp)
}

// GetAmountPayableDetails scarp data for amount payable page
func (app *application) GetAmountPayablePageDetails(w http.ResponseWriter, r *http.Request) {
	var resp struct {
		Error     bool               `json:"error,omitempty"`
		Message   string             `json:"message,omitempty"`
		Suppliers []*models.Supplier `json:"suppliers"`
		Customers []*models.Customer `json:"customers"`
	}

	suppliers, err := app.DB.GetActiveSuppliersIDAndName()
	if err != nil {
		app.badRequest(w, fmt.Errorf("ERROR: GetAmountPayableDetails => %w", err))
		return
	}
	customers, err := app.DB.GetActiveCustomersIDAndName()
	if err != nil {
		app.badRequest(w, fmt.Errorf("ERROR: GetAmountPayableDetails => %w", err))
		return
	}
	resp.Message = "Data fetched successfully"
	resp.Suppliers = suppliers
	resp.Customers = customers
	app.writeJSON(w, http.StatusOK, resp)
}

// CompleteAmountPayableProcess completes amount payable process
func (app *application) CompleteAmountPayableProcess(w http.ResponseWriter, r *http.Request) {
	var amountPayableSummary []*models.AmountPayable
	var resp struct {
		Error   bool   `json:"error,omitempty"`
		Message string `json:"message,omitempty"`
	}
	err := app.readJSON(w, r, &amountPayableSummary)
	if err != nil {
		app.badRequest(w, fmt.Errorf("unable to read JSON: %w", err))
		return
	}

	err = app.DB.CompleteAmountPayableTransactions(amountPayableSummary)
	if err != nil {
		app.badRequest(w, fmt.Errorf("unable to Insert data: %w", err))
		return
	}
	resp.Message = "Success"
	app.writeJSON(w, http.StatusOK, resp)
}

// GetAmountReceivableDetails scarp data for amount receivable page
func (app *application) GetAmountReceivablePageDetails(w http.ResponseWriter, r *http.Request) {
	var resp struct {
		Error     bool               `json:"error,omitempty"`
		Message   string             `json:"message,omitempty"`
		Suppliers []*models.Supplier `json:"suppliers"`
		Customers []*models.Customer `json:"customers"`
	}

	suppliers, err := app.DB.GetActiveSuppliersIDAndName()
	if err != nil {
		app.badRequest(w, fmt.Errorf("ERROR: GetAmountReceivableDetails => %w", err))
		return
	}
	customers, err := app.DB.GetActiveCustomersIDAndName()
	if err != nil {
		app.badRequest(w, fmt.Errorf("ERROR: GetAmountReceivableDetails => %w", err))
		return
	}
	resp.Message = "Data fetched successfully"
	resp.Suppliers = suppliers
	resp.Customers = customers
	app.writeJSON(w, http.StatusOK, resp)
}

// CompleteAmountReceivableProcess completes amount receivable process
func (app *application) CompleteAmountReceivableProcess(w http.ResponseWriter, r *http.Request) {
	var amountReceivableSummary []*models.AmountReceivable
	var resp struct {
		Error   bool   `json:"error,omitempty"`
		Message string `json:"message,omitempty"`
	}
	err := app.readJSON(w, r, &amountReceivableSummary)
	if err != nil {
		resp.Error = true
		resp.Message = "Unable to read JSON: " + err.Error()
		app.writeJSON(w, http.StatusAccepted, resp)
		return
	}

	err = app.DB.CompleteAmountReceivableTransactions(amountReceivableSummary)
	if err != nil {
		resp.Error = true
		resp.Message = err.Error()
		app.writeJSON(w, http.StatusAccepted, resp)
		return
	}
	resp.Message = "Success"
	app.writeJSON(w, http.StatusOK, resp)
}

// GetExpensesPageDetails scarp data for expenses page
func (app *application) GetExpensesPageDetails(w http.ResponseWriter, r *http.Request) {
	var resp struct {
		Error       bool                  `json:"error"`
		Message     string                `json:"message"`
		ExpenseList []*models.ExpenseType `json:"expense_list"`
		Accounts    []*models.HeadAccount `json:"head_accounts"`
	}

	accounts, err := app.DB.GetAvailableHeadAccountsByType("CASH & BANK ACCOUNTS")
	if err != nil {
		app.badRequest(w, fmt.Errorf("ERROR: GetExpensesPageDetails => %w", err))
		return
	}
	expList, err := app.DB.GetExpenseList()
	if err != nil {
		app.badRequest(w, fmt.Errorf("ERROR: GetExpensesPageDetails => %w", err))
		return
	}
	resp.Message = "Data fetched successfully"
	resp.ExpenseList = expList
	resp.Accounts = accounts
	app.writeJSON(w, http.StatusOK, resp)
}

// CompleteExpensesProcess completes expenses process
func (app *application) CompleteExpensesProcess(w http.ResponseWriter, r *http.Request) {
	var expenses []*models.Expense
	var resp struct {
		Error   bool   `json:"error,omitempty"`
		Message string `json:"message,omitempty"`
	}
	err := app.readJSON(w, r, &expenses)
	if err != nil {
		resp.Error = true
		resp.Message = "Unable to read JSON: " + err.Error()
		app.writeJSON(w, http.StatusAccepted, resp)
		return
	}

	err = app.DB.CompleteExpensesTransactions(expenses)
	if err != nil {
		resp.Error = true
		resp.Message = "Unable to complete expense transaction process: " + err.Error()
		app.writeJSON(w, http.StatusAccepted, resp)
		return
	}
	resp.Message = "Expense process completed successfully"
	app.writeJSON(w, http.StatusOK, resp)
}

// GetFundAcquisitionPageDetails scarp data for fund acquisition page
func (app *application) GetFundAcquisitionPageDetails(w http.ResponseWriter, r *http.Request) {

	var resp struct {
		Error        bool                  `json:"error"`
		Message      string                `json:"message"`
		StakeHolders []*models.StakeHolder `json:"stakeholders"`
		Accounts     []*models.HeadAccount `json:"head_accounts"`
	}

	acc, err := app.DB.GetAvailableHeadAccountsByType("CAPITAL ACCOUNTS")
	if err != nil {
		app.badRequest(w, fmt.Errorf("ERROR: GetAdjustmentPageDetails => %w", err))
		return
	}
	resp.Accounts = append(resp.Accounts, acc...)
	acc, err = app.DB.GetAvailableHeadAccountsByType("LOAN ACCOUNTS")
	if err != nil {
		app.badRequest(w, fmt.Errorf("ERROR: GetAdjustmentPageDetails => %w", err))
		return
	}
	resp.Accounts = append(resp.Accounts, acc...)
	stk, err := app.DB.GeActiveStakeHolderList()
	if err != nil {
		app.badRequest(w, fmt.Errorf("ERROR: GetAdjustmentPageDetails => %w", err))
		return
	}
	resp.Message = "Data fetched successfully"
	resp.StakeHolders = stk
	app.writeJSON(w, http.StatusOK, resp)
}

func (app *application) CompleteFundAcquisitionProcess(w http.ResponseWriter, r *http.Request) {
	var Summary []*models.FundAcquisition

	var resp struct {
		Error   bool   `json:"error,omitempty"`
		Message string `json:"message,omitempty"`
	}

	err := app.readJSON(w, r, &Summary)
	if err != nil {
		resp.Error = true
		resp.Message = "Unable to read JSON: " + err.Error()
		app.writeJSON(w, http.StatusAccepted, resp)
		return
	}
	err = app.DB.CompleteFundAcquisitionProcess(Summary)
	if err != nil {
		resp.Error = true
		resp.Message = "Unable to complete amount transfer process: " + err.Error()
		app.writeJSON(w, http.StatusAccepted, resp)
		return
	}
	//complete the process
	resp.Error = false
	resp.Message = "Fund Acquisition successful"
	app.writeJSON(w, http.StatusOK, resp)
}

// ClaimWarrantyBySerialID handles claiming warranty process for a specific product item with serial ID
func (app *application) ClaimWarrantyBySerialID(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		ReportedProblem string           `json:"reported_problem"`
		ContactNumber   string           `json:"contact_number"`
		ReceivedBy      string           `json:"received_by"`
		Product         *models.Product  `json:"product_item_details"`
		Customer        *models.Customer `json:"customer_info"`
	}

	err := app.readJSON(w, r, &payload)
	if err != nil {
		app.badRequest(w, fmt.Errorf("handler Error - ClaimWarrantyBySerialID => readJSON: %w", err))
		return
	}

	//set contact number if blank
	if payload.ContactNumber == "" {
		payload.ContactNumber = payload.Customer.Mobile
	}
	// 	when requested for warranty:
	// 	step-1: insert new row at warranty_history table with data from product_serial_number and status = warranty claim, product_serial_id = current_serial_id
	// 	step-2 : update latest_warranty_history_id = pkid of warranty_history, warranty_history_ids = concat{warranty_history_ids,pkid of warranty_history}, updated_at = time.Now() in product_serial_number

	//MM-WC-randomAlphanumeric(6)+CurrentIndexOfWarrantyHistory Table //Warranty Claimed
	//generate 6 digits random Alphanumeric
	randomCode, err := app.GenerateRandomAlphanumericCode(6)
	if err != nil {
		app.badRequest(w, fmt.Errorf("handler Error - ClaimWarrantyBySerialID => GenerateRandomAlphanumericCode: %w", err))
		return
	}
	memoPrefix := "MM-WC-" + randomCode
	id, err := app.DB.AddNewWarrantyClaim(memoPrefix, payload.Product.ProductMetadata[0].ID, payload.Product.ProductMetadata[0].SerialNumber, payload.ContactNumber, payload.ReportedProblem, payload.ReceivedBy, payload.Product.ProductMetadata[0].WarrantyHistoryIDs)
	if err != nil {
		app.badRequest(w, fmt.Errorf("handler Error - ClaimWarrantyBySerialID => AddNewWarrantyClaim: %w", err))
		return
	}
	var resp struct {
		Error    bool             `json:"error,omitempty"`
		Message  string           `json:"message,omitempty"`
		Item     *models.Product  `json:"product_item_info,omitempty"`
		Warranty *models.Warranty `json:"warranty_info,omitempty"`
		Customer *models.Customer `json:"customer,omitempty"`
	}

	resp.Error = false
	resp.Message = "Added for warranty"
	resp.Item = payload.Product
	resp.Customer = payload.Customer
	var warranty = models.Warranty{
		ID:               id,
		Status:           "claimed",
		MemoNo:           memoPrefix + strconv.Itoa(id),
		PreviousSerialNo: payload.Product.ProductMetadata[0].SerialNumber,
		ContactNumber:    payload.ContactNumber,
		RequestedDate:    time.Now().Format("01/02/2006"),
		ReportedProblem:  payload.ReportedProblem,
		ReceivedBy:       payload.ReceivedBy,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}
	resp.Warranty = &warranty

	app.writeJSON(w, http.StatusOK, resp)
}

// GetClaimWarrantyList writes a list of Claimed warranty data to the response body
func (app *application) GetClaimWarrantyList(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		SearchType string `json:"search_type"`
	}
	err := app.readJSON(w, r, &payload)
	if err != nil {
		app.badRequest(w, fmt.Errorf("ERROR => GetClaimWarranty: %w", err))
		return
	}

	warrantyHistory, err := app.DB.GetWarrantyList(payload.SearchType)
	if err != nil {
		app.badRequest(w, fmt.Errorf("ERROR => GetClaimWarranty: %w", err))
		return
	}
	var resp struct {
		Error           bool               `json:"error,omitempty"`
		Message         string             `json:"message,omitempty"`
		WarrantyHistory []*models.Warranty `json:"warranty_history"`
	}

	resp.Error = false
	resp.Message = "Data Fatched Successfully"
	resp.WarrantyHistory = warrantyHistory
	app.writeJSON(w, http.StatusOK, resp)
}

// CheckoutWarrantyProduct updates database tables for checkout process
func (app *application) CheckoutWarrantyProduct(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		WarrantyHistoryID int    `json:"warranty_history_id"`
		ProductSerialID   int    `json:"product_serial_id"`
		ArrivalDate       string `json:"checkout_date"`
		NewSerialNumber   string `json:"new_serial_number"`
		OldSerialNumber   string `json:"old_serial_number"`
		Comment           string `json:"comment"`
	}
	err := app.readJSON(w, r, &payload)
	if err != nil {
		app.badRequest(w, fmt.Errorf("ERROR => CheckoutWarrantyProduct: %w", err))
		return
	}
	//
	if strings.TrimSpace(payload.NewSerialNumber) == "" {
		payload.NewSerialNumber = payload.OldSerialNumber
	}

	err = app.DB.CheckoutWarrantyProduct(payload.WarrantyHistoryID, payload.ProductSerialID, payload.ArrivalDate, payload.NewSerialNumber, payload.Comment)
	if err != nil {
		app.badRequest(w, fmt.Errorf("ERROR => CheckoutWarrantyProduct: %w", err))
		return
	}
	var resp struct {
		Error   bool   `json:"error"`
		Message string `json:"message"`
	}

	resp.Error = false
	resp.Message = "product checked out successfully"
	app.writeJSON(w, http.StatusOK, resp)
}

// DeliveryWarrantyProduct updates database tables for delivery process
func (app *application) DeliveryWarrantyProduct(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		WarrantyHistoryID int    `json:"warranty_history_id,omitempty"`
		ProductSerialID   int    `json:"product_serial_id,omitempty"`
		DeliveredBy       string `json:"delivered_by,omitempty"`
	}
	err := app.readJSON(w, r, &payload)
	if err != nil {
		app.badRequest(w, fmt.Errorf("ERROR => DeliveryWarrantyProduct: %w", err))
		return
	}

	err = app.DB.DeliveryWarrantyProduct(payload.WarrantyHistoryID, payload.ProductSerialID, payload.DeliveredBy)
	if err != nil {
		app.badRequest(w, fmt.Errorf("ERROR => DeliveryWarrantyProduct: %w", err))
		return
	}
	var resp struct {
		Error   bool   `json:"error"`
		Message string `json:"message"`
	}

	resp.Error = false
	resp.Message = "product delivered to customer successfully"
	app.writeJSON(w, http.StatusOK, resp)
}

// .......................Inventory Reports.......................
// GetCategoryListReport retrieves the category list
func (app *application) GetCategoryListReport(w http.ResponseWriter, r *http.Request) {
	categories, err := app.DB.GetCategoryListReport()
	if err != nil {
		app.badRequest(w, fmt.Errorf("ERROR:GetCategoryList: %w", err))
		return
	}

	var resp struct {
		Error      bool               `json:"error"`
		Message    string             `json:"message"`
		Categories []*models.Category `json:"categories"`
	}

	resp.Error = false
	resp.Message = "Data fetched successfully"
	resp.Categories = categories

	app.writeJSON(w, http.StatusOK, resp)
}

// GetBrandListReport retrieves the brand list
func (app *application) GetBrandListReport(w http.ResponseWriter, r *http.Request) {
	brands, err := app.DB.GetBrandListReport()
	if err != nil {
		app.badRequest(w, fmt.Errorf("ERROR:GetBrandList: %w", err))
		return
	}

	var resp struct {
		Error   bool            `json:"error"`
		Message string          `json:"message"`
		Brands  []*models.Brand `json:"brands"`
	}

	resp.Error = false
	resp.Message = "Data fetched successfully"
	resp.Brands = brands

	app.writeJSON(w, http.StatusOK, resp)
}

// GetProductListReport retrieves the product list
func (app *application) GetProductListReport(w http.ResponseWriter, r *http.Request) {
	products, err := app.DB.GetProductListReport()
	if err != nil {
		app.badRequest(w, fmt.Errorf("ERROR:GetProductList: %w", err))
		return
	}

	var resp struct {
		Error    bool              `json:"error"`
		Message  string            `json:"message"`
		Products []*models.Product `json:"products"`
	}

	resp.Error = false
	resp.Message = "Data fetched successfully"
	resp.Products = products

	app.writeJSON(w, http.StatusOK, resp)
}

// GetLowStockProductReport retrieves the product list that marked as low stock
func (app *application) GetLowStockProductReport(w http.ResponseWriter, r *http.Request) {
	products, err := app.DB.GetLowStockProductReport()
	if err != nil {
		app.badRequest(w, fmt.Errorf("ERROR:GetLowStockProductReport: %w", err))
		return
	}

	var resp struct {
		Error    bool              `json:"error"`
		Message  string            `json:"message"`
		Products []*models.Product `json:"products"`
	}

	resp.Error = false
	resp.Message = "Data fetched successfully"
	resp.Products = products

	app.writeJSON(w, http.StatusOK, resp)
}

// GetServiceListReport retrieves the services list
func (app *application) GetServiceListReport(w http.ResponseWriter, r *http.Request) {
	services, err := app.DB.GetServiceListReport()
	if err != nil {
		app.badRequest(w, fmt.Errorf("ERROR:GetServicesList: %w", err))
		return
	}

	var resp struct {
		Error    bool              `json:"error"`
		Message  string            `json:"message"`
		Services []*models.Service `json:"services"`
	}

	resp.Error = false
	resp.Message = "Data fetched successfully"
	resp.Services = services

	app.writeJSON(w, http.StatusOK, resp)
}

// GetPurchaseHistoryReport retrieves the purchase history
func (app *application) GetPurchaseHistoryReport(w http.ResponseWriter, r *http.Request) {
	purchase, err := app.DB.GetPurchaseHistoryReport()
	if err != nil {
		app.badRequest(w, fmt.Errorf("ERROR: GetPurchaseHistoryReport => %w", err))
		return
	}
	var resp struct {
		Error    bool               `json:"error"`
		Message  string             `json:"message"`
		Purchase []*models.Purchase `json:"purchase_history"`
	}

	resp.Error = false
	resp.Message = "Data fetched successfully"
	resp.Purchase = purchase

	app.writeJSON(w, http.StatusOK, resp)
}

// // GetSalesHistoryReport retrieves the sales history
func (app *application) GetSalesHistoryReport(w http.ResponseWriter, r *http.Request) {
	sale, err := app.DB.GetSalesHistoryReport()
	if err != nil {
		app.badRequest(w, fmt.Errorf("ERROR: GetPurchaseHistoryReport => %w", err))
		return
	}
	var resp struct {
		Error    bool           `json:"error"`
		Message  string         `json:"message"`
		Purchase []*models.Sale `json:"sales_history"`
	}

	resp.Error = false
	resp.Message = "Data fetched successfully"
	resp.Purchase = sale

	app.writeJSON(w, http.StatusOK, resp)
}

// .....................Accounts Handlers......................
func (app *application) GetCustomerDueReport(w http.ResponseWriter, r *http.Request) {
	customers, err := app.DB.GetCustomerDueHistoryReport()
	if err != nil {
		app.badRequest(w, fmt.Errorf("ERROR:GetCustomerDueReport: %w", err))
		return
	}

	var resp struct {
		Error   bool           `json:"error"`
		Message string         `json:"message"`
		Report  []*models.Customer `json:"report"`
	}

	resp.Error = false
	resp.Message = "Data fetched successfully"
	resp.Report = customers

	app.writeJSON(w, http.StatusOK, resp)
}
func (app *application) GetSupplierDueReport(w http.ResponseWriter, r *http.Request) {
	suppliers, err := app.DB.GetSupplierDueHistoryReport()
	if err != nil {
		app.badRequest(w, fmt.Errorf("ERROR:GetSupplierDueReport: %w", err))
		return
	}

	var resp struct {
		Error   bool           `json:"error"`
		Message string         `json:"message"`
		Report  []*models.Supplier `json:"report"`
	}

	resp.Error = false
	resp.Message = "Data fetched successfully"
	resp.Report = suppliers

	app.writeJSON(w, http.StatusOK, resp)
}

func (app *application) GetTransactionsReport(w http.ResponseWriter, r *http.Request) {
	trx, err := app.DB.GetTransactionsHistoryReport()
	if err != nil {
		app.badRequest(w, fmt.Errorf("ERROR:GetTransactionsReport: %w", err))
		return
	}

	var resp struct {
		Error   bool                  `json:"error"`
		Message string                `json:"message"`
		Report  []*models.Transaction `json:"report"`
	}
	resp.Error = false
	resp.Message = "Data fetched successfully"
	resp.Report = trx

	fmt.Println(resp)
	app.writeJSON(w, http.StatusOK, resp)
}
func (app *application) GetCashBankStatement(w http.ResponseWriter, r *http.Request) {
	trx, err := app.DB.GetCashBankStatement()
	if err != nil {
		app.badRequest(w, fmt.Errorf("ERROR:GetCashBankStatement: %w", err))
		return
	}

	var resp struct {
		Error   bool                  `json:"error"`
		Message string                `json:"message"`
		Report  []*models.Transaction `json:"report"`
	}
	resp.Error = false
	resp.Message = "Data fetched successfully"
	resp.Report = trx

	fmt.Println(resp)
	app.writeJSON(w, http.StatusOK, resp)
}
func (app *application) GetLedgerBookDetails(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		AccountType string `json:"account_type"`
		AccountID   int    `json:"account_id"`
	}
	err := app.readJSON(w, r, &payload)
	if err != nil {
		app.badRequest(w, fmt.Errorf("ERROR:GetLedgerBookDetails: %w", err))
		return
	}
	trx, err := app.DB.GetLedgerBookDetails(payload.AccountType, payload.AccountID)
	if err != nil {
		app.badRequest(w, fmt.Errorf("ERROR:GetLedgerBookDetails: %w", err))
		return
	}
	cp, err := app.DB.GetCompanyProfile()
	if err != nil {
		app.badRequest(w, fmt.Errorf("ERROR:GetLedgerBookDetails: %w", err))
		return
	}

	var resp struct {
		Error          bool                   `json:"error"`
		Message        string                 `json:"message"`
		Report         []*models.Transaction  `json:"report"`
		CompanyProfile *models.CompanyProfile `json:"company_profile"`
	}
	resp.Error = false
	resp.Message = "Data fetched successfully"
	resp.Report = trx
	resp.CompanyProfile = &cp

	fmt.Println(resp)
	app.writeJSON(w, http.StatusOK, resp)
}
func (app *application) GetExpensesReport(w http.ResponseWriter, r *http.Request) {
	exp, err := app.DB.GetExpensesHistoryReport()
	if err != nil {
		app.badRequest(w, fmt.Errorf("ERROR:GetExpensesReport: %w", err))
		return
	}

	var resp struct {
		Error   bool                  `json:"error"`
		Message string                `json:"message"`
		Report  []*models.Transaction `json:"report"`
	}
	resp.Error = false
	resp.Message = "Data fetched successfully"
	resp.Report = exp

	fmt.Println(resp)
	app.writeJSON(w, http.StatusOK, resp)
}
func (app *application) GetIncomeStatementReport(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		StartDate string `json:"start_date"`
		EndDate   string `json:"end_date"`
	}

	err := app.readJSON(w, r, &payload)
	if err != nil {
		app.badRequest(w, fmt.Errorf("ERROR:GetIncomeStatementReport: %w", err))
		return
	}

	var resp struct {
		Error          bool                   `json:"error"`
		Message        string                 `json:"message"`
		Report         models.IncomeStatement `json:"report"`
		CompanyProfile models.CompanyProfile  `json:"company_profile"`
	}

	ins, err := app.DB.GetIncomeStatementData(payload.StartDate, payload.EndDate)
	if err != nil {
		app.badRequest(w, fmt.Errorf("ERROR:GetIncomeStatementReport: %w", err))
		return
	}
	cp, err := app.DB.GetCompanyProfile()
	if err != nil {
		app.badRequest(w, fmt.Errorf("ERROR:GetIncomeStatementReport: %w", err))
		return
	}
	resp.Error = false
	resp.Message = "Data fetched successfully"
	resp.Report = ins
	resp.CompanyProfile = cp

	fmt.Println(resp)
	app.writeJSON(w, http.StatusOK, resp)
}
func (app *application) GetTopSheetReport(w http.ResponseWriter, r *http.Request) {
	topSheetData, err := app.DB.GetTopSheetReport()
	if err != nil {
		app.badRequest(w, fmt.Errorf("ERROR:GetTopSheetReport: %w", err))
		return
	}

	cp, err := app.DB.GetCompanyProfile()
	if err != nil {
		app.badRequest(w, fmt.Errorf("ERROR:GetCompanyProfile: %w", err))
		return
	}

	var resp struct {
		Error          bool                  `json:"error"`
		Message        string                `json:"message"`
		Report         []*models.TopSheet    `json:"report"`
		CompanyProfile models.CompanyProfile `json:"company_profile"`
	}
	resp.Error = false
	resp.Message = "Data fetched successfully"
	resp.Report = topSheetData
	resp.CompanyProfile = cp

	fmt.Println(resp)
	app.writeJSON(w, http.StatusOK, resp)
}
func (app *application) GetBalanceSheetReport(w http.ResponseWriter, r *http.Request) {
	balanceSheet, err := app.DB.GetBalanceSheetReport()
	if err != nil {
		app.badRequest(w, fmt.Errorf("ERROR:GetTopSheetReport: %w", err))
		return
	}

	cp, err := app.DB.GetCompanyProfile()
	if err != nil {
		app.badRequest(w, fmt.Errorf("ERROR:GetCompanyProfile: %w", err))
		return
	}

	var resp struct {
		Error          bool                  `json:"error"`
		Message        string                `json:"message"`
		Report         models.BalanceSheet   `json:"report"`
		CompanyProfile models.CompanyProfile `json:"company_profile"`
	}
	resp.Error = false
	resp.Message = "Data fetched successfully"
	resp.Report = balanceSheet
	resp.CompanyProfile = cp

	fmt.Println(resp)
	app.writeJSON(w, http.StatusOK, resp)
}
func (app *application) GetTrialBalanceReport(w http.ResponseWriter, r *http.Request) {
	trialBalanceSheet, err := app.DB.GetTrialBalanceReport()
	if err != nil {
		app.badRequest(w, fmt.Errorf("ERROR:GetTopSheetReport: %w", err))
		return
	}

	cp, err := app.DB.GetCompanyProfile()
	if err != nil {
		app.badRequest(w, fmt.Errorf("ERROR:GetCompanyProfile: %w", err))
		return
	}

	var resp struct {
		Error          bool                  `json:"error"`
		Message        string                `json:"message"`
		Report         models.TrialBalance   `json:"report"`
		CompanyProfile models.CompanyProfile `json:"company_profile"`
	}
	resp.Error = false
	resp.Message = "Data fetched successfully"
	resp.Report = trialBalanceSheet
	resp.CompanyProfile = cp

	fmt.Println(resp)
	app.writeJSON(w, http.StatusOK, resp)
}
