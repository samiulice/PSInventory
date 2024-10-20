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

// GetActiveSuppliersIDAndName retruns a list of supplier's id and name
func (app *application) GetActiveSuppliersIDAndName(w http.ResponseWriter, r *http.Request) {
	//supplier
	var resp struct {
		Error     bool               `json:"error,omitempty"`
		Message   string             `json:"message,omitempty"`
		Suppliers []*models.Supplier `json:"suppliers,omitempty"`
	}

	//retrive suppliers from the database
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

	//retrive suppliers from the database
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

// FetchPurchaseMemoProductItems retrive purchased products list of a memo
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
	//retrive all product-serial of each product_id && purchase_id
	var products []*models.Product
	for _, v := range purchaseHistory {
		product, err := app.DB.GetInstockProductListByPurchaseIDAndProductID(v.ID, v.ProductID)
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
	resp.Message = "Data fetched succefully"
	resp.Product = products
	resp.PurchaseHistory = purchaseHistory
	app.writeJSON(w, http.StatusOK, resp)
}

// FetchPurchaseMemoProductItems retrive purchased products list of a memo
func (app *application) FetchSalesMemoProductItems(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		MemoNo string `json:"memo_no"`
	}
	err := app.readJSON(w, r, &payload) //read json body
	if err != nil {
		app.badRequest(w, err)
		return
	}
	//get purchase history associated with memo_no
	salesHistory, err := app.DB.GetSalesHistoryByMemoNo(payload.MemoNo)
	if err != nil {
		app.badRequest(w, fmt.Errorf("ErrorSalesHistory:: %v", err))
	}
	//Get product ids for this memo with associated purchase_id for the given memo
	//get detailed Product info for these ids
	//retrive all product-serial of each product_id && purchase_id
	var products []*models.Product
	for _, v := range salesHistory {
		product, err := app.DB.GetSoldProductListBySalesIDAndProductID(v.ID, v.SelectedItems[0].ProductID)
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
		Error        bool              `json:"error,omitempty"`
		Message      string            `json:"message,omitempty"`
		Products     []*models.Product `json:"products,omitempty"`
		SalesHistory []*models.Sale    `json:"sales_history,omitempty"`
	}
	resp.Error = false
	resp.Message = "Data fetched succefully"
	resp.Products = products
	// resp.SalesHistory = saleHistory
	app.writeJSON(w, http.StatusOK, resp)
}

// FetchProductItemsbyProductID retrive instock product items
func (app *application) FetchProductItemsbyProductID(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		ProductID int `json:"product_id"`
	}
	err := app.readJSON(w, r, &payload) //read json body
	if err != nil {
		app.badRequest(w, err)
		return
	}
	//get purchase history associated with memo_no
	productItems, err := app.DB.GetProductItemsListByProductID(payload.ProductID)
	if err != nil {
		app.badRequest(w, fmt.Errorf("ErrorPurchaseHistory:: %v", err))
	}

	var resp struct {
		Error        bool            `json:"error,omitempty"`
		Message      string          `json:"message,omitempty"`
		ProductItems *models.Product `json:"product_items,omitempty"`
	}
	resp.Error = false
	resp.Message = "Data fetched succefully"
	resp.ProductItems = productItems
	app.writeJSON(w, http.StatusOK, resp)
}

// FetchProductItembySerialNumber retrive product item by serial number
func (app *application) FetchInstockProductItembySerialNumber(w http.ResponseWriter, r *http.Request) {

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
		resp.Message = "Error retriving data from database: " + err.Error()
		app.writeJSON(w, http.StatusOK, resp)
		return
	}
	//Get product details for the product id

	resp.Error = false
	resp.Message = "Data fetched succefully"
	resp.ProductItems = productItem
	app.writeJSON(w, http.StatusOK, resp)
}

// FetchSoldProductItembySerialNumber retrive sold product item by serial number
func (app *application) FetchSoldProductItembySerialNumber(w http.ResponseWriter, r *http.Request) {

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
	productItem, err = app.DB.GetSoldItemDetailsBySerialNumber(payload.SerialNumber)
	if err == sql.ErrNoRows {
		resp.Message = "Product item not sold yet"
		app.writeJSON(w, http.StatusOK, resp)
		return
	} else if err != nil {
		app.badRequest(w, err)
		return
	}
	if len(productItem.ProductMetadata) > 0 {
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
	}
	resp.Error = false
	resp.Message = "Data fetched succefully"
	resp.ProductItem = productItem
	resp.SalesHistory = &salesHistory
	resp.Customer = &customer
	app.writeJSON(w, http.StatusOK, resp)
}

// FetchProductItembySerialNumber retrive product item by serial number
func (app *application) FetchProductItembySerialNumber(w http.ResponseWriter, r *http.Request) {

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
		resp.Message = "Error retriving data from database: " + err.Error()
		app.writeJSON(w, http.StatusOK, resp)
		return
	}
	//Get product details for the product id
	resp.Error = false
	resp.Message = "Data fetched succefully"
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

	//retrive suppliers from the database
	purchase, err := app.DB.GetMemoListBySupplierID(supplier.ID)
	if err == sql.ErrNoRows {
		resp.Message += "||No Memo Available For this selected supplier||"
	} else if err != nil {
		app.badRequest(w, err) //send error response
		return
	}

	//TODO: Retrive accounts and send to frontend
	resp.Error = false
	resp.Message += "data Succesfully fetched"
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

	//retrive memo list for given supplier id from sales_history talble
	sale, err := app.DB.GetMemoListByCustomerID(payload.CustomerID)
	if err == sql.ErrNoRows {
		resp.Message += "||No Memo Available For this selected customer||"
	} else if err != nil {
		app.badRequest(w, err) //send error response
		return
	}

	//TODO: Retrive accounts and send to frontend
	resp.Error = false
	resp.Message += "data Succesfully fetched"
	resp.Sale = sale
	app.writeJSON(w, http.StatusOK, resp)
}

