//renderSalesInvoicePage 
function renderSalesInvoicePage(invoiceData) {
    let productTable = '';
    invoiceData.selected_items.forEach((item, i) => {
        //Concat serial number
        let serial = '';
        if (item.serial_numbers.length === 1) {
            serial = `<span>${item.serial_numbers[0]}</span>`;
        } else {
            let snLn = item.serial_numbers.length
            for (let s = 0; s < snLn - 1; s++) {
                serial += `<span>${item.serial_numbers[s]} </span>`;
            }
            //ensuring that <br> tag is not used after the last serial number
            serial += `<span>${item.serial_numbers[snLn - 1]}</span>`
        }

        productTable += `
        <tr>
            <td class="text_center w_3">${i}</td>
            <td class="w_32">
            ${item.product_name}<br>
            <span class="product-deatils">
                <b>Category: </b>${item.category_name} &nbsp;&nbsp;
                <b>Brand: </b>${item.brand_name} &nbsp;&nbsp;<br>
                <b> Warranty: </b> ${item.warranty} days
            </span>

            </td>
            <td class="w_30 sn-line shrink-text">
            ${serial}
            </td>
            <td class="text_center w_5">${item.quantity}</td>
            <td class="text_right w_12">&nbsp;${item.mrp}</td>
            <!-- &#2547; is the unicode for BDT symbol -->
            <td class="text_right w_18">&nbsp;${item.sub_total}</td>
        </tr>
        `

    })

    document.getElementById("sale-page-content").innerHTML = `
        <div id="sales-invoice-content">
            <section>
                <div class="invoice_title">Sales Invoice</div>

                <div class="invoice">
                <div class="header" id="header">
                    <div class="i_row">
                    <div class="i_col text_left">
                        <!-- Company Logo here -->
                        <div class="invoice-header">
                        <div class="customer-info">
                            <div class="p_title">Customer Infomation</div>
                            <!-- <div class="divider"></div> -->
                            <p class="customer-id"><strong>ID:</strong> &nbsp;${invoiceData.customer_info.account_code}</p>
                            <p class="customer-name"><strong>Name:</strong> ${invoiceData.customer_info.account_name}</p>
                            <p class="contact-number"><strong>Contact:</strong> ${invoiceData.customer_info.mobile}</p>
                            <p class="customer-address"><strong>Address:</strong> ${invoiceData.customer_info.upazila}, ${invoiceData.customer_info.district}</p>
                        </div>
                        </div>

                    </div>
                    <div class="i_col text_left" style="padding-left: 10px;">
                        <!-- Company Logo here -->
                        <div class="i_logo">
                        <div class="company-info">
                            <!-- <img src="assets/icons/logo.jpeg" alt="Company Logo" class="logo"> -->
                            <div class="text-container">
                            <span class="text_right">ProjuktiSheba Inventory</span>
                            <label class="slogan">Track Smarter, Stock Better</label>
                            <div class="contact-info">
                                <div class="address">Netrakona, Mymensingh, Bangladesh</div>
                                <div class="phone">Phone: (123) 456-7890</div>
                                <div class="email">Email: info@yourcompany.com</div>
                            </div>
                            <hr>
                            <div class="i_number text_right">
                                <i>
                                <b>INVOICE NO: </b>${invoiceData.memo_no} &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; 
                                <b>Date:</b>${formatDate(getCurrentDate(), 'date', "-")}
                                </i>
                            </div>
                            </div>
                        </div>
                        </div>
                    </div>
                    </div>
                </div>
                <div class="body">
                    <table class="table table-bordered no-footer jambo_table table-container">
                    <thead>
                        <tr>
                        <th class="text_center w_3">SL.</th>
                        <th class="text-center w_32">Description</th>
                        <th class="text-center w_30">Serial Number</th>
                        <th class="text_center w_5">Quantity</th>
                        <th class="text_center w_12">Unit Price</th>
                        <th class="text_center w_18">Amount</th>
                        </tr>
                    </thead>
                    <tbody>
                        ${productTable}
                    </tbody>
                    </table>
                </div>
                <div class="footer">
                    <!-- Financial Sections -->
                    <div class="i_row">
                    <div class="w_75">
                        <div>
                        <span class="p_title">Payment History:</span>
                        <span style="font-weight: 400; margin-left: 20%;">
                            Total Paid:&nbsp;&nbsp;&#2547;&nbsp;${invoiceData.paid_amount} &nbsp;&nbsp;&nbsp;&nbsp; 
                            Due:&nbsp;&nbsp;&#2547;&nbsp;${invoiceData.total_amount - invoiceData.paid_amount}
                        </span>
                        </div>
                        <table class="table table-striped jambo_table">
                        <thead>
                            <tr>
                            <th class="text_center w_5">SL.</th>
                            <th class="text-center w_10">Date</th>
                            <th class="text-center w_10">Method</th>
                            <th class="text_center w_40">Details</th>
                            <th class="text_center w_10">Amount</th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr>
                            <td class="text_center w_5">1</td>
                            <td class="text-center w_15">${invoiceData.sale_date}</td>
                            <td class="text-center w_10">${invoiceData.head_account_info.account_name}</td>
                            <td class="text_center w_35">${invoiceData.chalan_no === "" ? "Cash Received" : invoiceData.chalan_no}</td>
                            <td class="text_center w_10">&#2547&nbsp;${invoiceData.paid_amount}</td>
                            </tr>
                        </tbody>
                        </table>
                    </div>
                    <div class="w_5"></div>
                    <div class="w_20 text_right">
                        <table class="table table-striped jambo_table">
                        <tbody>
                            <tr>
                            <td class="text_left">SUBTOTAL</td>
                            <td>
                                <b>&#2547; ${invoiceData.bill_amount}</b>
                            </td>
                            </tr>
                            <tr>
                            <td class="text_left">DISCOUNT</td>
                            <td><b>&#2547; ${invoiceData.bill_amount - invoiceData.total_amount}</b></td>
                            </tr>
                            <tr>
                            <td class="text_left">TOTAL</td>
                            <td><b>&#2547; ${invoiceData.total_amount}</b>
                            </tr>
                        </tbody>
                        </table>
                    </div>
                    </div>
                    <div class="i_row">
                    <div class="i_col w_10"></div>
                    <div class="i_col w_40 text_left">
                        Customer's Signature: <sub>....................................................</sub>
                    </div>
                    <div class="i_col w_40 text_right">
                        Seller's Signature: <sub>....................................................</sub>
                    </div>
                    <div class="i_col w_10"></div>
                    </div>
                    <div class="divider">

                    </div>
                    <div class="i_row">
                    <div class="i_col">
                        <p class="p_title">Terms and Conditions</p>
                        <ol type="i" style="margin-left: 20px;">
                        <li>No refunds or exchanges.</li>
                        <li>Warranty excludes damage from misuse, consumables, or unauthorized handling</li>
                        </ol>
                    </div>
                    </div>
                </div>
                </div>
            </section>
        </div>                        
        `

}

