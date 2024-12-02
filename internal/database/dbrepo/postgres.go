package dbrepo

import (
	"PSInventory/internal/models"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
	"time"
)

// GetDashBoardData retrieves dashboard data
func (p *postgresDBRepo) GetDashBoardData() (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	//total sale
	//total purchase
	//total employee
	//total customer
	//total supplier
	query := `
		SELECT 
    (SELECT COUNT(*) FROM public.employees) AS total_employee,
    (SELECT COUNT(*) FROM public.suppliers) AS total_customer,
    (SELECT COUNT(*) FROM public.customers) AS total_supplier,
    (SELECT COALESCE(SUM(total_purchases - total_purchase_returns), 0) FROM public.top_sheet) AS total_purchase_value,
    (SELECT COALESCE(SUM(total_sales - total_sale_returns), 0) FROM public.top_sheet) AS total_sale_value,
	(SELECT COALESCE(SUM(total_expenses), 0) FROM public.top_sheet) AS total_expense
`
	var data struct {
		TotalEmployee int `json:"total_employee"`
		TotalCustomer int `json:"total_customer"`
		TotalSupplier int `json:"total_supplier"`
		TotalPurchase int `json:"total_purchase"`
		TotalSale     int `json:"total_sale"`
		TotalExpense  int `json:"total_expense"`
	}
	err := p.DB.QueryRowContext(ctx, query).Scan(&data.TotalEmployee, &data.TotalSupplier, &data.TotalCustomer, &data.TotalPurchase, &data.TotalSale, &data.TotalExpense)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//.......................HR Management.......................

// AddHeadAccount inserts new head account information to the database
func (p *postgresDBRepo) AddHeadAccount(ha models.HeadAccount) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var id int

	stmt := `INSERT INTO public.head_accounts (account_code,account_name,account_status,current_balance,created_at,updated_at) 
				VALUES($1, $2, $3, $4, $5, $6) RETURNING id
	`
	err := p.DB.QueryRowContext(ctx, stmt,
		ha.AccountCode,
		ha.AccountName,
		ha.AccountStatus,
		ha.CurrentBalance,
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
			id, account_code, account_name, account_type, account_status, current_balance, created_at, updated_at
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
			&ha.AccountType,
			&ha.AccountStatus,
			&ha.CurrentBalance,
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

// GetCashBankHead
func (p *postgresDBRepo) GetAvailableHeadAccountsByType(accountType string) ([]*models.HeadAccount, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var headAccounts []*models.HeadAccount

	query := `
		SELECT 
			id, account_code, account_name, account_type, account_status, current_balance, created_at, updated_at
		FROM
			public.head_accounts
		WHERE 
			account_status = true AND account_type = $1
		`
	var rows *sql.Rows
	var err error

	rows, err = p.DB.QueryContext(ctx, query, accountType)
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
			&ha.AccountType,
			&ha.AccountStatus,
			&ha.CurrentBalance,
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

// UpdateHeadAccountBalance updates the current balance of an head account
func (p *postgresDBRepo) UpdateHeadAccountBalance(id, balance int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `
		UPDATE public.head_accounts
		SET current_balance = current_balance + $1
		WHERE id = $2`

	_, err := p.DB.ExecContext(ctx, stmt, balance, id)
	if err != nil {
		return fmt.Errorf("DBERROR: Unable to update current_balance in head_accounts table: %w", err)
	}
	return nil
}

//.......................Administrative Panel.....................

// AddNewStakeHolder inserts new stakeholder information to the database
func (p *postgresDBRepo) AddNewStakeHolder(stk models.StakeHolder) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var id int

	stmt := `
		INSERT INTO public.company_stakeholders (account_type,account_code,account_name,contact_person,division,district,upazila,area,mobile,email,account_status,joining_date,created_at,updated_at) 
		VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14) RETURNING id
	`
	err := p.DB.QueryRowContext(ctx, stmt,
		stk.AccountType,
		stk.AccountCode,
		stk.AccountName,
		stk.ContactPerson,
		stk.Division,
		stk.District,
		stk.Upazila,
		stk.Area,
		stk.Mobile,
		stk.Email,
		stk.AccountStatus,
		time.Now(),
		time.Now(),
		time.Now(),
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
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

// GetEmployeeListPaginated returns a chunks of employees
// if accountType == all, it will return list all employee accounts
// if accountType == active, it will return list of active employee accounts
// if accountType == inactive, it will return list of inactive employee accounts
// GetEmployeeListPaginated retrieves a paginated list of employees based on account type and pagination parameters
func (p *postgresDBRepo) GetEmployeeListPaginated(accountType string, pageSize, currentPageIndex int) ([]*models.Employee, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var employees []*models.Employee

	// Prepare the base query for selecting employees
	query := `
		SELECT 
			id, account_code, account_name, contact_person, division, district, upazila, area, mobile, email, account_status, monthly_salary, opening_balance, joining_date, created_at, updated_at
		FROM
			employees
		`

	// Prepare a separate query for counting total records
	newQuery := `
	SELECT 
		COUNT(id)
	FROM
		employees
	`

	// Initialize the limit and offset
	var limit int
	var offset int

	// Check for account type and adjust query accordingly
	if accountType == "all" {
		if pageSize == -1 {
			// If pageSize is -1, return all employees without limit
			query += ` ORDER BY id ASC`
			limit = 0 // No limit when fetching all records
		} else {
			// Set limit and offset for pagination
			limit = pageSize
			offset = (currentPageIndex - 1) * pageSize
			query += ` ORDER BY id ASC LIMIT $1 OFFSET $2`
		}
	} else if accountType == "active" {
		if pageSize == -1 {
			query += ` WHERE account_status = '1' ORDER BY id ASC`
			limit = 0 // No limit when fetching all records
			newQuery += ` WHERE account_status = '1'`
		} else {
			limit = pageSize
			offset = (currentPageIndex - 1) * pageSize
			query += ` WHERE account_status = '1' ORDER BY id ASC LIMIT $1 OFFSET $2`
			newQuery += ` WHERE account_status = '1'`
		}
	} else if accountType == "inactive" {
		if pageSize == -1 {
			query += ` WHERE account_status = '0' ORDER BY id ASC`
			limit = 0 // No limit when fetching all records
			newQuery += ` WHERE account_status = '0'`
		} else {
			limit = pageSize
			offset = (currentPageIndex - 1) * pageSize
			query += ` WHERE account_status = '0' ORDER BY id ASC LIMIT $1 OFFSET $2`
			newQuery += ` WHERE account_status = '0'`
		}
	} else {
		return employees, 0, errors.New("please enter correct parameter to get employees list")
	}

	// Execute the employee query
	var rows *sql.Rows
	var err error
	if limit > 0 {
		rows, err = p.DB.QueryContext(ctx, query, limit, offset)
	} else {
		rows, err = p.DB.QueryContext(ctx, query)
	}

	if err != nil {
		return employees, 0, err
	}
	defer rows.Close()

	// Scan the rows into employee struct
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

	// Get total records count
	var totalRecords int
	countRow := p.DB.QueryRowContext(ctx, newQuery)
	err = countRow.Scan(&totalRecords)
	if err != nil {
		return employees, 0, err
	}

	return employees, totalRecords, nil
}

// GetAllEmployeesList returns a list of all employees
func (p *postgresDBRepo) GetAllEmployeesList() ([]*models.Employee, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var employees []*models.Employee

	// Prepare the base query for selecting employees
	query := `
		SELECT 
			id, account_code, account_name, contact_person, division, district, upazila, area, mobile, email, account_status, monthly_salary, opening_balance, joining_date, created_at, updated_at
		FROM
			employees
		`
	rows, err := p.DB.QueryContext(ctx, query)

	if err != nil {
		return employees, err
	}
	defer rows.Close()

	// Scan the rows into employee struct
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
			return employees, err
		}
		employees = append(employees, &e)
	}

	return employees, nil
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

// GetCustomerByID search customer information by id from customers table
func (p *postgresDBRepo) GetCustomerByID(id int) (models.Customer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
		SELECT 
			id, account_code, account_name, contact_person, division, district, upazila, area, due_amount, mobile, whatsapp_account, email, account_status, discount, opening_balance, joining_date, created_at, updated_at
		FROM
			public.customers
		WHERE
			id=$1
	`
	var customer models.Customer
	err := p.DB.QueryRowContext(ctx, query, id).Scan(
		&customer.ID,
		&customer.AccountCode,
		&customer.AccountName,
		&customer.ContactPerson,
		&customer.Division,
		&customer.District,
		&customer.Upazila,
		&customer.Area,
		&customer.DueAmount,
		&customer.Mobile,
		&customer.WhatsappAccount,
		&customer.Email,
		&customer.AccountStatus,
		&customer.Discount,
		&customer.OpeningBalance,
		&customer.JoiningDate,
		&customer.CreatedAt,
		&customer.UpdatedAt,
	)

	if err != nil {
		return customer, fmt.Errorf("DBERROR--GetCustomerByID: %w", err)
	}
	return customer, nil
}

// GetActiveCustomersIDAndName returns a slice of customers name with id
func (p *postgresDBRepo) GetActiveCustomersIDAndName() ([]*models.Customer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var customers []*models.Customer

	query := `
		SELECT 
			id, account_code, account_name, division, district, upazila, area, mobile 
		FROM
			public.customers
		WHERE 
			account_status = '1'
		ORDER BY updated_at ASC
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
			&c.Division,
			&c.District,
			&c.Upazila,
			&c.Area,
			&c.Mobile,
		)
		if err != nil {
			return customers, err
		}
		customers = append(customers, &c)
	}
	return customers, nil
}

// GetCreditCustomersDetails returns a slice of customers details who have due amount(due_amount > 0) from the customers table
func (p *postgresDBRepo) GetCreditCustomersDetails() ([]*models.Customer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var customers []*models.Customer

	query := `
		SELECT 
			id, account_code, account_name, contact_person, due_amount, mobile
		FROM
			public.customers
		WHERE 
			due_amount > 0
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
			&c.ContactPerson,
			&c.DueAmount,
			&c.Mobile,
		)
		if err != nil {
			return customers, err
		}
		customers = append(customers, &c)
	}
	return customers, nil
}

// GetDebitCustomersDetails returns a slice of customers details who have amount receivable(due_amount > 0) from the customers table
func (p *postgresDBRepo) GetDebitCustomersDetails() ([]*models.Customer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var customers []*models.Customer

	query := `
		SELECT 
			id, account_code, account_name, contact_person, due_amount, mobile
		FROM
			public.customers
		WHERE 
			due_amount > 0
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
			&c.ContactPerson,
			&c.DueAmount,
			&c.Mobile,
		)
		if err != nil {
			return customers, err
		}
		customers = append(customers, &c)
	}
	return customers, nil
}