// ReturnProductsToSupplier read json and update product items state from purchase to purchase-return
func (app *application) ReturnProductsToSupplier(w http.ResponseWriter, r *http.Request) {
	var ReturnProductsInfo struct {
		JobID          string `json:"job_id"`
		ReturnedDate   string `json:"returned_date"`
		SupplierID     int    `json:"supplier_id"`
		ProductUnitsID []int  `json:"product_units_id"`
		TotalUnits     int    `json:"total_units"`
		TotalPrices    int    `json:"total_prices"`
	}
	err := app.readJSON(w, r, &ReturnProductsInfo)

	if err != nil {
		app.badRequest(w, err)
		return
	}
	//Update product_serial_numbers
	id, err := app.DB.ReturnProductUnitsToSupplier(ReturnProductsInfo.JobID, ReturnProductsInfo.ReturnedDate, ReturnProductsInfo.ProductUnitsID, ReturnProductsInfo.TotalUnits, ReturnProductsInfo.TotalPrices)
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

// SaleProducts sale products, update product items state from instock to sold in the database
func (app *application) SaleProducts(w http.ResponseWriter, r *http.Request) {
	var err error
	var saleDetails models.Sale
	defer func() {
		if err != nil {
			app.badRequest(w, err)
			return
		}
		var resp = JSONResponse{
			Error:   false,
			Message: "sale products Succesfully",
			Result:  saleDetails,
		}
		app.writeJSON(w, http.StatusOK, resp)
	}()

	err = app.readJSON(w, r, &saleDetails)
	if err != nil {
		return
	}
	fmt.Println(saleDetails)
	err = app.DB.SaleProducts(&saleDetails)
	if err != nil {
		return
	}

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

	//retrive suppliers from the database
	suppliers, err := app.DB.GetActiveSuppliersIDAndName()
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
		LastIndex    int                   `json:"last_index,omitempty"`
	}

	//retrive suppliers from the database
	customers, err := app.DB.GetActiveCustomersIDAndName()
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
	lastIndex, err := app.DB.LastIndex("sales_history")
	if err != nil {
		app.badRequest(w, err) //send error response
		return
	}

	//TODO: Retrive accounts and send to frontend
	resp.Error = false
	resp.Message += "||data Succesfully fetched||"
	resp.Customers = customers
	resp.Categories = categories
	resp.Brands = brands
	resp.Products = products
	resp.HeadAccounts = headAccounts
	resp.LastIndex = lastIndex

	app.writeJSON(w, http.StatusOK, resp)
}

// GetReceiveCollectionPageDetails scrape data from the database to initialize receive-collection page
func (app *application) GetReceiveCollectionPageDetails(w http.ResponseWriter, r *http.Request) {
	/* input : nill
	output :
	data = {
		accounts :[
		head_accounts = {
			id, account_name, account_code
		},
		customers = {
			id, account_name, account_code, amount_payable, amount_receivable
		}
		]
	}
	*/
	app.infoLog.Println("Hit receive-collection handler")
	var resp struct {
		Error        bool                  `json:"error,omitempty"`
		Message      string                `json:"message,omitempty"`
		Customers    []*models.Customer    `json:"customers,omitempty"`
		HeadAccounts []*models.HeadAccount `json:"head_accounts,omitempty"`
	}

	//retrive suppliers from the database
	customers, err := app.DB.GetCreditCustomersDetails()
	if err == sql.ErrNoRows {
		resp.Message += "||No Supplier Available||"
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
	resp.Customers = customers
	resp.HeadAccounts = headAccounts

	app.writeJSON(w, http.StatusOK, resp)
}

// ClaimWarrantyBySerialID handels claiming warranty process for a specific product item with serial ID
func (app *application) ClaimWarrantyBySerialID(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		ReportedProblem string          `json:"reported_problem"`
		ContactNumber   string          `json:"contact_number"`
		ReceivedBy      string          `json:"received_by"`
		Product         *models.Product `json:"product_item_details"`
	}

	err := app.readJSON(w, r, &payload)
	if err != nil {
		app.badRequest(w, fmt.Errorf("handler Error - ClaimWarrantyBySerialID => readJSON: %w", err))
		return
	}

	// 	when requested for warranty:
	// 	step-1: insert new row at warranty_history table with data from product_serial_number and status = warranty claim, product_serial_id = current_serial_id
	// 	step-2 : update latest_warranty_history_id = pkid of warranty_history, warranty_history_ids = concat{warranty_history_ids,pkid of warranty_history}, updated_at = time.Now() in product_serial_number

	err = app.DB.AddNewWarrantyClaim(payload.Product.ProductMetadata[0].ID, payload.Product.ProductMetadata[0].SerialNumber, payload.ContactNumber, payload.ReportedProblem, payload.ReceivedBy)
	if err != nil {
		app.badRequest(w, fmt.Errorf("handler Error - ClaimWarrantyBySerialID => AddNewWarrantyClaim: %w", err))
		return
	}
	var resp struct {
		Error   bool            `json:"error,omitempty"`
		Message string          `json:"message,omitempty"`
		Item    *models.Product `json:"product_item"`
	}

	resp.Error = false
	resp.Message = "Added for warranty"
	resp.Item = payload.Product

	app.writeJSON(w, http.StatusOK, resp)
}
