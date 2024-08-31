package dbrepo

import (
	"PSInventory/internal/models"
	"context"
	"database/sql"
	"errors"
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
func (p *postgresDBRepo) GetAvailableHeadAccounts()([]*models.HeadAccount, error) {
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
func (p *postgresDBRepo) GetSuppliersIDAndName() ([]*models.Supplier, error) {
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
func (p *postgresDBRepo) GetBrandList()([]*models.Brand, error) {
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
// GetAvailableBrands returns a list of active brands from the database
func (p *postgresDBRepo) GetAvailableBrands()([]*models.Brand, error) {
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
func (p *postgresDBRepo) GetCategoryList()([]*models.Category, error) {
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
func (p *postgresDBRepo) GetAvailableCategories()([]*models.Category, error) {
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
func (p *postgresDBRepo) AddItem(i models.Item) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var id int

	stmt := `INSERT INTO public.items (item_code, item_name, item_description, item_status, quantity, category_id, brand_id, discount, created_at, updated_at) 
				VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id
	`
	err := p.DB.QueryRowContext(ctx, stmt,
		i.ItemCode,
		i.ItemName,
		i.ItemDescription,
		i.ItemStatus,
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
// GetItemList returns a list of all items from the database
func (p *postgresDBRepo) GetItemList()([]*models.Item, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var items []*models.Item

	query := `
		SELECT 
			i.id, i.item_code, i.item_name, i.item_description, i.item_status, i.quantity, i.category_id, i.brand_id, i.discount, i.created_at, i.updated_at, b.id, b.name
		FROM 
			public.items i
			INNER JOIN brands b ON (b.id = i.category_id); 
		`
	var rows *sql.Rows
	var err error

	rows, err = p.DB.QueryContext(ctx, query)
	if err != nil {
		return items, err
	}
	defer rows.Close()

	for rows.Next() {
		var i models.Item
		err = rows.Scan(
			&i.ID,
			&i.ItemCode,
			&i.ItemName,
			&i.ItemDescription,
			&i.ItemStatus,
			&i.Quantity,
			&i.CategoryID,
			&i.BrandID,
			&i.Discount,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Brand.ID,
			&i.Brand.Name,
		)
		if err != nil {
			return items, err
		}
		items = append(items, &i)
	}
	return items, nil
}
// GetAvailableItems returns a list of in-stock and active item from the database
func (p *postgresDBRepo)GetAvailableItems()([]*models.Item, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var items []*models.Item

	query := `
		SELECT 
			i.id, i.item_code, i.item_name, i.quantity, i.category_id, i.discount, b.id, b.name
		FROM 
			public.items i
			INNER JOIN brands b ON (b.id = i.category_id)
		WHERE 
			item_status = true AND quantity > 0; 
		`
	var rows *sql.Rows
	var err error

	rows, err = p.DB.QueryContext(ctx, query)
	if err != nil {
		return items, err
	}
	defer rows.Close()

	for rows.Next() {
		var i models.Item
		err = rows.Scan(
			&i.ID,
			&i.ItemCode,
			&i.ItemName,
			&i.Quantity,
			&i.CategoryID,
			&i.Discount,
			&i.Brand.ID,
			&i.Brand.Name,
		)
		if err != nil {
			return items, err
		}
		items = append(items, &i)
	}
	return items, nil
}
// GetAvailableItems returns a list of details info in-stock and active item from the database
func (p *postgresDBRepo)GetAvailableItemsDetails()([]*models.Item, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var items []*models.Item

	query := `
		SELECT 
			i.id, i.item_code, i.item_name, i.item_description, i.item_status, i.quantity, i.category_id, i.brand_id, i.discount, i.created_at, i.updated_at, b.id, b.name
		FROM 
			public.items i
			INNER JOIN brands b ON (b.id = i.category_id)
		WHERE 
			item_status = true AND quantity > 0; 
		`
	var rows *sql.Rows
	var err error

	rows, err = p.DB.QueryContext(ctx, query)
	if err != nil {
		return items, err
	}
	defer rows.Close()

	for rows.Next() {
		var i models.Item
		err = rows.Scan(
			&i.ID,
			&i.ItemCode,
			&i.ItemName,
			&i.ItemDescription,
			&i.ItemStatus,
			&i.Quantity,
			&i.CategoryID,
			&i.BrandID,
			&i.Discount,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Brand.ID,
			&i.Brand.Name,
		)
		if err != nil {
			return items, err
		}
		items = append(items, &i)
	}
	return items, nil
}
// GetAvailableItemsByCategoryID returns a list of in-stock and active item that related to category ID from the database
func (p *postgresDBRepo)GetAvailableItemsByCategoryID(cat_id int)([]*models.Item, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var items []*models.Item

	query := `
		SELECT 
			i.id, i.item_code, i.item_name, i.item_description, i.item_status, i.quantity, i.category_id, i.brand_id, i.discount, i.created_at, i.updated_at, b.id, b.name
		FROM 
			public.items i
			INNER JOIN brands b ON (b.id = i.category_id)
		WHERE 
			item_status = true AND quantity > 0 AND category_id = $1; 
		`
	var rows *sql.Rows
	var err error

	rows, err = p.DB.QueryContext(ctx, query, cat_id)
	if err != nil {
		return items, err
	}
	defer rows.Close()

	for rows.Next() {
		var i models.Item
		err = rows.Scan(
			&i.ID,
			&i.ItemCode,
			&i.ItemName,
			&i.ItemDescription,
			&i.ItemStatus,
			&i.Quantity,
			&i.CategoryID,
			&i.BrandID,
			&i.Discount,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Brand.ID,
			&i.Brand.Name,
		)
		if err != nil {
			return items, err
		}
		items = append(items, &i)
	}
	return items, nil
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