// GetCreditCustomersDetails returns a slice of suppliers details who have due amount(due_amount > 0) from the supplier table
func (p *postgresDBRepo) GetCreditSuppliersDetails() ([]*models.Supplier, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var suppliers []*models.Supplier

	query := `
		SELECT 
			id, account_code, account_name, contact_person, due_amount, mobile
		FROM
			public.suppliers
		WHERE 
			due_amount < 0
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
			&s.ContactPerson,
			&s.DueAmount,
			&s.Mobile,
		)
		if err != nil {
			return suppliers, err
		}
		suppliers = append(suppliers, &s)
	}
	return suppliers, nil
}

// GetDebitSuppliersDetails returns a slice of suppliers details who have amount receivable(due_amount > 0) from the suppliers table
func (p *postgresDBRepo) GetDebitSuppliersDetails() ([]*models.Supplier, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var suppliers []*models.Supplier

	query := `
		SELECT 
			id, account_code, account_name, contact_person, due_amount, mobile
		FROM
			public.suppliers
		WHERE 
			due_amount > 0
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
			&s.ContactPerson,
			&s.DueAmount,
			&s.Mobile,
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
func (p *postgresDBRepo) GetAllCustomersList() ([]*models.Customer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var customers []*models.Customer

	query := `
		SELECT 
			id, account_code, account_name, contact_person, division, district, upazila, area, mobile, email, account_status, discount, due_amount, opening_balance, joining_date, created_at, updated_at
		FROM
			customers
		ORDER BY
			id asc`
	rows, err := p.DB.QueryContext(ctx, query)
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
			&c.ContactPerson,
			&c.Division,
			&c.District,
			&c.Upazila,
			&c.Area,
			&c.Mobile,
			&c.Email,
			&c.AccountStatus,
			&c.Discount,
			&c.DueAmount,
			&c.OpeningBalance,
			&c.JoiningDate,
			&c.CreatedAt,
			&c.UpdatedAt,
		)
		if err != nil {
			return customers, err
		}
		customers = append(customers, &c)
	}
	return customers, nil
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
		ORDER BY updated_at ASC
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
func (p *postgresDBRepo) GetAllSuppliersList() ([]*models.Supplier, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var suppliers []*models.Supplier

	query := `
		SELECT 
			id, account_code, account_name, contact_person, division, district, upazila, area, mobile, email, account_status, discount, due_amount, joining_date, created_at, updated_at
		FROM
			suppliers
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
			&s.ContactPerson,
			&s.Division,
			&s.District,
			&s.Upazila,
			&s.Area,
			&s.Mobile,
			&s.Email,
			&s.AccountStatus,
			&s.Discount,
			&s.DueAmount,
			&s.JoiningDate,
			&s.CreatedAt,
			&s.UpdatedAt,
		)
		if err != nil {
			return suppliers, err
		}
		suppliers = append(suppliers, &s)
	}
	return suppliers, nil
}

//.......................Inventory.......................

// AddBrand inserts new brand information to the database
func (p *postgresDBRepo) AddBrand(b models.Brand) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var id int

	stmt := `INSERT INTO public.brands (name) VALUES($1) RETURNING id`
	err := p.DB.QueryRowContext(ctx, stmt, b.Name).Scan(&id)

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
		SELECT * FROM categories ORDER BY id ASC
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
			&c.Status,
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

	stmt := `INSERT INTO public.products (product_code, product_name, product_description, product_status, quantity_purchased, quantity_sold, category_id, brand_id, stock_alert_level, created_at, updated_at) 
				VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id
	`
	err := p.DB.QueryRowContext(ctx, stmt,
		i.ProductCode,
		i.ProductName,
		i.Description,
		i.ProductStatus,
		i.QuantityPurchased,
		i.QuantitySold,
		i.CategoryID,
		i.BrandID,
		i.StockAlertLevel,
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
			i.id, i.product_code, i.product_name, i.product_description, i.product_status, i.quantity_purchased, i.quantity_sold, i.category_id, i.brand_id, i.stock_alert_level, i.created_at, i.updated_at, b.id, b.name, c.id, c.name
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
			&i.QuantityPurchased,
			&i.QuantitySold,
			&i.CategoryID,
			&i.BrandID,
			&i.StockAlertLevel,
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
			i.id, i.product_code, i.product_name, i.quantity_purchased, i.quantity_sold, i.category_id, i.stock_alert_level, b.id, b.name, c.id, c.name
		FROM 
			public.products i
			INNER JOIN brands b ON (b.id = i.brand_id) 
			INNER JOIN categories c ON (c.id = i.category_id)
		WHERE 
			i.product_status = true; 
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
			&i.QuantityPurchased,
			&i.QuantitySold,
			&i.CategoryID,
			&i.StockAlertLevel,
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
			i.id, i.product_code, i.product_name, i.quantity_purchased, i.quantity_sold, i.category_id, i.brand_id, i.stock_alert_level, b.id, b.name, c.id, c.name
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
		&product.QuantityPurchased,
		&product.QuantitySold,
		&product.CategoryID,
		&product.BrandID,
		&product.StockAlertLevel,
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
			i.id, i.product_code, i.product_name, i.product_description, i.product_status, i.quantity_purchased, i.quantity_sold, i.category_id, i.brand_id, i.stock_alert_level, i.created_at, i.updated_at, b.id, b.name, c.id, c.name
		FROM 
			public.products i
			INNER JOIN brands b ON (b.id = i.brand_id) 
			INNER JOIN categories c ON (c.id = i.category_id)
		WHERE 
			i.product_status = true AND i.quantity_purchased - i.quantity_sold > 0; 
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
			&i.QuantityPurchased,
			&i.QuantitySold,
			&i.CategoryID,
			&i.BrandID,
			&i.StockAlertLevel,
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
			i.id, i.product_code, i.product_name, i.product_description, i.product_status, i.quantity_purchased, i.quantity_sold, i.category_id, i.brand_id, i.stock_alert_level, i.created_at, i.updated_at, b.id, b.name, c.id, c.name
		FROM 
			public.products i
			INNER JOIN brands b ON (b.id = i.brand_id) 
			INNER JOIN categories c ON (c.id = i.category_id)
		WHERE 
			i.product_status = true AND i.quantity_purchased - i.quantity_sold, > 0 AND i.category_id = $1; 
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
			&i.QuantityPurchased,
			&i.QuantitySold,
			&i.CategoryID,
			&i.BrandID,
			&i.StockAlertLevel,
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

// GetAllProductsByCategoryID returns a list of all product that related to category ID from the database
func (p *postgresDBRepo) GetAllProductsByCategoryID(cat_id int) ([]*models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var products []*models.Product

	query := `
		SELECT 
			i.id, i.product_code, i.product_name, i.product_description, i.product_status, i.category_id, i.brand_id, i.stock_alert_level, i.created_at, i.updated_at, b.id, b.name, c.id, c.name
		FROM 
			public.products i
			INNER JOIN brands b ON (b.id = i.brand_id) 
			INNER JOIN categories c ON (c.id = i.category_id)
		WHERE 
			category_id = $1; 
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
			&i.CategoryID,
			&i.BrandID,
			&i.StockAlertLevel,
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

// GetInStockProductItemsListByProductID returns a in stock list of product items corresponds to productID
func (p *postgresDBRepo) GetInStockProductItemsListByProductID(productID int) (*models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var product *models.Product
	var metadata []*models.ProductMetadata

	//get metadata from product_serial_numbers
	query := `
			SELECT
				id, serial_number, product_id, purchase_history_id, status, warranty_period, max_retail_price, purchase_rate, created_at, updated_at
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
			&pm.WarrantyPeriod,
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

// GetInStockItemDetailsBySerialNumber returns in-stock product item details corresponds to serial number from product_serial_numbers table
func (p *postgresDBRepo) GetInStockItemDetailsBySerialNumber(serialNumber string) (*models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var product *models.Product

	//get metadata from product_serial_numbers
	query := `
		SELECT
			id, serial_number, product_id, purchase_history_id, status, warranty_period, max_retail_price, purchase_rate, created_at, updated_at
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
		&metadata.WarrantyPeriod,
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

// GetSoldItemDetailsBySerialNumber returns sold product item details corresponds to serial number from product_serial_numbers table
func (p *postgresDBRepo) GetSoldItemDetailsBySerialNumber(serialNumber string) (*models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var product *models.Product

	//get metadata from product_serial_numbers
	query := `
		SELECT
			id, serial_number, product_id, purchase_history_id, sales_history_id, status, warranty_status, warranty_period, latest_warranty_id, warranty_history_ids, max_retail_price, purchase_rate, created_at, updated_at
		FROM
			public.product_serial_numbers
		WHERE
			status = 'sold' AND serial_number = $1
	`
	var metadata models.ProductMetadata
	err := p.DB.QueryRowContext(ctx, query, serialNumber).Scan(
		&metadata.ID,
		&metadata.SerialNumber,
		&metadata.ProductID,
		&metadata.PurchaseHistoryID,
		&metadata.SalesHistoryID,
		&metadata.Status,
		&metadata.WarrantyStatus,
		&metadata.WarrantyPeriod,
		&metadata.LatesWarrantyHistoryID,
		&metadata.WarrantyHistoryIDs,
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

// GetItemDetailsBySerialNumber returns product item details corresponds to serial number
func (p *postgresDBRepo) GetItemDetailsBySerialNumber(serialNumber string) (*models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var product *models.Product

	//get metadata from product_serial_numbers
	query := `
		SELECT
			id, serial_number, product_id, purchase_history_id, sales_history_id, status, warranty_status, warranty_period, latest_warranty_id, warranty_history_ids, max_retail_price, purchase_rate, created_at, updated_at
		FROM
			public.product_serial_numbers
		WHERE
			serial_number = $1
	`
	var metadata models.ProductMetadata
	err := p.DB.QueryRowContext(ctx, query, serialNumber).Scan(
		&metadata.ID,
		&metadata.SerialNumber,
		&metadata.ProductID,
		&metadata.PurchaseHistoryID,
		&metadata.SalesHistoryID,
		&metadata.Status,
		&metadata.WarrantyStatus,
		&metadata.WarrantyPeriod,
		&metadata.LatesWarrantyHistoryID,
		&metadata.WarrantyHistoryIDs,
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
			ph.id, TO_CHAR(ph.purchase_date, 'MM/DD/YYYY') AS purchase_date_str, ph.supplier_id, ph.product_id, ph.account_id, ph.chalan_no, ph.memo_no, ph.note, ph.bill_amount, ph.discount, ph.total_amount, ph.paid_amount, ph.created_at, ph.updated_at
		FROM
			public.purchase_history ph
		WHERE
			ph.memo_no = $1
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
			&purchase.Supplier.ID,
			&purchase.Product.ID,
			&purchase.HeadAccount.ID,
			&purchase.ChalanNO,
			&purchase.MemoNo,
			&purchase.Note,
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

	//retrieves all product-serial of each product_id && purchase_is
	return PurchaseHistory, nil
}

// GetSalesHistoryByMemoNo returns sales history associated with memo_no from the sales_history table
func (p *postgresDBRepo) GetSalesHistoryByMemoNo(memo_no string) ([]*models.Sale, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var salesHistory []*models.Sale

	//Get product ids for this memo with associated purchase_id for the given memo from purchase_history table
	query := `
		SELECT
			id, sale_date, customer_id, account_id, chalan_no, memo_no, note, bill_amount, discount, total_amount, paid_amount, created_at, updated_at
		FROM
			public.sales_history 
		WHERE
			memo_no = $1
	`
	rows, err := p.DB.QueryContext(ctx, query, memo_no)
	if err != nil {
		return salesHistory, err
	}
	defer rows.Close()

	for rows.Next() {
		var si models.SelectedItems
		var sale models.Sale
		err = rows.Scan(
			&sale.ID,
			&sale.SaleDate,
			&sale.CustomerID,
			&sale.AccountID,
			&sale.ChalanNO,
			&sale.MemoNo,
			&sale.Note,
			&sale.BillAmount,
			&sale.Discount,
			&sale.TotalAmount,
			&sale.PaidAmount,
			&sale.CreatedAt,
			&sale.UpdatedAt,
		)
		if err != nil {
			return salesHistory, err
		}
		sale.SelectedItems = append(sale.SelectedItems, &si)
		salesHistory = append(salesHistory, &sale)
	}

	return salesHistory, nil
}

// GetSalesHistoryByID returns sales history by its id from the sales_history table
func (p *postgresDBRepo) GetSalesHistoryByID(id int) (models.Sale, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
		SELECT
			id, sale_date, customer_id, account_id, chalan_no, memo_no, note, bill_amount, discount, total_amount, paid_amount, created_at, updated_at
		FROM
			public.sales_history 
		WHERE
			id = $1
	`
	var si models.SelectedItems
	var salesHistory models.Sale
	err := p.DB.QueryRowContext(ctx, query, id).Scan(
		&salesHistory.ID,
		&salesHistory.SaleDate,
		&salesHistory.CustomerID,
		&salesHistory.AccountID,
		&salesHistory.ChalanNO,
		&salesHistory.MemoNo,
		&salesHistory.Note,
		&salesHistory.BillAmount,
		&salesHistory.Discount,
		&salesHistory.TotalAmount,
		&salesHistory.PaidAmount,
		&salesHistory.CreatedAt,
		&salesHistory.UpdatedAt,
	)
	if err != nil {
		return salesHistory, fmt.Errorf("GetSalesHistoryByID: %w", err)
	}
	salesHistory.SelectedItems = append(salesHistory.SelectedItems, &si)
	return salesHistory, nil
}

// GetInStockProductListByPurchaseIDAndProductID returns products list associated with purchaseID and ProductID
func (p *postgresDBRepo) GetInStockProductListByPurchaseIDAndProductID(purchaseID, productID int) (*models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var product *models.Product
	var metadata []*models.ProductMetadata

	//get metadata from product_serial_numbers
	query := `
			SELECT
				id, serial_number, product_id, purchase_history_id, status, warranty_period, max_retail_price, purchase_rate, created_at, updated_at
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
			&pm.WarrantyPeriod,
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

// GetProductListBySalesIDAndProductID returns products list associated with salesID and ProductID
func (p *postgresDBRepo) GetSoldProductListBySalesIDAndProductID(salesID, productID int) (*models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var product *models.Product
	var metadata []*models.ProductMetadata

	//get metadata from product_serial_numbers
	query := `
			SELECT
				id, serial_number, product_id, purchase_history_id, sales_history_id, status, warranty_period, max_retail_price, purchase_rate, created_at, updated_at
			FROM
				public.product_serial_numbers
			WHERE
				status = 'sold' AND sales_history_id = $1 AND product_id = $2
		`
	rows, err := p.DB.QueryContext(ctx, query, salesID, productID)
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
			&pm.SalesHistoryID,
			&pm.Status,
			&pm.WarrantyPeriod,
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

// GetProductListBySalesIDAndProductID returns products list associated with salesID and ProductID
func (p *postgresDBRepo) GetProductItemsDetailsBySalesHistoryID(id int) ([]*models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var products []*models.Product
	var metadata []*models.ProductMetadata

	//get metadata from product_serial_numbers
	query := `
		SELECT
			id, serial_number, product_id, purchase_history_id, sales_history_id, status, warranty_period, max_retail_price, purchase_rate, created_at, updated_at
		FROM
			public.product_serial_numbers
		WHERE
			sales_history_id = $1 AND status = 'sold'  AND warranty_status = 'no issue'
		ORDER BY
			product_id ASC
	`
	rows, err := p.DB.QueryContext(ctx, query, id)
	if err != nil {
		return products, err
	}
	for rows.Next() {
		var pm models.ProductMetadata
		err = rows.Scan(
			&pm.ID,
			&pm.SerialNumber,
			&pm.ProductID,
			&pm.PurchaseHistoryID,
			&pm.SalesHistoryID,
			&pm.Status,
			&pm.WarrantyPeriod,
			&pm.MaxRetailPrice,
			&pm.PurchaseRate,
			&pm.CreatedAt,
			&pm.UpdatedAt,
		)
		if err != nil {
			return products, err
		}
		metadata = append(metadata, &pm)
	}

	//Grouping product metadata with corresponding product
	idMap := make(map[int]models.Product)
	for _, v := range metadata {
		// Check if an element exists in the map
		_, exists := idMap[v.ProductID]
		if !exists {
			//get product info
			pr, err := p.GetProductByID(v.ProductID)
			if err != nil {
				return products, err
			}
			//push the product info
			idMap[v.ProductID] = pr
			products = append(products, &pr)
		}
	}
	for _, v := range metadata {
		for _, product := range products {
			if v.ProductID == product.ID {
				product.ProductMetadata = append(product.ProductMetadata, v)
			}
		}
	}
	return products, nil
}

// GetExpenseList retrieves the types of expenses from the database
func (p *postgresDBRepo) GetExpenseList() ([]*models.ExpenseType, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var expList []*models.ExpenseType

	query := `
		SELECT id, expense_name, total_expense, updated_at, created_at
		FROM public.expense_list
		ORDER BY id ASC
	`
	rows, err := p.DB.QueryContext(ctx, query)
	if err != nil {
		return expList, fmt.Errorf("DBERROR: Unable to retrieve data: %w", err)
	}
	for rows.Next() {
		var exp models.ExpenseType
		err = rows.Scan(&exp.ID, &exp.ExpenseName, &exp.TotalExpense, &exp.UpdatedAt, &exp.CreatedAt)
		if err != nil {
			return expList, fmt.Errorf("DBERROR: Unable to scan row: %w", err)
		}
		expList = append(expList, &exp)
	}
	return expList, nil
}

// GeActiveStakeHolderList retrieves stakeholders list from the database
func (p *postgresDBRepo) GeActiveStakeHolderList() ([]*models.StakeHolder, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var stkList []*models.StakeHolder

	query := `
		SELECT 
			id, account_type, account_code, account_name
		FROM
			public.company_stakeholders
		WHERE 
			account_status = true
		ORDER BY updated_at ASC
	`
	rows, err := p.DB.QueryContext(ctx, query)
	if err != nil {
		return stkList, fmt.Errorf("DBERROR: GeActiveStakeHolderList: Unable to retrieve data101: %w", err)
	}
	for rows.Next() {
		var stk models.StakeHolder
		err = rows.Scan(&stk.ID, &stk.AccountType, &stk.AccountCode, &stk.AccountName)
		if err != nil {
			return stkList, fmt.Errorf("DBERROR: GeActiveStakeHolderList: Unable to scan row: %w", err)
		}
		stkList = append(stkList, &stk)
	}
	return stkList, nil
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
		purchase.Supplier.ID,
		purchase.Product.ID,
		purchase.HeadAccount.ID,
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
		values = append(values, fmt.Sprintf("('%s',%d,%d,%d,'%s','%s')", serial_number, purchase.Product.ID, purchase.MaxRetailPrice, purchase.PurchaseRate, createdAt, updatedAt))
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
			supplier_id = $1 ; 
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

// GetMemoListByCustomerID returns a list of memo with sales id from the sales_history table
func (p *postgresDBRepo) GetMemoListByCustomerID(customerID int) ([]*models.Sale, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var sales []*models.Sale

	query := `
		SELECT id, customer_id, memo_no
		FROM 
			public.sales_history
		WHERE 
			customer_id = $1 ; 
		`

	rows, err := p.DB.QueryContext(ctx, query, customerID)
	if err != nil {
		return sales, err
	}
	defer rows.Close()

	for rows.Next() {
		var s models.Sale
		err = rows.Scan(
			&s.ID,
			&s.CustomerID,
			&s.MemoNo,
		)
		if err != nil {
			return sales, err
		}
		sales = append(sales, &s)
	}
	return sales, nil
}

// RestockProduct update product quantity, store purchase history and product serial numbers
func (p *postgresDBRepo) RestockProduct(purchase *models.Purchase) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	tx, err := p.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	//Tx-1: Increase quantity_purchased and purchase_cost in products table
	query := `
		UPDATE public.products
        SET quantity_purchased = quantity_purchased + $1, purchase_cost = purchase_cost+$2, purchase_discount = purchase_discount+$3, updated_at = CURRENT_TIMESTAMP
        WHERE id = $4;`

	// Execute the query with parameters
	_, err = tx.ExecContext(ctx, query, purchase.QuantityPurchased, purchase.TotalAmount, purchase.BillAmount-purchase.TotalAmount, purchase.Product.ID)
	if err != nil {
		tx.Rollback()
		return errors.New("SQLErrorRestockProduct(Update Quantity): " + err.Error())
	}

	//Tx-2: Insert data into purchase history table
	var purchase_id int

	query = `INSERT INTO public.purchase_history (purchase_date,supplier_id,product_id,account_id,chalan_no,memo_no,note,quantity_purchased,bill_amount,discount,total_amount,paid_amount,created_at,updated_at)
	VALUES (TO_DATE($1, 'MM/DD/YYYY'), $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14) RETURNING id
	`
	row := tx.QueryRowContext(ctx, query,
		&purchase.PurchaseDate,
		purchase.Supplier.ID,
		purchase.Product.ID,
		purchase.HeadAccount.ID,
		purchase.ChalanNO,
		purchase.MemoNo,
		purchase.Note,
		purchase.QuantityPurchased,
		purchase.BillAmount,
		purchase.Discount,
		purchase.TotalAmount,
		purchase.PaidAmount,
		time.Now(),
		time.Now(),
	)
	if err = row.Scan(&purchase_id); err != nil {
		tx.Rollback()
		return errors.New("SQLErrorRestockProduct(Insert purchase_history):" + err.Error())
	}

	//Tx-3: Insert data into financial_transactions table for purchase from supplier
	var financial_transactions_id int
	query = `INSERT INTO public.financial_transactions (transaction_type,source_type,source_id,destination_type,destination_id,current_balance,amount,transaction_date,description,voucher_no,created_at,updated_at)
	VALUES ('Purchase','suppliers',$1, 'head_accounts', $2, COALESCE((SELECT current_balance FROM Public.head_accounts WHERE id=$3), 0), $4, TO_DATE($5, 'MM/DD/YYYY'), $6, $7, $8, $9) RETURNING transaction_id
	`
	purchaseDescription := "cash payment / Bank Transfer"
	row = tx.QueryRowContext(ctx, query,
		&purchase.Supplier.ID,
		&purchase.HeadAccount.ID,
		&purchase.HeadAccount.ID,
		&purchase.TotalAmount,
		&purchase.PurchaseDate,
		&purchaseDescription,
		&purchase.MemoNo,
		time.Now(),
		time.Now(),
	)
	if err = row.Scan(&financial_transactions_id); err != nil {
		tx.Rollback()
		return errors.New("SQLErrorRestockProduct(Insert financial_transactions):" + err.Error())
	}
	//Tx-4 :Update head_accounts table : Decrease current_balance, increase amount_payable(total_supplier_due), increase earned_discount
	var current_balance int
	query = `
		UPDATE Public.head_accounts
		SET current_balance = current_balance - $1, amount_payable =  amount_payable + $2, earned_discount = earned_discount + $3, updated_at = CURRENT_TIMESTAMP
		WHERE id = $4 returning current_balance
	`
	err = tx.QueryRowContext(ctx, query, purchase.PaidAmount, purchase.TotalAmount-purchase.PaidAmount, purchase.BillAmount-purchase.TotalAmount, purchase.HeadAccount.ID).Scan(&current_balance)
	if err != nil {
		tx.Rollback()
		return errors.New("SQLErrorRestockProduct(Update head_accounts info):" + err.Error())
	}
	//Update current_assets
	query = `
		UPDATE Public.head_accounts
		SET current_balance = current_balance + $1, updated_at = CURRENT_TIMESTAMP
		WHERE id = 6;
	`
	_, err = tx.ExecContext(ctx, query, purchase.TotalAmount)
	if err != nil {
		tx.Rollback()
		return errors.New("SQLErrorRestockProduct(Update head_accounts info):" + err.Error())
	}
	//Tx-5: Insert data into financial_transactions table for payment to suppliers
	query = `INSERT INTO public.financial_transactions (transaction_type,source_type,source_id,destination_type,destination_id,current_balance,amount,transaction_date,description,voucher_no,created_at,updated_at)
			VALUES ('Payment','head_accounts',$1, 'suppliers', $2, $3, $4,  TO_DATE($5, 'MM/DD/YYYY'), $6, $7, $8, $9) RETURNING transaction_id
			`
	purchaseDescription = "Payment to supplier in product purchase"
	row = tx.QueryRowContext(ctx, query,
		&purchase.HeadAccount.ID,
		&purchase.Supplier.ID,
		&current_balance,
		&purchase.PaidAmount,
		&purchase.PurchaseDate,
		&purchaseDescription,
		&purchase.MemoNo,
		time.Now(),
		time.Now(),
	)
	if err = row.Scan(&financial_transactions_id); err != nil {
		tx.Rollback()
		return errors.New("SQLErrorRestockProduct(Insert financial_transactions for payment to supplier):" + err.Error())
	}

	//Tx-6: Update increase total_amount and due_amount(if available) in suppliers table
	query = `
			UPDATE Public.suppliers
			SET total_amount = total_amount + $1, due_amount = due_amount - $2, total_discount = total_discount + $3, updated_at = CURRENT_TIMESTAMP
			WHERE id = $4
		`
	_, err = tx.ExecContext(ctx, query, purchase.TotalAmount, purchase.TotalAmount-purchase.PaidAmount, purchase.BillAmount-purchase.TotalAmount, purchase.Supplier.ID)
	if err != nil {
		tx.Rollback()
		return errors.New("SQLErrorRestockProduct(Update suppliers):" + err.Error())
	}

	//Tx-7: Insert data into product_serial_numbers to register new products
	values := []string{}
	for _, serial_number := range purchase.ProductsSerialNo {
		values = append(values, fmt.Sprintf("('%s',%d,%d,0,%d,%d,%d)", serial_number, purchase.Product.ID, purchase_id, purchase.MaxRetailPrice, purchase.PurchaseRate, purchase.WarrantyPeriod))
	}

	query = "INSERT INTO public.product_serial_numbers (serial_number,product_id,purchase_history_id,sales_history_id,max_retail_price,purchase_rate,warranty_period) VALUES " + strings.Join(values, ",") + ";"
	// Execute the query
	_, err = tx.ExecContext(ctx, query)
	if err != nil {
		tx.Rollback()
		return errors.New("SQLErrorRestockProduct(Product Serial Number):" + err.Error())
	}

	//Tx-8: Insert into or update (if the sheet_date for purchase_date already exist) top_sheet table
	//(sheet_date, total_purchases, total_payments, purchases_discount, updated_at)
	purchase_discount := purchase.BillAmount - purchase.TotalAmount
	query = `
		INSERT INTO public.top_sheet (
			sheet_date, 
			total_purchases,
			total_payments, 
			purchases_discount,
			initial_stock_value,
			updated_at
		) 
		VALUES (
			 TO_DATE($1, 'MM/DD/YYYY'), $2, $3, $4, COALESCE((SELECT initial_stock_value FROM public.top_sheet ORDER BY sheet_date DESC LIMIT 1),0)+$5, $6
		)
		ON CONFLICT (sheet_date) 
		DO UPDATE SET 
			total_purchases = public.top_sheet.total_purchases + EXCLUDED.total_purchases,
			total_payments = public.top_sheet.total_payments + EXCLUDED.total_payments,
			purchases_discount = public.top_sheet.purchases_discount + EXCLUDED.purchases_discount,
			initial_stock_value = public.top_sheet.initial_stock_value + $7,
			updated_at = CURRENT_TIMESTAMP;
	`
	_, err = tx.ExecContext(ctx, query,
		&purchase.PurchaseDate,
		&purchase.BillAmount,
		&purchase.PaidAmount,
		&purchase_discount,
		&purchase.TotalAmount,
		time.Now(),
		&purchase.TotalAmount,
	)
	if err != nil {
		tx.Rollback()
		return errors.New("SQLErrorRestockProduct(Insert or Update top_sheet):" + err.Error())
	}
	// Commit transaction
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return errors.New("SQLErrorRestockProduct(Commit):" + err.Error())
	}
	return nil
}

// ReturnProductUnitsToSupplier updates database
func (p *postgresDBRepo) ReturnProductUnitsToSupplier(PurchaseHistory models.Purchase, JobID string, transactionDate string, ProductUnitsID []int, TotalUnits int, TotalPrices int) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var id, discount int
	tx, err := p.DB.BeginTx(ctx, nil)
	if err != nil {
		return id, fmt.Errorf("failed to start transaction: %w", err)
	}
	// Prepare the SQL update query
	q1 := `
		UPDATE public.product_serial_numbers
		SET status = 'purchase returned' 
		WHERE id = $1 returning product_id, purchase_rate`
	q2 := `UPDATE public.products SET quantity_purchased = quantity_purchased -1, purchase_cost = purchase_cost-$1, purchase_discount = purchase_discount-$2, updated_at = CURRENT_TIMESTAMP WHERE id = $3`

	// Execute updates within the transaction
	for _, unitsID := range ProductUnitsID {
		var productItemID, purchase_rate int
		err := tx.QueryRowContext(ctx, q1, unitsID).Scan(&productItemID, &purchase_rate)
		if err != nil {
			tx.Rollback() // Rollback on error
			return id, fmt.Errorf("failed to update record in product_serial_numbers table with id %d: %w", unitsID, err)
		}

		//update product amount in products table
		_, err = tx.ExecContext(ctx, q2, purchase_rate, purchase_rate*PurchaseHistory.Discount/100, productItemID)
		if err != nil {
			tx.Rollback() // Rollback on error
			return id, fmt.Errorf("failed to update record in products table with id %d: %w", unitsID, err)
		}
	}

	discount = int(math.Round(float64(TotalPrices) * float64(PurchaseHistory.Discount) / 100.0))
	//TODO: retrieve supplier due
	var due int
	err = tx.QueryRowContext(ctx, "SELECT ABS(due_amount) FROM public.suppliers WHERE id = $1", PurchaseHistory.Supplier.ID).Scan(&due)
	if err != nil {
		tx.Rollback()
		return id, fmt.Errorf("SQLErrorReturnProductUnitsToSupplier(retrieve supplier due): %w", err)
	}
	decreaseAmountPayable := 0
	IncreaseCurrentBalance := 0
	if TotalPrices > due {
		decreaseAmountPayable = due
		IncreaseCurrentBalance = TotalPrices - decreaseAmountPayable
	} else {
		decreaseAmountPayable = TotalPrices
	}

	fmt.Println("inc: ", IncreaseCurrentBalance, "Des: ", decreaseAmountPayable)
	//Update total_amount, due_amount(if available) in suppliers table
	query := `
		UPDATE Public.suppliers
		SET total_amount = total_amount - $1, due_amount = due_amount + $2,total_discount = total_discount - $3, updated_at = CURRENT_TIMESTAMP 
		WHERE id = $4
	`
	_, err = tx.ExecContext(ctx, query, TotalPrices, decreaseAmountPayable, discount, PurchaseHistory.Supplier.ID)
	if err != nil {
		tx.Rollback()
		return id, fmt.Errorf("SQLErrorReturnProductUnitsToSupplier(Update suppliers): %w", err)
	}

	//Insert Financial transaction for purchase return
	var financial_transactions_id int
	query = `INSERT INTO public.financial_transactions (transaction_type,source_type,source_id,destination_type,destination_id,amount,transaction_date,description, current_balance)
	VALUES ('Purchase Return','suppliers',$1, 'head_accounts', $2, $3, CURRENT_TIMESTAMP, $4, COALESCE((SELECT current_balance FROM Public.head_accounts WHERE id=$5), 0)) RETURNING transaction_id
	`
	description := ""
	err = tx.QueryRowContext(ctx, query,
		&PurchaseHistory.Supplier.ID,
		&PurchaseHistory.HeadAccount.ID,
		&TotalPrices,
		&description,
		&PurchaseHistory.HeadAccount.ID,
	).Scan(&financial_transactions_id)
	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("SQLErrorReturnProductUnitsToSupplier(Insert financial_transactions for purchase return): %w", err)
	}
	//Tx-3 :Update head_accounts info :: current_balance,  total_supplier_due
	var current_balance int
	query = `
		UPDATE Public.head_accounts
		SET current_balance = current_balance + $1, amount_payable =  amount_payable - $2, earned_discount = earned_discount - $3,  updated_at = CURRENT_TIMESTAMP 
		WHERE id = $4 returning current_balance
	`
	err = tx.QueryRowContext(ctx, query, IncreaseCurrentBalance, decreaseAmountPayable, discount, PurchaseHistory.HeadAccount.ID).Scan(&current_balance)
	if err != nil {
		tx.Rollback()
		return 0, errors.New("SQLErrorReturnProductUnitsToSupplier(Update head_accounts info):" + err.Error())
	}
	//Update current_assets
	query = `
		UPDATE Public.head_accounts
		SET current_balance = current_balance - $1, updated_at = CURRENT_TIMESTAMP
		WHERE id = 6;
	`
	_, err = tx.ExecContext(ctx, query, TotalPrices)
	if err != nil {
		tx.Rollback()
		return id, errors.New("SQLErrorReturnProductUnitsToSupplier(Update current_assets):" + err.Error())
	}
	//Insert Financial transaction for received amount
	query = `INSERT INTO public.financial_transactions (transaction_type,source_type,source_id,destination_type,destination_id,amount,transaction_date,description, current_balance)
	VALUES ('Refund','suppliers',$1, 'head_accounts', $2, $3, CURRENT_TIMESTAMP, $4, COALESCE((SELECT current_balance FROM Public.head_accounts WHERE id=$5), 0)) RETURNING transaction_id
	`
	description = "cash return due to returning product to the supplier"
	err = tx.QueryRowContext(ctx, query,
		&PurchaseHistory.Supplier.ID,
		&PurchaseHistory.HeadAccount.ID,
		&TotalPrices,
		&description,
		&PurchaseHistory.HeadAccount.ID,
	).Scan(&financial_transactions_id)
	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("SQLErrorReturnProductUnitsToSupplier(Insert financial_transactions for received amount): %w", err)
	}

	//Tx-8: Insert into or update (if the sheet_date for purchase_date already exist) top_sheet table
	//(sheet_date, total_purchases, total_payments, purchases_discount, updated_at)
	query = `
		INSERT INTO public.top_sheet (
			sheet_date, 
			total_purchase_returns,
			purchases_discount,
			initial_stock_value
		) 
		VALUES (
			CURRENT_TIMESTAMP, $1, $2, COALESCE((SELECT initial_stock_value FROM public.top_sheet ORDER BY sheet_date DESC LIMIT 1),0)-$3
		)
		ON CONFLICT (sheet_date) 
		DO UPDATE SET 
			total_purchase_returns = public.top_sheet.total_purchase_returns + EXCLUDED.total_purchase_returns,
			purchases_discount = public.top_sheet.purchases_discount + EXCLUDED.purchases_discount,
			initial_stock_value = EXCLUDED.initial_stock_value,
			updated_at = CURRENT_TIMESTAMP;
	`
	_, err = tx.ExecContext(ctx, query,
		&TotalPrices,
		(-1)*discount,
		&TotalPrices,
	)
	if err != nil {
		tx.Rollback()
		return id, errors.New("SQLErrorReturnProductUnitsToSupplier(Insert or Update top_sheet):" + err.Error())
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return id, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return id, nil
}

// SaleProductsToCustomer update database for sale process
func (p *postgresDBRepo) SaleProductsToCustomer(sale *models.SalesInvoice) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	tx, err := p.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	//Tx-1: Insert sales details to sales history table
	var sale_id int
	query := `
		INSERT INTO public.sales_history (sale_date,customer_id,account_id,chalan_no,memo_no,note,bill_amount,discount,total_amount,paid_amount,gross_profit)
		VALUES (TO_DATE($1,'MM/DD/YYYY'), $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id
	`
	row := tx.QueryRowContext(ctx, query,
		sale.SaleDate,
		sale.CustomerInfo.ID,
		sale.HeadAccountInfo.ID,
		sale.ChalanNo,
		sale.MemoNo,
		sale.Note,
		sale.BillAmount,
		sale.Discount,
		sale.TotalAmount,
		sale.PaidAmount,
		sale.GrossProfit,
	)
	if err = row.Scan(&sale_id); err != nil {
		tx.Rollback()
		return errors.New("SQLErrorSaleProductsToCustomer(INSERT sales_history): " + err.Error())
	}

	totalQuantity := 0
	totalPrice := 0
	totalPurchase := 0
	//loop over the SelectedProduct array
	for i, items := range sale.ProductItems {
		//Tx-2: Update product quantity
		query := `
			UPDATE public.products
			SET quantity_sold = quantity_sold + $1, sold_price = sold_price + $2, sold_discount = sold_discount + $3
			WHERE id = $4;
		  `
		_, err = tx.ExecContext(ctx, query, items.Quantity, items.SubTotal, items.SubTotal*sale.Discount, items.ProductID)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("SQLErrorSaleProductsToCustomer(Update products table): For id %d --%w", i, err)
		}
		totalQuantity += items.Quantity
		totalPrice += items.SubTotal
		//Tx-3: update product items status and sales_history_id
		for _, serialNumber := range items.SerialNumbers {
			query = `
				UPDATE public.product_serial_numbers
				SET status = 'sold', sales_history_id = $1 
				WHERE serial_number = $2 RETURNING purchase_rate
			`
			var purchaseRate int
			err := tx.QueryRowContext(ctx, query, sale_id, serialNumber).Scan(&purchaseRate)
			if err != nil {
				tx.Rollback()
				return fmt.Errorf("SQLErrorSaleProductsToCustomer(Update Product status And status):#serial-%s --%w", serialNumber, err)
			}
			totalPurchase += purchaseRate
		}
		//
	}

	due := sale.TotalAmount - sale.PaidAmount
	//Tx-5: Update total_amount, due_amount(if available) in customers table
	query = `
		UPDATE Public.customers
		SET total_amount = total_amount + $1, due_amount = due_amount + $2, total_discount = $3
		WHERE id = $4
	`
	_, err = tx.ExecContext(ctx, query, sale.TotalAmount, due, sale.BillAmount-sale.TotalAmount, sale.CustomerInfo.ID)
	if err != nil {
		tx.Rollback()
		return errors.New("SQLErrorSaleProductsToCustomer(Update customers):" + err.Error())
	}

	//Tx-7: store financial_transactions for sale
	var financial_transactions_id int
	query = `INSERT INTO public.financial_transactions (transaction_type,source_type,source_id,destination_type,destination_id,current_balance,amount,transaction_date,description,voucher_no,created_at,updated_at)
	VALUES ('Sale','customers',$1, 'head_accounts', $2, COALESCE((SELECT current_balance FROM Public.head_accounts WHERE id=$3), 0), $4, TO_DATE($5,'MM/DD/YYYY'), $6, $7, $8, $9) RETURNING transaction_id
	`
	description := "Cash Sale / Bank Transfer"
	row = tx.QueryRowContext(ctx, query,
		&sale.CustomerInfo.ID,
		&sale.HeadAccountInfo.ID,
		&sale.HeadAccountInfo.ID,
		&sale.TotalAmount,
		&sale.SaleDate,
		&description,
		&sale.MemoNo,
		time.Now(),
		time.Now(),
	)
	if err = row.Scan(&financial_transactions_id); err != nil {
		tx.Rollback()
		return errors.New("SQLErrorSaleProductsToCustomer(Insert financial_transactions):" + err.Error())
	}

	//Tx-4 :Update head_accounts info : current_balance,  total_customer_due
	var current_balance int
	query = `
		UPDATE Public.head_accounts
		SET current_balance = current_balance + $1, amount_receivable =  amount_receivable + $2, offered_discount = offered_discount + $3
		WHERE id = $4 RETURNING current_balance;
	`
	err = tx.QueryRowContext(ctx, query, sale.PaidAmount, sale.TotalAmount-sale.PaidAmount, sale.BillAmount-sale.TotalAmount, sale.HeadAccountInfo.ID).Scan(&current_balance)
	if err != nil {
		tx.Rollback()
		return errors.New("SQLErrorSaleProductsToCustomer(Update head_accounts info):" + err.Error())
	}
	query = `
		UPDATE Public.head_accounts
		SET current_balance = current_balance + $1
		WHERE id = 5;
	`
	_, err = tx.ExecContext(ctx, query, sale.BillAmount-totalPurchase)
	if err != nil {
		tx.Rollback()
		return errors.New("SQLErrorSaleProductsToCustomer(Update REVENUE ACCOUNTS info):" + err.Error())
	}
	query = `
		UPDATE Public.head_accounts
		SET current_balance = current_balance - $1
		WHERE id = 6;
	`
	_, err = tx.ExecContext(ctx, query, totalPurchase)
	if err != nil {
		tx.Rollback()
		return errors.New("SQLErrorSaleProductsToCustomer(Update CURRENT ASSETS info):" + err.Error())
	}

	//Tx-8: store financial_transactions for customer payment
	query = `
		INSERT INTO public.financial_transactions (transaction_type,source_type,source_id,destination_type,destination_id,current_balance,amount,transaction_date,description,voucher_no)
		VALUES ('Receive & Collection','customers',$1, 'head_accounts', $2, $3, $4, TO_DATE($5,'MM/DD/YYYY'), $6, $7) RETURNING transaction_id
	`

	description = "Received amount from customer"
	row = tx.QueryRowContext(ctx, query,
		&sale.CustomerInfo.ID,
		&sale.HeadAccountInfo.ID,
		&current_balance,
		&sale.PaidAmount,
		&sale.SaleDate,
		&description,
		&sale.MemoNo,
	)
	if err = row.Scan(&financial_transactions_id); err != nil {
		tx.Rollback()
		return errors.New("SQLErrorSaleProductsToCustomer(Insert financial_transactions for customer payment):" + err.Error())
	}
	//Tx-9: update top_sheet data if the sheet_date for sales date already exist,
	//Otherwise insert a new row  in the top_sheet table
	sale_discount := sale.BillAmount - sale.TotalAmount
	query = `
		INSERT INTO public.top_sheet (
			sheet_date, 
			total_sales,
			total_received_payments, 
			sales_discount,
			gross_profit,
			initial_stock_value
		) 
		VALUES (
			TO_DATE($1,'MM/DD/YYYY'), $2, $3, $4, $5, (SELECT initial_stock_value FROM public.top_sheet ORDER BY sheet_date DESC LIMIT 1)-$6
		)
		ON CONFLICT (sheet_date) 
		DO UPDATE SET 
			total_sales = public.top_sheet.total_sales + EXCLUDED.total_sales,
			total_received_payments = public.top_sheet.total_received_payments + EXCLUDED.total_received_payments,
			sales_discount = public.top_sheet.sales_discount + EXCLUDED.sales_discount,
			gross_profit = public.top_sheet.gross_profit + EXCLUDED.gross_profit,
			initial_stock_value = EXCLUDED.initial_stock_value,
			updated_at = CURRENT_TIMESTAMP;
	`

	_, err = tx.ExecContext(ctx, query,
		&sale.SaleDate,
		&sale.BillAmount,
		&sale.PaidAmount,
		&sale_discount,
		&sale.GrossProfit,
		&totalPurchase,
	)
	if err != nil {
		tx.Rollback()
		return errors.New("SQLErrorSaleProductsToCustomer(Insert or Update top_sheet):" + err.Error())
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return errors.New("SQLErrorSaleProductsToCustomer(Commit):" + err.Error())
	}
	return nil
}

// SaleReturnDB updates database sale return process
func (p *postgresDBRepo) SaleReturnDB(SalesHistory *models.Sale, SelectedItemsID []int, SaleReturnDate string, ReturnItemsCount int, ReturnAmount int, MemoNo string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	//tx start
	tx, err := p.DB.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("DBERROR:=>SaleReturnDB: Unable to begin transaction %w", err)
	}
	//step:1 iterate over the SelectedItemsID slice and update the associated row(id) , set status="in stock", updated_at = time.Now()
	for _, id := range SelectedItemsID {
		var productItemID, mrp int
		query := `
			UPDATE 
				public.product_serial_numbers
			SET
				status='in stock', updated_at=$1
			WHERE
				id = $2
			RETURNING product_id, max_retail_price		
		`
		err := tx.QueryRowContext(ctx, query, time.Now(), id).Scan(&productItemID, &mrp)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("DBERROR:=>SaleReturnDB: Unable to execute UPDATE in product_serial_numbers table SQL %w", err)
		}
		query = `
			UPDATE 
				public.products
			SET
				quantity_sold = quantity_sold - 1, sold_price = sold_price - $1
			WHERE
				id = $2		
		`
		_, err = tx.ExecContext(ctx, query, int(mrp-mrp*SalesHistory.Discount/100), productItemID)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("DBERROR:=>SaleReturnDB: Unable to execute UPDATE in products table SQL %w", err)
		}
	}

	//step:2 insert sale return data to the sale_return_history
	//get the last index of sales_return_history Table
	lastID, err := p.LastIndex("sales_return_history")
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("DBERROR:=>SaleReturnDB: Unable to get last index of sales_return_history table:  %w", err)
	}
	MemoNo = MemoNo + strconv.Itoa(int(lastID+1)) //update memo no
	//return_items_id in string, separated by '-'
	returnItemsID := strconv.Itoa(SelectedItemsID[0])
	ln := len(SelectedItemsID)
	for i := 1; i < ln; i++ {
		returnItemsID += "-" + strconv.Itoa(SelectedItemsID[i])
	}
	query := `
		INSERT INTO public.sales_return_history(sale_return_date,customer_id ,sales_history_id ,memo_no,returned_product_ids,total_returned_count,total_returned_amount, created_at, updated_at)
		VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`
	result, err := tx.ExecContext(ctx, query, SaleReturnDate, SalesHistory.CustomerID, SalesHistory.ID, MemoNo, returnItemsID, ReturnItemsCount, ReturnAmount, time.Now(), time.Now())
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("DBERROR:=>SaleReturnDB: Unable to execute UPDATE SQL %w", err)
	}
	if n, err := result.RowsAffected(); err != nil || n != 1 {
		return fmt.Errorf("DBERROR:=>SaleReturnDB: Number of affected row is not equal to 1:  %w", err)
	}

	due := 0
	discount := (SalesHistory.BillAmount - SalesHistory.TotalAmount) * ReturnAmount / SalesHistory.TotalAmount
	if ReturnAmount >= SalesHistory.TotalAmount-SalesHistory.PaidAmount {
		due = SalesHistory.TotalAmount - SalesHistory.PaidAmount
	} else {
		due = ReturnAmount
	}
	//Update total_amount, due_amount(if available) in customers table
	query = `
		UPDATE Public.customers
		SET total_amount = total_amount - $1, due_amount = due_amount - $2
		WHERE id = $3
	`
	_, err = tx.ExecContext(ctx, query, ReturnAmount, due, SalesHistory.CustomerID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("SQLErrorSaleReturnDB(Update customer): %w", err)
	}

	//Tx-3 :Update head_accounts info :: current_balance,  total_customer_due
	var current_balance int
	query = `
		UPDATE Public.head_accounts
		SET current_balance = current_balance - $1, amount_receivable =  amount_receivable - $2, offered_discount = offered_discount - $3
		WHERE id = $4 returning current_balance
	`
	err = tx.QueryRowContext(ctx, query, ReturnAmount, due, discount, SalesHistory.AccountID).Scan(&current_balance)
	if err != nil {
		tx.Rollback()
		return errors.New("SQLErrorSaleReturnDB(Update head_accounts info):" + err.Error())
	}
	//Insert Financial transaction for sale return amount
	var financial_transactions_id int
	query = `INSERT INTO public.financial_transactions (transaction_type,source_type,source_id,destination_type,destination_id,amount,transaction_date,description,current_balance,created_at,updated_at)
	VALUES ('Sale Return','head_accounts',$1, 'customers', $2, $3, $4, $5, $6, $7, $8) RETURNING transaction_id
	`
	description := "sales return from customer"
	err = tx.QueryRowContext(ctx, query,
		&SalesHistory.AccountID,
		&SalesHistory.CustomerID,
		&ReturnAmount,
		time.Now(),
		&description,
		&current_balance,
		time.Now(),
		time.Now(),
	).Scan(&financial_transactions_id)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("SQLErrorSaleReturn(Insert financial_transactions): %w", err)
	}

	//Insert Financial transaction for sale return Repayment
	query = `INSERT INTO public.financial_transactions (transaction_type,source_type,source_id,destination_type,destination_id,amount,transaction_date,description,current_balance,created_at,updated_at)
	VALUES ('Sale Return','head_accounts',$1, 'customers', $2, $3, $4, $5, $6, $7) RETURNING transaction_id
	`
	description = "cash return to customer due to returning product"
	err = tx.QueryRowContext(ctx, query,
		&SalesHistory.AccountID,
		&SalesHistory.CustomerID,
		&ReturnAmount,
		time.Now(),
		&description,
		&current_balance,
		time.Now(),
		time.Now(),
	).Scan(&financial_transactions_id)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("SQLErrorSaleReturn(Insert financial_transactions): %w", err)
	}

	//tx commit
	err = tx.Commit()

	if err != nil {
		tx.Rollback()
		return fmt.Errorf("DBERROR:=>SaleReturnDB: Unable to commit transaction:  %w", err)
	}
	return nil
}

