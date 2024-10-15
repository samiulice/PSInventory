package dbrepo

import (
	"PSInventory/internal/models"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"
)

//.......................HR Management.......................

// AddHeadAccount inserts new head account information to the database
func (p *postgresDBRepo) AddHeadAccount(ha models.HeadAccount) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var id int

	stmt := `INSERT INTO public.head_accounts (account_code,account_name,account_status,current_amount,created_at,updated_at) 
				VALUES($1, $2, $3, $4, $5, $6) RETURNING id
	`
	err := p.DB.QueryRowContext(ctx, stmt,
		ha.AccountCode,
		ha.AccountName,
		ha.AccoutnStatus,
		ha.CurrentAmount,
		time.Now(),
		time.Now(),
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

// GetAvailableBrands returns a list of active brands from the database
func (p *postgresDBRepo) GetAvailableHeadAccounts() ([]*models.HeadAccount, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var headAccounts []*models.HeadAccount

	query := `
		SELECT 
			id, account_code, account_name, current_amount, created_at, updated_at
		FROM
			public.head_accounts
		WHERE 
			account_status = true
		`
	var rows *sql.Rows
	var err error

	rows, err = p.DB.QueryContext(ctx, query)
	if err != nil {
		return headAccounts, errors.New("100: " + err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var ha models.HeadAccount
		err = rows.Scan(
			&ha.ID,
			&ha.AccountCode,
			&ha.AccountName,
			&ha.CurrentAmount,
			&ha.CreatedAt,
			&ha.UpdatedAt,
		)
		if err != nil {
			return headAccounts, err
		}
		headAccounts = append(headAccounts, &ha)
	}
	return headAccounts, nil
}

//.......................HR Management.......................

// AddEmployee inserts new employee information to the database
func (p *postgresDBRepo) AddEmployee(employee models.Employee) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var id int

	stmt := `INSERT INTO public.employees (account_code,account_name,contact_person,division,district,upazila,area,mobile,email,account_status,monthly_salary,opening_balance,joining_date,created_at,updated_at) 
				VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15) RETURNING id
	`
	err := p.DB.QueryRowContext(ctx, stmt,
		employee.AccountCode,
		employee.AccountName,
		employee.ContactPerson,
		employee.Division,
		employee.District,
		employee.Upazila,
		employee.Area,
		employee.Mobile,
		employee.Email,
		employee.AccountStatus,
		employee.MonthlySalary,
		employee.OpeningBalance,
		time.Now(),
		time.Now(),
		time.Now(),
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

// GetEmployeeDetails retrive detailed info about an employee
func (p *postgresDBRepo) GetEmployeeByID(id int) (models.Employee, error) {
	// 	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	// 	defer cancel()

	var employee models.Employee
	var err error

	// 	query := `
	// 		SELECT
	// 			id, user_name, first_name, last_name, address, email, fb_id, whatsapp_id, x_id, linkedin_id, github_id, mobile, image_link, account_status,
	// 			credits, task_completed, task_cancelled, rating, kyc_upload_status, kyc_verification_status, created_at, updated_at, es.id, es.name, es.description
	// 		FROM
	// 			employees e
	// 			LEFT JOIN employee_status es on (account_status = es.id)
	// 		WHERE
	// 			id = $1
	// 		`

	// 	err = p.DB.QueryRowContext(ctx, query, id).Scan(
	// 		&employee.ID,
	// 		&employee.UserName,
	// 		&employee.FirstName,
	// 		&employee.LastName,
	// 		&employee.Address,
	// 		&employee.Email,
	// 		&employee.FacebookID,
	// 		&employee.WhatsappID,
	// 		&employee.TwitterID,
	// 		&employee.LinkedinID,
	// 		&employee.GithubID,
	// 		&employee.Mobile,
	// 		&employee.ImageLink,
	// 		&employee.AccountStatusID,
	// 		&employee.Credits,
	// 		&employee.TaskCompleted,
	// 		&employee.TaskCancelled,
	// 		&employee.Rating,
	// 		&employee.KYCUploadStatus,
	// 		&employee.KYCVerificationStatus,
	// 		&employee.CreatedAt,
	// 		&employee.UpdatedAt,
	// 		&employee.AccountStatus.ID,
	// 		&employee.AccountStatus.Name,
	// 		&employee.AccountStatus.Description,
	// 	)

	return employee, err
}

// GetEmployeeListPaginated returns a chunks of employees
// if accountType == all, it will return list all employee accounts
// if accountType == active, it will return list of active employee accounts
// if accountType == inactive, it will return list of inactive employee accounts
func (p *postgresDBRepo) GetEmployeeListPaginated(accountType string, pageSize, currentPageIndex int) ([]*models.Employee, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	offset := (currentPageIndex - 1) * pageSize
	var employees []*models.Employee

	query := `
		SELECT 
			id, account_code, account_name, contact_person, division, district, upazila, area, mobile, email, account_status, monthly_salary, opening_balance, joining_date, created_at, updated_at
		FROM
			employees
		`

	trails := `
		 ORDER BY
			id asc
		LIMIT $1 OFFSET $2`
	newQuery := `
	SELECT 
		COUNT(id)
	FROM
		employees
	`
	var rows *sql.Rows
	var err error

	if accountType == "all" {
		query = query + trails
	} else if accountType == "active" {
		query += ` WHERE account_status = '1'` + trails
		newQuery += ` WHERE account_status = '1'`
	} else if accountType == "inactive" {
		query += ` WHERE account_status = '0'` + trails
		newQuery += ` WHERE account_status = '0'`
	} else {
		return employees, 0, errors.New("please enter correct parameter to get employees list")
	}

	rows, err = p.DB.QueryContext(ctx, query, pageSize, offset)
	if err != nil {
		return employees, 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var e models.Employee
		err = rows.Scan(
			&e.ID,
			&e.AccountCode,
			&e.AccountName,
			&e.ContactPerson,
			&e.Division,
			&e.District,
			&e.Upazila,
			&e.Area,
			&e.Mobile,
			&e.Email,
			&e.AccountStatus,
			&e.MonthlySalary,
			&e.OpeningBalance,
			&e.JoiningDate,
			&e.CreatedAt,
			&e.UpdatedAt,
		)
		if err != nil {
			return employees, 0, err
		}
		employees = append(employees, &e)
	}

	var totalRecords int
	countRow := p.DB.QueryRowContext(ctx, newQuery)
	err = countRow.Scan(&totalRecords)
	if err != nil {
		return employees, 0, err
	}
	return employees, totalRecords, nil
}

// .......................MIS.......................
// AddCustomer inserts new customer information to the database
func (p *postgresDBRepo) AddCustomer(customer models.Customer) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var id int

	stmt := `INSERT INTO public.customers (account_code,account_name,contact_person,division,district,upazila,area,mobile,email,account_status,discount,opening_balance,joining_date,created_at,updated_at) 
				VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15) RETURNING id
	`
	err := p.DB.QueryRowContext(ctx, stmt,
		customer.AccountCode,
		customer.AccountName,
		customer.ContactPerson,
		customer.Division,
		customer.District,
		customer.Upazila,
		customer.Area,
		customer.Mobile,
		customer.Email,
		customer.AccountStatus,
		customer.Discount,
		customer.OpeningBalance,
		time.Now(),
		time.Now(),
		time.Now(),
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

// GetActiveCustomersIDAndName returns a slice of customers name with id
func (p *postgresDBRepo) GetActiveCustomersIDAndName() ([]*models.Customer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var customers []*models.Customer

	query := `
		SELECT 
			id, account_code, account_name
		FROM
			public.customers
		WHERE 
			account_status = '1'
		`
	var rows *sql.Rows
	var err error

	rows, err = p.DB.QueryContext(ctx, query)
	if err != nil {
		return customers, err
	}
	defer rows.Close()

	for rows.Next() {
		var c models.Customer
		err = rows.Scan(
			&c.ID,
			&c.AccountCode,
			&c.AccountName,
		)
		if err != nil {
			return customers, err
		}
		customers = append(customers, &c)
	}
	return customers, nil
}

// GetCustomerPaginated returns a chunks of customers
// if accountType == all, it will return list all customer accounts
// if accountType == active, it will return list of active customer accounts
// if accountType == inactive, it will return list of inactive customer accounts
func (p *postgresDBRepo) GetCustomerListPaginated(accountType string, pageSize, currentPageIndex int) ([]*models.Customer, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	offset := (currentPageIndex - 1) * pageSize
	var customers []*models.Customer

	query := `
		SELECT 
			id, account_code, account_name, contact_person, division, district, upazila, area, mobile, email, account_status, discount, opening_balance, joining_date, created_at, updated_at
		FROM
			customers
		`

	trails := `
		 ORDER BY
			id asc
		LIMIT $1 OFFSET $2`
	newQuery := `
	SELECT 
		COUNT(id)
	FROM
		customers
	`
	var rows *sql.Rows
	var err error

	if accountType == "all" {
		query = query + trails
	} else if accountType == "active" {
		query += ` WHERE account_status = '1'` + trails
		newQuery += ` WHERE account_status = '1'`
	} else if accountType == "inactive" {
		query += ` WHERE account_status = '0'` + trails
		newQuery += ` WHERE account_status = '0'`
	} else {
		return customers, 0, errors.New("please enter correct parameter to get employees list")
	}

	rows, err = p.DB.QueryContext(ctx, query, pageSize, offset)
	if err != nil {
		return customers, 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var c models.Customer
		err = rows.Scan(
			&c.ID,
			&c.AccountCode,
			&c.AccountName,
			&c.ContactPerson,
			&c.Division,
			&c.District,
			&c.Upazila,
			&c.Area,
			&c.Mobile,
			&c.Email,
			&c.AccountStatus,
			&c.Discount,
			&c.OpeningBalance,
			&c.JoiningDate,
			&c.CreatedAt,
			&c.UpdatedAt,
		)
		if err != nil {
			return customers, 0, err
		}
		customers = append(customers, &c)
	}

	var totalRecords int
	countRow := p.DB.QueryRowContext(ctx, newQuery)
	err = countRow.Scan(&totalRecords)
	if err != nil {
		return customers, 0, err
	}
	return customers, totalRecords, nil
}

// AddSupplier inserts new supplier information to the database
func (p *postgresDBRepo) AddSupplier(supplier models.Supplier) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var id int

	stmt := `INSERT INTO public.suppliers (account_code,account_name,contact_person,division,district,upazila,area,mobile,email,account_status,discount,joining_date,created_at,updated_at) 
				VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14) RETURNING id
	`
	err := p.DB.QueryRowContext(ctx, stmt,
		supplier.AccountCode,
		supplier.AccountName,
		supplier.ContactPerson,
		supplier.Division,
		supplier.District,
		supplier.Upazila,
		supplier.Area,
		supplier.Mobile,
		supplier.Email,
		supplier.AccountStatus,
		supplier.Discount,
		time.Now(),
		time.Now(),
		time.Now(),
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

// GetSuppliersIDAndName returns a slice of suppliers name with id
func (p *postgresDBRepo) GetActiveSuppliersIDAndName() ([]*models.Supplier, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var suppliers []*models.Supplier

	query := `
		SELECT 
			id, account_code, account_name
		FROM
			suppliers
		WHERE 
			account_status = '1'
		`
	var rows *sql.Rows
	var err error

	rows, err = p.DB.QueryContext(ctx, query)
	if err != nil {
		return suppliers, err
	}
	defer rows.Close()

	for rows.Next() {
		var s models.Supplier
		err = rows.Scan(
			&s.ID,
			&s.AccountCode,
			&s.AccountName,
		)
		if err != nil {
			return suppliers, err
		}
		suppliers = append(suppliers, &s)
	}
	return suppliers, nil
}

// GetCustomerPaginated returns a chunks of customers
// if accountType == all, it will return list all customer accounts
// if accountType == active, it will return list of active customer accounts
// if accountType == inactive, it will return list of inactive customer accounts
func (p *postgresDBRepo) GetSupplierListPaginated(accountType string, pageSize, currentPageIndex int) ([]*models.Supplier, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	offset := (currentPageIndex - 1) * pageSize
	var suppliers []*models.Supplier

	query := `
		SELECT 
			id, account_code, account_name, contact_person, division, district, upazila, area, mobile, email, account_status, discount, joining_date, created_at, updated_at
		FROM
			suppliers
		`

	trails := `
		 ORDER BY
			id asc
		LIMIT $1 OFFSET $2`
	newQuery := `
	SELECT 
		COUNT(id)
	FROM
		suppliers
	`
	var rows *sql.Rows
	var err error

	if accountType == "all" {
		query = query + trails
	} else if accountType == "active" {
		query += ` WHERE account_status = '1'` + trails
		newQuery += ` WHERE account_status = '1'`
	} else if accountType == "inactive" {
		query += ` WHERE account_status = '0'` + trails
		newQuery += ` WHERE account_status = '0'`
	} else {
		return suppliers, 0, errors.New("please enter correct parameter to get employees list")
	}

	rows, err = p.DB.QueryContext(ctx, query, pageSize, offset)
	if err != nil {
		return suppliers, 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var s models.Supplier
		err = rows.Scan(
			&s.ID,
			&s.AccountCode,
			&s.AccountName,
			&s.ContactPerson,
			&s.Division,
			&s.District,
			&s.Upazila,
			&s.Area,
			&s.Mobile,
			&s.Email,
			&s.AccountStatus,
			&s.Discount,
			&s.JoiningDate,
			&s.CreatedAt,
			&s.UpdatedAt,
		)
		if err != nil {
			return suppliers, 0, err
		}
		suppliers = append(suppliers, &s)
	}

	var totalRecords int
	countRow := p.DB.QueryRowContext(ctx, newQuery)
	err = countRow.Scan(&totalRecords)
	if err != nil {
		return suppliers, 0, err
	}
	return suppliers, totalRecords, nil
}

//.......................Inventory.......................

// AddBrand inserts new brand information to the database
func (p *postgresDBRepo) AddBrand(b models.Brand) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var id int

	stmt := `INSERT INTO public.brands (name,created_at,updated_at) 
				VALUES($1, $2, $3, $4) RETURNING id
	`
	err := p.DB.QueryRowContext(ctx, stmt,
		b.Name,
		time.Now(),
		time.Now(),
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

// GetBrandList returns a list of all brands from the database
func (p *postgresDBRepo) GetBrandList() ([]*models.Brand, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var brands []*models.Brand

	query := `
		SELECT 
			id, name, created_at, updated_at
		FROM
			public.brands
		`
	var rows *sql.Rows
	var err error

	rows, err = p.DB.QueryContext(ctx, query)
	if err != nil {
		return brands, err
	}
	defer rows.Close()

	for rows.Next() {
		var b models.Brand
		err = rows.Scan(
			&b.ID,
			&b.Name,
			&b.CreatedAt,
			&b.UpdatedAt,
		)
		if err != nil {
			return brands, err
		}
		brands = append(brands, &b)
	}
	return brands, nil
}

// GetActiveBrands returns a list of active brands from the database
func (p *postgresDBRepo) GetActiveBrands() ([]*models.Brand, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var brands []*models.Brand

	query := `
		SELECT 
			id, name, created_at, updated_at
		FROM
			public.brands
		WHERE 
			status = '1'
		`
	var rows *sql.Rows
	var err error

	rows, err = p.DB.QueryContext(ctx, query)
	if err != nil {
		return brands, err
	}
	defer rows.Close()

	for rows.Next() {
		var b models.Brand
		err = rows.Scan(
			&b.ID,
			&b.Name,
			&b.CreatedAt,
			&b.UpdatedAt,
		)
		if err != nil {
			return brands, err
		}
		brands = append(brands, &b)
	}
	return brands, nil
}

// AddCategory inserts new product category to the database
func (p *postgresDBRepo) AddCategory(c models.Category) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var id int

	stmt := `INSERT INTO public.categories (name,created_at,updated_at) 
				VALUES($1, $2, $3) RETURNING id
	`
	err := p.DB.QueryRowContext(ctx, stmt,
		c.Name,
		time.Now(),
		time.Now(),
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

// GetCategoryList returns a list of all categories from the database
func (p *postgresDBRepo) GetCategoryList() ([]*models.Category, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var categories []*models.Category

	query := `
		SELECT 
			id, name, created_at, updated_at
		FROM
			categories
		`
	var rows *sql.Rows
	var err error

	rows, err = p.DB.QueryContext(ctx, query)
	if err != nil {
		return categories, err
	}
	defer rows.Close()

	for rows.Next() {
		var c models.Category
		err = rows.Scan(
			&c.ID,
			&c.Name,
			&c.CreatedAt,
			&c.UpdatedAt,
		)
		if err != nil {
			return categories, err
		}
		categories = append(categories, &c)
	}
	return categories, nil
}

// GetAvailableCategories returns a list of active categories from the database
func (p *postgresDBRepo) GetActiveCategoryList() ([]*models.Category, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var categories []*models.Category

	query := `
		SELECT 
			id, name, created_at, updated_at
		FROM
			categories
		WHERE 
			status = '1'
		`
	var rows *sql.Rows
	var err error

	rows, err = p.DB.QueryContext(ctx, query)
	if err != nil {
		return categories, err
	}
	defer rows.Close()

	for rows.Next() {
		var c models.Category
		err = rows.Scan(
			&c.ID,
			&c.Name,
			&c.CreatedAt,
			&c.UpdatedAt,
		)
		if err != nil {
			return categories, err
		}
		categories = append(categories, &c)
	}
	return categories, nil
}

// AddCategory inserts new product category to the database
func (p *postgresDBRepo) AddProduct(i models.Product) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var id int

	stmt := `INSERT INTO public.products (product_code, product_name, product_description, product_status, quantity, category_id, brand_id, discount, created_at, updated_at) 
				VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id
	`
	err := p.DB.QueryRowContext(ctx, stmt,
		i.ProductCode,
		i.ProductName,
		i.Description,
		i.ProductStatus,
		i.Quantity,
		i.CategoryID,
		i.BrandID,
		i.Discount,
		time.Now(),
		time.Now(),
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

// GetProductList returns a list of all products from the database
func (p *postgresDBRepo) GetProductList() ([]*models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var products []*models.Product

	query := `
		SELECT 
			i.id, i.product_code, i.product_name, i.product_description, i.product_status, i.quantity, i.category_id, i.brand_id, i.discount, i.created_at, i.updated_at, b.id, b.name, c.id, c.name
		FROM 
			public.products i
			INNER JOIN brands b ON (b.id = i.brand_id) 
			INNER JOIN categories c ON (c.id = i.category_id); 
		`
	var rows *sql.Rows
	var err error

	rows, err = p.DB.QueryContext(ctx, query)
	if err != nil {
		return products, err
	}
	defer rows.Close()

	for rows.Next() {
		var i models.Product
		err = rows.Scan(
			&i.ID,
			&i.ProductCode,
			&i.ProductName,
			&i.Description,
			&i.ProductStatus,
			&i.Quantity,
			&i.CategoryID,
			&i.BrandID,
			&i.Discount,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Brand.ID,
			&i.Brand.Name,
			&i.Category.ID,
			&i.Category.Name,
		)
		if err != nil {
			return products, err
		}
		products = append(products, &i)
	}
	return products, nil
}

// GetAvailableProducts returns a list of in-stock and active product from the database
func (p *postgresDBRepo) GetActiveProducts() ([]*models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var products []*models.Product

	query := `
		SELECT 
			i.id, i.product_code, i.product_name, i.quantity, i.category_id, i.discount, b.id, b.name, c.id, c.name
		FROM 
			public.products i
			INNER JOIN brands b ON (b.id = i.brand_id) 
			INNER JOIN categories c ON (c.id = i.category_id)
		WHERE 
			product_status = true; 
		`
	var rows *sql.Rows
	var err error

	rows, err = p.DB.QueryContext(ctx, query)
	if err != nil {
		return products, err
	}
	defer rows.Close()

	for rows.Next() {
		var i models.Product
		err = rows.Scan(
			&i.ID,
			&i.ProductCode,
			&i.ProductName,
			&i.Quantity,
			&i.CategoryID,
			&i.Discount,
			&i.Brand.ID,
			&i.Brand.Name,
			&i.Category.ID,
			&i.Category.Name,
		)
		if err != nil {
			return products, err
		}
		products = append(products, &i)
	}
	return products, nil
}

// GetProductByID returns product info from the database
func (p *postgresDBRepo) GetProductByID(id int) (models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var product models.Product

	query := `
		SELECT 
			i.id, i.product_code, i.product_name, i.quantity, i.category_id, i.discount, b.id, b.name, c.id, c.name
		FROM 
			public.products i
			INNER JOIN brands b ON (b.id = i.brand_id) 
			INNER JOIN categories c ON (c.id = i.category_id)
		WHERE 
			i.id = $1; 
		`
	err := p.DB.QueryRowContext(ctx, query, id).Scan(
		&product.ID,
		&product.ProductCode,
		&product.ProductName,
		&product.Quantity,
		&product.CategoryID,
		&product.Discount,
		&product.Brand.ID,
		&product.Brand.Name,
		&product.Category.ID,
		&product.Category.Name,
	)

	return product, err
}

// GetAvailableProducts returns a list of details info in-stock and active product from the database
func (p *postgresDBRepo) GetAvailableProductsDetails() ([]*models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var products []*models.Product

	query := `
		SELECT 
			i.id, i.product_code, i.product_name, i.product_description, i.product_status, i.quantity, i.category_id, i.brand_id, i.discount, i.created_at, i.updated_at, b.id, b.name, c.id, c.name
		FROM 
			public.products i
			INNER JOIN brands b ON (b.id = i.brand_id) 
			INNER JOIN categories c ON (c.id = i.category_id)
		WHERE 
			product_status = true AND quantity > 0; 
		`
	var rows *sql.Rows
	var err error

	rows, err = p.DB.QueryContext(ctx, query)
	if err != nil {
		return products, err
	}
	defer rows.Close()

	for rows.Next() {
		var i models.Product
		err = rows.Scan(
			&i.ID,
			&i.ProductCode,
			&i.ProductName,
			&i.Description,
			&i.ProductStatus,
			&i.Quantity,
			&i.CategoryID,
			&i.BrandID,
			&i.Discount,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Brand.ID,
			&i.Brand.Name,
			&i.Category.ID,
			&i.Category.Name,
		)
		if err != nil {
			return products, err
		}
		products = append(products, &i)
	}
	return products, nil
}

// GetAvailableProductsByCategoryID returns a list of in-stock and active product that related to category ID from the database
func (p *postgresDBRepo) GetAvailableProductsByCategoryID(cat_id int) ([]*models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var products []*models.Product

	query := `
		SELECT 
			i.id, i.product_code, i.product_name, i.product_description, i.product_status, i.quantity, i.category_id, i.brand_id, i.discount, i.created_at, i.updated_at, b.id, b.name, c.id, c.name
		FROM 
			public.products i
			INNER JOIN brands b ON (b.id = i.brand_id) 
			INNER JOIN categories c ON (c.id = i.category_id)
		WHERE 
			product_status = true AND quantity > 0 AND category_id = $1; 
		`
	var rows *sql.Rows
	var err error

	rows, err = p.DB.QueryContext(ctx, query, cat_id)
	if err != nil {
		return products, err
	}
	defer rows.Close()

	for rows.Next() {
		var i models.Product
		err = rows.Scan(
			&i.ID,
			&i.ProductCode,
			&i.ProductName,
			&i.Description,
			&i.ProductStatus,
			&i.Quantity,
			&i.CategoryID,
			&i.BrandID,
			&i.Discount,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Brand.ID,
			&i.Brand.Name,
			&i.Category.ID,
			&i.Category.Name,
		)
		if err != nil {
			return products, err
		}
		products = append(products, &i)
	}
	return products, nil
}

// GetProductItemsListByProductID returns a list of product items corresponds to productID
func (p *postgresDBRepo) GetProductItemsListByProductID(productID int) (*models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var product *models.Product
	var metadata []*models.ProductMetadata

	//get metadata from product_serial_numbers
	query := `
			SELECT
				id, serial_number, product_id, purchase_history_id, status, warranty, max_retail_price, purchase_rate, created_at, updated_at
			FROM
				public.product_serial_numbers
			WHERE
				status = 'in stock' AND product_id = $1
		`
	rows, err := p.DB.QueryContext(ctx, query, productID)
	if err != nil {
		return product, err
	}

	for rows.Next() {
		var pm models.ProductMetadata
		err = rows.Scan(
			&pm.ID,
			&pm.SerialNumber,
			&pm.ProductID,
			&pm.PurchaseHistoryID,
			&pm.Status,
			&pm.Warranty,
			&pm.MaxRetailPrice,
			&pm.PurchaseRate,
			&pm.CreatedAt,
			&pm.UpdatedAt,
		)
		if err != nil {
			return product, err
		}
		metadata = append(metadata, &pm)
	}

	log.Println(metadata)
	//get product info
	pr, err := p.GetProductByID(productID)
	if err != nil {
		return product, err
	}
	pr.ProductMetadata = metadata
	product = &pr
	return product, nil
}

// GetProductItemDetailsBySerialNumber returns product item details corresponds to serial number
func (p *postgresDBRepo) GetProductItemDetailsBySerialNumber(serialNumber string) (*models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var product *models.Product

	//get metadata from product_serial_numbers
	query := `
		SELECT
			id, serial_number, product_id, purchase_history_id, status, warranty, max_retail_price, purchase_rate, created_at, updated_at
		FROM
			public.product_serial_numbers
		WHERE
			status = 'in stock' AND serial_number = $1
	`
	var metadata models.ProductMetadata
	err := p.DB.QueryRowContext(ctx, query, serialNumber).Scan(
		&metadata.ID,
		&metadata.SerialNumber,
		&metadata.ProductID,
		&metadata.PurchaseHistoryID,
		&metadata.Status,
		&metadata.Warranty,
		&metadata.MaxRetailPrice,
		&metadata.PurchaseRate,
		&metadata.CreatedAt,
		&metadata.UpdatedAt,
	)
	if err != nil {
		return product, err
	}

	//get product info
	pr, err := p.GetProductByID(metadata.ProductID)
	if err != nil {
		return product, err
	}
	pr.ProductMetadata = append(pr.ProductMetadata, &metadata)
	product = &pr
	return product, nil
}

// GetPurchaseHistoryByMemoNo returns purchase history associated with memo_no from database
func (p *postgresDBRepo) GetPurchaseHistoryByMemoNo(memo_no string) ([]*models.Purchase, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// var products []*models.Product
	var PurchaseHistory []*models.Purchase

	//Get product ids for this memo with associated purchase_id for the given memo from purchase_history table
	query := `
		SELECT
			ph.id, ph.purchase_date, ph.supplier_id, ph.product_id, ph.account_id, ph.chalan_no, ph.memo_no, ph.note, ph.quantity_purchased, ph.quantity_sold, ph.bill_amount, ph.discount, ph.total_amount, ph.paid_amount, ph.created_at, ph.updated_at
		FROM
			public.purchase_history ph
		WHERE
			ph.quantity_purchased>ph.quantity_sold AND ph.memo_no = $1
	`
	rows, err := p.DB.QueryContext(ctx, query, memo_no)
	if err != nil {
		return PurchaseHistory, err
	}
	defer rows.Close()

	for rows.Next() {
		var purchase models.Purchase
		err = rows.Scan(
			&purchase.ID,
			&purchase.PurchaseDate,
			&purchase.SupplierID,
			&purchase.ProductID,
			&purchase.AccountID,
			&purchase.ChalanNO,
			&purchase.MemoNo,
			&purchase.Note,
			&purchase.QuantityPurchased,
			&purchase.QuantitySold,
			&purchase.BillAmount,
			&purchase.Discount,
			&purchase.TotalAmount,
			&purchase.PaidAmount,
			&purchase.CreatedAt,
			&purchase.UpdatedAt,
		)
		if err != nil {
			return PurchaseHistory, err
		}
		PurchaseHistory = append(PurchaseHistory, &purchase)
	}
	//get detailed Product info for these ids

	//retrive all product-serial of each product_id && purchase_is
	return PurchaseHistory, nil
}

// GetProductListByPurchaseIDAndProductID returns products list associated with purchaseID and ProductID
func (p *postgresDBRepo) GetProductListByPurchaseIDAndProductID(purchaseID, productID int) (*models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var product *models.Product
	var metadata []*models.ProductMetadata

	//get metadata from product_serial_numbers
	query := `
			SELECT
				id, serial_number, product_id, purchase_history_id, status, warranty, max_retail_price, purchase_rate, created_at, updated_at
			FROM
				public.product_serial_numbers
			WHERE
				status = 'in stock' AND purchase_history_id = $1 AND product_id = $2
		`
	rows, err := p.DB.QueryContext(ctx, query, purchaseID, productID)
	if err != nil {
		return product, err
	}

	for rows.Next() {
		var pm models.ProductMetadata
		err = rows.Scan(
			&pm.ID,
			&pm.SerialNumber,
			&pm.ProductID,
			&pm.PurchaseHistoryID,
			&pm.Status,
			&pm.Warranty,
			&pm.MaxRetailPrice,
			&pm.PurchaseRate,
			&pm.CreatedAt,
			&pm.UpdatedAt,
		)
		if err != nil {
			return product, err
		}
		metadata = append(metadata, &pm)
	}

	//get product info
	pr, err := p.GetProductByID(productID)
	if err != nil {
		return product, err
	}
	pr.ProductMetadata = metadata
	product = &pr
	return product, nil
}

// UpdateProductQuantityByProductID increases product quantity, to reduce product quantity pass negetive quantity number
func (p *postgresDBRepo) UpdateProductQuantityByProductID(quantity, productID int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	stmt := `UPDATE public.products
        SET quantity = quantity + $1
        WHERE id = $2;`
	// Execute the query
	_, err := p.DB.ExecContext(ctx, stmt, quantity, productID)
	if err != nil {
		return errors.New("SQLErrorUpdateProductQuantity:" + err.Error())
	}
	return nil
}

// UpdateProductItemStatusByProductSerialNumber updates the status of the product unit in product_serial_numbers Table
func (p *postgresDBRepo) UpdateProductItemStatusByProductUnitsID(productUnitsID, status int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	stmt := `UPDATE public.product_serial_numbers
        SET status = $1
        WHERE id = $2;`
	// Execute the query
	_, err := p.DB.ExecContext(ctx, stmt, status, productUnitsID)
	if err != nil {
		return errors.New("SQLErrorUpdateProductItemStatusByProductSerialNumber:" + err.Error())
	}
	return nil
}

// AddToPurchaseHistory insets purchase history info to the database
func (p *postgresDBRepo) AddToPurchaseHistory(purchase *models.Purchase) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var id int

	stmt := `INSERT INTO public.purchase_history (purchase_date,supplier_id,product_id,account_id,chalan_no,memo_no,note,bill_amount,discount,total_amount,paid_amount,created_at,updated_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) RETURNING id
	`
	err := p.DB.QueryRowContext(ctx, stmt,
		purchase.PurchaseDate,
		purchase.SupplierID,
		purchase.ProductID,
		purchase.AccountID,
		purchase.ChalanNO,
		purchase.MemoNo,
		purchase.Note,
		purchase.BillAmount,
		purchase.Discount,
		purchase.TotalAmount,
		purchase.PaidAmount,
		time.Now(),
		time.Now(),
	).Scan(&id)

	if err != nil {
		return 0, err
	}
	return id, nil
}

// AddProductSerialNumbers insets product serial numbers to the database
func (p *postgresDBRepo) AddProductSerialNumbers(purchase *models.Purchase) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	values := []string{}
	now := time.Now()
	purchase.CreatedAt = now
	purchase.UpdatedAt = now

	// Format timestamps with time zone in a PostgreSQL-friendly format
	createdAt := purchase.CreatedAt.Format("2006-01-02 15:04:05 -07:00")
	updatedAt := purchase.UpdatedAt.Format("2006-01-02 15:04:05 -07:00")

	for _, serial_number := range purchase.ProductsSerialNo {
		values = append(values, fmt.Sprintf("('%s',%d,%d,%d,'%s','%s')", serial_number, purchase.ProductID, purchase.MaxRetailPrice, purchase.PurchaseRate, createdAt, updatedAt))
	}

	query := "INSERT INTO public.product_serial_numbers (serial_number,product_id,max_retail_price,purchase_rate,created_at,updated_at) VALUES " + strings.Join(values, ",") + ";"
	// Execute the query
	_, err := p.DB.ExecContext(ctx, query)
	if err != nil {
		return errors.New("SQLErrorAddProductSerialNumbers:" + err.Error())
	}
	return nil
}

// GetMemoListBySupplierID returns a list of memo with purchase id from the database
func (p *postgresDBRepo) GetMemoListBySupplierID(supplierID int) ([]*models.Purchase, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var purchases []*models.Purchase

	query := `
		SELECT DISTINCT memo_no
		FROM 
			public.purchase_history
		WHERE 
			quantity_purchased > quantity_sold AND supplier_id = $1 ; 
		`

	rows, err := p.DB.QueryContext(ctx, query, supplierID)
	if err != nil {
		return purchases, err
	}
	defer rows.Close()

	for rows.Next() {
		var p models.Purchase
		err = rows.Scan(
			&p.MemoNo,
		)
		if err != nil {
			return purchases, err
		}
		purchases = append(purchases, &p)
	}
	return purchases, nil
}

// RestockProduct update product quantity, store purchase history and product serial numbers
func (p *postgresDBRepo) RestockProduct(purchase *models.Purchase) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	tx, err := p.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	log.Println(purchase.Quantity)

	//Tx-1: Update product quantity
	//Set quantity += newQuantity
	query := `UPDATE public.products
          SET quantity = quantity + $1
          WHERE id = $2;`

	// Execute the query with parameters
	_, err = tx.ExecContext(ctx, query, purchase.Quantity, purchase.ProductID)
	if err != nil {
		tx.Rollback()
		return errors.New("SQLErrorRestockProduct(Update Quantity): " + err.Error())
	}

	//Tx-2: Insert data to purchase history table
	//.........insert the following data into purchase_history table.........
	// PurchaseDate     string
	// SupplierID       int
	// ProductID        int
	// AccountID        int
	// ChalanNO         string
	// MemoNo           string
	// Note             string
	// BillAmount       int
	// Discount         int
	// TotalAmount      int
	// PaidAmount       int
	var purhcase_id int
	query = `INSERT INTO public.purchase_history (purchase_date,supplier_id,product_id,account_id,chalan_no,memo_no,note,quantity_purchased,bill_amount,discount,total_amount,paid_amount,created_at,updated_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14) RETURNING id
	`
	row := tx.QueryRowContext(ctx, query,
		purchase.PurchaseDate,
		purchase.SupplierID,
		purchase.ProductID,
		purchase.AccountID,
		purchase.ChalanNO,
		purchase.MemoNo,
		purchase.Note,
		purchase.Quantity,
		purchase.BillAmount,
		purchase.Discount,
		purchase.TotalAmount,
		purchase.PaidAmount,
		time.Now(),
		time.Now(),
	)
	if err = row.Scan(&purhcase_id); err != nil {
		tx.Rollback()
		return errors.New("SQLErrorRestockProduct(Insert purchase_history):" + err.Error())
	}

	//Tx-3: store product serial numbers
	//.........insert the following data into product_serial_numbers table............//
	//ProductID        int
	// ProductsSerialNo []string
	// MaxRetailPrice        int
	// PurchaseRate              int

	values := []string{}
	now := time.Now()
	purchase.CreatedAt = now
	purchase.UpdatedAt = now

	// Format timestamps with time zone in a PostgreSQL-friendly format
	createdAt := purchase.CreatedAt.Format("2006-01-02 15:04:05 -07:00")
	updatedAt := purchase.UpdatedAt.Format("2006-01-02 15:04:05 -07:00")

	for _, serial_number := range purchase.ProductsSerialNo {
		values = append(values, fmt.Sprintf("('%s',%d,%d,%d,%d,%d,'%s','%s')", serial_number, purchase.ProductID, purhcase_id, purchase.MaxRetailPrice, purchase.PurchaseRate, purchase.Warranty, createdAt, updatedAt))
	}

	query = "INSERT INTO public.product_serial_numbers (serial_number,product_id,purchase_history_id,max_retail_price,purchase_rate,warranty,created_at,updated_at) VALUES " + strings.Join(values, ",") + ";"
	// Execute the query
	_, err = tx.ExecContext(ctx, query)
	if err != nil {
		tx.Rollback()
		return errors.New("SQLErrorRestockProduct(Product Serial Number):" + err.Error())
	}
	// Commit transaction
	if err := tx.Commit(); err != nil {
		return errors.New("SQLErrorRestockProduct(Commit):" + err.Error())
	}
	return nil
}

// RestockProduct update product quantity, store purchase history and product serial numbers
func (p *postgresDBRepo) SaleProducts(sale *models.Sale) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	tx, err := p.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	//Tx-1: Insert data to sales history table70
	var sale_id int
	query := `INSERT INTO public.sales_history (sale_date,customer_id,account_id,chalan_no,memo_no,note,bill_amount,discount,total_amount,paid_amount,created_at,updated_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) RETURNING id
	`
	date, err := time.Parse("01/02/2006", sale.SaleDate)
	if err != nil {
		return err
	}
	row := tx.QueryRowContext(ctx, query,
		date,
		sale.CustomerID,
		sale.AccountID,
		sale.ChalanNO,
		sale.MemoNo,
		sale.Note,
		sale.BillAmount,
		sale.Discount,
		sale.TotalAmount,
		sale.PaidAmount,
		time.Now(),
		time.Now(),
	)
	if err = row.Scan(&sale_id); err != nil {
		tx.Rollback()
		return errors.New("SQLErrorSaleProducts(INSERT sales_history): " + err.Error())
	}
	fmt.Println("sale ID: ", sale_id)

	//Tx-2:
	//step-1: Update product quantity, Set quantity -= newQuantity where id = {affected row id}
	//step-2: update product items status and sales_history_id, returnin id affected row id

	//loop over the SelectedProduct array
	for i, items := range sale.SelectedItems {
		//Update product quantity
		query := `
			UPDATE public.products
			SET quantity = quantity - $1
			WHERE id = $2;
		  `
		// Execute the query with parameters
		_, err = tx.ExecContext(ctx, query, len(items.SerialNumbers), items.ProductID)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("SQLErrorSaleProducts(Update Product Quantity):#%d --%w", i, err)
		}

		//update product items status and sales_history_id
		for _, serialNumber := range items.SerialNumbers {
			query = `
				UPDATE public.product_serial_numbers
				SET status = 'sold', sales_history_id = $1 
				WHERE serial_number = $2
			`
			_, err := tx.ExecContext(ctx, query, sale_id, serialNumber)
			if err != nil {
				tx.Rollback()
				return fmt.Errorf("SQLErrorSaleProducts(Update Product status And status):#serial-%s --%w", serialNumber, err)
			}
		}
		//
	}

	//Tx-3 insert summary about the sales in the inventory_transaction_logs table
	var inv_tx_log_id int
	query = `INSERT INTO public.inventory_transaction_logs(job_id, transaction_type, quantity, price, transaction_date, created_at, updated_at)
	VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING id`
	err = tx.QueryRowContext(ctx, query, "S-"+sale.MemoNo, "sale", len(sale.ProductsSerialNo), sale.TotalAmount, date, time.Now(), time.Now()).Scan(&inv_tx_log_id)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("SQLErrorSaleProducts(INSERT inventory_transaction_logs: %w", err)
	}
	// Commit transaction
	if err := tx.Commit(); err != nil {
		return errors.New("SQLErrorSaleProducts(Commit):" + err.Error())
	}
	return nil
}

// /ReturnProductUnitsToSupplier updates database
func (p *postgresDBRepo) ReturnProductUnitsToSupplier(JobID string, transactionDate time.Time, ProductUnitsID []int, TotalUnits int, TotalPrices int) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var id int
	tx, err := p.DB.BeginTx(ctx, nil)
	if err != nil {
		return id, fmt.Errorf("failed to start transaction: %w", err)
	}

	// Prepare the SQL update query
	query := `UPDATE public.product_serial_numbers SET status = $1 WHERE id = $2`

	// Execute updates within the transaction
	for _, unitsID := range ProductUnitsID {
		_, err := tx.ExecContext(ctx, query, "purchase-returned", unitsID)
		if err != nil {
			tx.Rollback() // Rollback on error
			return id, fmt.Errorf("failed to update record with id %d: %w", unitsID, err)
		}
	}

	query = `INSERT INTO public.inventory_transaction_logs (job_id, transaction_type, quantity, price, transaction_date, created_at, updated_at)
			VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING id
		`

	err = tx.QueryRowContext(ctx, query,
		JobID,
		"purchase-returned",
		TotalUnits,
		TotalPrices,
		transactionDate,
		time.Now(),
		time.Now(),
	).Scan(&id)

	if err != nil {
		tx.Rollback() // Rollback on error
		return id, fmt.Errorf("failed to insert record to inventory_transaction_logs: %w", err)
	}
	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return id, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return id, nil
}

// Helper functions
// CountTotalEntries counts total number of rows in given the table
func (p *postgresDBRepo) CountRows(tableName string) (int, error) {
	ctx, cancle := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancle()

	var c int
	query := "SELECT COUNT(id) FROM " + tableName
	err := p.DB.QueryRowContext(ctx, query).Scan(&c)
	return c, err
}
