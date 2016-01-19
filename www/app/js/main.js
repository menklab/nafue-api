if (window.File && window.FileReader && window.FileList && window.Blob) {

    var dropZone = document.getElementById('drop');
    var list = document.getElementById('list');
    var password = document.getElementById('password');
    var passCont = document.getElementById('passCont');
    var busy = document.getElementById('busy');
    var showLink = document.getElementById('showLink');
    var linkToShare = document.getElementById('linkToShare');
    var fileSize = document.getElementById('totalSize');
    var busyMessage = document.getElementById('busyMessage');
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
};



