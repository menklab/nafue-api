class FileHeader {

    constructor(prop) {
        this.name = prop.name;
        this.nameAry = this.getNameAry();
        this.fileSize = prop.size;
        this.secureFileSize = this.getTotalSecuredSize();
    }

    getTotalSecuredSize() {
        var t = this.fileSize;
        t += Crypt.config().KEY_SIZE;
        t += this.getNameAry().length;
        return t;
    }

    getNameAry() {
        return Utility.bytesToUint8(this.name);
    }
}