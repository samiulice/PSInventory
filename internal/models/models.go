package models

import (
	"time"
)

//MM-PR-randomAlphanumeriac(6)+CurrentIndexOfPurchaseHistory Table //Purchase
//MM-SL-randomAlphanumeriac(6)+CurrentIndexOfSalesHistory Table //Sale
//MM-WC-randomAlphanumeriac(6)+CurrentIndexOfWarrantyHistory Table //Warranty Claimed
//MM-WCO-randomAlphanumeriac(6)+CurrentIndexOfWarrantyHistory Table //Warranty Checked Out
//MM-WD-randomAlphanumeriac(6)+CurrentIndexOfWarrantyHistory Table //Warranty Delivered
//MM-RR-randomAlphanumeriac(6)+CurrentIndexOfWarrantyHistory Table //Repair Receive
//MM-RC-randomAlphanumeriac(6)+CurrentIndexOfWarrantyHistory Table //Repair Completed
//MM-RD-randomAlphanumeriac(6)+CurrentIndexOfWarrantyHistory Table //Repair Delivered

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
	ID        int       `json:"id"`
	Name      string    `json:"name,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// Category is the type for Category
type Category struct {
	ID        int       `json:"id"`
	Name      string    `json:"name,omitempty"`
	Status      string    `json:"status,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// Employee is the type for employees
type Employee struct {
	ID            int    `json:"id"`
	AccountCode   string `json:"account_code,omitempty"`
	AccountName   string `json:"account_name,omitempty"`
	ContactPerson string `json:"contact_person,omitempty"`
	// Gender                string         `json:"gender,omitempty"`
	// DateOfBith            time.Time      `json:"date_of_birth,omitempty"`
	// Experties        string         `json:"exparties,omitempty"`
	Division         string    `json:"division,omitempty"`
	District         string    `json:"district,omitempty"`
	Upazila          string    `json:"upazila,omitempty"`
	Area             string    `json:"area,omitempty"`
	Mobile           string    `json:"mobile,omitempty"`
	AmountPayable    int       `json:"amount_payable"`
	AmountReceivable int       `json:"amount_receivable"`
	WhatsappAccount  string    `json:"whatsapp_account,omitempty"`
	Email            string    `json:"email,omitempty"`
	ImageLink        string    `json:"image_link,omitempty"` //username_profile_id_yy-mm-dd_hh-mm-ss.jpg
	AccountStatus    bool      `json:"account_status"`       //Active = true, Inactive = false
	MonthlySalary    int       `json:"monthly_salary"`
	OpeningBalance   int       `json:"opening_balance"`
	CVLink           string    `json:"cv_link,omitempty"`
	NIDLink          string    `json:"nid_link,omitempty"`
	JoiningDate      time.Time `json:"joining_date,omitempty"`
	CreatedAt        time.Time `json:"created_at,omitempty"`
	UpdatedAt        time.Time `json:"updated_at,omitempty"`
}

// Customer is the type for customers
type Customer struct {
	ID               int       `json:"id"`
	AccountCode      string    `json:"account_code,omitempty"`
	AccountName      string    `json:"account_name,omitempty"`
	ContactPerson    string    `json:"contact_person,omitempty"`
	Division         string    `json:"division,omitempty"`
	District         string    `json:"district,omitempty"`
	Upazila          string    `json:"upazila,omitempty"`
	Area             string    `json:"area,omitempty"`
	Mobile           string    `json:"mobile,omitempty"`
	AmountPayable    int       `json:"amount_payable"`
	AmountReceivable int       `json:"amount_receivable"`
	Email            string    `json:"email,omitempty"`
	WhatsappAccount  string    `json:"whatsapp_account,omitempty"`
	AccountStatus    bool      `json:"account_status"` //Active = true, Inactive = false
	Discount         int       `json:"discount"`
	OpeningBalance   int       `json:"opening_balance"`
	JoiningDate      time.Time `json:"joining_date,omitempty"`
	CreatedAt        time.Time `json:"created_at,omitempty"`
	UpdatedAt        time.Time `json:"updated_at,omitempty"`
}

// Supplier is the type for suppliers
type Supplier struct {
	ID               int       `json:"id"`
	AccountCode      string    `json:"account_code,omitempty"`
	AccountName      string    `json:"account_name,omitempty"`
	ContactPerson    string    `json:"contact_person,omitempty"`
	Division         string    `json:"division,omitempty"`
	District         string    `json:"district,omitempty"`
	Upazila          string    `json:"upazila,omitempty"`
	Area             string    `json:"area,omitempty"`
	Mobile           string    `json:"mobile,omitempty"`
	AmountPayable    int       `json:"amount_payable"`
	AmountReceivable int       `json:"amount_receivable"`
	Email            string    `json:"email,omitempty"`
	WhatsappAccount  string    `json:"whatsapp_account,omitempty"`
	AccountStatus    bool      `json:"account_status"` //Active = true, Inactive = false
	Discount         int       `json:"discount"`
	JoiningDate      time.Time `json:"joining_date,omitempty"`
	CreatedAt        time.Time `json:"created_at,omitempty"`
	UpdatedAt        time.Time `json:"updated_at,omitempty"`
}

