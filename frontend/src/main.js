//Global JS function for greeting
function greet() {
    //Call Go Greet function
    window.go.main.App.Greet().then(result => {
        //Display result from Go
        document.getElementById("greet-text").innerText = result;
    }).catch(err => {
        console.log(err);
    }).finally(() => {
        console.log("finished!")
    });
}