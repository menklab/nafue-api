class FileHeader {

    constructor(prop) {
        this.salt = prop.salt || Cipher.makeSalt();
        this.name = prop.file.name;
        this.nameAry = this.getNameAry();
        this.fileSize = prop.file.size;
        this.secureFileSize = this.getTotalSecuredSize();
    }

    getTotalSecuredSize() {
        var t = this.fileSize;
        t += Cipher.config().KEY_SIZE;
        t += this.getNameAry().length;
        return t;
    }

    getNameAry() {
        return Utility.bytesToUint8(this.name);
    }
}