// Product is the type for products
type HeadAccount struct {
	ID            int       `json:"id"`
	AccountCode   string    `json:"account_code,omitempty"`
	AccountName   string    `json:"account_name,omitempty"`
	AccoutnStatus bool      `json:"account_status"`
	CurrentAmount int       `json:"current_amount,omitempty"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
	UpdatedAt     time.Time `json:"updated_at,omitempty"`
}

// Product is the type for products
type Product struct {
	ID              int                `json:"id"`
	ProductCode     string             `json:"product_code,omitempty"`
	ProductName     string             `json:"product_name,omitempty"`
	Description     string             `json:"product_description,omitempty"`
	ProductStatus   bool               `json:"product_status"`
	Quantity        int                `json:"quantity"`
	CategoryID      int                `json:"category_id"`
	BrandID         int                `json:"brand_id"`
	Discount        int                `json:"discount"`
	CreatedAt       time.Time          `json:"created_at,omitempty"`
	UpdatedAt       time.Time          `json:"updated_at,omitempty"`
	Category        Category           `json:"category"`
	Brand           Brand              `json:"brand,omitempty"`
	ProductMetadata []*ProductMetadata `json:"product_metadata,omitempty"`
}

// ProductMetadata holds products meta data
type ProductMetadata struct {
	ID                     int       `json:"id"`
	SerialNumber           string    `json:"serial_number,omitempty"`
	ProductID              int       `json:"product_id"`
	PurchaseHistoryID      int       `json:"purchase_history_id"`
	SalesHistoryID         int       `json:"sales_history_id,omitempty"`
	Status                 string    `json:"status,,omitempty"`
	WarrantyPeriod         int       `json:"warranty"`
	WarrantyStatus         string    `json:"warranty_status,omitempty"`
	LatesWarrantyHistoryID int       `json:"latest_warranty_history_id"`
	WarrantyHistoryIDs     string    `json:"warranty_history_ids,omitempty"`
	MaxRetailPrice         int       `json:"max_retail_price"` //max_retail_price = total_amount+profit/quantity_purchased
	PurchaseRate           int       `json:"purchase_rate"`    //purchase_rate = total_amount/quantity_purchased
	SaleRate               int       `json:"sale_rate"`
	CreatedAt              time.Time `json:"created_at,omitempty"`
	UpdatedAt              time.Time `json:"updated_at,omitempty"`
}

// Purchase is the type for purchase info
type Purchase struct {
	ID                int       `json:"id"`
	PurchaseDate      string    `json:"purchase_date,omitempty"`
	SupplierID        int       `json:"supplier_id"`
	ProductID         int       `json:"product_id"`
	Quantity          int       `json:"quantity"`
	QuantityPurchased int       `json:"quantity_purchased"`
	QuantitySold      int       `json:"quantity_sold"`
	ProductsSerialNo  []string  `json:"products_serial_no,omitempty"`
	AccountID         int       `json:"account_id"`
	ChalanNO          string    `json:"chalan_no,omitempty"`
	MemoNo            string    `json:"memo_no,omitempty"` //MM-P-randomAlphanumeriac(6)CurrentIndexOfPurchaseHistoryDB
	Note              string    `json:"note,omitempty"`
	MaxRetailPrice    int       `json:"max_retail_price"`
	PurchaseRate      int       `json:"purchase_rate"`
	WarrantyPeriod    int       `json:"warranty"`
	BillAmount        int       `json:"bill_amount"`
	Discount          int       `json:"discount"`
	TotalAmount       int       `json:"total_amount"`
	PaidAmount        int       `json:"paid_amount"`
	CreatedAt         time.Time `json:"created_at,omitempty"`
	UpdatedAt         time.Time `json:"updated_at,omitempty"`
}

type SelectedItems struct {
	ProductID     int      `json:"product_id"`
	SerialNumbers []string `json:"serialNumbers,omitempty"`
}

// Sale is the type for purchase info
type Sale struct {
	ID               int              `json:"id"`
	SaleDate         string           `json:"sale_date,omitempty"`
	CustomerID       int              `json:"customer_id"`
	ProductsSerialNo []string         `json:"products_serial_no,omitempty"`
	SelectedItems    []*SelectedItems `json:"selected_items,omitempty"`
	AccountID        int              `json:"account_id"`
	ChalanNO         string           `json:"chalan_no,omitempty"`
	MemoNo           string           `json:"memo_no,omitempty"`
	Note             string           `json:"note,omitempty"`
	BillAmount       int              `json:"bill_amount"`
	Discount         int              `json:"discount"`
	TotalAmount      int              `json:"total_amount"`
	PaidAmount       int              `json:"paid_amount"`
	CreatedAt        time.Time        `json:"created_at,omitempty"`
	UpdatedAt        time.Time        `json:"updated_at,omitempty"`
}

type Warranty struct {
	ID               int       `json:"id"`
	Status           string    `json:"status,omitempty"`
	ProductSerialID  int       `json:"product_serial_id"`
	PreviousSerialNo string    `json:"previous_serial_no,omitempty"`
	NewSerialNo      string    `json:"new_serial_no,omitempty"`
	MemoNo           string    `json:"memo_no,omitempty"`
	ContactNumber    string    `json:"contact_number,omitempty"`
	RequestedDate    string    `json:"requested_date,omitempty"`
	ReportedProblem  string    `json:"reported_problem,omitempty"`
	ReceivedBy       string    `json:"received_by,omitempty"`
	CheckoutDate     string    `json:"checkout_date,omitempty"`
	DeliveryDate     string    `json:"delivery_date,omitempty"`
	DeliveredBy      string    `json:"delivered_by,omitempty"`
	Comment          string    `json:"comment,omitempty"`
	CreatedAt        time.Time `json:"created_at,omitempty"`
	UpdatedAt        time.Time `json:"updated_at,omitempty"`
	ProductInfo      Product   `json:"product_info,omitempty"`
	SalesInfo        Sale      `json:"sales_info,omitempty"`
	CustomerInfo     Customer  `json:"customer_info,omitempty"`
	SupplierInfo     Supplier  `json:"supplier_info,omitempty"`
}
