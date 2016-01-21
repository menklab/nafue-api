function handleFileSelect(e) {
    setContent(dom.busyMessage, "Processing");
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
        console.log("only 1 file at a time!");
        return;
    }
    var file = files[0];
    var type = file.type;
    var name = file.name;
    var reader = new FileReader();

    // closure to capture file
    reader.onload = (function (f) {
        return function (e) {
            // Render thumbnail.
            g.binData = btoa(JSON.stringify({
                content: e.target.result,
                type: type,
                name: name
            }));
            console.log(g.binData);
            hide(dom.dropZone.hidden);
            show(dom.passCont.hidden);
            dom.password.focus();
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

function downloadFile(file, cb) {
    hide(dom.upload.hidden);
    show(dom.downloadBtn.hidden);
    setContent(dom.busyMessage, "Downloading");
    show(dom.busy.hidden);

    http.get(services + "/api/files/" + file)
        .success(function (res) {
            http.get(res.downloadUrl)
                .success(function (eData) {
                    res.ct = eData;
                    var ct = {
                        ct: sjcl.codec.base64.toBits(eData),
                        p: {
                            adata: sjcl.codec.base64.toBits(res.aData),
                            iv: sjcl.codec.base64.toBits(res.iv),
                            salt: sjcl.codec.base64.toBits(res.salt)
                        }
                    };
                    cb(ct);
                })
                .error(function (err) {
                    console.log("error");
                })
        })
        .error(function (err) {
            console.log("error");
        });
}

function shareFile() {
    hide(dom.passCont.hidden);
    show(busy.hidden);

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
                        setContent(dom.linkToShare,services + "?file=" + fileData.shortUrl);
                        dom.reset.focus();
                    })
                    .error(function (err) {
                        console.log("fail!");
                    });
            }, 500);

        })
        .error(function (err) {

        });
}

function tryPasswordToDecrpyt() {
    g.mode = "download";
    hide(dom.busy);
    show(dom.passCont);
    hide(dom.passwordReqs);
    hide(dom.hidden);
    show(dom.downloadBtn);
    dom.password.focus();
}

function decryptFile() {
    var data = JSON.parse(atob(sjcl.codec.base64.fromBits(doDecrypt(dom.password.value, g.ct))));
    var blob = new Blob([data.content], {type: data.type});
    saveAs(blob, data.name);
    show(dom.doneDownloading);
    hide(dom.passCont);
    timeout(function() {
        hide(dom.doneDownloading);
        hide(dom.upload);
        reset_ui();
    },1000 * 10);
}
