package models

import (
	"time"
)

const (
	// AustraliaRegex matches Australian mobile numbers with or without country code (+61)
	AustraliaRegex = `^(\+?61|0)4\d{8}$`

	// BangladeshRegex matches Bangladeshi mobile numbers with or without country code (+880)
	BangladeshRegex = `^\+?(880)?1[3-9]\d{8}$`

	// CanadaRegex matches Canadian phone numbers in various formats
	CanadaRegex = `^(\+?1)?[-.\s]?\(?\d{3}\)?[-.\s]?\d{3}[-.\s]?\d{4}$`

	// FranceRegex matches French phone numbers with or without country code (+33)
	FranceRegex = `^(?:(?:\+|00)33|0)\s*[1-9](?:[\s.-]*\d{2}){4}$`

	// GermanyRegex matches German phone numbers with or without country code (+49)
	GermanyRegex = `^(\+?49|0)(\d{3,4})?[ -]?(\d{3,4})?[ -]?(\d{4,6})$`

	// IndiaRegex matches Indian mobile numbers with or without country code (+91)
	IndiaRegex = `^\+?(91)?\d{10}$`

	// JapanRegex matches Japanese phone numbers with or without country code (+81)
	JapanRegex = `^\+?81[-.\s]?\d{1,4}[-.\s]?\d{1,4}[-.\s]?\d{4}$`

	// PakistanRegex matches Pakistani mobile numbers with or without country code (+92)
	PakistanRegex = `^\+?(92)?\d{10}$`

	// SriLankaRegex matches Sri Lankan mobile numbers with or without country code (+94)
	SriLankaRegex = `^\+?(94)?\d{9}$`

	// UKRegex matches UK phone numbers including landline, mobile, and toll-free numbers
	UKRegex = `^(?:(?:\+44\s?|0)(?:\d{5}\s?\d{5}|\d{4}\s?\d{4}\s?\d{4}|\d{3}\s?\d{3}\s?\d{4}|\d{2}\s?\d{4}\s?\d{4}|\d{4}\s?\d{4}|\d{4}\s?\d{3})|\d{5}\s?\d{4}\s?\d{4}|0800\s?\d{3}\s?\d{4})$`

	// USRegex matches US phone numbers in various formats
	USRegex = `^\+?1?[-.\s]?\(?\d{3}\)?[-.\s]?\d{3}[-.\s]?\d{4}$`
)

