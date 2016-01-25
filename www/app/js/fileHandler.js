function handleFileSelect(e) {
    setContent(dom.busyMessage, "Processing");
    show(dom.homeLink);

    show(dom.busy);
    hide(dom.dropZone);
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
        error("Only 1 file at a time.");
        return;
    }

    var file = files[0];
    var type = file.type;
    var name = file.name;
    var reader = new FileReader();

    // only 50 MB
    if (file.size/1024/1024 > 50) {
        reader.abort();
        error('The uploaded file cannot be greater than 50MB');
    }

    // closure to capture file
    reader.onload = (function (f) {
        return function (e) {
            if (file.size/1024/1024 <= 50) {
                var data = {
                    content: encodeAb(e.target.result),
                    type: type,
                    name: name
                };

                //var blob = new Blob([decodeAb(data.content)]); //new Blob([btoa(data.content)], {type: data.type});
                //saveAs(blob, data.name);

                g.binData = btoa(JSON.stringify(data));
                hide(dom.busy);
                show(dom.passCont);
                dom.password.focus();
            }
            else {
                e = null;
            }
        };
    })(file);

    // Read in the image file as a data URL.

    reader.readAsArrayBuffer(file);

}

function handleDragOver(e) {
    e.stopPropagation();
    e.preventDefault();
    e.dataTransfer.dropEffect = 'copy'; // Explicitly show this is a copy.
}

function downloadFile(file) {
    hide(dom.dropZone);
    show(dom.downloadBtn);
    setContent(dom.busyMessage, "Downloading");
    show(dom.busy);

    http.get(services + "/api/files/" + file)
        .success(function (res) {
            http.get(res.downloadUrl)
                .success(function (eData) {
                    res.ct = eData;
                    g.ct = {
                        ct: sjcl.codec.base64.toBits(eData),
                        p: {
                            adata: sjcl.codec.base64.toBits(res.aData),
                            iv: sjcl.codec.base64.toBits(res.iv),
                            salt: sjcl.codec.base64.toBits(res.salt)
                        }
                    };

                    // move to decrypt
                    decryptScreen();
                })
                .error(function (err) {
                    error("The file could not be access or no longer exists.");
                })
        })
        .error(function (err) {
            error("The file could not be access or no longer exists.");
        });
}

function shareFile() {
    hide(dom.passCont);
    show(dom.busy);

    // do encryption
    setContent(dom.busyMessage, "Encrypting");
    var ct = doEncrypt(dom.password.value, sjcl.codec.base64.toBits(g.binData));
    // make api request for saving file
    var payload = {
        iv: sjcl.codec.base64.fromBits(ct.p.iv),
        salt: sjcl.codec.base64.fromBits(ct.p.salt),
        aData: sjcl.codec.base64.fromBits(ct.p.adata)
    };
    // make upload request
    request = JSON.stringify(payload);
    http.post(services + "/api/files", request)
        .success(function (fileData) {
            setTimeout(function () {
                setContent(dom.busyMessage.innerHTML, "Uploading");
                // do upload to aws s3
                var data = sjcl.codec.base64.fromBits(ct.ct);
                http.put(fileData.uploadUrl, data, {contentType: "text/plain;charset=UTF-8"})
                    .success(function (res) {
                        hide(dom.busy);
                        show(dom.showLink);
                        setContent(dom.linkToShare, services + "?file=" + fileData.shortUrl);
                    })
                    .error(function (err) {
                       error(err.message);
                    });
            }, 500);

        })
        .error(function (err) {
            error(err.message);
        });
}

function decryptScreen() {
    g.mode = "download";
    hide(dom.busy);
    show(dom.passCont);
    hide(dom.passwordReqs);
    hide(dom.share);
    show(dom.downloadBtn);
    dom.password.focus();
}

function decryptFile() {
    var data = JSON.parse(atob(sjcl.codec.base64.fromBits(doDecrypt(dom.password.value, g.ct))));
    var blob = new Blob([decodeAb(data.content)], {type: data.type});

    saveAs(blob, data.name);

    show(dom.doneDownloading);
    hide(dom.passCont);
    setTimeout(function () {
        hide(dom.doneDownloading);
        reset_ui();
    }, 1000 * 10);
}
