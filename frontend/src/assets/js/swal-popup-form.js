//addNewBrand show a popup form and then make an api call to insert brand data to the database table
function addNewBrand(page, brand) {
  Swal.fire({
    title: 'Add New Brand',
    width: 600,
    html: `
        <div class="x_panel">
            <div class="x_content">
                <form id="add-brand" class="needs-validation" novalidate>
                    <!-- Brand Name -->
                    <div class="col-6 form-group has-feedback">
                        <input type="text" class="form-control has-feedback-left" id="name" name="name"
                            placeholder="Brand Name" autocomplete="off" required>
                        <div class="invalid-feedback d-none text-danger">Please enter the brand name.</div>
                        <span style="color: rgba(0, 0, 0, 1); transform:translate(-40%,-10%)" class="form-control-feedback left glyphicon glyphicon-tags" aria-hidden="true"></span>
                    </div>
                    <div class="form-group">
                      <div id="btns" class="col-4">
                          <br>
                          <button type="submit" class="btn btn-round btn-success">Submit</button>
                      </div>
                    </div>
                </form>
            </div>
        </div>
          `,
    showCloseButton: true,
    showConfirmButton: false,
    showCancelButton: false,
    allowOutsideClick: false,
    preConfirm: () => {
      return new Promise((resolve) => {
        const form = document.getElementById('add-brand');
        const formFields = form.querySelectorAll('.form-control, .form-select');
        let isValid = true;

        // Reset all feedback messages
        form.querySelectorAll('.invalid-feedback').forEach(feedback => {
          feedback.classList.add('d-none');
        });

        // Check each field
        formFields.forEach(field => {
          if (!field.checkValidity()) {
            isValid = false;
            const feedback = field.nextElementSibling;
            if (feedback && feedback.classList.contains('invalid-feedback')) {
              feedback.classList.remove('d-none');
            }
          }
        });

        if (isValid) {
          resolve({
            name: form.name.value,
          });
        } else {
          Swal.showValidationMessage('Please correct the errors in the form.');
        }
      });
    },
    willOpen: () => {
      const form = document.getElementById('add-brand');
      form.addEventListener('submit', function (event) {
        event.preventDefault();
        event.stopPropagation();
        document.querySelectorAll('.invalid-feedback').forEach(feedback => {
          feedback.classList.add('d-none');
        });
        let isValid = true;
        form.querySelectorAll('.form-control, .form-select').forEach(field => {
          if (!field.checkValidity()) {
            isValid = false;
            const feedback = field.nextElementSibling;
            if (feedback && feedback.classList.contains('invalid-feedback')) {
              feedback.classList.remove('d-none');
            }
          }
        });
        if (isValid) {
          Swal.getConfirmButton().click();
        } else {
          Swal.showValidationMessage('Please correct the errors in the form.');
        }
      });
    }
  }).then((result) => {
    if (result.isConfirmed) {
      const data = result.value;
      let brand = {
        name: data.name,
      }
      const requestOptions = {
        method: 'post',
        headers: {
          'Accept': 'application/json',
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(brand),
      }

      fetch('http://localhost:4321/api/inventory/add-brand', requestOptions)
        .then(response => response.json())
        .then(data => {
          if (data.error === true) {
            showErrorMessage(data.message)
          } else {
            showSuccessMessage(data.message);
            brand.push(data.result)
            if (page === "purchase") {
              document.getElementById("brand").innerHTML = '';
              document.getElementById("brand").innerHTML = `<option value="${brand.length - 1}" selected>${data.result.name}</option>`;;
              document.getElementById("brand").disabled = true;
            }
          }
        });
    }
  });
}
//addNewProduct show a popup form and then make an api call to insert iten data to the database table
function addNewProduct(page, brands, categories, products) {
  let brandList = '';
  let categoryList = '';
  if (brands) {
    brandList += `<select id="brand" class="form-control form-select has-feedback-left" required>
                          <option value="" selected disabled>Select brand</option>`
    brands.forEach(b => {
      brandList += `<option value="${b.id}">${b.name}</option>`;
    })
    brandList += '</select>';
  }
  if (categories) {
    categoryList += `<select id="category" class="form-control form-select has-feedback-left" required>
                          <option value="" selected disabled>Select category</option>`
    categories.forEach(c => {
      categoryList += `<option value="${c.id}">${c.name}</option>`;
    })
    categoryList += '</select>';
  }
  let htmlContent = `
        <div class="x_panel">
            <div class="x_content">
                <form id="add-product" class="needs-validation" novalidate>
                    <!-- Product Name -->
                    <div class="col-6 form-group has-feedback">
                        <input type="text" class="form-control has-feedback-left" id="name" name="name"
                            placeholder="Product Name" autocomplete="off" required>
                        <div class="invalid-feedback d-none text-danger">Please enter the product name.</div>
                        <span style="color: rgba(0, 0, 0, 1); transform:translate(-40%,-10%)" class="form-control-feedback left glyphicon glyphicon-plus-sign" aria-hidden="true"></span>
                    </div>
                    <div class="col-6 form-group has-feedback">
                        <input type="text" class="form-control has-feedback-left" id="description" name="description"
                            placeholder="Product Description(Optional)" autocomplete="off">
                        <div class="invalid-feedback d-none text-danger">Please enter the product name.</div>
                        <span style="color: rgba(0, 0, 0, 1); transform:translate(-40%,-10%)" class="form-control-feedback left glyphicon glyphicon-plus-sign" aria-hidden="true"></span>
                    </div>
                    <div class="col-6 form-group has-feedback">` + categoryList +
    `<span style="color: rgba(0, 0, 0, 1); transform:translate(-40%,-10%)" class="form-control-feedback left glyphicon glyphicon-list" aria-hidden="true"></span>
                    </div>
                    <div class="col-6 form-group has-feedback">` + brandList +
    `<span style="color: rgba(0, 0, 0, 1); transform:translate(-40%,-10%)" class="form-control-feedback left glyphicon glyphicon-tags" aria-hidden="true"></span>
                    </div>
                    <!-- Sale Discount -->
                    <div class="col-4 form-group has-feedback">
                        <input type="number" class="form-control has-feedback-left" id="discount" name="discount"
                            placeholder="discount(%)" min="0" max="100" autocomplete="off">
                        <div class="invalid-feedback d-none text-danger">Enter value between 0 to 100.</div>
                        <span style="color: rgba(0, 0, 0, 1); transform:translate(-40%,-10%)" class="form-control-feedback left glyphicon  glyphicon-gift" aria-hidden="true"></span>
                    </div>
                    <div class="form-group">
                      <div id="btns" class="col-4">
                          <br>
                          <button type="submit" class="btn btn-round btn-success">Submit</button>
                      </div>
                    </div>
                </form>
            </div>
        </div>
          `;
  Swal.fire({
    title: 'Add New product',
    width: 600,
    html: htmlContent,
    showCloseButton: true,
    showConfirmButton: false,
    showCancelButton: false,
    allowOutsideClick: false,
    preConfirm: () => {
      return new Promise((resolve) => {
        const form = document.getElementById('add-product');
        const formFields = form.querySelectorAll('.form-control, .form-select');
        let isValid = true;

        // Reset all feedback messages
        form.querySelectorAll('.invalid-feedback').forEach(feedback => {
          feedback.classList.add('d-none');
        });

        // Check each field
        formFields.forEach(field => {
          if (!field.checkValidity()) {
            isValid = false;
            const feedback = field.nextElementSibling;
            if (feedback && feedback.classList.contains('invalid-feedback')) {
              feedback.classList.remove('d-none');
            }
          }
        });

        if (isValid) {
          resolve({
            name: form.name.value,
            description: form.description.value,
            brand_id: form.brand.value,
            category_id: form.category.value,
            category_name: form.category.options[form.category.selectedIndex].text,
            discount: form.discount.value,
          });
        } else {
          Swal.showValidationMessage('Please correct the errors in the form.');
        }
      });
    },
    willOpen: () => {
      const form = document.getElementById('add-product');
      form.addEventListener('submit', function (event) {
        event.preventDefault();
        event.stopPropagation();
        document.querySelectorAll('.invalid-feedback').forEach(feedback => {
          feedback.classList.add('d-none');
        });
        let isValid = true;
        form.querySelectorAll('.form-control, .form-select').forEach(field => {
          if (!field.checkValidity()) {
            isValid = false;
            const feedback = field.nextElementSibling;
            if (feedback && feedback.classList.contains('invalid-feedback')) {
              feedback.classList.remove('d-none');
            }
          }
        });
        if (isValid) {
          Swal.getConfirmButton().click();
        } else {
          Swal.showValidationMessage('Please correct the errors in the form.');
        }
      });
    }
  }).then((result) => {
    if (result.isConfirmed) {
      const data = result.value;
      categoryName = data.category_name;
      let product = {
        product_name: data.name,
        product_description: data.description,
        brand_id: parseInt(data.brand_id),
        category_id: parseInt(data.category_id),
        discount: parseInt(data.discount),
      }
      const requestOptions = {
        method: 'post',
        headers: {
          'Accept': 'application/json',
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(product),
      }
      console.log(product)

      fetch('http://localhost:4321/api/inventory/add-product', requestOptions)
        .then(response => response.json())
        .then(data => {
          console.log(data)
          if (data.error === true) {
            showErrorMessage(data.message)
          } else {
            showSuccessMessage(data.message);
            products.push(data.result)
            if (page === "purchase") {
              categories.forEach((i, item) => {
                if (item.id = product.category_id) {
                  document.getElementById("category").innerHTML = '';
                  document.getElementById("category").innerHTML = `<option value="${i}" selected>${categoryName}</option>`;
                  document.getElementById("category").disabled = true;

                }
              })
              document.getElementById("product").innerHTML = '';
              document.getElementById("product").innerHTML = `<option value="${products.length-1}" selected>${data.result.product_name}</option>`;
              document.getElementById("product").disabled = true;
            }
          }
        });
    }
  });
}
//addNewCategory show a popup form and then make an api call to insert category data to the database table
function addNewCategory(page, categories) {
  Swal.fire({
    title: 'Add Product Category',
    width: 600,
    html: `
        <div class="x_panel">
            <div class="x_content">
                <form id="add-category" class="needs-validation" novalidate>
                    <!-- Category Name -->
                    <div class="col-6 form-group has-feedback">
                        <input type="text" class="form-control has-feedback-left" id="name" name="name"
                            placeholder="Category Name" autocomplete="off" required>
                        <div class="invalid-feedback d-none text-danger">Please enter the category name.</div>
                        <span style="color: rgba(0, 0, 0, 1); transform:translate(-40%,-10%)" class="form-control-feedback left glyphicon glyphicon-plus" aria-hidden="true"></span>
                    </div>
                    <div class="form-group">
                      <div id="btns" class="col-4">
                          <br>
                          <button type="submit" class="btn btn-round btn-success">Submit</button>
                      </div>
                    </div>
                </form>
            </div>
        </div>
          `,
    showCloseButton: true,
    showConfirmButton: false,
    showCancelButton: false,
    allowOutsideClick: false,
    preConfirm: () => {
      return new Promise((resolve) => {
        const form = document.getElementById('add-category');
        const formFields = form.querySelectorAll('.form-control, .form-select');
        let isValid = true;

        // Reset all feedback messages
        form.querySelectorAll('.invalid-feedback').forEach(feedback => {
          feedback.classList.add('d-none');
        });

        // Check each field
        formFields.forEach(field => {
          if (!field.checkValidity()) {
            isValid = false;
            const feedback = field.nextElementSibling;
            if (feedback && feedback.classList.contains('invalid-feedback')) {
              feedback.classList.remove('d-none');
            }
          }
        });

        if (isValid) {
          resolve({
            name: form.name.value,
          });
        } else {
          Swal.showValidationMessage('Please correct the errors in the form.');
        }
      });
    },
    willOpen: () => {
      const form = document.getElementById('add-category');
      form.addEventListener('submit', function (event) {
        event.preventDefault();
        event.stopPropagation();
        document.querySelectorAll('.invalid-feedback').forEach(feedback => {
          feedback.classList.add('d-none');
        });
        let isValid = true;
        form.querySelectorAll('.form-control, .form-select').forEach(field => {
          if (!field.checkValidity()) {
            isValid = false;
            const feedback = field.nextElementSibling;
            if (feedback && feedback.classList.contains('invalid-feedback')) {
              feedback.classList.remove('d-none');
            }
          }
        });
        if (isValid) {
          Swal.getConfirmButton().click();
        } else {
          Swal.showValidationMessage('Please correct the errors in the form.');
        }
      });
    }
  }).then((result) => {
    if (result.isConfirmed) {
      const data = result.value;
      let category = {
        name: data.name,
      }
      const requestOptions = {
        method: 'post',
        headers: {
          'Accept': 'application/json',
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(category),
      }

      fetch('http://localhost:4321/api/inventory/add-category', requestOptions)
        .then(response => response.json())
        .then(data => {
          if (data.error === true) {
            showErrorMessage(data.message)
          } else {
            categories.push(data.result);
            showSuccessMessage(data.message);
            if (page === "purchase") {
              document.getElementById("category").innerHTML = '';
              document.getElementById("category").innerHTML = `<option value="${categories.length-1}" selected>${data.result.name}</option>`;;
              document.getElementById("category").disabled = true;
            }
          }
        });
    }
  });
}
//addNewCustomer show a popup form and then make an api call to insert customer data to the database table
function addNewCustomer(page, customers) {
  Swal.fire({
    title: 'Add Customer',
    width: 400,
    html: `
      <div class="x_panel">
          <div class="x_content">
              <form id="add-customer" class="needs-validation" novalidate>
                  <!-- Account code -->
                  <div class="col-4 form-group has-feedback">
                      <input type="text" class="form-control has-feedback-left" id="account_code" name="account_code"
                          placeholder="Account Code(Autofill: Random)" autocomplete="off">
                      <span style="color: rgba(0, 0, 0, 1); transform:translate(-40%,-10%)" class="form-control-feedback left glyphicon  glyphicon-info-sign" aria-hidden="true"></span>
                  </div>

                  <!-- Account Name -->
                  <div class="col-4 form-group has-feedback">
                      <input type="text" class="form-control has-feedback-left" id="account_name" name="account_name"
                          placeholder="Account Name" autocomplete="off" required>
                      <div class="invalid-feedback d-none text-danger">Please enter the account name.</div>
                      <span style="color: rgba(0, 0, 0, 1); transform:translate(-40%,-10%)" class="form-control-feedback left glyphicon  glyphicon-user" aria-hidden="true"></span>
                  </div>

                  <!-- Contact Person Name -->
                  <div class="col-4 form-group has-feedback">
                      <input type="text" class="form-control has-feedback-left" id="contact_person" name="contact_person"
                          placeholder="Contact Person Name" autocomplete="off" required>
                      <div class="invalid-feedback d-none text-danger">Please enter the contact person name.</div>
                      <span style="color: rgba(0, 0, 0, 1); transform:translate(-40%,-10%)" class="form-control-feedback left glyphicon  glyphicon-user" aria-hidden="true"></span>
                  </div>

                  <!-- Mobile Number -->
                  <div class="col-4 form-group has-feedback">
                      <input type="tel" pattern="[0]{1}[1]{1}[3-9]{1}[0-9]{8}" class="form-control has-feedback-left" id="mobile" name="mobile"
                          placeholder="Mobile Number" autocomplete="off" required>
                      <div class="invalid-feedback d-none text-danger">Please enter a valid mobile number.</div>
                      <span style="color: rgba(0, 0, 0, 1); transform:translate(-40%,-10%)" class="form-control-feedback left glyphicon  glyphicon-phone" aria-hidden="true"></span>
                  </div>

                  <!-- Email -->
                  <div class="col-4 form-group has-feedback">
                      <input type="email" class="form-control has-feedback-left" id="email" name="email"
                          placeholder="Email" autocomplete="off">
                      <div class="invalid-feedback d-none text-danger">Please enter a valid email address.</div>
                      <span style="color: rgba(0, 0, 0, 1); transform:translate(-40%,-10%)" class="form-control-feedback left glyphicon  glyphicon-envelope" aria-hidden="true"></span>
                  </div>

                  <!-- Sale Discount -->
                  <div class="col-4 form-group has-feedback">
                      <input type="number" class="form-control has-feedback-left" id="discount" name="discount"
                          placeholder="discount(%)" min="0" max="100" autocomplete="off">
                      <div class="invalid-feedback d-none text-danger">Enter value between 0 to 100.</div>
                      <span style="color: rgba(0, 0, 0, 1); transform:translate(-40%,-10%)" class="form-control-feedback left glyphicon  glyphicon-gift" aria-hidden="true"></span>
                  </div>

                  <!-- Opening Balance -->
                  <div class="col-4 form-group has-feedback">
                      <input type="number" class="form-control has-feedback-left" id="opening_balance" name="opening_balance"
                          placeholder="Opening Balance(Optional)" autocomplete="off">
                      <span style="color: rgba(0, 0, 0, 1); transform:translate(-40%,-10%)" class="form-control-feedback left glyphicon  glyphicon-plus" aria-hidden="true"></span>
                  </div>
                  <!-- Division -->
                  <div class="col-4 form-group has-feedback">
                      <select id="division" class="form-control form-select has-feedback-left" 
                        onchange="updateChildList('division','district', 'divisionToDistrict')" required>
                          <option value="" selected disabled>Select Division</option>
                          <option value="Barisal">Barisal</option>
                          <option value="Chattogram">Chattogram</option>
                          <option value="Dhaka">Dhaka</option>
                          <option value="Khulna">Khulna</option>
                          <option value="Mymensingh">Mymensingh</option>
                          <option value="Rajshahi">Rajshahi</option>
                          <option value="Rangpur">Rangpur</option>
                          <option value="Sylhet">Sylhet</option>
                      </select>
                      <span style="color: rgba(0, 0, 0, 1); transform:translate(-40%,-10%)"
                        class="form-control-feedback left glyphicon glyphicon-home" aria-hidden="true"></span>
                      <div class="invalid-feedback d-none text-danger">Please select a division.</div>
                  </div>

                  <!-- District -->
                  <div class="col-4 form-group has-feedback">
                      <select id="district" class="form-control form-select has-feedback-left"
                        onchange="updateChildList('district', 'upazila', 'districtToUpazila')" required>
                          <option value="" selected disabled>Select District</option>
                      </select>
                      <span style="color: rgba(0, 0, 0, 1); transform:translate(-40%,-10%)"
                        class="form-control-feedback left glyphicon glyphicon-home" aria-hidden="true"></span>
                      <div class="invalid-feedback d-none text-danger">Please select a district.</div>
                  </div>

                  <!-- Upazila -->
                  <div class="col-4 form-group has-feedback">
                      <select id="upazila" class="form-control form-select has-feedback-left" required>
                          <option value="" selected disabled>Select Upazila</option>
                      </select>
                      <span style="color: rgba(0, 0, 0, 1); transform:translate(-40%,-10%)"
                        class="form-control-feedback left glyphicon glyphicon-home" aria-hidden="true"></span>
                      <div class="invalid-feedback d-none text-danger">Please select an upazila.</div>
                  </div>
                  <!-- Area -->
                  <div class="col-4 form-group has-feedback">
                      <input type="text" class="form-control has-feedback-left" id="area" name="area"
                          placeholder="Road/House No.(Optional)" autocomplete="off">
                      <span style="color: rgba(0, 0, 0, 1); transform:translate(-40%,-10%)" class="form-control-feedback left glyphicon glyphicon-home" aria-hidden="true"></span>
                  </div>
                  <div class="form-group">
                    <div id="btns" class="col-4">
                        <br>
                        <button type="submit" class="btn btn-round btn-success">Submit</button>
                    </div>
                  </div>
              </form>
          </div>
      </div>
    `,
    showCloseButton: true,
    showConfirmButton: false,
    showCancelButton: false,
    allowOutsideClick: false,
    preConfirm: () => {
      return new Promise((resolve) => {
        const form = document.getElementById('add-customer');
        const formFields = form.querySelectorAll('.form-control, .form-select');
        let isValid = true;

        // Reset all feedback messages
        form.querySelectorAll('.invalid-feedback').forEach(feedback => {
          feedback.classList.add('d-none');
        });

        // Check each field
        formFields.forEach(field => {
          if (!field.checkValidity()) {
            isValid = false;
            const feedback = field.nextElementSibling;
            if (feedback && feedback.classList.contains('invalid-feedback')) {
              feedback.classList.remove('d-none');
            }
          }
        });

        if (isValid) {
          resolve({
            account_code: form.account_code.value,
            account_name: form.account_name.value,
            contact_person: form.contact_person.value,
            mobile: form.mobile.value,
            email: form.email.value,
            discount: form.discount.value,
            opening_balance: form.opening_balance.value,
            division: form.division.value,
            district: form.district.value,
            upazila: form.upazila.value,
            area: form.area.value
          });
        } else {
          Swal.showValidationMessage('Please correct the errors in the form.');
        }
      });
    },
    willOpen: () => {
      const form = document.getElementById('add-customer');
      form.addEventListener('submit', function (event) {
        event.preventDefault();
        event.stopPropagation();
        document.querySelectorAll('.invalid-feedback').forEach(feedback => {
          feedback.classList.add('d-none');
        });
        let isValid = true;
        form.querySelectorAll('.form-control, .form-select').forEach(field => {
          if (!field.checkValidity()) {
            isValid = false;
            const feedback = field.nextElementSibling;
            if (feedback && feedback.classList.contains('invalid-feedback')) {
              feedback.classList.remove('d-none');
            }
          }
        });
        if (isValid) {
          Swal.getConfirmButton().click();
        } else {
          Swal.showValidationMessage('Please correct the errors in the form.');
        }
      });
    }
  }).then((result) => {
    if (result.isConfirmed) {
      const data = result.value;
      dis = parseInt(data.discount, 10);
      opBalance = parseInt(data.opening_balance, 10);
      let customer = {
        account_code: data.account_code,
        account_name: data.account_name,
        contact_person: data.contact_person,
        mobile: data.mobile,
        email: data.email,
        discount: dis,
        opening_balance: opBalance,
        division: data.division,
        district: data.district,
        upazila: data.upazila,
        area: data.area
      }
      const requestOptions = {
        method: 'post',
        headers: {
          'Accept': 'application/json',
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(customer),
      }

      fetch('http://localhost:4321/api/mis/add-customer', requestOptions)
        .then(response => response.json())
        .then(data => {
          if (data.error === true) {
            showErrorMessage(data.message)
          } else {
            showSuccessMessage(data.message);
            customers.push(data.result)
            if (page === "sale") {
              document.getElementById("customer").innerHTML = '';
              document.getElementById("customer").innerHTML = `<option value="${customers.length-1}" selected>${data.result.account_name} (${data.result.account_code})</option>`;;
              document.getElementById("customer").disabled = true;
            } else {
              setTimeout(function () {
                location.reload();
              }, 3000); // Adjust the delay as needed 
            }
          }
        });
    }
  });
}
//addNewEmployee show a popup form and then make an api call to insert employee data to the database table
function addNewEmployee(page, employees) {
  Swal.fire({
    title: 'Add Employee',
    width: 400,
    html: `
              <div class="x_panel">
                  <div class="x_content">
                      <form id="add-employee" class="needs-validation" novalidate>
                          <!-- Account code -->
                          <div class="col-4 form-group has-feedback">
                              <input type="text" class="form-control has-feedback-left" id="account_code" name="account_code"
                                  placeholder="Account Code(Autofill: Random)" autocomplete="off">
                              <span style="color: rgba(0, 0, 0, 1); transform:translate(-40%,-10%)" class="form-control-feedback left glyphicon  glyphicon-info-sign" aria-hidden="true"></span>
                          </div>
  
                          <!-- Account Name -->
                          <div class="col-4 form-group has-feedback">
                              <input type="text" class="form-control has-feedback-left" id="account_name" name="account_name"
                                  placeholder="Account Name" autocomplete="off" required>
                              <div class="invalid-feedback d-none text-danger">Please enter the account name.</div>
                              <span style="color: rgba(0, 0, 0, 1); transform:translate(-40%,-10%)" class="form-control-feedback left glyphicon  glyphicon-user" aria-hidden="true"></span>
                          </div>
  
                          <!-- Contact Person Name -->
                          <div class="col-4 form-group has-feedback">
                              <input type="text" class="form-control has-feedback-left" id="contact_person" name="contact_person"
                                  placeholder="Contact Person Name" autocomplete="off" required>
                              <div class="invalid-feedback d-none text-danger">Please enter the contact person name.</div>
                              <span style="color: rgba(0, 0, 0, 1); transform:translate(-40%,-10%)" class="form-control-feedback left glyphicon  glyphicon-user" aria-hidden="true"></span>
                          </div>
  
                          <!-- Mobile Number -->
                          <div class="col-4 form-group has-feedback">
                              <input type="tel" pattern="[0]{1}[1]{1}[3-9]{1}[0-9]{8}" class="form-control has-feedback-left" id="mobile" name="mobile"
                                  placeholder="Mobile Number" autocomplete="off" required>
                              <div class="invalid-feedback d-none text-danger">Please enter a valid mobile number.</div>
                              <span style="color: rgba(0, 0, 0, 1); transform:translate(-40%,-10%)" class="form-control-feedback left glyphicon  glyphicon-phone" aria-hidden="true"></span>
                          </div>
  
                          <!-- Email -->
                          <div class="col-4 form-group has-feedback">
                              <input type="email" class="form-control has-feedback-left" id="email" name="email"
                                  placeholder="Email" autocomplete="off">
                              <div class="invalid-feedback d-none text-danger">Please enter a valid email address.</div>
                              <span style="color: rgba(0, 0, 0, 1); transform:translate(-40%,-10%)" class="form-control-feedback left glyphicon  glyphicon-envelope" aria-hidden="true"></span>
                          </div>
  
                          <!-- Monthly Salary -->
                          <div class="col-4 form-group has-feedback">
                              <input type="number" class="form-control has-feedback-left" id="monthly_salary" name="monthly_salary"
                                  placeholder="Monthly Salary" autocomplete="off" required>
                              <div class="invalid-feedback d-none text-danger">Please enter the monthly salary.</div>
                              <span style="color: rgba(0, 0, 0, 1); transform:translate(-40%,-10%)" class="form-control-feedback left glyphicon  glyphicon-gift" aria-hidden="true"></span>
                          </div>
  
                          <!-- Opening Balance -->
                          <div class="col-4 form-group has-feedback">
                              <input type="number" class="form-control has-feedback-left" id="opening_balance" name="opening_balance"
                                  placeholder="Opening Balance" autocomplete="off">
                              <span style="color: rgba(0, 0, 0, 1); transform:translate(-40%,-10%)" class="form-control-feedback left glyphicon  glyphicon-plus" aria-hidden="true"></span>
                          </div>
                          <!-- Division -->
                          <div class="col-4 form-group has-feedback">
                              <select id="division" class="form-control form-select has-feedback-left" 
                                onchange="updateChildList('division','district', 'divisionToDistrict')" required>
                                  <option value="" selected disabled>Select Division</option>
                                  <option value="Barisal">Barisal</option>
                                  <option value="Chattogram">Chattogram</option>
                                  <option value="Dhaka">Dhaka</option>
                                  <option value="Khulna">Khulna</option>
                                  <option value="Mymensingh">Mymensingh</option>
                                  <option value="Rajshahi">Rajshahi</option>
                                  <option value="Rangpur">Rangpur</option>
                                  <option value="Sylhet">Sylhet</option>
                              </select>
                              <span style="color: rgba(0, 0, 0, 1); transform:translate(-40%,-10%)"
                                class="form-control-feedback left glyphicon glyphicon-home" aria-hidden="true"></span>
                              <div class="invalid-feedback d-none text-danger">Please select a division.</div>
                          </div>
  
                          <!-- District -->
                          <div class="col-4 form-group has-feedback">
                              <select id="district" class="form-control form-select has-feedback-left"
                                onchange="updateChildList('district', 'upazila', 'districtToUpazila')" required>
                                  <option value="" selected disabled>Select District</option>
                              </select>
                              <span style="color: rgba(0, 0, 0, 1); transform:translate(-40%,-10%)"
                                class="form-control-feedback left glyphicon glyphicon-home" aria-hidden="true"></span>
                              <div class="invalid-feedback d-none text-danger">Please select a district.</div>
                          </div>
  
                          <!-- Upazila -->
                          <div class="col-4 form-group has-feedback">
                              <select id="upazila" class="form-control form-select has-feedback-left" required>
                                  <option value="" selected disabled>Select Upazila</option>
                              </select>
                              <span style="color: rgba(0, 0, 0, 1); transform:translate(-40%,-10%)"
                                class="form-control-feedback left glyphicon glyphicon-home" aria-hidden="true"></span>
                              <div class="invalid-feedback d-none text-danger">Please select an upazila.</div>
                          </div>
                          <!-- Area -->
                          <div class="col-4 form-group has-feedback">
                              <input type="text" class="form-control has-feedback-left" id="area" name="area"
                                  placeholder="Road/House No." autocomplete="off">
                              <span style="color: rgba(0, 0, 0, 1); transform:translate(-40%,-10%)" class="form-control-feedback left glyphicon glyphicon-home" aria-hidden="true"></span>
                          </div>
                          <div class="form-group">
                            <div id="btns" class="col-4">
                                <br>
                                <button type="submit" class="btn btn-round btn-success">Submit</button>
                            </div>
                          </div>
                      </form>
                  </div>
              </div>
          `,
    showCloseButton: true,
    showConfirmButton: false,
    showCancelButton: false,
    allowOutsideClick: false,
    preConfirm: () => {
      return new Promise((resolve) => {
        const form = document.getElementById('add-employee');
        const formFields = form.querySelectorAll('.form-control, .form-select');
        let isValid = true;

        // Reset all feedback messages
        form.querySelectorAll('.invalid-feedback').forEach(feedback => {
          feedback.classList.add('d-none');
        });

        // Check each field
        formFields.forEach(field => {
          if (!field.checkValidity()) {
            isValid = false;
            const feedback = field.nextElementSibling;
            if (feedback && feedback.classList.contains('invalid-feedback')) {
              feedback.classList.remove('d-none');
            }
          }
        });

        if (isValid) {
          resolve({
            account_code: form.account_code.value,
            account_name: form.account_name.value,
            contact_person: form.contact_person.value,
            mobile: form.mobile.value,
            email: form.email.value,
            monthly_salary: form.monthly_salary.value,
            opening_balance: form.opening_balance.value,
            division: form.division.value,
            district: form.district.value,
            upazila: form.upazila.value,
            area: form.area.value
          });
        } else {
          Swal.showValidationMessage('Please correct the errors in the form.');
        }
      });
    },
    willOpen: () => {
      const form = document.getElementById('add-employee');
      form.addEventListener('submit', function (event) {
        event.preventDefault();
        event.stopPropagation();
        document.querySelectorAll('.invalid-feedback').forEach(feedback => {
          feedback.classList.add('d-none');
        });
        let isValid = true;
        form.querySelectorAll('.form-control, .form-select').forEach(field => {
          if (!field.checkValidity()) {
            isValid = false;
            const feedback = field.nextElementSibling;
            if (feedback && feedback.classList.contains('invalid-feedback')) {
              feedback.classList.remove('d-none');
            }
          }
        });
        if (isValid) {
          Swal.getConfirmButton().click();
        } else {
          Swal.showValidationMessage('Please correct the errors in the form.');
        }
      });
    }
  }).then((result) => {
    if (result.isConfirmed) {
      const data = result.value;
      salary = parseInt(data.monthly_salary, 10);
      opBalance = parseInt(data.opening_balance, 10);
      let employee = {
        account_code: data.account_code,
        account_name: data.account_name,
        contact_person: data.contact_person,
        mobile: data.mobile,
        email: data.email,
        monthly_salary: salary,
        opening_balance: opBalance,
        division: data.division,
        district: data.district,
        upazila: data.upazila,
        area: data.area
      }
      const requestOptions = {
        method: 'post',
        headers: {
          'Accept': 'application/json',
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(employee),
      }

      fetch('http://localhost:4321/api/hr/add-employee', requestOptions)
        .then(response => response.json())
        .then(data => {
          if (page === "") {
            setTimeout(function () {
              location.reload();
            }, 3000); // Adjust the delay as needed 
          }
          if (data.error === true) {
            showErrorMessage(data.message)
          } else {
            employees.push(data.result)
            showSuccessMessage(data.message);
          }
        });
    }
  });
}
//addNewSupplier show a popup form and then make an api call to insert supplier data to the database table
function addNewSupplier(page, suppliers) {
  Swal.fire({
    title: 'Add Supplier',
    width: 400,
    html: `
      <div class="x_panel">
          <div class="x_content">
              <form id="add-supplier" class="needs-validation" novalidate>
                  <!-- Account code -->
                  <div class="col-4 form-group has-feedback">
                      <input type="text" class="form-control has-feedback-left" id="account_code" name="account_code"
                          placeholder="Account Code(Autofill: Random)" autocomplete="off">
                      <span style="color: rgba(0, 0, 0, 1); transform:translate(-40%,-10%)" class="form-control-feedback left glyphicon  glyphicon-info-sign" aria-hidden="true"></span>
                  </div>

                  <!-- Account Name -->
                  <div class="col-4 form-group has-feedback">
                      <input type="text" class="form-control has-feedback-left" id="account_name" name="account_name"
                          placeholder="Account Name" autocomplete="off" required>
                      <div class="invalid-feedback d-none text-danger">Please enter the account name.</div>
                      <span style="color: rgba(0, 0, 0, 1); transform:translate(-40%,-10%)" class="form-control-feedback left glyphicon  glyphicon-user" aria-hidden="true"></span>
                  </div>

                  <!-- Contact Person Name -->
                  <div class="col-4 form-group has-feedback">
                      <input type="text" class="form-control has-feedback-left" id="contact_person" name="contact_person"
                          placeholder="Contact Person Name" autocomplete="off" required>
                      <div class="invalid-feedback d-none text-danger">Please enter the contact person name.</div>
                      <span style="color: rgba(0, 0, 0, 1); transform:translate(-40%,-10%)" class="form-control-feedback left glyphicon  glyphicon-user" aria-hidden="true"></span>
                  </div>

                  <!-- Mobile Number -->
                  <div class="col-4 form-group has-feedback">
                      <input type="tel" pattern="[0]{1}[1]{1}[3-9]{1}[0-9]{8}" class="form-control has-feedback-left" id="mobile" name="mobile"
                          placeholder="Mobile Number" autocomplete="off" required>
                      <div class="invalid-feedback d-none text-danger">Please enter a valid mobile number.</div>
                      <span style="color: rgba(0, 0, 0, 1); transform:translate(-40%,-10%)" class="form-control-feedback left glyphicon  glyphicon-phone" aria-hidden="true"></span>
                  </div>

                  <!-- Email -->
                  <div class="col-4 form-group has-feedback">
                      <input type="email" class="form-control has-feedback-left" id="email" name="email"
                          placeholder="Email" autocomplete="off">
                      <div class="invalid-feedback d-none text-danger">Please enter a valid email address.</div>
                      <span style="color: rgba(0, 0, 0, 1); transform:translate(-40%,-10%)" class="form-control-feedback left glyphicon  glyphicon-envelope" aria-hidden="true"></span>
                  </div>

                  <!-- Purchase Discount -->
                  <div class="col-4 form-group has-feedback">
                      <input type="number" class="form-control has-feedback-left" id="discount" name="discount"
                          placeholder="discount(%)" min="0" max="100" autocomplete="off" >
                      <div class="invalid-feedback d-none text-danger">Enter value between 0 to 100..</div>
                      <span style="color: rgba(0, 0, 0, 1); transform:translate(-40%,-10%)" class="form-control-feedback left glyphicon  glyphicon-gift" aria-hidden="true"></span>
                  </div>

                  <!-- Division -->
                  <div class="col-4 form-group has-feedback">
                      <select id="division" class="form-control form-select has-feedback-left" 
                        onchange="updateChildList('division','district', 'divisionToDistrict')" required>
                          <option value="" selected disabled>Select Division</option>
                          <option value="Barisal">Barisal</option>
                          <option value="Chattogram">Chattogram</option>
                          <option value="Dhaka">Dhaka</option>
                          <option value="Khulna">Khulna</option>
                          <option value="Mymensingh">Mymensingh</option>
                          <option value="Rajshahi">Rajshahi</option>
                          <option value="Rangpur">Rangpur</option>
                          <option value="Sylhet">Sylhet</option>
                      </select>
                      <span style="color: rgba(0, 0, 0, 1); transform:translate(-40%,-10%)"
                        class="form-control-feedback left glyphicon glyphicon-home" aria-hidden="true"></span>
                      <div class="invalid-feedback d-none text-danger">Please select a division.</div>
                  </div>

                  <!-- District -->
                  <div class="col-4 form-group has-feedback">
                      <select id="district" class="form-control form-select has-feedback-left"
                        onchange="updateChildList('district', 'upazila', 'districtToUpazila')" required>
                          <option value="" selected disabled>Select District</option>
                      </select>
                      <span style="color: rgba(0, 0, 0, 1); transform:translate(-40%,-10%)"
                        class="form-control-feedback left glyphicon glyphicon-home" aria-hidden="true"></span>
                      <div class="invalid-feedback d-none text-danger">Please select a district.</div>
                  </div>

                  <!-- Upazila -->
                  <div class="col-4 form-group has-feedback">
                      <select id="upazila" class="form-control form-select has-feedback-left" required>
                          <option value="" selected disabled>Select Upazila</option>
                      </select>
                      <span style="color: rgba(0, 0, 0, 1); transform:translate(-40%,-10%)"
                        class="form-control-feedback left glyphicon glyphicon-home" aria-hidden="true"></span>
                      <div class="invalid-feedback d-none text-danger">Please select an upazila.</div>
                  </div>
                  <!-- Area -->
                  <div class="col-4 form-group has-feedback">
                      <input type="text" class="form-control has-feedback-left" id="area" name="area"
                          placeholder="Road/House No.(Optional)" autocomplete="off">
                      <span style="color: rgba(0, 0, 0, 1); transform:translate(-40%,-10%)" class="form-control-feedback left glyphicon glyphicon-home" aria-hidden="true"></span>
                  </div>
                  <div class="form-group">
                    <div id="btns" class="col-4">
                        <br>
                        <button type="submit" class="btn btn-round btn-success">Submit</button>
                    </div>
                  </div>
              </form>
          </div>
      </div>
    `,
    showCloseButton: true,
    showConfirmButton: false,
    showCancelButton: false,
    allowOutsideClick: false,
    preConfirm: () => {
      return new Promise((resolve) => {
        const form = document.getElementById('add-supplier');
        const formFields = form.querySelectorAll('.form-control, .form-select');
        let isValid = true;

        // Reset all feedback messages
        form.querySelectorAll('.invalid-feedback').forEach(feedback => {
          feedback.classList.add('d-none');
        });

        // Check each field
        formFields.forEach(field => {
          if (!field.checkValidity()) {
            isValid = false;
            const feedback = field.nextElementSibling;
            if (feedback && feedback.classList.contains('invalid-feedback')) {
              feedback.classList.remove('d-none');
            }
          }
        });

        if (isValid) {
          resolve({
            account_code: form.account_code.value,
            account_name: form.account_name.value,
            contact_person: form.contact_person.value,
            mobile: form.mobile.value,
            email: form.email.value,
            discount: form.discount.value,
            division: form.division.value,
            district: form.district.value,
            upazila: form.upazila.value,
            area: form.area.value
          });
        } else {
          Swal.showValidationMessage('Please correct the errors in the form.');
        }
      });
    },
    willOpen: () => {
      const form = document.getElementById('add-supplier');
      form.addEventListener('submit', function (event) {
        event.preventDefault();
        event.stopPropagation();
        document.querySelectorAll('.invalid-feedback').forEach(feedback => {
          feedback.classList.add('d-none');
        });
        let isValid = true;
        form.querySelectorAll('.form-control, .form-select').forEach(field => {
          if (!field.checkValidity()) {
            isValid = false;
            const feedback = field.nextElementSibling;
            if (feedback && feedback.classList.contains('invalid-feedback')) {
              feedback.classList.remove('d-none');
            }
          }
        });
        if (isValid) {
          Swal.getConfirmButton().click();
        } else {
          Swal.showValidationMessage('Please correct the errors in the form.');
        }
      });
    }
  }).then((result) => {
    if (result.isConfirmed) {
      const data = result.value;
      dis = parseInt(data.discount, 10);
      let supplier = {
        account_code: data.account_code,
        account_name: data.account_name,
        contact_person: data.contact_person,
        mobile: data.mobile,
        email: data.email,
        discount: dis,
        division: data.division,
        district: data.district,
        upazila: data.upazila,
        area: data.area
      }
      const requestOptions = {
        method: 'post',
        headers: {
          'Accept': 'application/json',
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(supplier),
      }

      fetch('http://localhost:4321/api/mis/add-supplier', requestOptions)
        .then(response => response.json())
        .then(data => {
          if (data.error === true) {
            showErrorMessage(data.message)
          } else {
            suppliers.push(data.result)
            showSuccessMessage(data.message);
            if (page === "purchase") {
              document.getElementById("supplier").innerHTML = '';
              document.getElementById("supplier").innerHTML = `<option value="${suppliers.length-1}" selected>${data.result.account_name} (${data.result.account_code})</option>`;;
              document.getElementById("supplier").disabled = true;
            } else {
              setTimeout(function () {
                location.reload();
              }, 3000); // Adjust the delay as needed 
            }
          }
        });
    }
  });
}
//checkoutWarrantyProducts show a popup checkout warranty for and make an api call to update database table for checkout process
function checkoutWarrantyProducts(warrantyHistoryID, productSerialNo, productSerialID) {

  let htmlContent = `
        <div class="x_panel">
            <div class="x_content">
                <form id="checkout-wp" class="needs-validation" novalidate>
                    <!-- Arrival Date -->
                    <div class="col-6 form-group has-feedback">
                        <input type="date" class="form-control has-feedback-left"  id="arrival_date" name="arrival_date"
                            placeholder="Enter date(mm-dd-yyyy)" autocomplete="off" required>
                        <div class="invalid-feedback d-none text-danger">Please correct date.</div>
                        <span style="color: rgba(0, 0, 0, 1); transform:translate(-40%,-10%)" class="form-control-feedback left glyphicon glyphicon-calendar" aria-hidden="true"></span>
                    </div>
                    <!-- New SN Name -->
                    <div class="col-6 form-group has-feedback">
                        <input type="text" class="form-control has-feedback-left" id="new_sn" name="new_sn"
                            placeholder="New S/N (enter previous S/N if warranty repair)" autocomplete="off" required>
                        <div class="invalid-feedback d-none text-danger">For old items, enter previous serial; for new, enter a new serial."</div>
                        <span style="color: rgba(0, 0, 0, 1); transform:translate(-40%,-10%)" class="form-control-feedback left glyphicon glyphicon-barcode" aria-hidden="true"></span>
                    </div>
                    <div class="col-6 form-group has-feedback">
                        <input type="text" class="form-control has-feedback-left" id="comment" name="comment"
                            placeholder="Comment" autocomplete="off">
                        <div class="invalid-feedback d-none text-danger">Please enter the product name.</div>
                        <span style="color: rgba(0, 0, 0, 1); transform:translate(-40%,-10%)" class="form-control-feedback left fa fa-file-text-o" aria-hidden="true"></span>
                    </div>
                    <div class="form-group">
                      <div id="btns" class="col-4">
                          <br>
                          <button type="submit" class="btn btn-danger">Submit</button>
                      </div>
                    </div>
                </form>
            </div>
        </div>
          `;
  Swal.fire({
    title: 'Checkout Warranty',
    width: 600,
    html: htmlContent,
    showCloseButton: true,
    showConfirmButton: false,
    showCancelButton: false,
    allowOutsideClick: false,
    preConfirm: () => {
      return new Promise((resolve) => {
        const form = document.getElementById('checkout-wp');
        const formFields = form.querySelectorAll('.form-control, .form-select');
        let isValid = true;

        // Reset all feedback messages
        form.querySelectorAll('.invalid-feedback').forEach(feedback => {
          feedback.classList.add('d-none');
        });

        // Check each field
        formFields.forEach(field => {
          if (!field.checkValidity()) {
            isValid = false;
            const feedback = field.nextElementSibling;
            if (feedback && feedback.classList.contains('invalid-feedback')) {
              feedback.classList.remove('d-none');
            }
          }
        });

        if (isValid) {
          resolve({
            arrival_date: form.arrival_date.value,
            new_s_n: form.new_sn.value,
            comment: form.comment.value,
          });
        } else {
          Swal.showValidationMessage('Please correct the errors in the form.');
        }
      });
    },
    willOpen: () => {
      const form = document.getElementById('checkout-wp');
      form.addEventListener('submit', function (event) {
        event.preventDefault();
        event.stopPropagation();
        document.querySelectorAll('.invalid-feedback').forEach(feedback => {
          feedback.classList.add('d-none');
        });
        let isValid = true;
        form.querySelectorAll('.form-control, .form-select').forEach(field => {
          if (!field.checkValidity()) {
            isValid = false;
            const feedback = field.nextElementSibling;
            if (feedback && feedback.classList.contains('invalid-feedback')) {
              feedback.classList.remove('d-none');
            }
          }
        });
        if (isValid) {
          Swal.getConfirmButton().click();
        } else {
          Swal.showValidationMessage('Please correct the errors in the form.');
        }
      });
    }
  }).then((result) => {
    if (result.isConfirmed) {
      const data = result.value;
      let wpData = {
        warranty_history_id: warrantyHistoryID,
        product_serial_id: productSerialID,
        old_serial_number: productSerialNo,
        checkout_date: data.arrival_date,
        new_serial_number: data.new_s_n,
        comment: data.comment,
      }
      const requestOptions = {
        method: 'post',
        headers: {
          'Accept': 'application/json',
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(wpData),
      }
      console.log(wpData)

      fetch('http://localhost:4321/api/inventory/products/warranty/checkout', requestOptions)
        .then(response => response.json())
        .then(data => {
          console.log(data)
          if (data.error === true) {
            showErrorMessage(data.message)
          } else {
            showSuccessMessage(data.message);
          }
        });
    }
  });
}
//Function to show SweetAlert2 prompt and handle the delivery process for a warranty product
function confirmWarrantyDeliveryProcess(warrantyHistoryID, productSerialID) {
  // Display a SweetAlert2 confirmation prompt to proceed with delivery
  Swal.fire({
    title: 'Proceed with delivery?',   // Prompt message
    showCancelButton: true,            // Show the cancel button
    confirmButtonText: 'Yes',          // Text for the confirm button
    cancelButtonText: 'Cancel',        // Text for the cancel button
    icon: 'warning',                   // Icon type to indicate a warning action
  }).then((result) => {
    // Check if the user confirmed by clicking "Yes"
    if (result.isConfirmed) {
      // Data object containing warranty history ID and product serial ID
      let wpData = {
        warranty_history_id: warrantyHistoryID,
        product_serial_id: productSerialID,
      };

      // Fetch API request options, including method, headers, and body (as JSON)
      const requestOptions = {
        method: 'post',
        headers: {
          'Accept': 'application/json',
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(wpData),  // Convert data object to a JSON string
      };

      // Log the data being sent to the API (for debugging purposes)
      console.log(wpData);

      // Send a POST request to the API endpoint for warranty delivery
      fetch('http://localhost:4321/api/inventory/products/warranty/delivery', requestOptions)
        .then(response => response.json())  // Parse the response as JSON
        .then(data => {
          // Check if the API returned an error
          if (data.error === true) {
            showErrorMessage(data.message);  // Show error message if there's an issue
          } else {
            showSuccessMessage(data.message);  // Show success message on success
            setTimeout(function () {
              location.reload();
            }, 3000); // Adjust the delay as needed 
          }
        })
        .catch(error => {
          // Handle any errors that occur during the fetch request
          console.error('Error during the request:', error);
          showErrorMessage('An error occurred during the delivery process.');
        });
    }
  });
}
//viewWarrantyHistory shows warranty history
function viewWarrantyHistory(warrantyHistory) {
  let mm = warrantyHistory.memo_no.split("-").pop() //get the last part of the memo
  // Display a SweetAlert2 confirmation prompt to proceed with delivery
  Swal.fire({
    width: 1024,
    title: 'Warranty Information',   // Prompt message
    showCancelButton: false,            // Show the cancel button
    confirmButtonText: 'Ok',          // Text for the confirm button
    // cancelButtonText: 'Cancel',        // Text for the cancel button
    allowOutsideClick: false,          //disable outside click

    html: `<div class="row">
          <div class="col-md-6 col-sm-6 col-xs-12">
            <!-- Table row -->
            <table class="table table-striped">
              <caption>Product Reception</caption>
              <tbody>
                <tr>
                  <td>Invoice No:</td>
                  <td>
                    <b>MM-WC-${mm}</b>
                  </td>
                </tr>
                <tr>
                  <td>Old S/N:</td>
                  <td><b>${warrantyHistory.previous_serial_no}</b></td>
                </tr>
                <tr>
                  <td>Received By:</td>
                  <td><b>${warrantyHistory.received_by ? warrantyHistory.received_by : ""}</b></td>
                </tr>
                <tr>
                  <td>Receiption Date:</td>
                  <td><b>${formatDate(warrantyHistory.requested_date, "date", "-")}</b></td>
                </tr>
                <tr>
                  <td>Complain:</td>
                  <td><b>${warrantyHistory.reported_problem}</b></td>
                </tr>
                <tr>
                  <td></td>
                  <td></td>
                </tr>
              </tbody>
            </table>
          </div>
  
          <div class="col-md-6 col-sm-6 col-xs-12">
            <!-- Table row -->
            <table class="table table-striped">
              <caption>Delivery Information</caption>
              <tbody>
                <tr>
                  <td>Invoice No:</td>
                  <td>
                    <b>MM-WD-${mm}</b>
                  </td>
                </tr>
                 <tr>
                  <td>New S/N:</td>
                  <td><b>${warrantyHistory.new_serial_no}</b></td>
                </tr>
                <tr>
                  <td>Delivered By:</td>
                  <td><b>${warrantyHistory.delivered_by ? warrantyHistory.delivered_by : ""}
                </tr>
                <tr>
                  <td>Delivery Date:</td>
                  <td><b>${formatDate(warrantyHistory.delivery_date, "date", "-")}</b></td>
                </tr>
                <tr>
                  <td>Comment:</td>
                  <td><b>${warrantyHistory.comment}</b></td>
                </tr>
                
                <tr>
                  <td></td>
                  <td></td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      `
  }).then((result) => {
  });
}
// Function to initialize the SweetAlert2 form with date range picker
function showDateRangePickerPopup() {
  Swal.fire({
    title: 'Select Date Range',
    html: '<div id="daterange-container" style="cursor: pointer; padding: 5px; border: 1px solid #ccc; width: 100%;">' +
      '<span>Click to select date range</span> <i class="fa fa-calendar"></i></div>',
    showCancelButton: true,
    confirmButtonText: 'Submit',
    width: 360,
    preConfirm: () => {
      const daterangepicker = $('#daterange-container').data('daterangepicker');
      const startDate = daterangepicker.startDate.format('MM/DD/YYYY');
      const endDate = daterangepicker.endDate.format('MM/DD/YYYY');
      return { startDate, endDate };
    }
  }).then((result) => {
    if (result.isConfirmed) {
      const { startDate, endDate } = result.value;
      console.log("Pop:", startDate, endDate)
      showFilteredReport(startDate, endDate);
    }
  });

  // Initialize the date range picker in the SweetAlert popup
  DateRangePicker_Cal('daterange-container');
}