// Brand is the type for Brand
type Brand struct {
	ID        int       `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// Category is the type for Category
type Category struct {
	ID        int       `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// Customer is the type for customers
type Customer struct {
	ID             int       `json:"id,omitempty"`
	AccountCode    string    `json:"account_code,omitempty"`
	AccountName    string    `json:"account_name,omitempty"`
	ContactPerson  string    `json:"contact_person,omitempty"`
	Division       string    `json:"division,omitempty"`
	District       string    `json:"district,omitempty"`
	Upazila        string    `json:"upazila,omitempty"`
	Area           string    `json:"area,omitempty"`
	Mobile         string    `json:"mobile,omitempty"`
	Email          string    `json:"email,omitempty"`
	AccountStatus  bool      `json:"account_status,omitempty"` //Active = true, Inactive = false
	Discount       int       `json:"discount,omitempty"`
	OpeningBalance int       `json:"opening_balance,omitempty"`
	JoiningDate    time.Time `json:"joining_date,omitempty"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
	UpdatedAt      time.Time `json:"updated_at,omitempty"`
}

// Employee is the type for employees
type Employee struct {
	ID            int    `json:"id,omitempty"`
	AccountCode   string `json:"account_code,omitempty"`
	AccountName   string `json:"account_name,omitempty"`
	ContactPerson string `json:"contact_person,omitempty"`
	// Gender                string         `json:"gender,omitempty"`
	// DateOfBith            time.Time      `json:"date_of_birth,omitempty"`
	// Experties        string         `json:"exparties,omitempty"`
	Division string `json:"division,omitempty"`
	District string `json:"district,omitempty"`
	Upazila  string `json:"upazila,omitempty"`
	Area     string `json:"area,omitempty"`
	Mobile   string `json:"mobile,omitempty"`
	Email    string `json:"email,omitempty"`
	// ImageLink             string         `json:"image_link,omitempty"`        //username_profile_id_yy-mm-dd_hh-mm-ss.jpg
	AccountStatus  bool `json:"account_status,omitempty"` //Active = true, Inactive = false
	MonthlySalary  int  `json:"monthly_salary,omitempty"`
	OpeningBalance int  `json:"opening_balance,omitempty"`
	// CVLink               string         `json:"cv_link,omitempty"`
	// NIDLink               string         `json:"nid_link,omitempty"`
	// KYCUploadStatus       string         `json:"kyc_upload_status,omitempty"`       // incomplete // completed
	// KYCVerificationStatus string         `json:"kyc_verification_status,omitempty"` //unchecked;Verified; Info Mismatch
	JoiningDate time.Time `json:"joining_date,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

// Product is the type for products
type HeadAccount struct {
	ID            int       `json:"id,omitempty"`
	AccountCode   string    `json:"account_code,omitempty"`
	AccountName   string    `json:"account_name,omitempty"`
	AccoutnStatus bool      `json:"account_status,omitempty"`
	CurrentAmount int       `json:"current_amount,omitempty"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
	UpdatedAt     time.Time `json:"updated_at,omitempty"`
}

// Product is the type for products
type Product struct {
	ID              int                `json:"id,omitempty"`
	ProductCode     string             `json:"product_code,omitempty"`
	ProductName     string             `json:"product_name,omitempty"`
	Description     string             `json:"product_description,omitempty"`
	ProductStatus   bool               `json:"product_status,omitempty"`
	Quantity        int                `json:"quantity,omitempty"`
	CategoryID      int                `json:"category_id,omitempty"`
	BrandID         int                `json:"brand_id,omitempty"`
	Discount        int                `json:"discount,omitempty"`
	CreatedAt       time.Time          `json:"created_at,omitempty"`
	UpdatedAt       time.Time          `json:"updated_at,omitempty"`
	Category        Category           `json:"category"`
	Brand           Brand              `json:"brand,omitempty"`
	ProductMetadata []*ProductMetadata `json:"product_metadata,omitempty"`
}

// ProductMetadata holds products meta data
type ProductMetadata struct {
	ID                int       `json:"id,omitempty"`
	SerialNumber      string    `json:"serial_number,omitempty"`
	ProductID         int       `json:"product_id,omitempty"`
	PurchaseHistoryID int       `json:"purchase_history_id,omitempty"`
	Status            string    `json:"status,omitempty"`
	Warranty          int       `json:"warranty,omitempty"`
	MaxRetailPrice    int       `json:"max_retail_price,omitempty"` //max_retail_price = total_amount+profit/quantity_purchased
	PurchaseRate      int       `json:"purchase_rate,omitempty"`    //purchase_rate = total_amount/quantity_purchased
	CreatedAt         time.Time `json:"created_at,omitempty"`
	UpdatedAt         time.Time `json:"updated_at,omitempty"`
}

// Purchase is the type for purchase info
type Purchase struct {
	ID                int       `json:"id"`
	PurchaseDate      string    `json:"purchase_date,omitempty"`
	SupplierID        int       `json:"supplier_id,omitempty"`
	ProductID         int       `json:"product_id,omitempty"`
	Quantity          int       `json:"quantity,omitempty"`
	QuantityPurchased int       `json:"quantity_purchased,omitempty"`
	QuantitySold      int       `json:"quantity_sold,omitempty"`
	ProductsSerialNo  []string  `json:"products_serial_no,omitempty"`
	AccountID         int       `json:"account_id,omitempty"`
	ChalanNO          string    `json:"chalan_no,omitempty"`
	MemoNo            string    `json:"memo_no,omitempty"`
	Note              string    `json:"note,omitempty"`
	MaxRetailPrice    int       `json:"max_retail_price,omitempty"`
	PurchaseRate      int       `json:"purchase_rate,omitempty"`
	Warranty          int       `json:"warranty,omitempty"`
	BillAmount        int       `json:"bill_amount,omitempty"`
	Discount          int       `json:"discount,omitempty"`
	TotalAmount       int       `json:"total_amount,omitempty"`
	PaidAmount        int       `json:"paid_amount,omitempty"`
	CreatedAt         time.Time `json:"created_at,omitempty"`
	UpdatedAt         time.Time `json:"updated_at,omitempty"`
}

type SelectedItems struct {
	ProductID     int      `json:"productId"`
	SerialNumbers []string `json:"serialNumbers"`
}

// Sale is the type for purchase info
type Sale struct {
	ID               int             `json:"id"`
	SaleDate         string          `json:"sale_date,omitempty"`
	CustomerID       int             `json:"customer_id,omitempty"`
	ProductsSerialNo []string        `json:"products_serial_no,omitempty"`
	SelectedItems    []SelectedItems `json:"selected_items,omitempty"`
	AccountID        int             `json:"account_id,omitempty"`
	ChalanNO         string          `json:"chalan_no,omitempty"`
	MemoNo           string          `json:"memo_no,omitempty"`
	Note             string          `json:"note,omitempty"`
	BillAmount       int             `json:"bill_amount,omitempty"`
	Discount         int             `json:"discount,omitempty"`
	TotalAmount      int             `json:"total_amount,omitempty"`
	PaidAmount       int             `json:"paid_amount,omitempty"`
	CreatedAt        time.Time       `json:"created_at,omitempty"`
	UpdatedAt        time.Time       `json:"updated_at,omitempty"`
}

// Supplier is the type for suppliers
type Supplier struct {
	ID            int       `json:"id,omitempty"`
	AccountCode   string    `json:"account_code,omitempty"`
	AccountName   string    `json:"account_name,omitempty"`
	ContactPerson string    `json:"contact_person,omitempty"`
	Division      string    `json:"division,omitempty"`
	District      string    `json:"district,omitempty"`
	Upazila       string    `json:"upazila,omitempty"`
	Area          string    `json:"area,omitempty"`
	Mobile        string    `json:"mobile,omitempty"`
	Email         string    `json:"email,omitempty"`
	AccountStatus bool      `json:"account_status,omitempty"` //Active = true, Inactive = false
	Discount      int       `json:"discount,omitempty"`
	JoiningDate   time.Time `json:"joining_date,omitempty"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
	UpdatedAt     time.Time `json:"updated_at,omitempty"`
}

//................................................................

// Order is the type for all orders
type Order struct {
	ID            int       `json:"id,omitempty"`
	DatesID       int       `json:"dates_id,omitempty"`
	TransactionID int       `json:"transaction_id,omitempty"`
	CustomerID    int       `json:"customer_id,omitempty"`
	StatusID      int       `json:"status_id,omitempty"` //Processing=1, Completed=2, Cancelled = 3
	Quantity      int       `json:"quantity,omitempty"`
	Amount        int       `json:"amount,omitempty"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
	UpdatedAt     time.Time `json:"updated_at,omitempty"`
	// Transaction   Transaction `json:"transaction,omitempty"`
	Customer Customer `json:"customer,omitempty"`
}

// ................JSON Response model for invoice microservice........................
// Order holds the necessary info to build invoice
type Invoice struct {
	ID        int              `json:"id,omitempty"`
	FirstName string           `json:"first_name,omitempty"`
	LastName  string           `json:"last_name,omitempty"`
	Email     string           `json:"email,omitempty"`
	CreatedAt time.Time        `json:"created_at,omitempty"`
	Products  []InvoiceProduct `json:"items,omitempty"`
}
type InvoiceProduct struct {
	ID       int    `json:"product_id,omitempty"`
	Name     string `json:"product_name,omitempty"`
	Quantity int    `json:"quantity,omitempty"`
	Amount   int    `json:"amount,omitempty"`
}

//........................................................................................
