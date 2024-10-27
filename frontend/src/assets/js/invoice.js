//renderSalesInvoicePage 
function renderSalesInvoicePage(InvoiceData) {
    let productTable = '';
    InvoiceData.selected_items.forEach((item, i) => {
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
                            <p class="customer-id"><strong>ID:</strong> &nbsp;${InvoiceData.customer_info.account_code}</p>
                            <p class="customer-name"><strong>Name:</strong> ${InvoiceData.customer_info.account_name}</p>
                            <p class="contact-number"><strong>Contact:</strong> ${InvoiceData.customer_info.mobile}</p>
                            <p class="customer-address"><strong>Address:</strong> ${InvoiceData.customer_info.upazila}, ${InvoiceData.customer_info.district}</p>
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
                                <b>INVOICE NO: </b>${InvoiceData.memo_no} &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; 
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
                            Total Paid:&nbsp;&nbsp;&#2547;&nbsp;${InvoiceData.paid_amount} &nbsp;&nbsp;&nbsp;&nbsp; 
                            Due:&nbsp;&nbsp;&#2547;&nbsp;${InvoiceData.total_amount - InvoiceData.paid_amount}
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
                            <td class="text-center w_15">${InvoiceData.sale_date}</td>
                            <td class="text-center w_10">${InvoiceData.head_account_info.account_name}</td>
                            <td class="text_center w_35">${InvoiceData.chalan_no === "" ? "Cash Received" : InvoiceData.chalan_no}</td>
                            <td class="text_center w_10">&#2547&nbsp;${InvoiceData.paid_amount}</td>
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
                                <b>&#2547; ${InvoiceData.bill_amount}</b>
                            </td>
                            </tr>
                            <tr>
                            <td class="text_left">DISCOUNT</td>
                            <td><b>&#2547; ${InvoiceData.bill_amount - InvoiceData.total_amount}</b></td>
                            </tr>
                            <tr>
                            <td class="text_left">TOTAL</td>
                            <td><b>&#2547; ${InvoiceData.total_amount}</b>
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

