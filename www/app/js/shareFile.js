function shareFile() {
    passCont.hidden = true;
    busy.hidden = false;

    // do encryption
    var ct = doEncrypt(password.value, binStr);
    console.log("ct: ", ct);
    // make api request for saving file
    var payload = {
        iv: sjcl.codec.base64.fromBits(ct.p.iv),
        salt: sjcl.codec.base64.fromBits(ct.p.salt),
        aData: ct.p.adata
    };
    // make upload request
    request = JSON.stringify(payload);
    busyMessage.innerHTML = "Encrypting";
    http.post(services + "/api/files", request)
        .success(function (fileData) {
            setTimeout(function () {
                busyMessage.innerHTML = "Uploading";
                // do upload to aws s3
                http.put(fileData.uploadUrl, sjcl.codec.base64.fromBits(ct.ct), {contentType: "text/plain;charset=UTF-8"})
                    .success(function (res) {
                        busy.hidden = true;
                        showLink.hidden = false;
                        linkToShare.innerHTML = services + "?file=" + fileData.shortUrl;
                        reset.focus();
                    })
                    .error(function (err) {
                        console.log("fail!");
                    });
            }, 500);

        })
        .error(function (err) {

        });
}