/**
 * renderPurchaseInvoice render the purchase details page and invoice printing facilities
 * 
 * param{invoiceData} object - the data in specific format that is need to deploy the invoice page
 * 
 * The function handles the following:
 * - populate an invoice div
 * - change the innerHTML of the page content section with the populated invoice div
 * - a print invoice button for opening printing window
 * 
 * Dependency: htmlStringToPDF function
 */
function renderPurchaseInvoice(invoiceData, dpi) {
    let tableBody = `
        <tr>
            <td>1</td>
            <td>${invoiceData.product_info.product_name}
                </br>Category: ${invoiceData.product_info.category.name}} &nbsp;&nbsp;&nbsp;&nbsp; ${invoiceData.product_info.brand.name}
                </br>Category: ${invoiceData.warranty}
            </td>
            <td>${invoiceData.quantity_purchased}</td>
            <td>${invoiceData.max_retail_price}</td>
            <td>${invoiceData.total_amount}</td>
        </tr>
    `

    let htmlContent = `
    <!DOCTYPE html>
    <html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>Purchase Invoice</title>
        <style>
            /* Basic Reset */
            * {
                box-sizing: border-box;
                margin: 0;
                padding: 0;
            }
            body {
                font-family: 'Arial', sans-serif;
                color: #333;
                background-color: #f4f4f9;
                padding: 20px;
            }
            .invoice-container {
                max-width: 800px;
                margin: auto;
                padding: 30px;
                background-color: #fff;
                border-radius: 8px;
                box-shadow: 0 4px 10px rgba(0, 0, 0, 0.15);
            }
            .header {
                display: flex;
                flex-direction: column;
                align-items: center;
                margin-bottom: 20px;
                text-align: center;
                border-bottom: 2px solid #4a90e2;
                padding-bottom: 10px;
            }
            .company-info {
                font-family: 'Georgia', serif;
                font-size: 1.2em;
                color: #222;
                font-weight: bold;
                line-height: 1.5;
                display: flex;
                align-items: center;
                gap: 10px;
            }
            .company-info img {
                height: 50px;
                width: 50px;
                border-radius: 5px;
            }
            .vendor-invoice-info {
                display: flex;
                justify-content: space-between;
                align-items: flex-start;
                margin-top: 20px;
            }
            .vendor-section h2 {
                font-family: 'Verdana', sans-serif;
                font-size: 1.1em;
                color: #4a90e2;
                border-left: 5px solid #4a90e2;
                padding-left: 10px;
                margin-bottom: 10px;
            }
            .vendor-details, .invoice-details {
                font-size: 0.95em;
                color: #333;
                line-height: 1.5;
                font-weight: 500;
            }
            .invoice-details span {
                display: block;
                margin-bottom: 5px;
                color: #4a90e2;
                font-family: 'Arial', sans-serif;

            }
            .items-table {
                width: 100%;
                border-collapse: collapse;
                margin-top: 20px;
                font-size: 0.95em;
                font-family: 'Tahoma', sans-serif;
            }
            .items-table th, .items-table td {
                padding: 10px;
                border: 1px solid #333;
                color: #222;
            }
            .items-table th {
                background-color: #eaf2fb;
                color: #4a90e2;
                font-weight: 600;
                text-align: left;
            }
            .items-table td {
                text-align: left;
            }
            .items-table td:first-child, .items-table th:first-child {
                text-align: left;
            }
            .totals-section {
                display: flex;
                justify-content: space-between;
                gap: 20px;
                margin-top: 15px;
            }
            .totals, .payment-info {
                flex: 1;
                font-size: 0.9em;
                color: #333;
                font-family: 'Verdana', sans-serif;
            }
            .totals .total-row, .payment-info .info-row {
                display: flex;
                justify-content: space-between;
                padding: 5px 0;
            }
            .totals .total-row:last-child {
                font-size: 1em;
                color: #4a90e2;
            }
            .totals .total-label, .payment-info .info-label {
                color: #222;
            }
            .payment-info {
                background-color: #f1f1f1;
                padding: 8px;
                border-radius: 5px;
                border: 1px solid #333;
            }
            .remarks {
                font-family: 'Georgia', serif;
                font-size: 0.9em;
                color: #333;
                line-height: 1.6;
                background-color: #f9f9f9;
                padding: 15px;
                border-radius: 5px;
                border-left: 3px solid #4a90e2;
                margin-top: 20px;
                font-weight: 400;
            }

            /* Print Styling for A4 Page */
            @media print {
                body {
                    background-color: #ffffff;
                    padding: 0;
                }
                .invoice-container {
                    max-width: none;
                    width: 210mm;
                    height: 297mm;
                    margin: 0 auto;
                    padding: 20mm;
                    box-shadow: none;
                    border-radius: 0;
                    page-break-after: avoid;
                }
                .header, .vendor-invoice-info, .totals-section, .remarks {
                    margin-bottom: 10mm;
                }
                .items-table th, .items-table td {
                    padding: 8px;
                    font-size: 0.9em;
                    page-break-before: auto;
                    page-break-after: auto;
                    page-break-inside: avoid;
                }
            }
        </style>
    </head>
    <body>
        <div class="invoice-container">
            <!-- Header with Company Logo and Info Centered, with underline -->
            <div class="header">
                <div class="company-info">
                    <img src="your-logo.png" alt="Company Logo">
                    <div>
                        Projukti Sheba Inventory<br>
                        <span style="font-size: 0.85em; font-weight: normal;">Teribazar, Netrakona Shadar, Netrakon | Mobile: +8801742135093</span>
                    </div>
                </div>
            </div>

            <!-- Vendor and Invoice Details Side by Side -->
            <div class="vendor-invoice-info">
                <!-- Vendor Section -->
                <div class="vendor-section">
                    <h2>Vendor Details</h2>
                    <div class="vendor-details">
                        Name: ${invoiceData.supplier_info.account_name}<br>
                        Address: ${invoiceData.supplier_info.area ? invoiceData.supplier_info.area+",":""} ${invoiceData.supplier_info.upazila ? invoiceData.supplier_info.upazila+",":""}, ${invoiceData.supplier_info.district ? invoiceData.supplier_info.district:""}<br>
                        Mobile: ${invoiceData.supplier_info.mobile ? invoiceData.supplier_info.mobile : ""}
                    </div>
                </div>
                <!-- Invoice Section -->
                <div class="invoice-details">
                    <span><strong>INVOICE NO:</strong> #${invoiceData.memo_no}</span>
                    <span><strong>DATE:</strong> ${formatDate(invoiceData.purchase_date, "date", "-")}</span>
                </div>
            </div>

            <!-- Items Table -->
            <table class="items-table">
                <thead>
                    <tr>
                        <th>S/N</th>
                        <th>PRODUCT DESCRIPTION</th>
                        <th>QUANTITY</th>
                        <th>UNIT PRICE</th>
                        <th>SUB TOTAL</th>
                    </tr>
                </thead>
                <tbody>
                    ${tableBody}
                </tbody>
            </table>

            <!-- Totals and Payment Info Section -->
            <div class="totals-section">
                <!-- Totals Section -->
                <div class="totals">
                    <div class="total-row"><span class="total-label">BILL AMOUNT</span><span>&#2547 ${invoiceData.bill_amount}</span></div>
                    <div class="total-row"><span class="total-label">DISCOUNT</span><span>${invoiceData.bill_amount-invoiceData.total_amount}</span></div>
                    <div class="total-row"><span class="total-label">TOTAL AMOUNT</span><span>${invoiceData.total_amount}</span></div>
                    <div class="total-row"><strong>Paid Amount</strong><span><strong>${invoiceData.paid_amount}</strong></span></div>
                    <div class="total-row"><strong>Amount Due</strong><span><strong>${invoiceData.total_amount-invoiceData.paid_amount}</strong></span></div>
                </div>
                <!-- Payment Info Section -->
                <div class="payment-info">
                    <div class="info-row"><span class="info-label">Payment Method</span><span>${invoiceData.chalan_no ? "Bank Transfer" : "Cash"}</span></div>
                    <div class="info-row"><span class="info-label">Transferred Amount</span><span>${invoiceData.paid_amount}</span></div>
                    <div class="info-row"><span class="info-label">Note</span><span>${invoiceData.note}</span></div>
                </div>
            </div>

            <!-- Remarks Section -->
            <div class="remarks">
                <strong>Remarks / Payment Instructions:</strong><br>
                Please make payments by the due date. Reference Purchase Invoice #${invoiceData.memo_no} when submitting payment.
            </div>
        </div>
    </body>
    </html>`

    document.getElementById("purchase-form-content").innerHTML = htmlContent

}

