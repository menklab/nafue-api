var dom, modal, g, s = {
    adata: sjcl.codec.base64.toBits(btoa('Sparticus')),
    mode: 'ccm',
    cipher: 'aes',
    tagSize: 128,
    keySize: 256,
    iterations: 1000
};


function init() {
    dom = domInit();
    modal = modalInit();
    g = {
        mode: "upload",
        binData: null
    };

    if (window.File && window.FileReader && window.FileList && window.Blob) {
        // handle history

        // check for decryption file
        var dFile;
        var parse = parseURL(window.location.href);
        if (parse.pathname[1] == "file") {
            dFile = parse.pathname[2];
        }
        if (!!dFile) {
            window.history.pushState('', '', '/');
            downloadFile(dFile);
        }

        if (!!dom.dropZone) {
            dom.dropZone.addEventListener('dragover', handleDragOver, false);
            dom.dropZone.addEventListener('drop', handleFileSelect, false);
        }
        if (!!dom.file) {
            dom.file.addEventListener('change', handleFileSelect, false);
            dom.file.focus();
        }
        // setup password check listener
        window.onload = function () {
            if (!!dom.password) {
                dom.password.onkeyup = checkPassword;
            }
        };
    }
    else {
        hide(dom.dropZone);
        show(dom.unsupported);
    }

}

function getParameterByName(name) {
    name = name.replace(/[\[]/, "\\[").replace(/[\]]/, "\\]");
    var regex = new RegExp("[\\?&]" + name + "=([^&#]*)"),
        results = regex.exec(location.search);
    return results === null ? "" : decodeURIComponent(results[1].replace(/\+/g, " "));
}

init();



