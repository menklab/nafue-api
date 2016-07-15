function handleFileSelect(e) {
    setContent(dom.busyMessage, "Processing");

    show(dom.busy);
    hide(dom.dropZone);
    e.stopPropagation();
    e.preventDefault();

    var files;
    if (!!e.dataTransfer && !!e.dataTransfer.files) {
        files = e.dataTransfer.files;
    }

    if (!!e.target && !!e.target.files) {
        files = e.target.files;
    }

    // only 1 file at a time
    if (files.length > 1) {
        error("Only 1 file at a time.");
        return;
    }

    var file = files[0];
    var totalfileSize = file.size + 32;
    var name = file.name;
    var nameAry = bytesToUint8(name);
    totalfileSize += name.length;

    var reader = new FileReader();

    // only 50 MB
    if (file.size / 1024 / 1024 > 50) {
        reader.abort();
        error('The uploaded file cannot be greater than 50MB');
    }

    var chunkSize = 32000; // 32kb
    var tChunks = Math.ceil(totalfileSize / chunkSize); // +32 for iv
    console.log("num of chunks: ", tChunks);
    console.log("filesize: ", totalfileSize);


    var readChunk = function (curChunk, fileHeaderId) {

        reader.onloadend = function (evt) {
            if (evt.target.readyState == FileReader.DONE) {

                // get data from read
                var buffer = evt.target.result;
                var data = new Uint8Array(buffer);

                // encrypt data from read
                cipher.update(new forge.util.ByteBuffer(data));
                var eData = bytesToUint8(cipher.output.getBytes());
                var addedChunkEvt = g.db.chunks().add(
                    {
                        fileHeaderId: fileHeaderId,
                        data: eData
                    }
                );

                addedChunkEvt.onsuccess = function () {
                    // if there are more chunks read the next one
                    if (curChunk < (tChunks - 1)) {
                        curChunk++;
                        readChunk(curChunk, fileHeaderId)
                    }
                    // otherwise finish last block with extra data and close
                    else {
                        var paddedFileName = new Uint8Array(255);
                        paddedFileName.fill(0);
                        paddedFileName.set(nameAry);
                        cipher.update(new forge.util.ByteBuffer(paddedFileName));
                        eData = bytesToUint8(cipher.output.getBytes());
                        g.db.chunks().add(
                            {
                                fileHeaderId: fileHeaderId,
                                data: eData
                            }
                        );
                        cipher.finish();
                    }
                };
            }
        };

        // calc chunk start/end
        var start = chunkSize * curChunk;
        var end = (tChunks * curChunk) + chunkSize;
        if (curChunk == (tChunks - 1)) { // if on last chunk end == last byte
            end = totalfileSize;
        }
        var chunk = file.slice(start, end);
        reader.readAsArrayBuffer(chunk)
    };

    var iv = forge.random.getBytesSync(32);
    var salt = forge.random.getBytesSync(32);
    var key = forge.pkcs5.pbkdf2('password', salt, 1000, 32);
    var cipher = forge.cipher.createCipher('AES-CTR', key);
    cipher.start({iv: iv});

    // create file header
    var fileHeader = {
        size: totalfileSize,
        salt: salt
    };

    // add header to db
    var req = g.db.headers().add(fileHeader);
    req.onsuccess = function (event) {
        var fileHeaderId = event.target.result;
        console.log("file header: ", fileHeaderId);

        readChunk(0, fileHeaderId);
    };


}

function bytesToUint8(buf) {
    var u8 = new Uint8Array(buf.split("").map(function (c) {
        return c.charCodeAt(0);
    }));
    return u8;
}

function uint8ToBytes(buf) {
    return String.fromCharCode.apply(null, buf)
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

    http.get(api_services + "/api/files/" + file)
        .success(function (res) {
            http.get(res.downloadUrl)
                .success(function (eData) {
                    res.ct = eData;
                    g.ct = {
                        ct: sjcl.codec.base64.toBits(eData),
                        p: {
                            hmac: sjcl.codec.base64.toBits(res.hmac),
                            salt: sjcl.codec.base64.toBits(res.salt)
                        }
                    };

                    // move to decrypt
                    decryptScreen();
                })
                .error(function (err) {
                    error("The file could not be access or no longer exists.");
                });
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
    console.log("salt: ", ct.p.salt);
    console.log("salt b64: ", sjcl.codec.base64.fromBits(ct.p.salt));
    console.log("iv: ", ct.p.iv);
    console.log("iv b64: ", sjcl.codec.base64.fromBits(ct.p.iv));


    // make upload request
    request = JSON.stringify(payload);
    http.post(api_services + "/api/files", request)
        .success(function (fileData) {
            setTimeout(function () {
                setContent(dom.busyMessage.innerHTML, "Uploading");
                // do upload to aws s3
                var data = sjcl.codec.base64.fromBits(ct.ct);
                http.put(fileData.uploadUrl, data, {contentType: "text/plain;charset=UTF-8"})
                    .success(function (res) {
                        hide(dom.busy);
                        show(dom.showLink);
                        setContent(dom.linkToShare, www_services + "/file/" + fileData.shortUrl);
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
    hide(dom.share);
    show(dom.downloadBtn);
    dom.password.focus();
}

function decryptFile() {
    var data = JSON.parse(atob(sjcl.codec.base64.fromBits(doDecrypt(dom.password.value, g.ct))));
    var blob = new Blob([decodeAb(data.content)], {type: "application/json"});

    saveAs(blob, data.name);

    show(dom.doneDownloading);
    hide(dom.passCont);
    setTimeout(function () {
        hide(dom.doneDownloading);
        reset_ui();
    }, 1000 * 10);
}