/**
 * Converts a html string to a pdf and show a printing window
 *
 * param {content} string - The html content that is converted to a pdf.
 * param {dpi} int - set the pdf resolution.
 *
 * The function handles the following:
 * - create a html div from the param content
 * - convert the div into pdf using html2pdf library
 * - shows a printing window 
 * - higher dpi takes higher space in the disk for pdf resulting high quality pdf. For faster processing less dpi value is preferable(e.g. 2)
 * 
 * Dependency: <script src="https://cdnjs.cloudflare.com/ajax/libs/html2pdf.js/0.10.1/html2pdf.bundle.min.js"></script>
 */
function htmlStringToPDF(content, dpi) {
    // Create a temporary element to hold the HTML
    const tempElement = document.createElement('div');
    tempElement.innerHTML = content;

    // Use html2pdf to convert the HTML to PDF
    html2pdf()
        .from(tempElement)
        .set({
            margin: 1,
            filename: 'document.pdf',
            html2canvas: { scale: dpi },
            jsPDF: { unit: 'in', format: 'letter', orientation: 'portrait' }
        })
        .toPdf()
        .get('pdf')
        .then(function (pdf) {
            // Open the PDF in the print dialog
            pdf.autoPrint();
            window.open(pdf.output('bloburl'), '_blank');
        })
        .finally(() => {
            // Clean up the temporary element
            tempElement.remove();
        });
}