// GetWarrantyList retrieves a slice of warranty history from warranty_history_table
func (p *postgresDBRepo) GetWarrantyList(searchType string) ([]*models.Warranty, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var warrantyHistory []*models.Warranty

	query := `
			SELECT
				id, status, memo_no, product_serial_id, previous_serial_number, new_serial_number, contact_number, request_date, reported_problem, received_by, checkout_date, delivery_date, delivered_by, comment, created_at, updated_at
			FROM
				public.warranty_history
			WHERE
				status = $1
		`
	rows, err := p.DB.QueryContext(ctx, query, searchType)
	if err != nil {
		return warrantyHistory, fmt.Errorf("DBERROR:GetWarrantyList => %w", err)
	}

	for rows.Next() {
		var wh models.Warranty
		err = rows.Scan(
			&wh.ID,
			&wh.Status,
			&wh.MemoNo,
			&wh.ProductSerialID,
			&wh.PreviousSerialNo,
			&wh.NewSerialNo,
			&wh.ContactNumber,
			&wh.RequestedDate,
			&wh.ReportedProblem,
			&wh.ReceivedBy,
			&wh.CheckoutDate,
			&wh.DeliveryDate,
			&wh.DeliveredBy,
			&wh.Comment,
			&wh.CreatedAt,
			&wh.UpdatedAt,
		)
		if err != nil {
			return warrantyHistory, err
		}
		warrantyHistory = append(warrantyHistory, &wh)
	}
	return warrantyHistory, nil
}

