
//prevent user mouse right click
// document.addEventListener('contextmenu', function (e) {
//   e.preventDefault();
// });

function paginator(currPageIndex, pageSize, totalRecords) {
  let pNav = document.getElementById("paginate_nav");
  let pInfo = document.getElementById("paginator_info");
  if (pNav && pInfo) {
    let startID = (currPageIndex - 1) * pageSize + 1;
    let endID = Math.min(startID + pageSize - 1, totalRecords);
    pInfo.innerHTML = `Showing <strong>${startID}</strong> to <strong>${endID}</strong> of <strong>${totalRecords}</strong> entries`;

    let htmlTmpl = ``;


    if (currPageIndex > 1) {
      htmlTmpl += `<li class="page-product"><a class="page-link" href="#" onclick="updatePage(${pageSize}, ${currPageIndex - 1})">Previous</a></li>`;
    } else {
      htmlTmpl += `<li class="page-product disabled"><a class="page-link" href="#">Previous</a></li>`;
    }
    pages = Math.ceil(totalRecords / pageSize)
    for (let i = 1; i <= pages; i++) {
      htmlTmpl += `<li class="page-product ${i === currPageIndex ? 'active' : ''}"><a class="page-link" href="#" onclick="updatePage(${pageSize}, ${i})">${i}</a></li>`;
    }

    if (currPageIndex == pages) {
      htmlTmpl += `<li class="page-product disabled"><a class="page-link" href="#">Next</a></li>`;
    } else {
      htmlTmpl += `<li class="page-product"><a class="page-link" href="#" onclick="updatePage(${pageSize}, ${currPageIndex + 1})">Next</a></li>`;
    }

    pNav.innerHTML = htmlTmpl;
  }
}
/* DATA TABLES */

function init_DataTables() {

  console.log('run_datatables');

  if (typeof ($.fn.DataTable) === 'undefined') { return; }
  console.log('init_DataTables');

  var handleDataTableButtons = function () {
    if ($("#datatable-buttons").length) {
      $("#datatable-buttons").DataTable({
        dom: "Bfrtip",
        buttons: [
          {
            extend: "copy",
            className: "btn-sm"
          },
          {
            extend: "csv",
            className: "btn-sm"
          },
          {
            extend: "excel",
            className: "btn-sm"
          },
          {
            extend: "pdfHtml5",
            className: "btn-sm"
          },
          {
            extend: "print",
            className: "btn-sm"
          },
        ],
        responsive: true
      });
    }
  };

  TableManageButtons = function () {
    "use strict";
    return {
      init: function () {
        handleDataTableButtons();
      }
    };
  }();

  $('#datatable').dataTable();

  $('#datatable-keytable').DataTable({
    keys: true
  });

  $('#datatable-responsive').DataTable();

  // $('#warranty-inprogress-table').DataTable();

  $('#datatable-scroller').DataTable({
    ajax: "js/datatables/json/scroller-demo.json",
    deferRender: true,
    scrollY: 380,
    scrollCollapse: true,
    scroller: true
  });

  $('#datatable-fixed-header').DataTable({
    fixedHeader: true
  });

  var $datatable = $('#datatable-checkbox');

  $datatable.dataTable({
    'order': [[1, 'asc']],
    'columnDefs': [
      { orderable: false, targets: [0] }
    ]
  });
  $datatable.on('draw.dt', function () {
    $('checkbox input').iCheck({
      checkboxClass: 'icheckbox_flat-green'
    });
  });

  TableManageButtons.init();

};

/**
 * Generates a memo number in the format MM-memo_type-rand(digit6 int)LastIndexOfDBTable.
 * 
 * The memo number consists of:
 * - rand(digit6 int): A random 6-digit integer
 *
 * @returns {string} The formatted memo number.
 */
function generateMemoNumber(lastIndex, memoType) {
  // Generate a random 6-digit number (from 100000 to 999999)
  const randomSixDigit = String(Math.floor(100000 + Math.random() * 900000));

  // Combine all parts into the desired format
  const memoNumber = `MM-${memoType}-${randomSixDigit}${lastIndex}`;

  return memoNumber; // Return the formatted serial number
}


