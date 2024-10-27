/**
 * Converts a number into its word representation based on the Indian numbering system.
 * Supports numbers up to 99,99,99,999 (99 crores).
 *
 * param {number} num - The number to be converted into words.
 * returns {string} - The word representation of the given number.
 *
 * The function handles the following:
 * - Converts numbers below 20 and multiples of 10 using predefined word arrays.
 * - Utilizes place values specific to the Indian numbering system: thousand, lakh, and crore.
 * - Recursively breaks down the number into groups (hundreds, thousands, lakhs, crores)
 *   and converts each group to words.
 * - Combines each group's word representation with the corresponding place value.
 *
 * Example:
 *   numberToWords(12345678) returns "one crore twenty-three lakh forty-five thousand six hundred seventy-eight"
 */

function numberToWords(num) {
  // Return "zero" immediately if the number is zero
  if (num === 0) return "zero";

  // Arrays for words of numbers below 20 and multiples of ten
  const belowTwenty = [
      "", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", 
      "ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", 
      "sixteen", "seventeen", "eighteen", "nineteen"
  ];
  const tens = [
      "", "", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"
  ];

  // Array for place values in the Indian numbering system (thousand, lakh, crore)
  const units = ["", "thousand", "lakh", "crore"];

  // Helper function to convert numbers below 1000 to words
  function helper(n) {
      if (n < 20) 
          return belowTwenty[n]; // For numbers 0-19, directly return from `belowTwenty` array
      else if (n < 100) 
          return tens[Math.floor(n / 10)] + (n % 10 ? " " + belowTwenty[n % 10] : ""); // For numbers 20-99
      else 
          return belowTwenty[Math.floor(n / 100)] + " hundred" + (n % 100 ? " " + helper(n % 100) : ""); // For numbers 100-999
  }

  let word = "";       // Variable to store the final word result
  let unitIndex = 0;   // Index for tracking the place value (thousand, lakh, crore)

  // Loop through the number in chunks based on Indian numbering system
  while (num > 0) {
      // Extract current chunk: hundreds for thousand's place, thousands for lakh and crore
      let chunk = num % 100;
      if (unitIndex === 1) chunk = num % 1000; // Special handling for thousand's place (3 digits)

      // If the chunk is not zero, convert it to words and add place value
      if (chunk > 0) {
          word = helper(chunk) + (units[unitIndex] ? " " + units[unitIndex] : "") + (word ? " " + word : "");
      }

      // Remove the processed chunk from `num`
      num = Math.floor(num / (unitIndex === 1 ? 1000 : 100));

      // Move to the next place value (thousand, lakh, crore)
      unitIndex++;
  }

  return word.trim(); // Return the final word result, trimmed of any extra spaces
}



// GenerateRandomAlphanumericCode generates a random alphanumeric string of the specified length.
//
// The string consists of a mix of uppercase letters, lowercase letters, and digits.
// It uses cryptographic randomness to ensure that the generated string is secure.
//
// Parameters:
//   - length: The desired length of the generated string (should be a positive integer).
//
// Returns:
//   - A random alphanumeric string of the specified length.
function generateRandomAlphanumericCode(length) {
  const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789";
  const charsetLength = charset.length;
  let randomCode = '';

  for (let i = 0; i < length; i++) {
      const index = Math.floor(Math.random() * charsetLength);
      randomCode += charset[index];
  }

  return randomCode;
}


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



/**
 * Generates a PDF from an HTML element and opens the print dialog.
 * @param {string} id - The ID of the HTML element to convert to PDF.
 * @param {number} dpi - The quality of the PDF rendering.
 */
function generatePDF(id, dpi) {
  const element = document.getElementById(id);

  // Configure and generate the PDF with html2pdf
  html2pdf()
      .from(element)
      .set({
          margin: [10,5,10,5],  // Adjust margin as needed (in mm)
          filename: 'myContent.pdf',
          html2canvas: { scale: dpi },
          jsPDF: { format: 'a4', orientation: 'portrait' },
          pagebreak: { mode: ['css', 'legacy'] }  // Page breaks managed by CSS rules
      })
      .outputPdf('bloburl')
      .then((pdfUrl) => {
          const pdfWindow = window.open(pdfUrl, '_blank');
          pdfWindow.print();
      });
}



