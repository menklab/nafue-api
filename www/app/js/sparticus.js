if (window.File && window.FileReader && window.FileList && window.Blob) {

    var dropZone = document.getElementById('drop');
    var list = document.getElementById('list');
    var password = document.getElementById('password');
    var passCont = document.getElementById('passCont');
    var fileSize = document.getElementById('totalSize');

    var binStr;

    dropZone.addEventListener('dragover', handleDragOver, false);
    dropZone.addEventListener('drop', handleFileSelect, false);
    document.getElementById('file').addEventListener('change', handleFileSelect, false);

} else {
    document.getElementById('status').innerHTML = 'Your browser does not support the HTML5 FileReader.';
}


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
    var p = doPbkdf2(password.value);
    console.log("password: ", password.value);
    console.log('p: ', p);
    var ct = doEncrypt(password.value, binStr);
    console.log("sData: ", ct);

}
