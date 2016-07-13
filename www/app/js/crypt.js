function makeSecure() {
//    sjcl.random.startCollectors();
//    sjcl.random.getProgress(100);
}
//
///* Encrypt a message */
function doEncrypt(password, d) {
//    var ct = {};
//    var p = {
//        mode: s.mode,
//        ts: s.tagSize,
//        ks: s.keySize,
//        iter: s.iterations,
//        iv: makeIv(),
//        adata: s.adata,
//        cipher: s.cipher
//    };
//    ct.p = p;
//
//    // setup password as key
//    var tmp = sjcl.misc.cachedPbkdf2(password, p);
//    password = tmp.key.slice(0, p.ks / 32);
//    var prp = new sjcl.cipher[p.cipher](password);
//    p.salt = tmp.salt;
//
//    // encrypt
//    ct.ct = sjcl.mode[p.mode].encrypt(prp, d, p.iv, p.adata, p.ts);
//    return ct;
}
//
function doDecrypt(password, ct) {
//    var p = {
//        mode: s.mode,
//        ts: s.tagSize,
//        ks: s.keySize,
//        iter: s.iterations,
//        iv: ct.p.iv,
//        adata: ct.p.adata,
//        salt: ct.p.salt,
//        cipher: s.cipher
//    };
//
//    // setup password as key
//    var tmp = sjcl.misc.cachedPbkdf2(password, p);
//    password = tmp.key.slice(0, p.ks / 32);
//    var prp = new sjcl.cipher[p.cipher](password);
//
//    try {
//        data = sjcl.mode[p.mode].decrypt(prp, ct.ct, p.iv, p.adata, p.ts);
//    }
//    catch (e) {
//        badPassword();
//        return;
//    }
//    return data;
}
//
function makeIv() {
//    return sjcl.random.randomWords(8,0);
//    //return sjcl.random.randomWords(8,100);
}
//
//
function resetPassword() {
//    dom.password.value = "";
}