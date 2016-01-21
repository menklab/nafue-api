
function loaded() {
    sjcl.random.startCollectors();
}

/* Encrypt a message */
function doEncrypt(password, d) {
    var ct = {};
    var p = {
        mode: s.mode,
        ts: s.tagSize,
        ks: s.keySize,
        iter: s.iterations,
        iv: makeIv(),
        adata: s.adata,
        cipher: s.cipher
    };
    ct.p = p;

    // setup password as key
    var tmp = sjcl.misc.cachedPbkdf2(password, p);
    password = tmp.key.slice(0, p.ks / 32);
    var prp = new sjcl.cipher[p.cipher](password);
    p.salt = tmp.salt;

    // encrypt
    ct.ct = sjcl.mode[p.mode].encrypt(prp, d, p.iv, p.adata, p.ts);
    return ct;
}

function doDecrypt(password, ct) {
    var p = {
        mode: s.mode,
        ts: s.tagSize,
        ks: s.keySize,
        iter: s.iterations,
        iv: ct.p.iv,
        adata: ct.p.adata,
        salt: ct.p.salt,
        cipher: s.cipher
    };

    // setup password as key
    var tmp = sjcl.misc.cachedPbkdf2(password, p);
    password = tmp.key.slice(0, p.ks / 32);
    var prp = new sjcl.cipher[p.cipher](password);

    try {
        data = sjcl.mode[p.mode].decrypt(prp, ct.ct, p.iv, p.adata, p.ts);
    }
    catch (e) {
        console.log("error: ", e);
        return;
    }
    return data;
}

function makeIv() {
    return sjcl.random.randomWords(4, 0);
}

function checkPassword(e) {
    if (e.keyCode == 13) {
        if (mode == "upload"){
            share.click();
        }
        else {
            downloadBtn.click();
            return;
        }
    }

    var pass = password.value;

    var p = {};

    if (pass.length > 7) {
        p.length = true;
        len.className = "present";
    }
    else {
        p.length = false;
        len.className = "missing";
    }

    if (!!pass.match(/[0-9]/)) {
        p.num = true;
        num.className = "present"
    }
    else {
        p.num = false;
        num.className = "missing"
    }

    if (!!pass.match("[a-z]")) {
        p.lower = true;
        low.className = "present";
    }
    else {
        p.lower = false;
        low.className = "missing";
    }

    if (!!pass.match("[A-Z]")) {
        p.upper = true;
        upp.className = "present";
    }
    else {
        p.upper = false;
        upp.className = "missing";
    }

    if (!!pass.match(/[!,@,#,$,%,^,&,*,?,_,~,-,(,),\s]/)) {
        p.special = true;
        spc.className = "present";
    }
    else {
        p.special = false;
        spc.className = "missing";
    }


    // update button when password is good enough
    var strength = 0;
    strength += 1 ? p.num : 0;
    strength += 1 ? p.lower : 0;
    strength += 1 ? p.upper : 0;
    strength += 1 ? p.special : 0;
    if (strength >= 3 && p.length) {
        share.disabled = false;
    }
    else {
        share.disabled = true;
    }
}