// AddNewWarrantyClaim handles database operations for completing claim warranty procedures
func (p *postgresDBRepo) AddNewWarrantyClaim(memoPrefix string, serialID int, serialNumber, contactNumber, reportedProblem, receivedBy, warrantyHistoryIds string) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var id int
	//begin transaction
	tx, err := p.DB.Begin()
	if err != nil {
		return id, fmt.Errorf("DBERROR: AddNewWarrantyClaim => Begin Tx => %w", err)
	}
	// Get the current time
	currentTime := time.Now()

	// 	step-1: insert new row at warranty_history table with data from product_serial_number and status = warranty claim, product_serial_id = current_serial_id
	query := `INSERT INTO public.warranty_history(memo_no, product_serial_id, previous_serial_number, contact_number, request_date, reported_problem, received_by, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`

	err = tx.QueryRowContext(ctx, query, memoPrefix, serialID, serialNumber, contactNumber, currentTime.Format("01/02/2006"), reportedProblem, receivedBy, currentTime, currentTime).Scan(&id)

	if err != nil {
		tx.Rollback()
		return id, fmt.Errorf("DBERROR: AddNewWarrantyClaim =>Insert_warranty_history Table => %w", err)
	}
	// 	step-2 : update latest_warranty_history_id = pkid of warranty_history, warranty_history_ids = concat{warranty_history_ids,pkid of warranty_history}, updated_at = time.Now() in product_serial_number
	query = `
		UPDATE 
			public.product_serial_numbers
		SET 
			latest_warranty_id = $1,
			warranty_status = 'in progress',
			warranty_history_ids = warranty_history_ids || $2, 
			updated_at = $3
		WHERE 
			id = $4
	`
	// Execute the query, passing the correct values in the right order
	result, err := p.DB.ExecContext(ctx, query, id, strconv.Itoa(id), currentTime, serialID)
	if err != nil {
		tx.Rollback()
		return id, fmt.Errorf("DBERROR: AddNewWarrantyClaim => Update_product_serial_numbers: %w", err)
	}

	// Optionally, you can check the result (rows affected) if needed
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return id, fmt.Errorf("DBERROR: RowsAffected: %w", err)
	}

	if rowsAffected != 1 {
		tx.Rollback()
		return id, errors.New("DBERROR:AddNewWarrantyClaim::no row affected when trying to update product_serial_numbers table")
	}
	// Commit the transaction
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return id, fmt.Errorf("DBERROR: AddNewWarrantyClaim => failed to commit transaction: %w", err)
	}
	return id, nil
}

