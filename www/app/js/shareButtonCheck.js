function doCheck(){
    var allFilled = true;

    var password = document.getElementsByTagName('password');
    for(var i=0; i<password.length; i++){
        if(password[i].type == "text" && password[i].value == ''){
            allFilled = false;
            break;
        }
    }

    document.getElementById("share").disabled = !allFilled;
}

window.onload = function() {
    var password = document.getElementById('password');
    for (var i = 0; i < password.length; i++) {
        if (password[i].type == "text") {
            password[i].onkeyup = doCheck;
            password[i].onblur = doCheck;
        }
    }
};