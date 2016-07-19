class Utility {

    static uint8ToBytes(buf) {
        return String.fromCharCode.apply(null, buf)
    }

    static bytesToUint8(buf) {
        var u8 = new Uint8Array(buf.split("").map(function (c) {
            return c.charCodeAt(0);
        }));
        return u8;
    }

}