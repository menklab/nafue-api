function tryPasswordToDecrpyt() {
    busy.hidden = true;
    skipCheck = true;
    passCont.hidden = false;
    passwordReqs.hidden = true;
    share.hidden = true;
    downloadBtn.hidden = false;
    //aData: "c3BhcnRpY3Vz"
    //downloadUrl: "https://s3.amazonaws.com/sfds.menklab.com/files/d4b0baf3-b858-429b-42ad-29fe5225c59d?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAIJBIQ6WXSSKJRWXQ%2F20160120%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20160120T130600Z&X-Amz-Expires=900&X-Amz-SignedHeaders=host&X-Amz-Signature=c55ad365bb86887747027969f9bebc456c9d976306d389305860135b5f296c4b"
    //iv: "ATsWI+N9iS80SZe69xxJBQ=="
    //salt: "ZQrix1NJtz0="
    //shortUrl: "4d8a41e7-c480-490d-592a-f337ee22c8d5"
}

function decrypt() {
    console.log("password: ", password.value);
    doDecrypt(password.value, ct);
}