/*formatDate returns formattedDate for the given time, 
if format = "date", returns date only, 
if format = "time", returns time only 
and returns both date and time for format = "" */
function formatDate(time, format, separator) {
  // Parse the input string into a Date object
  const date = new Date(time);

  // Define arrays for month names and zero-padding for formatting
  const months = ["January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"];
  const pad = (num) => num.toString().padStart(2, '0');

  // Extract date components
  const day = pad(date.getDate());
  const month = months[date.getMonth()];
  const year = date.getFullYear();
  const hours = pad(date.getHours());
  const minutes = pad(date.getMinutes());
  const seconds = pad(date.getSeconds());

  //fix separator
  if (separator === "") {
    separator = `/`
  }
  // Format the date
  let formattedDate
  if (format === "") {
    formattedDate = `${day}${separator}${month}${separator}${year} ${hours}:${minutes}:${seconds}`;
  } else if (format == "date") {
    formattedDate = `${day}${separator}${month}${separator}${year}`
  } else if (format == "time") {
    formattedDate = `${hours}:${minutes}:${seconds}`
  }
  return formattedDate
}

function stringToDate(dateString, separator = '-', format = 'mm-dd-yyyy') {
  // Split the date string using the specified separator
  const parts = dateString.split(separator);

  let day, month, year;

  // Assign values based on the specified format
  if (format === 'mm-dd-yyyy') {
    month = parseInt(parts[0]) - 1; // Month (0-based)
    day = parseInt(parts[1]);
    year = parseInt(parts[2]);
  } else if (format === 'dd-mm-yyyy') {
    day = parseInt(parts[0]);
    month = parseInt(parts[1]) - 1; // Month (0-based)
    year = parseInt(parts[2]);
  } else if (format === 'yyyy-mm-dd') {
    year = parseInt(parts[0]);
    month = parseInt(parts[1]) - 1; // Month (0-based)
    day = parseInt(parts[2]);
  } else {
    throw new Error('Unsupported date format. Use mm-dd-yyyy, dd-mm-yyyy, or yyyy-mm-dd.');
  }

  // Create and return a Date object
  return new Date(year, month, day);
}

// setCurrentDate sets current date(mm-dd-yyyy) to the date input by its id
function setCurrentDate(dateInputField) {
  document.getElementById(dateInputField).value = getCurrentDate(); // Set value of the input
}
// getCurrentDate returns current date(mm-dd-yyyy) 
function getCurrentDate() {
  const today = new Date();
  const month = String(today.getMonth() + 1).padStart(2, '0'); // Months are zero-based
  const day = String(today.getDate()).padStart(2, '0');
  const year = today.getFullYear();

  return `${month}-${day}-${year}`; // Format as MM-DD-YYYY
}

// Example usage:
// const soldDate = '01-10-2022'; // mm-dd-yyyy format
// const warrantyPeriod = 365; // 1 year warranty
// const result = checkWarranty(soldDate, warrantyPeriod, '-', 'mm-dd-yyyy');
// console.log(result);
function checkWarranty(soldDate, warrantyPeriodInDays, separator = '-', format = 'mm-dd-yyyy') {
  // Convert the sold date string to a Date object
  const parts = soldDate.split(separator);
  let day, month, year;

  // Assign values based on the specified format
  if (format === 'mm-dd-yyyy') {
    month = parseInt(parts[0]) - 1; // Month (0-based)
    day = parseInt(parts[1]);
    year = parseInt(parts[2]);
  } else if (format === 'dd-mm-yyyy') {
    day = parseInt(parts[0]);
    month = parseInt(parts[1]) - 1; // Month (0-based)
    year = parseInt(parts[2]);
  } else if (format === 'yyyy-mm-dd') {
    year = parseInt(parts[0]);
    month = parseInt(parts[1]) - 1; // Month (0-based)
    day = parseInt(parts[2]);
  } else {
    throw new Error('Unsupported date format. Use mm-dd-yyyy, dd-mm-yyyy, or yyyy-mm-dd.');
  }

  const soldDateObj = new Date(year, month, day);

  // Calculate warranty expiration date
  const warrantyExpirationDate = new Date(soldDateObj);
  warrantyExpirationDate.setDate(soldDateObj.getDate() + warrantyPeriodInDays);

  // Get today's date at midnight
  const today = new Date();
  today.setHours(0, 0, 0, 0); // Set time to midnight

  // Calculate the difference in milliseconds
  const differenceInTime = warrantyExpirationDate - today;
  const differenceInDays = Math.ceil(differenceInTime / (1000 * 3600 * 24));

  warrantyAvailability = false;
  // Check if warranty is available
  if (differenceInDays > 0) {
    warrantyAvailability = true;
  }
  return [warrantyAvailability, differenceInDays]
}




