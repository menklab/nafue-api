function shareFile() {
    passCont.hidden = true;
    busy.hidden = false;

    // do encryption
    var ct = JSON.parse(doEncrypt(password.value, binStr));

    // make api request for saving file
    var payload = {
        iv: ct.iv,
        salt: ct.salt,
        aData: ct.adata
    };
    // make upload request
    request = JSON.stringify(payload);
    busyMessage.innerHTML = "Encrypting";
    http.post(services + "/api/files", request)
        .success(function (fileData) {
            setTimeout(function () {
                busyMessage.innerHTML = "Uploading";
                // do upload to aws s3
                http.put(fileData.uploadUrl, ct.ct, {contentType: "text/plain;charset=UTF-8"})
                    .success(function (res) {
                        setTimeout(function () {
                            busy.hidden = true;
                            showLink.hidden = false;
                            linkToShare.innerHTML = services + "/api/getFile/" + fileData.shortUrl;
                        }, 1000);
                    })
                    .error(function (err) {
                        console.log("fail!");
                    });
            }, 1000);

        })
        .error(function (err) {

        });
}
