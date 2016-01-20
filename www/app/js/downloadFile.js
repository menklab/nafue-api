var ct = {};
function downloadFile(file) {
    upload.hidden = true;
    downloadBtn.hidden = false;
    busyMessage.innerHTML = "Downloading";
    busy.hidden = false;

    http.get(services + "/api/files/" + file)
        .success(function (res) {
            http.get(res.downloadUrl)
                .success(function (eData) {
                    res.ct = eData;
                    ct = res;
                    tryPasswordToDecrpyt();
                })
                .error(function (err) {
                    console.log("error");
                })
        })
        .error(function (err) {
            console.log("error");
        });
}