//CheckoutWarrantyProduct update database for checkout product
/*	step-1: set new_serial_number = NewSerialNumber, status = "checked out", checkout_date = ArrivalDate, comment=Comment, updated_at = time.Now()
	where id = WarrantyHistoryID in warranty_history table
	step-2: set serial_number = NewSerialNumber, warranty_status = 'no issue', updated_at = time.Now()
	where id = productSerialID in product_serial_numbers Table
*/
func (p *postgresDBRepo) CheckoutWarrantyProduct(warrantyHistoryID, productSerialID int, arrivalDate, newSerialNumber, comment string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	//begin a transaction
	tx, err := p.DB.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("DBERROR: CheckoutWarrantyProduct: Unable to begin transaction => %w", err)
	}
	//update warranty_history Table
	query := `
		UPDATE 
			public.warranty_history
		SET 
			new_serial_number = $1, status = 'checked out', checkout_date = $2, comment= $3, updated_at = $4
		WHERE 
			id = $5;		
	`
	result, err := tx.ExecContext(ctx, query, newSerialNumber, arrivalDate, comment, time.Now(), warrantyHistoryID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("DBERROR: CheckoutWarrantyProduct: Unable to update warranty_history Table => %w", err)
	}
	n, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("DBERROR: CheckoutWarrantyProduct: Error counting affected rows in warranty_history Table => %w", err)
	}
	if n != 1 {
		tx.Rollback()
		return fmt.Errorf("DBERROR: CheckoutWarrantyProduct: Number of affected rows in warranty_history Table is undesired => %w", err)
	}
	//update product_serial_numbers
	query = `
		UPDATE 
			public.product_serial_numbers
		SET 
			serial_number = $1, warranty_status = 'delivery ready', updated_at = $2
		WHERE 
			id = $3;		
	`
	result, err = tx.ExecContext(ctx, query, newSerialNumber, time.Now(), productSerialID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("DBERROR: CheckoutWarrantyProduct: Unable to update product_serial_numbers Table => %w", err)
	}
	n, err = result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("DBERROR: CheckoutWarrantyProduct: Error counting affected rows in product_serial_numbers Table => %w", err)
	}
	if n != 1 {
		tx.Rollback()
		return fmt.Errorf("DBERROR: CheckoutWarrantyProduct: Number of affected rows in product_serial_numbers Table is undesired => %w", err)
	}

	//commit the changes
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("DBERROR: CheckoutWarrantyProduct: Unable to commit changes => %w", err)
	}
	return nil
}

//DeliveryWarrantyProduct update database to delivery product
/*	step-1: set set status='delivered', delivery_date=string(time.now), delivered_by=deliveredBy, updated_at=time.Now()
	where id=warrantyHistoryID in warranty_history table
	step-2: set warranty_status='no issue', updated_at=time.Now()
	where id=productSerialID in product_serial_numbers Table
*/
func (p *postgresDBRepo) DeliveryWarrantyProduct(warrantyHistoryID, productSerialID int, deliveredBy string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	//begin a transaction
	tx, err := p.DB.Begin()
	if err != nil {
		return fmt.Errorf("DBERROR: DeliveryWarrantyProduct: Unable to begin transaction => %w", err)
	}
	//update warranty_history Table
	query := `
		UPDATE 
			public.warranty_history
		SET 
			status = 'delivered', delivery_date = $1, delivered_by= $2, updated_at = $3
		WHERE 
			id = $4;		
	`
	curentDate := time.Now()
	result, err := tx.ExecContext(ctx, query, curentDate.Format("01/02/2006"), deliveredBy, curentDate, warrantyHistoryID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("DBERROR: DeliveryWarrantyProduct: Unable to update warranty_history Table => %w", err)
	}
	n, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("DBERROR: DeliveryWarrantyProduct: Error counting affected rows in warranty_history Table => %w", err)
	}
	if n != 1 {
		tx.Rollback()
		return fmt.Errorf("DBERROR: DeliveryWarrantyProduct: Number of affected rows in warranty_history Table is undesired => %w", err)
	}

	//update product_serial_numbers
	query = `
		UPDATE 
			public.product_serial_numbers
		SET 
			warranty_status = 'no issue', updated_at = $1
		WHERE 
			id = $2;		
	`
	result, err = tx.ExecContext(ctx, query, time.Now(), productSerialID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("DBERROR: DeliveryWarrantyProduct: Unable to update product_serial_numbers Table => %w", err)
	}
	n, err = result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("DBERROR: DeliveryWarrantyProduct: Error counting affected rows in product_serial_numbers Table => %w", err)
	}
	if n != 1 {
		tx.Rollback()
		return fmt.Errorf("DBERROR: DeliveryWarrantyProduct: Number of affected rows in product_serial_numbers Table is undesired => %w", err)
	}

	//commit the changes
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("DBERROR: DeliveryWarrantyProduct: Unable to commit changes => %w", err)
	}
	return nil
}

// .......................Accounts.......................
func (p *postgresDBRepo) CompleteReceiveCollectionTransactions(summary []*models.Reception) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	//Begin transaction
	tx, err := p.DB.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("DBERROR: CompleteReceiveCollectionTransactions => Unable to begin transaction: %w", err)
	}

	total_received_payments := 0
	for _, sum := range summary {
		total_received_payments += sum.ReceivedAmount
		//update Cash-Bank account
		//set current_balance += received_amount
		var current_balance int
		stmt := `
			UPDATE public.head_accounts
			SET current_balance = current_balance + $1, amount_receivable = amount_receivable - $2, updated_at = $3
			WHERE id = $4 RETURNING current_balance;
		`
		err = tx.QueryRowContext(ctx, stmt, sum.ReceivedAmount, sum.ReceivedAmount, time.Now(), sum.DestinationAccount.ID).Scan(&current_balance)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("DBERROR: CompleteReceiveCollectionTransactions => Unable to update current_balance in head_accounts table: %w", err)
		}

		//insert to financial transactions
		stmt = `
			INSERT INTO public.financial_transactions(transaction_type, source_type, source_id, destination_type, destination_id, amount, transaction_date, voucher_no, description,carrier,cheque_no, current_balance)
			VALUES('Receive & Collection', 'customers', $1, 'head_accounts', $2, $3, TO_DATE($4,'MM/DD/YYYY'), $5, $6, $7, $8, $9)
		`
		_, err = tx.ExecContext(ctx, stmt, sum.SourceAccount.ID, sum.DestinationAccount.ID, sum.ReceivedAmount, sum.ReceivedDate, sum.VoucherNo, sum.Description, sum.Carrier, sum.ChequeNo, current_balance)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("DBERROR: CompleteReceiveCollectionTransactions => Unable Insert into financial_transaction table: %w", err)
		}

		//Tx: update top_sheet data if the sheet_date for entry for current date already exist,
		//Otherwise insert a new row  in the top_sheet table
		query := `
		INSERT INTO public.top_sheet (
			sheet_date, 
			total_received_payments, 
			initial_stock_value
		) 
		VALUES (
			 TO_DATE($1, 'MM/DD/YYYY'), $2, COALESCE((SELECT initial_stock_value FROM public.top_sheet ORDER BY sheet_date DESC LIMIT 1),0)
		)
		ON CONFLICT (sheet_date) 
		DO UPDATE SET 
			total_received_payments = public.top_sheet.total_received_payments + EXCLUDED.total_received_payments,
			updated_at = CURRENT_TIMESTAMP;
	`
		_, err = tx.ExecContext(ctx, query,
			&sum.ReceivedDate,
			&sum.ReceivedAmount,
		)
		if err != nil {
			tx.Rollback()
			return errors.New("CompleteReceiveCollection(Insert or Update top_sheet):" + err.Error())
		}
		//update customer account
		//set due_amount -= received_amount
		stmt = `
			UPDATE public.customers
			SET due_amount = due_amount - $1, updated_at = $2
			WHERE id = $3`

		_, err = tx.ExecContext(ctx, stmt, sum.ReceivedAmount, time.Now(), sum.SourceAccount.ID)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("DBERROR: CompleteReceiveCollectionTransactions => Unable to update due_amount in customers table: %w", err)
		}
	}

	//commit transaction
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("DBERROR: CompleteReceiveCollectionTransactions => Unable to commit: %w", err)
	}
	return nil
}
func (p *postgresDBRepo) CompletePaymentTransactions(summary []*models.Payment) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	//Begin transaction
	tx, err := p.DB.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("DBERROR: CompletePaymentTransactions => Unable to begin transaction: %w", err)
	}

	//update Cash-Bank account
	//set current_balance -= paid_amount
	//update supplier account
	//set due_mount -= paid_amount

	for _, sum := range summary {
		//update Cash-Bank account
		//set current_balance += received_amount
		stmt := `
			UPDATE Public.head_accounts
			SET current_balance = current_balance - $1, amount_payable =  amount_payable - $2, updated_at = $3
			WHERE id = $4 returning current_balance`

		var current_balance int
		err = tx.QueryRowContext(ctx, stmt, sum.PaidAmount, sum.PaidAmount, time.Now(), sum.SourceAccount.ID).Scan(&current_balance)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("DBERROR: CompletePaymentTransactions => Unable to update current_balance in head_accounts table: %w", err)
		}

		//convert receivedDate into go supported date
		txDate, err := time.Parse("01/02/2006", sum.PaymentDate)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("DBERROR: CompletePaymentTransactions => Unable parse time in go supported format: %w", err)
		}
		//insert to financial transactions
		stmt = `
			INSERT INTO public.financial_transactions(transaction_type, source_type, source_id, destination_type, destination_id, amount, transaction_date, voucher_no, description,carrier,cheque_no, current_balance)
			VALUES('Payment', 'head_accounts', $1, 'suppliers', $2, $3, $4, $5, $6, $7, $8, $9)
		`
		_, err = tx.ExecContext(ctx, stmt, sum.SourceAccount.ID, sum.DestinationAccount.ID, sum.PaidAmount, txDate, sum.VoucherNo, sum.Description, sum.Carrier, sum.ChequeNo, current_balance)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("DBERROR: CompletePaymentTransactions => Unable Insert into financial_transaction table: %w", err)
		}

		//Tx: update top_sheet data if the sheet_date for entriy for current date already exist,
		//Otherwise insert a new row  in the top_sheet table
		query := `
		INSERT INTO public.top_sheet (
			sheet_date, 
			total_payments, 
			initial_stock_value
		) 
		VALUES (
			 TO_DATE($1, 'MM/DD/YYYY'), $2, COALESCE((SELECT initial_stock_value FROM public.top_sheet ORDER BY sheet_date DESC LIMIT 1),0)
		)
		ON CONFLICT (sheet_date) 
		DO UPDATE SET 
			total_payments = public.top_sheet.total_payments + EXCLUDED.total_payments,
			updated_at = CURRENT_TIMESTAMP;
		`
		_, err = tx.ExecContext(ctx, query,
			&sum.PaymentDate,
			&sum.PaidAmount,
		)
		if err != nil {
			tx.Rollback()
			return errors.New("CompletePaymentTransactions(Insert or Update top_sheet):" + err.Error())
		}
		//update suppliers account
		//set due_mount -= received_amount
		stmt = `
			UPDATE public.suppliers
			SET due_amount = due_amount + $1, updated_at = $2
			WHERE id = $3`

		_, err = tx.ExecContext(ctx, stmt, sum.PaidAmount, time.Now(), sum.DestinationAccount.ID)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("DBERROR: CompletePaymentTransactions => Unable to update due_amount in suppliers table: %w", err)
		}
	}

	//commit transaction
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("DBERROR: CompletePaymentTransactions => Unable to commit: %w", err)
	}
	return nil
}

func (p *postgresDBRepo) CompleteAmountTransferTransactions(summary []*models.AmountTransfer) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	//Begin transaction
	tx, err := p.DB.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("DBERROR: CompleteAmountTransferTransactions => Unable to begin transaction: %w", err)
	}

	for _, sum := range summary {
		//update Source account
		//set current_balance -= received_amount
		stmt := `
			UPDATE public.head_accounts
			SET current_balance = current_balance - $1, updated_at = $2
			WHERE id = $3`

		_, err = tx.ExecContext(ctx, stmt, sum.TransferAmount, time.Now(), sum.SourceAccount.ID)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("DBERROR: CompleteAmountTransferTransactions => Unable to update current_balance in source_account table: %w", err)
		}
		//update Destination account
		//set current_balance -= received_amount
		stmt = `
			UPDATE public.head_accounts
			SET current_balance = current_balance + $1, updated_at = $2
			WHERE id = $3`

		_, err = tx.ExecContext(ctx, stmt, sum.TransferAmount, time.Now(), sum.DestinationAccount.ID)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("DBERROR: CompleteAmountTransferTransactions => Unable to update current_balance in destination_account table: %w", err)
		}

		//convert receivedDate into go supported date
		txDate, err := time.Parse("01/02/2006", sum.TransactionDate)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("DBERROR: CompleteAmountTransferTransactions => Unable parse time in go supported format: %w", err)
		}
		//insert to financial transactions
		stmt = `
			INSERT INTO public.financial_transactions(transaction_type, source_type, source_id, destination_type, destination_id, amount, transaction_date, voucher_no, description)
			VALUES('Cash Payment', 'head_accounts', $1, 'head_accounts', $2, $3, $4, $5, $6)
		`
		_, err = tx.ExecContext(ctx, stmt, sum.SourceAccount.ID, sum.DestinationAccount.ID, sum.TransferAmount, txDate, sum.VoucherNo, sum.Description)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("DBERROR: CompleteAmountTransferTransactions => Unable Insert into financial_transaction table: %w", err)
		}
	}

	//commit transaction
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("DBERROR: CompleteAmountTransferTransactions => Unable to commit: %w", err)
	}
	return nil
}

