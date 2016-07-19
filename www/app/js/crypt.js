class Cipher {

    static config() {
        return {
            ITERATIONS: 1000,
            KEY_SIZE: 32
        };
    }

    // create cipher
    constructor(prop) {

        this.salt = prop.salt || Cipher.makeSalt();
        this.iv = prop.iv || Cipher.makeIv();
        this.key = Cipher.genPbkdf2Key(prop.password, this.salt);
        this.cipher = forge.cipher.createCipher('AES-CTR-', this.key);
        this.cipher.start({iv: this.iv});
    }

    // secure destroy cipher
    destroy() {
        this.cipher.finish();
        delete this.salt;
        delete this.iv;
        delete this.key;
        delete this.cipher;
    }

    // xor encrypt/decrypt in Uint8Array
    xorUintAry(data) {
        this.cipher.update(new forge.util.ByteBuffer(data));
    }

    // make an iv
    static makeIv() {
        return forge.random.getBytesSync(this.KEY_SIZE);
    }

    // make salt
    static makeSalt() {
        return forge.random.getBytesSync(this.KEY_SIZE);
    }

    // gen a key
    static genPbkdf2Key(password, salt) {
        return forge.pkcs5.pbkdf2(password, salt, this.ITERATIONS, this.KEY_SIZE);
    }

}




