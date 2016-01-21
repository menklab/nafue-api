var dom, g, s = {
    adata: sjcl.codec.base64.toBits(btoa('Sparticus')),
    mode: 'ccm',
    cipher: 'aes',
    tagSize: 128,
    keySize: 256,
    iterations: 1000
};


function init() {
    dom = domInit();
    g = {
        mode: "upload",
        binData: null
    };

    if (window.File && window.FileReader && window.FileList && window.Blob) {

        // check for decryption file
        var dFile = getParameterByName("file");
        if (!!dFile) {
            window.history.pushState('', '', '/');
            downloadFile(dFile);
        }

        dom.dropZone.addEventListener('dragover', handleDragOver, false);
        dom.dropZone.addEventListener('drop', handleFileSelect, false);
        dom.file.addEventListener('change', handleFileSelect, false);
        dom.file.focus();

        // setup password check listener
        window.onload = function () {
            dom.password.onkeyup = checkPassword;
        };
    }
    else {
        dom.upload.hidden = true;
        dom.unsupported.hidden = false;
    }
}

function getParameterByName(name) {
    name = name.replace(/[\[]/, "\\[").replace(/[\]]/, "\\]");
    var regex = new RegExp("[\\?&]" + name + "=([^&#]*)"),
        results = regex.exec(location.search);
    return results === null ? "" : decodeURIComponent(results[1].replace(/\+/g, " "));
}

init();