func (p *postgresDBRepo) CompleteAmountPayableTransactions(summary []*models.AmountPayable) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	//Begin transaction
	tx, err := p.DB.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("DBERROR: CompleteAmountPayableTransactions => Unable to begin transaction: %w", err)
	}

	for _, sum := range summary {
		//update customers/suppliers/employees accounts
		//update head_accounts
		stmt := fmt.Sprintf(`
			UPDATE public.%s
			SET due_amount = due_amount - $1, updated_at = CURRENT_TIMESTAMP
			WHERE id = $2
		`, sum.AccountType)
		fmt.Println("account id : ", sum.AccountID)
		_, err = tx.ExecContext(ctx, stmt, sum.PayableAmount, sum.AccountID)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("DBERROR: CompleteAmountPayableTransactions => Unable to update due_amount in %s table: %w", sum.AccountType, err)
		}
		var current_balance int
		stmt = `
			UPDATE public.head_accounts
			SET amount_payable = amount_payable + $1, updated_at = $2
			WHERE id = $3 RETURNING current_balance`

		if sum.HeadAccount.ID == 0 {
			sum.HeadAccount.ID = 1
		}
		err = tx.QueryRowContext(ctx, stmt, sum.PayableAmount, time.Now(), sum.HeadAccount.ID).Scan(&current_balance)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("DBERROR: CompleteAmountPayableTransactions => Unable to update due_amount in head_accounts table: %w", err)
		}

		//insert to financial transactions
		stmt = `
			INSERT INTO public.financial_transactions(transaction_type, source_type, source_id, destination_type, destination_id, amount, transaction_date, voucher_no, description,carrier,cheque_no, current_balance)
			VALUES('Amount Receivable', $1, $2, 'head_accounts', $3, $4, TO_DATE($5, 'MM/DD/YYYY'), $6, $7, $8, $9,$10)
		`
		_, err = tx.ExecContext(ctx, stmt, sum.AccountType, sum.AccountID, sum.HeadAccount.ID, sum.PayableAmount, sum.Date, sum.VoucherNo, sum.Description, sum.Carrier, sum.ChequeNo, current_balance)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("DBERROR: CompleteAmountPayableTransactions => Unable Insert into financial_transaction table(type-2): %w", err)
		}
	}

	//commit transaction
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("DBERROR: CompleteAmountPayableTransactions => Unable to commit: %w", err)
	}
	return nil
}

func (p *postgresDBRepo) CompleteAmountReceivableTransactions(summary []*models.AmountReceivable) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	//Begin transaction
	tx, err := p.DB.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("DBERROR: CompleteAmountReceivableTransactions => Unable to begin transaction: %w", err)
	}

	for _, sum := range summary {
		//update customers/suppliers/employees accounts
		//update head_accounts
		stmt := `
			UPDATE public.` + sum.AccountType + `
			SET due_amount = due_amount + $1, updated_at = $2
			WHERE id = $3
		`
		_, err = tx.ExecContext(ctx, stmt, sum.ReceivableAmount, time.Now(), sum.AccountID)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("DBERROR: CompleteAmountReceivableTransactions => Unable to update due_amount in %s table: %w", sum.AccountType, err)
		}
		var current_balance int
		stmt = `
			UPDATE public.head_accounts
			SET amount_receivable = amount_receivable + $1, updated_at = $2
			WHERE id = $3 RETURNING current_balance`
		if sum.HeadAccount.ID == 0 {
			sum.HeadAccount.ID = 1
		}
		err = tx.QueryRowContext(ctx, stmt, sum.ReceivableAmount, time.Now(), sum.HeadAccount.ID).Scan(&current_balance)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("DBERROR: CompleteAmountReceivableTransactions => Unable to update head_accounts table: %w", err)
		}
		//insert to financial transactions
		stmt = `
			INSERT INTO public.financial_transactions(transaction_type, source_type, source_id, destination_type, destination_id, amount, transaction_date, voucher_no, description,carrier,cheque_no, current_balance)
			VALUES('Amount Receivable', $1, $2, 'head_accounts', $3, $4, TO_DATE($5,'MM/DD/YYYY'), $6, $7, $8, $9,$10)
		`
		_, err = tx.ExecContext(ctx, stmt, sum.AccountType, sum.AccountID, sum.HeadAccount.ID, sum.ReceivableAmount, sum.Date, sum.VoucherNo, sum.Description, sum.Carrier, sum.ChequeNo, current_balance)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("DBERROR: CompleteAmountReceivableTransactions => Unable Insert into financial_transaction table(type-2): %w", err)
		}
	}

	//commit transaction
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("DBERROR: CompleteAmountReceivableTransactions => Unable to commit: %w", err)
	}
	return nil
}

func (p *postgresDBRepo) CompleteExpensesTransactions(summary []*models.Expense) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	//Begin transaction
	tx, err := p.DB.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("DBERROR: CompleteExpensesTransactions => Unable to begin transaction: %w", err)
	}

	for _, sum := range summary {
		//update Cash-Bank account
		//set current_balance += received_amount
		var current_balance int
		stmt := `
			UPDATE public.head_accounts
			SET current_balance = current_balance - $1, updated_at = $2
			WHERE id = $3 RETURNING current_balance`

		err = tx.QueryRowContext(ctx, stmt, sum.ExpenseAmount, time.Now(), sum.SourceAccount.ID).Scan(&current_balance)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("DBERROR: CompleteExpensesTransactions => Unable to update current_balance in head_accounts table: %w", err)
		}

		//insert to financial transactions
		stmt = `
			INSERT INTO public.financial_transactions(transaction_type, source_type, source_id, destination_type, destination_id, amount, transaction_date, voucher_no, description,carrier,cheque_no, current_balance)
			VALUES('Expense', 'head_accounts', $1, 'expenses', 0, $2, TO_DATE($3, 'MM/DD/YYYY'), $4, $5, $6, $7, $8)
		`
		_, err = tx.ExecContext(ctx, stmt, sum.SourceAccount.ID, sum.ExpenseAmount, sum.ExpenseDate, sum.VoucherNo, sum.Description, sum.Carrier, sum.ChequeNo, current_balance)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("DBERROR: CompleteExpensesTransactions => Unable Insert into financial_transaction table: %w", err)
		}

		//insert to expense_history table
		stmt = `
			INSERT INTO public.expense_history(expense_type, voucher_no, account, expense_date, expense_amount, description, carrier, cheque_no)
			VALUES($1, $2, $3, TO_DATE($4,'MM/DD/YYY'), $5, $6, $7, $8)
		`
		_, err = tx.ExecContext(ctx, stmt, sum.ExpenseType.ID, sum.VoucherNo, sum.SourceAccount.ID, sum.ExpenseDate, sum.ExpenseAmount, sum.Description, sum.Carrier, sum.ChequeNo)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("DBERROR: CompleteExpensesTransactions => Unable Insert into expense_history table: %w", err)
		}

		exp_type := []string{
			"",
			"miscellaneous",
			"rent",
			"utilities",
			"salaries_and_wages",
			"advertising_and_promotions",
			"maintenance_and_repairs",
			"office_supplies",
			"insurance",
			"delivery_and_freight_charges",
			"depreciation",
			"taxes_and_licenses",
			"inventory_costs",
			"office_expense",
			"travel_expense",
			"training_and_development",
			"bank_charges_and_fees",
			"interest_on_loans",
			"software_and_subscriptions",
			"security_costs",
			"waste_disposal",
			"non_operating_income",
			"non_operating_expense",
		}
		expense_type := exp_type[sum.ExpenseType.ID]
		//update top_sheet
		stmt = fmt.Sprintf(`
		INSERT INTO public.top_sheet (
			sheet_date,
			total_expenses,
			%s,
			initial_stock_value
		)
		VALUES (
			 TO_DATE($1, 'MM/DD/YYYY'), $2, $3, COALESCE((SELECT initial_stock_value FROM public.top_sheet ORDER BY sheet_date DESC LIMIT 1),0)
		)
		ON CONFLICT (sheet_date)
		DO UPDATE SET
			total_expenses = public.top_sheet.total_expenses + EXCLUDED.total_expenses,
			%s = public.top_sheet.%s + EXCLUDED.%s,
			updated_at = CURRENT_TIMESTAMP;
		`, expense_type, expense_type, expense_type, expense_type)

		_, err = tx.ExecContext(ctx, stmt, sum.ExpenseDate, sum.ExpenseAmount, sum.ExpenseAmount)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("DBERROR: CompleteExpensesTransactions => failed to update or  Insert into top_sheet table: %w", err)
		}
	}

	//commit transaction
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("DBERROR: CompleteExpensesTransactions => Unable to commit: %w", err)
	}
	return nil
}

func (p *postgresDBRepo) CompleteFundAcquisitionProcess(summary []*models.FundAcquisition) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	//Begin transaction
	tx, err := p.DB.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("DBERROR: CompleteFundAcquisitionProcess => Unable to begin transaction: %w", err)
	}

	for _, sum := range summary {
		//update Source account
		//set current_balance -= received_amount
		stmt := `
			UPDATE public.company_stakeholders
			SET total_investment = total_investment + $1, updated_at = CURRENT_TIMESTAMP
			WHERE id = $2`

		_, err = tx.ExecContext(ctx, stmt, sum.TransferAmount, sum.SourceAccount.ID)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("DBERROR: CompleteFundAcquisitionProcess => Unable to update total_investment in company_stakeholders table: %w", err)
		}
		//update Destination account
		//set current_balance += received_amount
		var current_balance int
		stmt = `
			UPDATE public.head_accounts
			SET current_balance = current_balance + $1, updated_at = CURRENT_TIMESTAMP
			WHERE id = 1 RETURNING current_balance`

		err = tx.QueryRowContext(ctx, stmt, sum.TransferAmount).Scan(&current_balance)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("DBERROR: CompleteFundAcquisitionProcess => Unable to update current_balance in head_accounts table: %w", err)
		}
		//set current_balance += received_amount
		stmt = `
			UPDATE public.head_accounts
			SET current_balance = current_balance + $1, updated_at = CURRENT_TIMESTAMP
			WHERE id = $2`

		_, err = tx.ExecContext(ctx, stmt, sum.TransferAmount, sum.DestinationAccount.ID)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("DBERROR: CompleteFundAcquisitionProcess => Unable to update current_balance in head_accounts table: %w", err)
		}

		//insert to financial transactions
		stmt = `
			INSERT INTO public.financial_transactions(transaction_type, source_type, source_id, destination_type, destination_id, amount, transaction_date, voucher_no, description,carrier,cheque_no, current_balance)
			VALUES('Fund Acquisition', 'company_stakeholders', $1, 'head_accounts', $2, $3,  TO_DATE($4,'MM/DD/YYYY'), $5, $6, $7, $8,$9)
		`
		_, err = tx.ExecContext(ctx, stmt, sum.SourceAccount.ID, sum.DestinationAccount.ID, sum.TransferAmount, sum.TransactionDate, sum.VoucherNo, sum.Description, sum.Carrier, sum.ChequeNo, current_balance)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("DBERROR: CompleteFundAcquisitionProcess => Unable Insert into financial_transaction table: %w", err)
		}
	}

	//commit transaction
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("DBERROR: CompleteFundAcquisitionProcess => Unable to commit: %w", err)
	}
	return nil
}

