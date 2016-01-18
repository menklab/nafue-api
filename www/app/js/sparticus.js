if (window.File && window.FileReader && window.FileList && window.Blob) {

    var dropZone = document.getElementById('drop');
    var list = document.getElementById('list');
    var password = document.getElementById('password');
    var passCont = document.getElementById('passCont');
    var encrypting = document.getElementById('encrypting');
    var fileSize = document.getElementById('totalSize');
    var share = document.getElementById("share");
    var len = document.getElementById("len");
    var low = document.getElementById("low");
    var upp = document.getElementById("upp");
    var num = document.getElementById("num");
    var spc = document.getElementById("spc");

    var binStr;

    dropZone.addEventListener('dragover', handleDragOver, false);
    dropZone.addEventListener('drop', handleFileSelect, false);
    document.getElementById('file').addEventListener('change', handleFileSelect, false);


} else {
    document.getElementById('status').innerHTML = 'Your browser does not support the HTML5 FileReader.';
}

// on load
window.onload = function () {
    password.onkeyup = checkPassword;
    password.onblur = checkPassword;
};


function handleFileSelect(e) {
    e.stopPropagation();
    e.preventDefault();

    var files;
    if (!!e.dataTransfer && !!e.dataTransfer.files) {
        files = e.dataTransfer.files
    }

    if (!!e.target && !!e.target.files) {
        files = e.target.files
    }

    // only 1 file at a time
    if (files.length > 1) {
        console.log("only 1 file at a time!");
        return;
    }
    var file = files[0];

    var reader = new FileReader();

    // closure to capture file
    reader.onload = (function (f) {
        return function (e) {
            // Render thumbnail.
            console.log("file loaded: ", f);
            binStr = e.target.result;
            //console.log("binStr: ", binStr);
            dropZone.hidden = true;
            passCont.hidden = false;
            password.focus();
        };
    })(file);

    // Read in the image file as a data URL.

    reader.readAsBinaryString(file);

}

function handleDragOver(e) {
    e.stopPropagation();
    e.preventDefault();
    e.dataTransfer.dropEffect = 'copy'; // Explicitly show this is a copy.
}

function shareFile() {
    passCont.hidden = true;
    encrypting.hidden = false;

    // do encryption
    var ct = JSON.parse(doEncrypt(password.value, binStr));

    // make api request for saving file
    var payload = {
        iv: ct.iv,
        salt: ct.salt,
        aData: ct.adata
    };
    console.log('ct: ', ct);
    request = JSON.stringify(payload);
    http.post("/api/files", request)
        .success(function (res) {})
        .error(function (err) {});


}
function checkPassword() {
    var pass = password.value;

    var p={};

    if (pass.length > 7) {
        p.length = true;
        len.className = "present";
    }
    else {
        p.length = false;
        len.className = "missing";
    }

    if (!!pass.match(/[0-9]/)) {
        p.num = true;
        num.className = "present"
    }
    else {
        p.num = false;
        num.className = "missing"
    }

    if (!!pass.match("[a-z]")) {
        p.lower = true;
        low.className = "present";
    }
    else {
        p.lower = false;
        low.className = "missing";
    }

    if (!!pass.match("[A-Z]")) {
        p.upper = true;
        upp.className = "present";
    }
    else {
        p.upper = false;
        upp.className = "missing";
    }

    if(!!pass.match(/[!,@,#,$,%,^,&,*,?,_,~,-,(,),\s]/)) {
        p.special = true;
        spc.className = "present";
    }
    else {
        p.special = false;
        spc.className = "missing";
    }


    // update button when password is good enough
    var strength = 0;
    strength += 1? p.num:0;
    strength += 1? p.lower:0;
    strength += 1? p.upper:0;
    strength += 1? p.special:0;
    if (strength >=3 && p.length) {
        share.disabled = false;
    }
    else {
        share.disabled = true;
    }
}

