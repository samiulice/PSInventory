package dbrepo

import (
	"PSInventory/internal/models"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

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

// GetEmployeeDetails retrieves detailed info about an employee
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
			id, account_code, account_name, contact_person, division, district, upazila, area, mobile, email, account_status, discount, opening_balance, joining_date, created_at, updated_at
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
			id, account_code, account_name, contact_person, division, district, upazila, area, mobile, email, account_status, discount, joining_date, created_at, updated_at
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

	stmt := `INSERT INTO public.products (product_code, product_name, product_description, product_status, quantity_purchased, quantity_sold, category_id, brand_id, discount, created_at, updated_at) 
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
			i.id, i.product_code, i.product_name, i.product_description, i.product_status, i.quantity_purchased, i.quantity_sold, i.category_id, i.brand_id, i.discount, i.created_at, i.updated_at, b.id, b.name, c.id, c.name
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
			i.id, i.product_code, i.product_name, i.quantity_purchased, i.quantity_sold, i.category_id, i.discount, b.id, b.name, c.id, c.name
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
			i.id, i.product_code, i.product_name, i.quantity_purchased, i.quantity_sold, i.category_id, i.brand_id, i.discount, b.id, b.name, c.id, c.name
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
			i.id, i.product_code, i.product_name, i.product_description, i.product_status, i.quantity_purchased, i.quantity_sold, i.category_id, i.brand_id, i.discount, i.created_at, i.updated_at, b.id, b.name, c.id, c.name
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
			i.id, i.product_code, i.product_name, i.product_description, i.product_status, i.quantity_purchased, i.quantity_sold, i.category_id, i.brand_id, i.discount, i.created_at, i.updated_at, b.id, b.name, c.id, c.name
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

// GetAllProductsByCategoryID returns a list of all product that related to category ID from the database
func (p *postgresDBRepo) GetAllProductsByCategoryID(cat_id int) ([]*models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var products []*models.Product

	query := `
		SELECT 
			i.id, i.product_code, i.product_name, i.product_description, i.product_status, i.category_id, i.brand_id, i.discount, i.created_at, i.updated_at, b.id, b.name, c.id, c.name
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

// GetInStockProductItemsListByProductID returns a instock list of product items corresponds to productID
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
			ph.id, ph.purchase_date, ph.supplier_id, ph.product_id, ph.account_id, ph.chalan_no, ph.memo_no, ph.note, ph.bill_amount, ph.discount, ph.total_amount, ph.paid_amount, ph.created_at, ph.updated_at
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
	log.Println(purchase.Quantity)

	//Tx-1: Update product quantity
	//Set quantity += newQuantity
	query := `UPDATE public.products
          SET quantity_purchased = quantity_purchased + $1, purchase_cost = purchase_cost+$2 
          WHERE id = $3;`

	// Execute the query with parameters
	_, err = tx.ExecContext(ctx, query, purchase.QuantityPurchased, purchase.TotalAmount, purchase.Product.ID)
	if err != nil {
		tx.Rollback()
		return errors.New("SQLErrorRestockProduct(Update Quantity): " + err.Error())
	}

	//Tx-2: Update total_amount, due_amount(if available) in suppliers table
	query = `
		UPDATE Public.suppliers
		SET total_amount = total_amount + $1, due_amount = due_amount + $2
		WHERE id = $3
	`
	_, err = tx.ExecContext(ctx, query, purchase.TotalAmount, purchase.TotalAmount-purchase.PaidAmount, purchase.Supplier.ID)
	if err != nil {
		tx.Rollback()
		return errors.New("SQLErrorRestockProduct(Update suppliers):" + err.Error())
	}

	//Tx-3: Insert data to purchase history table
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
	var purchase_id int
	query = `INSERT INTO public.purchase_history (purchase_date,supplier_id,product_id,account_id,chalan_no,memo_no,note,bill_amount,discount,total_amount,paid_amount,created_at,updated_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) RETURNING id
	`
	row := tx.QueryRowContext(ctx, query,
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
	)
	if err = row.Scan(&purchase_id); err != nil {
		tx.Rollback()
		return errors.New("SQLErrorRestockProduct(Insert purchase_history):" + err.Error())
	}

	//Tx-4: store financial_transactions
	var financial_transactions_id int
	query = `INSERT INTO public.financial_transactions (transaction_type,source_type,source_id,destination_type,destination_id,amount,transaction_date,description,created_at,updated_at)
	VALUES ('Payment','Account',$1, 'Supplier', $2, $3, $4, $5, $6, $7) RETURNING transaction_id
	`
	purchaseDate, err := time.Parse("01/02/2006", purchase.PurchaseDate)
	if err != nil {
		tx.Rollback()
		return errors.New("SQLErrorRestockProduct(unable to parse time from string):" + err.Error())
	}
	purchaseDescription := "cash payment / Bank Transfer"
	row = tx.QueryRowContext(ctx, query,
		&purchase.HeadAccount.ID,
		&purchase.Supplier.ID,
		&purchase.PaidAmount,
		&purchaseDate,
		&purchaseDescription,
		time.Now(),
		time.Now(),
	)
	if err = row.Scan(&financial_transactions_id); err != nil {
		tx.Rollback()
		return errors.New("SQLErrorRestockProduct(Insert financial_transactions):" + err.Error())
	}
	//Tx-5: store product serial numbers
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
		values = append(values, fmt.Sprintf("('%s',%d,%d,0,%d,%d,%d,'%s','%s')", serial_number, purchase.Product.ID, purchase_id, purchase.MaxRetailPrice, int(purchase.PurchaseRate-purchase.PurchaseRate*purchase.Discount/100), purchase.WarrantyPeriod, createdAt, updatedAt))
	}

	query = "INSERT INTO public.product_serial_numbers (serial_number,product_id,purchase_history_id,sales_history_id,max_retail_price,purchase_rate,warranty_period,created_at,updated_at) VALUES " + strings.Join(values, ",") + ";"
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
func (p *postgresDBRepo) SaleProducts(sale *models.SalesInvoice) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//get the last index of sales_history table
	id, err := p.LastIndex("sales_history")
	if err != nil {
		return err
	}
	//append id to the memo_no
	sale.MemoNo = sale.MemoNo + strconv.FormatInt(id+int64(1), 10)

	tx, err := p.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	//step-1: Insert sales details to sales history table
	//step-2: Update product quantity, Set quantity -= newQuantity where id = {affected row id}
	//step-3: update product items status and sales_history_id, returning id affected row id

	//step-1: Insert sales details to sales history table
	var sale_id int
	query := `
		INSERT INTO public.sales_history (sale_date,customer_id,account_id,chalan_no,memo_no,note,bill_amount,discount,total_amount,paid_amount,created_at,updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) RETURNING id
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
		time.Now(),
		time.Now(),
	)
	if err = row.Scan(&sale_id); err != nil {
		tx.Rollback()
		return errors.New("SQLErrorSaleProducts(INSERT sales_history): " + err.Error())
	}

	totalQuantity := 0
	totalPrice := 0
	//loop over the SelectedProduct array
	for i, items := range sale.ProductItems {
		//Step-2: Update product quantity
		query := `
			UPDATE public.products
			SET quantity_sold = quantity_sold + $1, sold_price = sold_price + $2
			WHERE id = $3;
		  `
		_, err = tx.ExecContext(ctx, query, items.Quantity, items.SubTotal, items.ProductID)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("SQLErrorSaleProducts(Update Product Quantity):#%d --%w", i, err)
		}
		totalQuantity += items.Quantity
		totalPrice += items.SubTotal
		//Step-3: update product items status and sales_history_id
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

	//Tx-2: Update total_amount, due_amount(if available) in customers table
	query = `
		UPDATE Public.customers
		SET total_amount = total_amount + $1, due_amount = due_amount + $2
		WHERE id = $3
	`
	_, err = tx.ExecContext(ctx, query, sale.TotalAmount, sale.TotalAmount-sale.PaidAmount, sale.CustomerInfo.ID)
	if err != nil {
		tx.Rollback()
		return errors.New("SQLErrorSaleProducts(Update cutomers):" + err.Error())
	}

	//Tx-3 insert summary about the sales in the inventory_transaction_logs table
	var inv_tx_log_id int
	query = `INSERT INTO public.inventory_transaction_logs(job_id, transaction_type, quantity, price, transaction_date, created_at, updated_at)
	VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING id`
	err = tx.QueryRowContext(ctx, query, "JOB-"+sale.MemoNo+strconv.Itoa(sale_id), "sale", totalQuantity, sale.TotalAmount, sale.SaleDate, time.Now(), time.Now()).Scan(&inv_tx_log_id)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("SQLErrorSaleProducts(INSERT inventory_transaction_logs: %w", err)
	}

	//Tx-4: store financial_transactions
	var financial_transactions_id int
	query = `INSERT INTO public.financial_transactions (transaction_type,source_type,source_id,destination_type,destination_id,amount,transaction_date,description,created_at,updated_at)
	VALUES ('Receive & Collection','Customer',$1, 'Account', $2, $3, $4, $5, $6, $7) RETURNING transaction_id
	`
	saleDate, err := time.Parse("01/02/2006", sale.SaleDate)
	if err != nil {
		tx.Rollback()
		return errors.New("SQLErrorSaleProducts(unable to parse time from string):" + err.Error())
	}
	description := "cash Sale / Bank Transfer"
	row = tx.QueryRowContext(ctx, query,
		&sale.CustomerInfo.ID,
		&sale.HeadAccountInfo.ID,
		&sale.PaidAmount,
		&saleDate,
		&description,
		time.Now(),
		time.Now(),
	)
	if err = row.Scan(&financial_transactions_id); err != nil {
		tx.Rollback()
		return errors.New("SQLErrorSaleProducts(Insert financial_transactions):" + err.Error())
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		return errors.New("SQLErrorSaleProducts(Commit):" + err.Error())
	}
	return nil
}

// ReturnProductUnitsToSupplier updates database
func (p *postgresDBRepo) ReturnProductUnitsToSupplier(PurchaseHistory models.Purchase, JobID string, transactionDate string, ProductUnitsID []int, TotalUnits int, TotalPrices int) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var id int
	tx, err := p.DB.BeginTx(ctx, nil)
	if err != nil {
		return id, fmt.Errorf("failed to start transaction: %w", err)
	}

	// Prepare the SQL update query
	q1 := `DELETE FROM public.product_serial_numbers WHERE id = $1 returning product_id, purchase_rate`
	q2 := `UPDATE public.products SET quantity_purchased = quantity_purchased - 1, purchase_cost = purchase_cost-$1 WHERE id = $2`

	// Execute updates within the transaction
	for _, unitsID := range ProductUnitsID {
		var productItemID, purchase_rate int
		err := tx.QueryRowContext(ctx, q1, unitsID).Scan(&productItemID, &purchase_rate)
		if err != nil {
			tx.Rollback() // Rollback on error
			return id, fmt.Errorf("failed to update record in product_serial_numbers table with id %d: %w", unitsID, err)
		}
		//update product amount in products table
		_, err = tx.ExecContext(ctx, q2, purchase_rate, productItemID)
		if err != nil {
			tx.Rollback() // Rollback on error
			return id, fmt.Errorf("failed to update record in products table with id %d: %w", unitsID, err)
		}
	}

	//Update total_amount, due_amount(if available) in suppliers table
	query := `
		UPDATE Public.suppliers
		SET total_amount = total_amount - $1, due_amount = due_amount - $2
		WHERE id = $3
	`
	_, err = tx.ExecContext(ctx, query, TotalPrices, PurchaseHistory.TotalAmount-PurchaseHistory.PaidAmount, PurchaseHistory.Supplier.ID)
	if err != nil {
		tx.Rollback()
		return id, fmt.Errorf("SQLErrorRestockProduct(Update suppliers): %w", err)
	}

	//Insert Financial transaction
	var financial_transactions_id int
	query = `INSERT INTO public.financial_transactions (transaction_type,source_type,source_id,destination_type,destination_id,amount,transaction_date,description,created_at,updated_at)
	VALUES ('Cash Transfer','Supplier',$1, 'Account', $2, $3, $4, $5, $6, $7) RETURNING transaction_id
	`
	description := "cash return due to returning product to the supplier"
	err = tx.QueryRowContext(ctx, query,
		&PurchaseHistory.Supplier.ID,
		&PurchaseHistory.HeadAccount.ID,
		&TotalPrices,
		time.Now(),
		&description,
		time.Now(),
		time.Now(),
	).Scan(&financial_transactions_id)
	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("SQLErrorRestockProduct(Insert financial_transactions): %w", err)
	}

	//Insert inventory transaction logs
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
			return fmt.Errorf("DBERROR:=>SaleReturnDB: Unable to execute UPDATE in products table SQL %w", err)
		}
	}

	//step:2 insert sale return data to the sale_return_history
	//get the last index of sales_return_history Table
	lastID, err := p.LastIndex("sales_return_history")
	if err != nil {
		return fmt.Errorf("DBERROR:=>SaleReturnDB: Unable to get last index of sales_return_history table:  %w", err)
	}
	MemoNo = MemoNo + strconv.FormatInt(lastID+1, 10) //update memo no
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
		return fmt.Errorf("DBERROR:=>SaleReturnDB: Unable to execute UPDATE SQL %w", err)
	}
	if n, err := result.RowsAffected(); err != nil || n != 1 {
		return fmt.Errorf("DBERROR:=>SaleReturnDB: Number of affected row is not equal to 1:  %w", err)
	}

	//Update total_amount, due_amount(if available) in customers table
	query = `
		UPDATE Public.customers
		SET total_amount = total_amount - $1, due_amount = due_amount - $2
		WHERE id = $3
	`
	_, err = tx.ExecContext(ctx, query, ReturnAmount, SalesHistory.TotalAmount-SalesHistory.PaidAmount, SalesHistory.CustomerID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("SQLErrorSaleReturnDB(Update customer): %w", err)
	}

	//Insert Financial transaction
	var financial_transactions_id int
	query = `INSERT INTO public.financial_transactions (transaction_type,source_type,source_id,destination_type,destination_id,amount,transaction_date,description,created_at,updated_at)
	VALUES ('Cash Transfer','Account',$1, 'Customer', $2, $3, $4, $5, $6, $7) RETURNING transaction_id
	`
	description := "cash return due to returning product from the customer"
	err = tx.QueryRowContext(ctx, query,
		&SalesHistory.AccountID,
		&SalesHistory.CustomerID,
		&ReturnAmount,
		time.Now(),
		&description,
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
	tx, err := p.DB.Begin()
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
			i.id, i.product_code, i.product_name, i.product_description, i.product_status, i.quantity_purchased, i.purchase_cost, i.quantity_sold, i.sold_price, i.category_id, i.brand_id, i.discount, i.created_at, i.updated_at, b.id, b.name, c.id, c.name
		FROM 
			public.products i
			INNER JOIN brands b ON (b.id = i.brand_id) 
			INNER JOIN categories c ON (c.id = i.category_id)
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
			&p.Discount,
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

// Helper functions
// CountTotalEntries counts total number of rows in given the table
func (p *postgresDBRepo) CountRows(tableName string) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var c int
	query := "SELECT COUNT(id) FROM " + tableName
	err := p.DB.QueryRowContext(ctx, query).Scan(&c)
	fmt.Println(c)
	return c, err
}

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
