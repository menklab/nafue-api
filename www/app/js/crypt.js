class Crypt {

    static config() {
        return {
            ITERATIONS: 1000,
            KEY_SIZE: 32
        };
    }

    // create cipher
    constructor(prop) {

        this.salt = prop.salt || Crypt.makeSalt();
        this.iv = prop.iv || Crypt.makeIv();
        this.key = Crypt.genPbkdf2Key(prop.password, this.salt);
        this.cipher = forge.cipher.createCipher('AES-CTR', this.key);
        this.cipher.start({iv: this.iv});
        this.hmac = forge.hmac.create();
        this.hmac.start('sha256', this.key);
    }

    // secure destroy cipher
    destroy() {
        this.cipher.finish();
        delete this.salt;
        delete this.iv;
        delete this.key;
        delete this.cipher;
        var h = this.hmac.digest().toHex();
        delete this.hmac;
        return h;
    }

    // xor encrypt/decrypt in Uint8Array
    xorUintAry(data) {
        this.cipher.update(new forge.util.ByteBuffer(data));
    }

    // make an iv
    static makeIv() {
        return forge.random.getBytesSync(Crypt.config().KEY_SIZE);
    }

    // make salt
    static makeSalt() {
        return forge.random.getBytesSync(Crypt.config().KEY_SIZE);
    }

    // gen a key
    static genPbkdf2Key(password, salt) {
        return forge.pkcs5.pbkdf2(password, salt, Crypt.config().ITERATIONS, Crypt.config().KEY_SIZE);
    }

}




