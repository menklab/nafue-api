/* enter actions */
var s = {
    adata: 'sparticus',
    mode: 'ccm',
    tagSize: 128,
    keySize: 256,
    iterations: 1000
};

function loaded() {
    sjcl.random.startCollectors();
}

/* compute PBKDF2 on the password. */
function doPbkdf2(password, decrypting) {
    var salt = !!decrypting ? decrypting.salt : makeSalt();

    if (password.length == 0) {
        if (decrypting) {
            error("Can't decrypt: need a password!");
        }
        return;
    }

    if (salt.length === 0 && decrypting) {
        error("Can't decrypt: need a salt for PBKDF2!");
        return;
    }

    var p = {
        iter: 1000,
        salt: salt
    };


    p = sjcl.misc.cachedPbkdf2(password, p);
    console.log("generated pbkdf2: ", p);
    return p;
}

/* Encrypt a message */
function doEncrypt(password, d) {
    var p = {
        mode: s.mode,
        ts: s.tagSize,
        ks: s.keySize,
        iter: s.iterations,
        iv: makeIv(),
        adata: s.adata
    };
    var rp = {};
    var ct = sjcl.encrypt(password, d, p, rp); //.replace(/,/g, ",\n");
    return ct;
}

/* Decrypt a message */
function doDecrypt() {
    var v = form.get(), iv = v.iv, key = v.key, adata = v.adata, aes, ciphertext = v.ciphertext, rp = {};

    if (ciphertext.length === 0) {
        return;
    }
    if (!v.password && !v.key.length) {
        error("Can't decrypt: need a password or key!");
        return;
    }

    if (ciphertext.match("{")) {
        /* it's jsonized */
        try {
            v.plaintext = sjcl.decrypt(v.password || v.key, ciphertext, {}, rp);
        } catch (e) {
            error("Can't decrypt: " + e);
            return;
        }
        v.mode = rp.mode;
        v.iv = rp.iv;
        v.adata = rp.adata;
        if (v.password) {
            v.salt = rp.salt;
            v.iter = rp.iter;
            v.keysize = rp.ks;
            v.tag = rp.ts;
        }
        v.key = rp.key;
        v.ciphertext = "";
        document.getElementById('plaintext').select();
    } else {
        /* it's raw */
        ciphertext = sjcl.codec.base64.toBits(ciphertext);
        if (iv.length === 0) {
            error("Can't decrypt: need an IV!");
            return;
        }
        if (key.length === 0) {
            if (v.password.length) {
                doPbkdf2(true);
                key = v.key;
            }
        }
        aes = new sjcl.cipher.aes(key);

        try {
            v.plaintext = sjcl.codec.utf8String.fromBits(sjcl.mode[v.mode].decrypt(aes, ciphertext, iv, v.adata, v.tag));
            v.ciphertext = "";
            document.getElementById('plaintext').select();
        } catch (e) {
            error("Can't decrypt: " + e);
        }
    }
    form.set(v);
}

function extendKey(size) {
    form.key.set(form._extendedKey.slice(0, size));
}

function makeSalt() {
    return sjcl.random.randomWords(2, 0);
}
function makeIv() {
    return sjcl.random.randomWords(4, 0);
}