//Define Division List
// const divisionList = ["Barisal", "Chattogram", "Dhaka", "Khulna", "Mymensingh", "Rajshahi", "Rangpur", "Sylhet"]

//Define division to district map
const divisionToDistrict = {
  Barisal: ["Barguna", "Barisal", "Bhola", "Jhalokathi", "Patuakhali", "Pirojpur"],
  Chattogram: ["Bandarban", "Brahmanbaria", "Chandpur", "Chattogram", "Comilla", "CoxsBazar", "Feni", "Khagrachari", "Lakshmipur", "Noakhali", "Rangamati"],
  Dhaka: ["Dhaka", "Faridpur", "Gazipur", "Gopalganj", "Kishoreganj", "Madaripur", "Manikganj", "Munshiganj", "Narayanganj", "Narsingdi", "Rajbari", "Shariatpur", "Tangail"],
  Khulna: ["Bagerhat", "Chuadanga", "Jashore", "Jhenidah", "Khulna", "Kushtia", "Magura", "Meherpur", "Narail", "Satkhira"],
  Mymensingh: ["Jamalpur", "Mymensingh", "Netrokona", "Sherpur"],
  Rajshahi: ["Bogura", "Chapainawabganj", "Joypurhat", "Naogaon", "Natore", "Pabna", "Rajshahi", "Sirajganj"],
  Rangpur: ["Dinajpur", "Gaibandha", "Kurigram", "Lalmonirhat", "Nilphamari", "Panchagarh", "Rangpur", "Thakurgaon"],
  Sylhet: ["Habiganj", "Moulvibazar", "Sunamganj", "Sylhet"]
};

