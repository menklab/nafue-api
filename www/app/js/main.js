if (window.File && window.FileReader && window.FileList && window.Blob) {

    var dropZone = document.getElementById('drop');
    var list = document.getElementById('list');
    var file = document.getElementById('file');
    var password = document.getElementById('password');
    var passwordReqs = document.getElementById('passwordReqs');
    var passCont = document.getElementById('passCont');
    var busy = document.getElementById('busy');
    var showLink = document.getElementById('showLink');
    var linkToShare = document.getElementById('linkToShare');
    var busyMessage = document.getElementById('busyMessage');
    var share = document.getElementById("share");
    var downloadBtn = document.getElementById("downloadBtn");
    var len = document.getElementById("len");
    var low = document.getElementById("low");
    var upp = document.getElementById("upp");
    var num = document.getElementById("num");
    var spc = document.getElementById("spc");
    var reset = document.getElementById("reset");
    var upload = document.getElementById("upload");

    var skipCheck = false;
    var binStr;

    dropZone.addEventListener('dragover', handleDragOver, false);
    dropZone.addEventListener('drop', handleFileSelect, false);
    document.getElementById('file').addEventListener('change', handleFileSelect, false);

    // check for decryption file
    var dFile = getParameterByName("file");
    if (!!dFile) {
        window.history.pushState('','','/');
        downloadFile(dFile);
    }

} else {
    document.getElementById('status').innerHTML = 'Your browser does not support the HTML5 FileReader.';
}

// on load
window.onload = function () {
    password.onkeyup = checkPassword;
};