// .......................Inventory Reports.......................
// GetCategoryListReport returns a list of all categories with detailed info from categories table
func (p *postgresDBRepo) GetCategoryListReport() ([]*models.Category, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var categories []*models.Category

	query := `
		SELECT * FROM public.categories ORDER BY id ASC
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
			&c.Status,
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

// GetBrandListReport returns a list of all brands with detailed info from brands table
func (p *postgresDBRepo) GetBrandListReport() ([]*models.Brand, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var brands []*models.Brand

	query := `
		SELECT * FROM public.brands
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
			&b.Status,
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

// GetProductListReport returns a list of all products with detailed info from products table
func (p *postgresDBRepo) GetProductListReport() ([]*models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var product []*models.Product

	query := `
		SELECT 
			i.id, i.product_code, i.product_name, i.product_description, i.product_status, i.quantity_purchased, i.purchase_cost, i.quantity_sold, i.sold_price, i.category_id, i.brand_id, i.stock_alert_level, i.created_at, i.updated_at, b.id, b.name, c.id, c.name
		FROM 
			public.products i
			INNER JOIN brands b ON (b.id = i.brand_id) 
			INNER JOIN categories c ON (c.id = i.category_id)
		ORDER BY c.name ASC
	`
	var rows *sql.Rows
	var err error

	rows, err = p.DB.QueryContext(ctx, query)
	if err != nil {
		return product, err
	}
	defer rows.Close()

	for rows.Next() {
		var p models.Product
		err = rows.Scan(
			&p.ID,
			&p.ProductCode,
			&p.ProductName,
			&p.Description,
			&p.ProductStatus,
			&p.QuantityPurchased,
			&p.PurchaseCost,
			&p.QuantitySold,
			&p.SoldPrice,
			&p.CategoryID,
			&p.BrandID,
			&p.StockAlertLevel,
			&p.CreatedAt,
			&p.UpdatedAt,
			&p.Brand.ID,
			&p.Brand.Name,
			&p.Category.ID,
			&p.Category.Name,
		)
		if err != nil {
			return product, err
		}
		product = append(product, &p)
	}
	return product, nil
}

// GetLowStockProductReport returns a list of all products with detailed info from products table where active status = true and stock_alert_level > quantity_purchased - quantity_sold
func (p *postgresDBRepo) GetLowStockProductReport() ([]*models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var product []*models.Product

	query := `
		SELECT 
			i.id, i.product_code, i.product_name, i.product_description, i.product_status, i.quantity_purchased, i.purchase_cost, i.quantity_sold, i.sold_price, i.category_id, i.brand_id, i.stock_alert_level, i.created_at, i.updated_at, b.id, b.name, c.id, c.name
		FROM 
			public.products i
			INNER JOIN brands b ON (b.id = i.brand_id) 
			INNER JOIN categories c ON (c.id = i.category_id)
		WHERE i.product_status = true AND i.stock_alert_level > i.quantity_purchased - i.quantity_sold
		ORDER BY c.name ASC
	`
	var rows *sql.Rows
	var err error

	rows, err = p.DB.QueryContext(ctx, query)
	if err != nil {
		return product, err
	}
	defer rows.Close()

	for rows.Next() {
		var p models.Product
		err = rows.Scan(
			&p.ID,
			&p.ProductCode,
			&p.ProductName,
			&p.Description,
			&p.ProductStatus,
			&p.QuantityPurchased,
			&p.PurchaseCost,
			&p.QuantitySold,
			&p.SoldPrice,
			&p.CategoryID,
			&p.BrandID,
			&p.StockAlertLevel,
			&p.CreatedAt,
			&p.UpdatedAt,
			&p.Brand.ID,
			&p.Brand.Name,
			&p.Category.ID,
			&p.Category.Name,
		)
		if err != nil {
			return product, err
		}
		product = append(product, &p)
	}
	return product, nil
}

// GetServiceListReport returns a list of all products with detailed info from products table
func (p *postgresDBRepo) GetServiceListReport() ([]*models.Service, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var service []*models.Service

	query := `
		SELECT * FROM public.services ORDER BY id ASC
	`
	var rows *sql.Rows
	var err error

	rows, err = p.DB.QueryContext(ctx, query)
	if err != nil {
		return service, err
	}
	defer rows.Close()

	for rows.Next() {
		var p models.Service
		err = rows.Scan(
			&p.ID,
			&p.ServiceCode,
			&p.ServiceName,
			&p.Description,
			&p.BaseFee,
			&p.Discount,
			&p.TrackRecord,
			&p.ServiceStatus,
			&p.CreatedAt,
			&p.UpdatedAt,
		)
		if err != nil {
			return service, err
		}
		service = append(service, &p)
	}
	return service, nil
}

// GetPurchaseHistoryReport returns a list of all purchases with detailed info from purchase_history table
func (p *postgresDBRepo) GetPurchaseHistoryReport() ([]*models.Purchase, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var purchase []*models.Purchase

	query := `
		SELECT
			ph.id, TO_CHAR(ph.purchase_date, 'MM/DD/YYYY') AS purchase_date_str, ph.supplier_id, ph.product_id, ph.account_id, ph.chalan_no, ph.memo_no, ph.note, ph.quantity_purchased, ph.bill_amount,
			ph.discount, ph.total_amount, ph.paid_amount, ph.created_at, ph.updated_at, s.account_name, s.account_code, 
			p.product_code,p.product_name, p.category_id, p.brand_id, pc.name, pb.name
		FROM
			public.purchase_history ph	
			INNER JOIN public.suppliers as s ON (s.id = ph.supplier_id)	
			INNER JOIN public.products as p ON(p.id = ph.product_id)
			INNER JOIN public.categories as pc ON(pc.id =p.category_id)	
			INNER JOIN public.brands as pb ON(pb.id = p.brand_id)
		ORDER BY id ASC
	`
	var rows *sql.Rows
	var err error

	rows, err = p.DB.QueryContext(ctx, query)
	if err != nil {
		return purchase, err
	}
	defer rows.Close()

	for rows.Next() {
		var ph models.Purchase
		err = rows.Scan(
			&ph.ID,
			&ph.PurchaseDate,
			&ph.Supplier.ID,
			&ph.Product.ID,
			&ph.HeadAccount.ID,
			&ph.ChalanNO,
			&ph.MemoNo,
			&ph.Note,
			&ph.QuantityPurchased,
			&ph.BillAmount,
			&ph.Discount,
			&ph.TotalAmount,
			&ph.PaidAmount,
			&ph.CreatedAt,
			&ph.UpdatedAt,
			&ph.Supplier.AccountName,
			&ph.Supplier.AccountCode,
			&ph.Product.ProductCode,
			&ph.Product.ProductName,
			&ph.Product.Category.ID,
			&ph.Product.Brand.ID,
			&ph.Product.Category.Name,
			&ph.Product.Brand.Name,
		)
		if err != nil {
			return purchase, err
		}
		purchase = append(purchase, &ph)
	}
	return purchase, nil
}

// GetSalesHistoryReport returns a list of all sales with detailed info from sales_history table
func (p *postgresDBRepo) GetSalesHistoryReport() ([]*models.Sale, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var sales []*models.Sale

	query := `
		SELECT
			sl.id, sl.sale_date, sl.customer_id, sl.account_id, sl.chalan_no, sl.memo_no, sl.note, sl.bill_amount,
			sl.discount, sl.total_amount, sl.paid_amount, sl.created_at, sl.updated_at, cm.id, cm.account_name, cm.account_code 
		FROM
			public.sales_history sl	
			INNER JOIN public.customers as cm ON (cm.id = sl.customer_id)	
		ORDER BY sl.id ASC
	`
	var rows *sql.Rows
	var err error

	rows, err = p.DB.QueryContext(ctx, query)
	if err != nil {
		return sales, err
	}
	defer rows.Close()

	for rows.Next() {
		var sl models.Sale
		err = rows.Scan(
			&sl.ID,
			&sl.SaleDate,
			&sl.Customer.ID,
			&sl.AccountID,
			&sl.ChalanNO,
			&sl.MemoNo,
			&sl.Note,
			&sl.BillAmount,
			&sl.Discount,
			&sl.TotalAmount,
			&sl.PaidAmount,
			&sl.CreatedAt,
			&sl.UpdatedAt,
			&sl.Customer.ID,
			&sl.Customer.AccountName,
			&sl.Customer.AccountCode,
		)
		if err != nil {
			return sales, err
		}
		sales = append(sales, &sl)
	}
	return sales, nil
}

// .......................Accounts Reports.......................
func (p *postgresDBRepo) GetCustomerDueHistoryReport() ([]*models.Sale, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var sales []*models.Sale

	query := `
		SELECT sl.sale_date, sl.customer_id, sl.memo_no, sl.bill_amount, sl.total_amount, sl.paid_amount, c.id, c.account_code, c.account_name, c.due_amount, c.mobile
		FROM public.sales_history as sl
		INNER JOIN public.customers as c ON (sl.customer_id = c.id)
		WHERE sl.total_amount > sl.paid_amount
	`
	rows, err := p.DB.QueryContext(ctx, query)
	if err != nil {
		return sales, fmt.Errorf("DBERROR: GetCustomerDueHistoryReport => %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var sale models.Sale
		err = rows.Scan(
			&sale.SaleDate,
			&sale.CustomerID,
			&sale.MemoNo,
			&sale.BillAmount,
			&sale.TotalAmount,
			&sale.PaidAmount,
			&sale.Customer.ID,
			&sale.Customer.AccountCode,
			&sale.Customer.AccountName,
			&sale.Customer.DueAmount,
			&sale.Customer.Mobile,
		)
		if err != nil {
			return sales, fmt.Errorf("DBERROR: GetCustomerDueHistoryReport => %w", err)
		}
		sales = append(sales, &sale)
	}
	return sales, nil
}
func (p *postgresDBRepo) GetTransactionsHistoryReport() ([]*models.Transaction, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var transactions []*models.Transaction

	query := `
		SELECT transaction_id, voucher_no, transaction_type, source_type, source_id, destination_type, destination_id, amount, current_balance, transaction_date, description
		FROM public.financial_transactions ORDER BY transaction_id DESC
	`
	rows, err := p.DB.QueryContext(ctx, query)
	if err != nil {
		return transactions, fmt.Errorf("DBERROR: GetTransactionsHistoryReport => %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var trx models.Transaction
		err = rows.Scan(
			&trx.ID,
			&trx.VoucherNo,
			&trx.TransactionType,
			&trx.SourceType,
			&trx.SourceID,
			&trx.DestinationType,
			&trx.DestinationID,
			&trx.Amount,
			&trx.CurrentBalance,
			&trx.TransactionDate,
			&trx.Description,
		)
		if err != nil {
			return transactions, fmt.Errorf("DBERROR: GetTransactionsHistoryReport => %w", err)
		}
		//retrieve source account name
		var account_name string
		query = fmt.Sprintf("SELECT account_name FROM public.%s WHERE id = $1", trx.SourceType)
		err = p.DB.QueryRowContext(ctx, query, trx.SourceID).Scan(&account_name)
		if err != nil {
			return transactions, fmt.Errorf("DBERROR: GetTransactionsHistoryReport (Unable to retrieve %s account name)=> %w", trx.SourceType, err)
		}
		trx.SourceAccountName = account_name
		if trx.DestinationType == "expenses" {
			trx.DestinationAccountName = trx.DestinationType
		} else {
			//retrieve destination account name
			query = fmt.Sprintf("SELECT account_name FROM public.%s WHERE id = $1", trx.DestinationType)
			err = p.DB.QueryRowContext(ctx, query, trx.DestinationID).Scan(&account_name)
			if err != nil {
				return transactions, fmt.Errorf("DBERROR: GetTransactionsHistoryReport (Unable to retrieve %s account name)=> %w", trx.DestinationType, err)
			}
			trx.DestinationAccountName = account_name
		}

		transactions = append(transactions, &trx)
	}
	return transactions, nil
}

func (p *postgresDBRepo) GetCashBankStatement() ([]*models.Transaction, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var transactions []*models.Transaction

	query := `
		SELECT transaction_id, voucher_no, transaction_type, source_type, source_id, destination_type, destination_id, amount, current_balance, transaction_date, description
		FROM public.financial_transactions 
		WHERE transaction_type IN('Refund','Repayment','Receive & Collection', 'Payment', 'Cash Transfer', 'Expense', 'Cash Adjustment')
		ORDER BY transaction_id DESC
	`
	rows, err := p.DB.QueryContext(ctx, query)
	if err != nil {
		return transactions, fmt.Errorf("DBERROR: GetCashBankStatement => %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var trx models.Transaction
		err = rows.Scan(
			&trx.ID,
			&trx.VoucherNo,
			&trx.TransactionType,
			&trx.SourceType,
			&trx.SourceID,
			&trx.DestinationType,
			&trx.DestinationID,
			&trx.Amount,
			&trx.CurrentBalance,
			&trx.TransactionDate,
			&trx.Description,
		)
		if err != nil {
			return transactions, fmt.Errorf("DBERROR: GetCashBankStatement => %w", err)
		}
		//retrieve source account name
		var account_name string
		query = fmt.Sprintf("SELECT account_name FROM public.%s WHERE id = $1", trx.SourceType)
		err = p.DB.QueryRowContext(ctx, query, trx.SourceID).Scan(&account_name)
		if err != nil {
			return transactions, fmt.Errorf("DBERROR: GetCashBankStatement (Unable to retrieve %s account name)=> %w", trx.SourceType, err)
		}
		trx.SourceAccountName = account_name
		if trx.DestinationType == "expenses" {
			trx.DestinationAccountName = trx.DestinationType
		} else {
			//retrieve destination account name
			query = fmt.Sprintf("SELECT account_name FROM public.%s WHERE id = $1", trx.DestinationType)
			err = p.DB.QueryRowContext(ctx, query, trx.DestinationID).Scan(&account_name)
			if err != nil {
				return transactions, fmt.Errorf("DBERROR: GetTransactionHistoryReport (Unable to retrieve %s account name)=> %w", trx.DestinationType, err)
			}
			trx.DestinationAccountName = account_name
		}

		transactions = append(transactions, &trx)
	}
	return transactions, nil
}

func (p *postgresDBRepo) GetLedgerBookDetails(account_type string, account_id int) ([]*models.Transaction, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var transactions []*models.Transaction

	query := `
		SELECT transaction_id, voucher_no, transaction_type, source_type, source_id, destination_type, destination_id, amount, current_balance, transaction_date, description
		FROM public.financial_transactions 
		WHERE (source_type = $1 AND source_id = $2) OR (destination_type = $1 AND destination_id = $2)
		ORDER BY transaction_id ASC
	`
	rows, err := p.DB.QueryContext(ctx, query, account_type, account_id)
	if err != nil {
		return transactions, fmt.Errorf("DBERROR: GetLedgerBookDetails => %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var trx models.Transaction
		err = rows.Scan(
			&trx.ID,
			&trx.VoucherNo,
			&trx.TransactionType,
			&trx.SourceType,
			&trx.SourceID,
			&trx.DestinationType,
			&trx.DestinationID,
			&trx.Amount,
			&trx.CurrentBalance,
			&trx.TransactionDate,
			&trx.Description,
		)
		if err != nil {
			return transactions, fmt.Errorf("DBERROR: GetLedgerBookDetails => %w", err)
		}
		//retrieve source account name
		var account_name string
		query = fmt.Sprintf("SELECT account_name FROM public.%s WHERE id = $1", trx.SourceType)
		err = p.DB.QueryRowContext(ctx, query, trx.SourceID).Scan(&account_name)
		if err != nil {
			return transactions, fmt.Errorf("DBERROR: GetLedgerBookDetails (Unable to retrieve %s account name)=> %w", trx.SourceType, err)
		}
		trx.SourceAccountName = account_name
		//retrieve source account name
		query = fmt.Sprintf("SELECT account_name FROM public.%s WHERE id = $1", trx.DestinationType)
		err = p.DB.QueryRowContext(ctx, query, trx.DestinationID).Scan(&account_name)
		if err != nil {
			return transactions, fmt.Errorf("DBERROR: GetLedgerBookDetails (Unable to retrieve %s account name)=> %w", trx.DestinationType, err)
		}
		trx.DestinationAccountName = account_name

		transactions = append(transactions, &trx)
	}
	return transactions, nil
}

func (p *postgresDBRepo) GetExpensesHistoryReport() ([]*models.Transaction, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var transactions []*models.Transaction

	query := `
		SELECT transaction_id, voucher_no, transaction_type, source_type, source_id, destination_type, destination_id, amount, current_balance, transaction_date, description
		FROM public.financial_transactions 
		WHERE transaction_type IN ('Payment','Expense')
		ORDER BY transaction_id DESC

	`

	rows, err := p.DB.QueryContext(ctx, query)
	if err != nil {
		return transactions, fmt.Errorf("DBERROR: GetExpensesHistoryReport => %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var trx models.Transaction
		err = rows.Scan(
			&trx.ID,
			&trx.VoucherNo,
			&trx.TransactionType,
			&trx.SourceType,
			&trx.SourceID,
			&trx.DestinationType,
			&trx.DestinationID,
			&trx.Amount,
			&trx.CurrentBalance,
			&trx.TransactionDate,
			&trx.Description,
		)
		if err != nil {
			return transactions, fmt.Errorf("DBERROR: GetExpensesHistoryReport => %w", err)
		}
		//retrieve source account name
		var account_name string
		query = fmt.Sprintf("SELECT account_name FROM public.%s WHERE id = $1", trx.SourceType)
		err = p.DB.QueryRowContext(ctx, query, trx.SourceID).Scan(&account_name)
		if err != nil {
			return transactions, fmt.Errorf("DBERROR: GetExpensesHistoryReport (Unable to retrieve %s account name)=> %w", trx.SourceType, err)
		}
		trx.SourceAccountName = account_name
		if trx.DestinationType == "expenses" {
			trx.DestinationAccountName = trx.DestinationType
		} else {
			//retrieve destination account name
			query = fmt.Sprintf("SELECT account_name FROM public.%s WHERE id = $1", trx.DestinationType)
			err = p.DB.QueryRowContext(ctx, query, trx.DestinationID).Scan(&account_name)
			if err != nil {
				return transactions, fmt.Errorf("DBERROR: GetExpensesHistoryReport (Unable to retrieve %s account name)=> %w", trx.DestinationType, err)
			}
			trx.DestinationAccountName = account_name
		}

		transactions = append(transactions, &trx)
	}
	return transactions, nil
}

func (p *postgresDBRepo) GetIncomeStatementData(startDate, endDate string) (models.IncomeStatement, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var ins models.IncomeStatement

	query := `
		WITH aggregated AS (
			SELECT 
				SUM(COALESCE(total_purchases, 0)) AS sum_total_purchases,
				SUM(COALESCE(total_sales, 0)) AS sum_total_sales,
				SUM(COALESCE(total_purchase_returns, 0)) AS sum_total_purchase_returns,
				SUM(COALESCE(total_sale_returns, 0)) AS sum_total_sale_returns,
				SUM(COALESCE(purchases_discount, 0)) AS sum_purchases_discount,
				SUM(COALESCE(sales_discount, 0)) AS sum_sales_discount,
				SUM(COALESCE(miscellaneous, 0)) AS sum_miscellaneous,
				SUM(COALESCE(rent, 0)) AS sum_rent,
				SUM(COALESCE(utilities, 0)) AS sum_utilities,
				SUM(COALESCE(salaries_and_wages, 0)) AS sum_salaries_and_wages,
				SUM(COALESCE(advertising_and_promotions, 0)) AS sum_advertising_and_promotions,
				SUM(COALESCE(maintenance_and_repairs, 0)) AS sum_maintenance_and_repairs,
				SUM(COALESCE(office_supplies, 0)) AS sum_office_supplies,
				SUM(COALESCE(insurance, 0)) AS sum_insurance,
				SUM(COALESCE(delivery_and_freight_charges, 0)) AS sum_delivery_and_freight_charges,
				SUM(COALESCE(depreciation, 0)) AS sum_depreciation,
				SUM(COALESCE(taxes_and_licenses, 0)) AS sum_taxes_and_licenses,
				SUM(COALESCE(inventory_costs, 0)) AS sum_inventory_costs,
				SUM(COALESCE(office_expense, 0)) AS sum_office_expense,
				SUM(COALESCE(travel_expense, 0)) AS sum_travel_expense,
				SUM(COALESCE(training_and_development, 0)) AS sum_training_and_development,
				SUM(COALESCE(bank_charges_and_fees, 0)) AS sum_bank_charges_and_fees,
				SUM(COALESCE(interest_on_loans, 0)) AS sum_interest_on_loans,
				SUM(COALESCE(software_and_subscriptions, 0)) AS sum_software_and_subscriptions,
				SUM(COALESCE(security_costs, 0)) AS sum_security_costs,
				SUM(COALESCE(waste_disposal, 0)) AS sum_waste_disposal,
				SUM(COALESCE(non_operating_income, 0)) AS sum_non_operating_income,
				SUM(COALESCE(non_operating_expense, 0)) AS sum_non_operating_expense
			FROM public.top_sheet
			WHERE sheet_date BETWEEN TO_DATE($1, 'MM/DD/YYYY') AND TO_DATE($2, 'MM/DD/YYYY')
		),
		last_value AS (
			SELECT initial_stock_value AS last_initial_stock_value
			FROM public.top_sheet
			WHERE sheet_date BETWEEN TO_DATE($1, 'MM/DD/YYYY') AND TO_DATE($2, 'MM/DD/YYYY')
			ORDER BY sheet_date DESC
			LIMIT 1
		)
		SELECT 
			aggregated.*,
			last_value.last_initial_stock_value
		FROM aggregated, last_value;
`

	err := p.DB.QueryRowContext(ctx, query, startDate, endDate).Scan(
		&ins.GoodsPurchased,
		&ins.GrossSales,
		&ins.PurchaseReturn,
		&ins.SalesReturn,
		&ins.PurchaseDiscount,
		&ins.SalesDiscount,
		&ins.ExpenseSection.Miscellaneous,
		&ins.ExpenseSection.Rent,
		&ins.ExpenseSection.Utilities,
		&ins.ExpenseSection.SalariesAndWages,
		&ins.ExpenseSection.AdvertisingAndPromotions,
		&ins.ExpenseSection.MaintenanceAndRepairs,
		&ins.ExpenseSection.OfficeSupplies,
		&ins.ExpenseSection.Insurance,
		&ins.ExpenseSection.DeliveryAndFreightCharges,
		&ins.ExpenseSection.Depreciation,
		&ins.ExpenseSection.TaxesAndLicenses,
		&ins.ExpenseSection.InventoryCosts,
		&ins.ExpenseSection.OfficeExpense,
		&ins.ExpenseSection.TravelExpense,
		&ins.ExpenseSection.TrainingAndDevelopment,
		&ins.ExpenseSection.BankChargesAndFees,
		&ins.ExpenseSection.InterestOnLoans,
		&ins.ExpenseSection.SoftwareAndSubscriptions,
		&ins.ExpenseSection.SecurityCosts,
		&ins.ExpenseSection.WasteDisposal,
		&ins.NonOperatingIncome,
		&ins.NonOperatingExpense,
		&ins.TotalAvailableGoods,
	)
	if err != nil {
		return ins, fmt.Errorf("DBERROR: GetIncomeStatementData => %w", err)
	}

	return ins, nil
}

func (p *postgresDBRepo) GetTopSheetReport() ([]*models.TopSheet, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var top_sheet []*models.TopSheet

	query := `
		SELECT id, sheet_date, total_purchases, total_payments, total_sales, total_received_payments, total_purchase_returns, total_sale_returns, purchases_discount, sales_discount, total_expenses, created_at, updated_at
		FROM public.top_sheet 
		ORDER BY id ASC
	`

	rows, err := p.DB.QueryContext(ctx, query)
	if err != nil {
		return top_sheet, fmt.Errorf("DBERROR: GetTopSheetReport => %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var ts models.TopSheet
		err = rows.Scan(
			&ts.ID,
			&ts.SheetDate,
			&ts.TotalPurchases,
			&ts.TotalPayments,
			&ts.TotalSales,
			&ts.TotalReceivedPayments,
			&ts.TotalPurchaseReturns,
			&ts.TotalSaleReturns,
			&ts.PurchasesDiscount,
			&ts.SalesDiscount,
			&ts.TotalExpenses,
			&ts.CreatedAt,
			&ts.UpdatedAt,
		)
		if err != nil {
			return top_sheet, fmt.Errorf("DBERROR: GetTopSheetReport => %w", err)
		}
		top_sheet = append(top_sheet, &ts)
	}
	return top_sheet, nil
}

func (p *postgresDBRepo) GetTrialBalanceReport() (models.TrialBalance, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var trialBalanceSheet models.TrialBalance
	// CapitalInvestment       int       `json:"capital_investment"`
	// CashBankAccounts        int       `json:"cash_bank_accounts"`
	// CurrentAssets           int       `json:"current_assets"`
	// CurrentLiabilities      int       `json:"current_liabilities"`
	// CustomerAccountsReceivable int       `json:"customer_accounts_payable"`
	// SupplierAccountsPayable int       `json:"supplier_accounts_payable"`
	// EmployeeAccountsPayable      int       `json:"employee_accounts_payable"`
	// ExpenseAccounts         int       `json:"expense_accounts"`
	// FixedAssets             int       `json:"fixed_assets"`
	// LoanAccounts            int       `json:"loan_accounts"`
	// RevenueAccounts         int       `json:"revenue_accounts"`

	var val int
	// CapitalInvestment
	err := p.DB.QueryRowContext(ctx, "SELECT COALESCE(SUM(current_balance), 0) FROM head_accounts WHERE account_type = 'CAPITAL ACCOUNTS'").Scan(&val)
	if err != nil {
		return trialBalanceSheet, fmt.Errorf("DBERROR: GetBalanceSheetReport => unable to retrieve CASH & BANK ACCOUNTS balance: %w", err)
	}
	trialBalanceSheet.CapitalInvestment = val
	// Cash & Bank
	err = p.DB.QueryRowContext(ctx, "SELECT COALESCE(SUM(current_balance), 0) FROM head_accounts WHERE account_type = 'CASH & BANK ACCOUNTS'").Scan(&val)
	if err != nil {
		return trialBalanceSheet, fmt.Errorf("DBERROR: GetBalanceSheetReport => unable to retrieve CAPITAL ACCOUNTS balance: %w", err)
	}
	trialBalanceSheet.CashBankAccounts = val
	// Current Assets
	err = p.DB.QueryRowContext(ctx, "SELECT COALESCE(SUM(current_balance), 0) FROM head_accounts WHERE account_type = 'CURRENT ASSETS'").Scan(&val)
	if err != nil {
		return trialBalanceSheet, fmt.Errorf("DBERROR: GetBalanceSheetReport => unable to retrieve CURRENT ASSETS balance: %w", err)
	}
	trialBalanceSheet.CurrentAssets = val

	// Supplier Accounts Payable
	err = p.DB.QueryRowContext(ctx, "SELECT COALESCE(SUM(due_amount), 0) FROM suppliers").Scan(&val)
	if err != nil {
		return trialBalanceSheet, fmt.Errorf("DBERROR: GetBalanceSheetReport => unable to retrieve Suppliers account payable: %w", err)
	}
	trialBalanceSheet.SupplierAccountsPayable = val
	// Customer Accounts Receivable
	err = p.DB.QueryRowContext(ctx, "SELECT COALESCE(SUM(due_amount), 0) FROM customers").Scan(&val)
	if err != nil {
		return trialBalanceSheet, fmt.Errorf("DBERROR: GetBalanceSheetReport => unable to retrieve customer account receivable: %w", err)
	}
	trialBalanceSheet.CustomerAccountsReceivable = val
	// Employee Accounts Payable
	err = p.DB.QueryRowContext(ctx, "SELECT COALESCE(SUM(due_amount), 0) FROM employees").Scan(&val)
	if err != nil {
		return trialBalanceSheet, fmt.Errorf("DBERROR: GetBalanceSheetReport => unable to retrieve employee account payable: %w", err)
	}
	trialBalanceSheet.EmployeeAccountsPayable = val
	// EXPENSE ACCOUNTS
	//product purchase + employee salary + other expense
	err = p.DB.QueryRowContext(ctx, "SELECT COALESCE(SUM(current_balance), 0) FROM head_accounts WHERE account_type = 'EXPENSE ACCOUNTS'").Scan(&val)
	if err != nil {
		return trialBalanceSheet, fmt.Errorf("DBERROR: GetBalanceSheetReport => unable to retrieve EXPENSE ACCOUNTS: %w", err)
	}
	trialBalanceSheet.ExpenseAccounts = val
	// Fixed Assets
	err = p.DB.QueryRowContext(ctx, "SELECT COALESCE(SUM(current_balance), 0) FROM head_accounts WHERE account_type = 'FIXED ASSETS'").Scan(&val)
	if err != nil {
		return trialBalanceSheet, fmt.Errorf("DBERROR: GetBalanceSheetReport => unable to retrieve FIXED ASSETS balance: %w", err)
	}
	trialBalanceSheet.FixedAssets = val
	// LOAN ACCOUNTS
	err = p.DB.QueryRowContext(ctx, "SELECT COALESCE(SUM(current_balance), 0) FROM head_accounts WHERE account_type = 'LOAN ACCOUNTS'").Scan(&val)
	if err != nil {
		return trialBalanceSheet, fmt.Errorf("DBERROR: GetBalanceSheetReport => unable to retrieve LOAN ACCOUNTS balance: %w", err)
	}
	trialBalanceSheet.LoanAccounts = val
	// REVENUE ACCOUNTS
	err = p.DB.QueryRowContext(ctx, "SELECT COALESCE(SUM(current_balance), 0) FROM head_accounts WHERE account_type = 'REVENUE ACCOUNTS'").Scan(&val)
	if err != nil {
		return trialBalanceSheet, fmt.Errorf("DBERROR: GetBalanceSheetReport => unable to retrieve REVENUE ACCOUNTS balance: %w", err)
	}
	trialBalanceSheet.RevenueAccounts = val

	return trialBalanceSheet, nil
}
func (p *postgresDBRepo) GetBalanceSheetReport() (models.BalanceSheet, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var balanceSheet models.BalanceSheet

	var val int
	// Cash & Bank
	err := p.DB.QueryRowContext(ctx, "SELECT COALESCE(SUM(current_balance), 0) FROM head_accounts WHERE account_type = 'CASH & BANK ACCOUNTS'").Scan(&val)
	if err != nil {
		return balanceSheet, fmt.Errorf("DBERROR: GetBalanceSheetReport => unable to retrieve CASH & BANK ACCOUNTS balance: %w", err)
	}
	balanceSheet.CashBankAccounts = val

	// Inventory Products
	err = p.DB.QueryRowContext(ctx, "SELECT COALESCE(initial_stock_value, 0) FROM top_sheet ORDER BY id DESC LIMIT 1").Scan(&val)
	if err != nil {
		return balanceSheet, fmt.Errorf("DBERROR: GetBalanceSheetReport => unable to retrieve current_assets(Inventory Stock Value): %w", err)
	}
	balanceSheet.CurrentAssets = val

	// Account Receivable(supplier)
	err = p.DB.QueryRowContext(ctx, "SELECT COALESCE(SUM(due_amount), 0) FROM suppliers WHERE due_amount > 0").Scan(&val)
	if err != nil {
		return balanceSheet, fmt.Errorf("DBERROR: GetBalanceSheetReport => unable to retrieve Supplies accounts receivable: %w", err)
	}
	balanceSheet.SupplierAccountsReceivable = val

	// Account Receivable(customer)
	err = p.DB.QueryRowContext(ctx, "SELECT COALESCE(SUM(due_amount), 0) FROM customers WHERE due_amount > 0").Scan(&val)
	if err != nil {
		return balanceSheet, fmt.Errorf("DBERROR: GetBalanceSheetReport => unable to retrieve Customers accounts receivable: %w", err)
	}
	balanceSheet.CustomerAccountsReceivable = val

	// Capital Investment
	err = p.DB.QueryRowContext(ctx, "SELECT COALESCE(SUM(current_balance), 0) FROM head_accounts WHERE account_type = 'CAPITAL ACCOUNTS'").Scan(&val)
	if err != nil {
		return balanceSheet, fmt.Errorf("DBERROR: GetBalanceSheetReport => unable to retrieve current liabilities: %w", err)
	}
	balanceSheet.CapitalInvestment = val

	// Capital Investment
	err = p.DB.QueryRowContext(ctx, "SELECT COALESCE(SUM(current_balance), 0) FROM head_accounts WHERE account_type = 'LOAN ACCOUNTS'").Scan(&val)
	if err != nil {
		return balanceSheet, fmt.Errorf("DBERROR: GetBalanceSheetReport => unable to retrieve current liabilities: %w", err)
	}
	balanceSheet.LoanAccounts = val

	// Account Payable(supplier)
	err = p.DB.QueryRowContext(ctx, "SELECT COALESCE(SUM(due_amount), 0) FROM suppliers WHERE due_amount < 0").Scan(&val)
	if err != nil {
		return balanceSheet, fmt.Errorf("DBERROR: GetBalanceSheetReport => unable to retrieve Suppliers account payable: %w", err)
	}
	balanceSheet.SupplierAccountsPayable = val

	// Account Payable(customer)
	err = p.DB.QueryRowContext(ctx, "SELECT COALESCE(SUM(due_amount), 0) FROM customers WHERE due_amount < 0").Scan(&val)
	if err != nil {
		return balanceSheet, fmt.Errorf("DBERROR: GetBalanceSheetReport => unable to retrieve Customers account payable: %w", err)
	}
	balanceSheet.CustomerAccountsPayable = val

	return balanceSheet, nil
}

// Helper functions
// LastIndex returns the last index of a given database table
func (p *postgresDBRepo) LastIndex(tableName string) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var lastID sql.NullInt64
	query := "SELECT MAX(id) AS last_id FROM " + tableName
	err := p.DB.QueryRowContext(ctx, query).Scan(&lastID)

	// If lastID is NULL, set it to 0
	id := int64(0) // Default to 0
	if lastID.Valid {
		id = lastID.Int64
	}
	return id, err
}

func (p *postgresDBRepo) GetCompanyProfile() (models.CompanyProfile, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var companyInfo models.CompanyProfile

	query := "SELECT * FROM public.company_profile WHERE id = 1"
	err := p.DB.QueryRowContext(ctx, query).Scan(
		&companyInfo.ID,
		&companyInfo.Name,
		&companyInfo.Description,
		&companyInfo.Slogan,
		&companyInfo.Mobile,
		&companyInfo.WhatsappAccount,
		&companyInfo.Telephone,
		&companyInfo.Email,
		&companyInfo.Division,
		&companyInfo.District,
		&companyInfo.Upazila,
		&companyInfo.Area,
		&companyInfo.PostalCode,
		&companyInfo.LogoLink,
		&companyInfo.OpeningDate,
		&companyInfo.CreatedAt,
		&companyInfo.UpdatedAt,
	)
	return companyInfo, err
}