//Define district to upazila map
const districtToUpazila = {
  Bagerhat: ['Bagerhat Sadar', 'Chitalmari', 'Fakirhat', 'Kachua', 'Mollahat', 'Mongla', 'Morrelganj', 'Rampal', 'Sarankhola'],
  Bandarban: ['Bandarban Sadar', 'Naikhongchhari', 'Rowangchhari', 'Lama', 'Ruma', 'Thanchi', 'Alikadam'],
  Barguna: ['Amtali', 'Bamna', 'Barguna Sadar', 'Betagi', 'Patharghata', 'Taltali'],
  Barisal: ['Agailjhara', 'Babuganj', 'Bakerganj', 'Banaripara', 'Gournadi', 'Hizla', 'Barisal Sadar', 'Mehendiganj', 'Muladi', 'Wazirpur'],
  Bhola: ['Bhola Sadar', 'Burhanuddin', 'Char Fasson', 'Daulatkhan', 'Lalmohan', 'Manpura', 'Tazumuddin'],
  Bogura: ['Adamdighi', 'Bogra Sadar', 'Dhunat', 'Dhupchanchia', 'Gabtali', 'Kahaloo', 'Nandigram', 'Sariakandi', 'Shajahanpur', 'Sherpur', 'Shibganj', 'Sonatala'],
  Brahmanbaria: ['Brahmanbaria Sadar', 'Ashuganj', 'Nasirnagar', 'Nabinagar', 'Bancharampur', 'Bijoynagar', 'Kasba', 'Sarail', 'Akhaura'],
  Chandpur: ['Chandpur Sadar', 'Faridganj', 'Haimchar', 'Hajiganj', 'Kachua', 'Matlab Dakshin', 'Matlab Uttar', 'Shahrasti'],
  Chattogram: ['Anwara', 'Banshkhali', 'Boalkhali', 'Chandanaish', 'Chattogram Sadar', 'Fatikchhari', 'Hathazari', 'Lohagara', 'Mirsharai', 'Patiya', 'Rangunia', 'Raozan', 'Sandwip', 'Satkania', 'Sitakunda'],
  Chuadanga: ['Alamdanga', 'Chuadanga Sadar', 'Damurhuda', 'Jibannagar'],
  Comilla: ['Barura', 'Brahmanpara', 'Burichong', 'Chandina', 'Chauddagram', 'Comilla Adarsha Sadar', 'Comilla Sadar Dakshin', 'Daudkandi', 'Debidwar', 'Homna', 'Laksam', 'Monohorgonj', 'Muradnagar', 'Nangalkot', 'Meghna', 'Titas'],
  CoxsBazar: ['Chakaria', 'Cox\'s Bazar Sadar', 'Kutubdia', 'Maheshkhali', 'Ramu', 'Teknaf', 'Ukhia', 'Pekua'],
  Dhaka: ['Dhamrai', 'Dohar', 'Keraniganj', 'Nawabganj', 'Savar'],
  Dinajpur: ['Birampur', 'Birganj', 'Biral', 'Bochaganj', 'Chirirbandar', 'Dinajpur Sadar', 'Ghoraghat', 'Hakimpur', 'Kaharole', 'Khansama', 'Nawabganj', 'Parbatipur', 'Phulbari'],
  Faridpur: ['Alfadanga', 'Bhanga', 'Boalmari', 'Charbhadrasan', 'Faridpur Sadar', 'Madhukhali', 'Nagarkanda', 'Sadarpur', 'Shaltha'],
  Feni: ['Chhagalnaiya', 'Daganbhuiyan', 'Feni Sadar', 'Parshuram', 'Fulgazi', 'Sonagazi'],
  Gaibandha: ['Fulchhari', 'Gaibandha Sadar', 'Gobindaganj', 'Palashbari', 'Sadullapur', 'Saghata', 'Sundarganj'],
  Gazipur: ['Gazipur Sadar', 'Kaliakair', 'Kapasia', 'Sreepur', 'Kaliganj'],
  Gopalganj: ['Gopalganj Sadar', 'Kashiani', 'Kotalipara', 'Muksudpur', 'Tungipara'],
  Habiganj: ['Ajmiriganj', 'Bahubal', 'Baniachang', 'Chunarughat', 'Habiganj Sadar', 'Lakhai', 'Madhabpur', 'Nabiganj', 'Shaistaganj'],
  Jhalokathi: ['Jhalokathi Sadar', 'Kathalia', 'Nalchity', 'Rajapur'],
  Jamalpur: ['Bakshiganj', 'Dewanganj', 'Islampur', 'Jamalpur Sadar', 'Madarganj', 'Melandaha', 'Sarishabari'],
  Jashore: ['Abhaynagar', 'Bagherpara', 'Chaugachha', 'Jashore Sadar', 'Jhikargachha', 'Keshabpur', 'Manirampur', 'Sharsha'],
  Jhenidah: ['Harinakunda', 'Jhenidah Sadar', 'Kaliganj', 'Kotchandpur', 'Maheshpur', 'Shailkupa'],
  Khagrachari: ['Dighinala', 'Khagrachhari Sadar', 'Lakshmichhari', 'Mahalchhari', 'Manikchhari', 'Matiranga', 'Panchhari', 'Ramgarh'],
  Khulna: ['Batiaghata', 'Dacope', 'Dumuria', 'Dighalia', 'Koyra', 'Paikgachha', 'Phultala', 'Rupsa', 'Terokhada', 'Khulna Sadar'],
  Kishoreganj: ['Austagram', 'Bajitpur', 'Bhairab', 'Hossainpur', 'Itna', 'Karimganj', 'Katiadi', 'Kishoreganj Sadar', 'Kuliarchar', 'Mithamoin', 'Nikli', 'Pakundia', 'Tarail'],
  Kurigram: ['Bhurungamari', 'Char Rajibpur', 'Chilmari', 'Kurigram Sadar', 'Nageshwari', 'Phulbari', 'Rajarhat', 'Raomari', 'Ulipur'],
  Kushtia: ['Bheramara', 'Daulatpur', 'Khoksa', 'Kumarkhali', 'Kushtia Sadar', 'Mirpur'],
  Lakshmipur: ['Lakshmipur Sadar', 'Raipur', 'Ramganj', 'Ramgati', 'Komol Nagar'],
  Lalmonirhat: ['Aditmari', 'Hatibandha', 'Kaliganj', 'Lalmonirhat Sadar', 'Patgram'],
  Madaripur: ['Kalkini', 'Madaripur Sadar', 'Rajoir', 'Shibchar'],
  Magura: ['Magura Sadar', 'Mohammadpur', 'Shalikha', 'Sreepur'],
  Manikganj: ['Daulatpur', 'Ghior', 'Harirampur', 'Manikganj Sadar', 'Saturia', 'Shivalaya', 'Singair'],
  Meherpur: ['Gangni', 'Meherpur Sadar', 'Mujibnagar'],
  Moulvibazar: ['Barlekha', 'Juri', 'Kamalganj', 'Kulaura', 'Moulvibazar Sadar', 'Rajnagar', 'Sreemangal'],
  Munshiganj: ['Gazaria', 'Lohajang', 'Munshiganj Sadar', 'Sirajdikhan', 'Sreenagar', 'Tongibari'],
  Mymensingh: ['Bhaluka', 'Dhobaura', 'Fulbaria', 'Gaffargaon', 'Gouripur', 'Haluaghat', 'Ishwarganj', 'Muktagacha', 'Mymensingh Sadar', 'Nandail', 'Phulpur', 'Tarakanda', 'Trishal'],
  Naogaon: ['Atrai', 'Badalgachhi', 'Manda', 'Dhamoirhat', 'Mohadevpur', 'Naogaon Sadar', 'Niamatpur', 'Patnitala', 'Porsha', 'Raninagar', 'Sapahar'],
  Narail: ['Kalia', 'Lohagara', 'Narail Sadar'],
  Narayanganj: ['Araihazar', 'Bandar', 'Narayanganj Sadar', 'Rupganj', 'Sonargaon'],
  Narsingdi: ['Belabo', 'Monohardi', 'Narsingdi Sadar', 'Palash', 'Raipura', 'Shibpur'],
  Natore: ['Bagatipara', 'Baraigram', 'Gurudaspur', 'Lalpur', 'Naldanga', 'Natore Sadar', 'Singra'],
  Netrokona: ['Atpara', 'Barhatta', 'Durgapur', 'Kalmakanda', 'Kendua', 'Khaliajuri', 'Madan', 'Mohanganj', 'Netrokona Sadar', 'Purbadhala'],
  Nilphamari: ['Dimla', 'Domar', 'Jaldhaka', 'Kishoreganj', 'Nilphamari Sadar', 'Saidpur'],
  Noakhali: ['Begumganj', 'Chatkhil', 'Companiganj', 'Hatiya', 'Noakhali Sadar', 'Senbagh', 'Sonaimuri', 'Subarnachar', 'Kabirhat'],
  Pabna: ['Atgharia', 'Bera', 'Bhangura', 'Chatmohar', 'Faridpur', 'Ishwardi', 'Pabna Sadar', 'Santhia', 'Sujanagar'],
  Panchagarh: ['Atwari', 'Boda', 'Debiganj', 'Panchagarh Sadar', 'Tetulia'],
  Patuakhali: ['Bauphal', 'Dashmina', 'Galachipa', 'Kalapara', 'Mirzaganj', 'Patuakhali Sadar', 'Rangabali', 'Dumki'],
  Pirojpur: ['Bhandaria', 'Kawkhali', 'Mathbaria', 'Nazirpur', 'Nesarabad (Swarupkathi)', 'Pirojpur Sadar', 'Zianagar'],
  Rajbari: ['Baliakandi', 'Goalandaghat', 'Pangsha', 'Rajbari Sadar', 'Kalukhali'],
  Rajshahi: ['Bagha', 'Bagmara', 'Charghat', 'Durgapur', 'Godagari', 'Mohanpur', 'Paba', 'Puthia', 'Tanore'],
  Rangamati: ['Bagaichhari', 'Barkal', 'Kawkhali', 'Belaichhari', 'Kaptai', 'Juraichhari', 'Langadu', 'Naniarchar', 'Rajasthali', 'Rangamati Sadar'],
  Rangpur: ['Badarganj', 'Gangachara', 'Kaunia', 'Rangpur Sadar', 'Mithapukur', 'Pirgacha', 'Pirganj', 'Taraganj'],
  Satkhira: ['Assasuni', 'Debhata', 'Kalaroa', 'Kaliganj', 'Satkhira Sadar', 'Shyamnagar', 'Tala'],
  Shariatpur: ['Bhedarganj', 'Damudya', 'Gosairhat', 'Naria', 'Shariatpur Sadar', 'Zajira', 'Shakhipur'],
  Sherpur: ['Jhenaigati', 'Nakla', 'Nalitabari', 'Sherpur Sadar', 'Sreebardi'],
  Sirajganj: ['Belkuchi', 'Chauhali', 'Kamarkhanda', 'Kazipur', 'Raiganj', 'Shahjadpur', 'Sirajganj Sadar', 'Tarash', 'Ullahpara'],
  Sunamganj: ['Bishwamvarpur', 'Chhatak', 'Derai', 'Dharamapasha', 'Dowarabazar', 'Jagannathpur', 'Jamalganj', 'Sullah', 'Sunamganj Sadar', 'Tahirpur'],
  Sylhet: ['Balaganj', 'Beanibazar', 'Bishwanath', 'Companiganj', 'Dakshin Surma', 'Fenchuganj', 'Golapganj', 'Gowainghat', 'Jaintiapur', 'Kanaighat', 'Osmani Nagar', 'Sylhet Sadar', 'Zakiganj'],
  Tangail: ['Basail', 'Bhuapur', 'Delduar', 'Dhanbari', 'Ghatail', 'Gopalpur', 'Kalihati', 'Madhupur', 'Mirzapur', 'Nagarpur', 'Sakhipur', 'Tangail Sadar'],
  Thakurgaon: ['Baliadangi', 'Haripur', 'Pirganj', 'Ranisankail', 'Thakurgaon Sadar']
};

// Function to update the subcategory dropdown based on the selected category
function updateChildList(parent, child, mapName) {
  const categorySelect = document.getElementById(parent);
  const subcategorySelect = document.getElementById(child);
  const selectedValue = categorySelect.value;

  // Clear existing options in the subcategory dropdown
  subcategorySelect.innerHTML = "";

  // Get the list of subcategory for the selected catergory
  let lists, firstLabel
  if (mapName === 'divisionToDistrict') {
    lists = divisionToDistrict[selectedValue];
    firstLabel = "Select District";
  } else if (mapName === 'districtToUpazila') {
    lists = districtToUpazila[selectedValue];
    firstLabel = "Select District";
  }

  const option = document.createElement("option");
  option.value = "";
  option.textContent = firstLabel;
  option.selected = true;
  option.disabled = true;
  subcategorySelect.appendChild(option);

  // Populate the subcategory dropdown with the fetched subcategory lists
  lists.forEach(product => {
    const option = document.createElement("option");
    option.value = product;
    option.textContent = product;
    subcategorySelect.appendChild(option);
  });
}

// Show success message function
function showSuccessMessage(message) {
  const Toast = Swal.mixin({
    toast: true,
    position: "top-end",
    showConfirmButton: false,
    timer: 2000,
    timerProgressBar: true,
    didOpen: (toast) => {
      toast.onmouseenter = Swal.stopTimer;
      toast.onmouseleave = Swal.resumeTimer;
    }
  });
  Toast.fire({
    icon: "success",
    title: message,
  });
}

// Show error message function
function showErrorMessage(message) {
  const Toast = Swal.mixin({
    toast: true,
    position: "top-end",
    showConfirmButton: false,
    timer: 2000,
    timerProgressBar: true,
    didOpen: (toast) => {
      toast.onmouseenter = Swal.stopTimer;
      toast.onmouseleave = Swal.resumeTimer;
    }
  });
  Toast.fire({
    icon: "error",
    title: message,
  });
}

//addNewBrand show a popup form and then make an api call to insert brand data to the database table
function addNewBrand(page) {
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
            if (page === "purchase") {
              document.getElementById("brand").innerHTML = '';
              document.getElementById("brand").innerHTML = `<option value="${data.result.id}" selected>${data.result.name}</option>`;;
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

            if (page === "purchase") {
              document.getElementById("category").innerHTML = '';
              document.getElementById("category").innerHTML = `<option value="${product.category_id}" selected>${categoryName}</option>`;
              document.getElementById("category").disabled = true;
              document.getElementById("product").innerHTML = '';
              document.getElementById("product").innerHTML = `<option value="${data.result.id}" selected>${data.result.product_name}</option>`;
              document.getElementById("product").disabled = true;
              products.push(data.result);
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
            showSuccessMessage(data.message);
            if (page === "purchase") {
              document.getElementById("category").innerHTML = '';
              document.getElementById("category").innerHTML = `<option value="${data.result.id}" selected>${data.result.name}</option>`;;
              document.getElementById("category").disabled = true;
              if (categories) {
                categories.push(data.result);
              }
            }
          }
        });
    }
  });
}
//addNewCustomer show a popup form and then make an api call to insert customer data to the database table
function addNewCustomer(page) {
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
          if (page === "") {
            setTimeout(function () {
              location.reload();
            }, 3000); // Adjust the delay as needed 
          }
          if (data.error === true) {
            showErrorMessage(data.message)
          } else {
            showSuccessMessage(data.message);
            if (page === "sale") {
              document.getElementById("customer").innerHTML = '';
              document.getElementById("customer").innerHTML = `<option value="${data.result.id}" selected>${data.result.account_name} (${data.result.account_code})</option>`;;
              document.getElementById("customer").disabled = true;
            }
          }
        });
    }
  });
}
//addNewEmployee show a popup form and then make an api call to insert employee data to the database table
function addNewEmployee(page) {
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
            showSuccessMessage(data.message);
          }
        });
    }
  });
}
//addNewSupplier show a popup form and then make an api call to insert supplier data to the database table
function addNewSupplier(page) {
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
          if (page === "") {
            setTimeout(function () {
              location.reload();
            }, 3000); // Adjust the delay as needed 
          }
          if (data.error === true) {
            showErrorMessage(data.message)
          } else {
            showSuccessMessage(data.message);
            if (page === "purchase") {
              document.getElementById("supplier").innerHTML = '';
              document.getElementById("supplier").innerHTML = `<option value="${data.result.id}" selected>${data.result.account_name} (${data.result.account_code})</option>`;;
              document.getElementById("supplier").disabled = true;
            }
          }
        });
    }
  });
}

//checkoutWarrantyProducts show a popup checkout warranty for and make an api call to update database table for checkout process
function checkoutWarrantyProducts(warrantyHistoryID, productSerialID) {

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