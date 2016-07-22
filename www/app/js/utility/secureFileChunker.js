class SecureFileChunker {

    static config() {
        return {
            fileSizeLimit: 1024 * 1024 * 50, // 50mb
            chunkSize: 3200 // 32kb
        }
    }

    constructor(prop) {
        this.file = prop;
        this.fileHeader = new FileHeader(this.file);
        this.reader = new FileReader();
        this.crypt = new Crypt({password: 'password'});


    }

    sealFile(cb) {
        var me = this;

        // calc total chunks
        var tChunks = Math.ceil(me.fileHeader.secureFileSize / SecureFileChunker.config().chunkSize);

        // add fileHeader to indexddb
        var req = g.db.headers().add(me.fileHeader);
        req.onsuccess = function (event) {
            me.fileHeader.id = event.target.result;

            // read first chunk
            me.readChunk(me, tChunks, 0, 0, function (tSize) {
                var hmac = me.crypt.destroy();
                cb(tSize, hmac);
            });
        };


    }

    readChunk(me, tChunks, curChunk, tSize, cb) {
        // calc chunk start/end
        var start = SecureFileChunker.config().chunkSize * curChunk;
        var end = (SecureFileChunker.config().chunkSize * curChunk) + SecureFileChunker.config().chunkSize;
        if (curChunk == (tChunks - 1)) { // if on last chunk end == last byte
            end = me.fileHeader.secureFileSize;
        }
        me.reader.onloadend = function (evt) {
            if (evt.target.readyState == FileReader.DONE) {

                // get data from read
                var buffer = evt.target.result;
                var data = new Uint8Array(buffer);

                // encrypt data from read
                me.crypt.cipher.update(new forge.util.ByteBuffer(data));
                var eData = Utility.bytesToUint8(me.crypt.cipher.output.getBytes());
                var addedChunkEvt = g.db.chunks().add(
                    {
                        fileHeaderId: me.fileHeader.id,
                        data: eData
                    }
                );

                addedChunkEvt.onsuccess = function () {
                    // sum data as it's added
                    tSize += eData.length;
                    me.crypt.hmac.update(eData);

                    // if there are more chunks read the next one
                    if (curChunk < tChunks) {
                        me.readChunk(me, tChunks, (curChunk + 1), tSize, cb);
                    }
                    // otherwise finish last block with extra data and close
                    else {
                        var paddedFileName = new Uint8Array(255);
                        paddedFileName.fill(0);
                        paddedFileName.set(me.fileHeader.nameAry);
                        me.crypt.cipher.update(new forge.util.ByteBuffer(paddedFileName));
                        eData = Utility.bytesToUint8(me.crypt.cipher.output.getBytes());

                        // take encrypted filename data and add the iv to it.
                        var eiData = new Uint8Array((255 + Crypt.config().KEY_SIZE));
                        eiData.set(eData);
                        var iv = Utility.bytesToUint8(me.crypt.iv);
                        eiData.set(iv, 255);
                        var addFinalChunkEvt = g.db.chunks().add(
                            {
                                fileHeaderId: me.fileHeader.id,
                                data: eiData
                            }
                        );

                        addFinalChunkEvt.onsuccess = function() {
                            me.crypt.hmac.update(eiData);
                            cb(tSize + eiData.length);
                        };
                    }
                };
            }
        };
       // do read
        var chunkToRead = me.file.slice(start, end);
        me.reader.readAsArrayBuffer(chunkToRead);
    }

    checkSizeLimit() {
        if (this.fileHeader.fileSize / 1024 / 1024 > 50) {
            this.reader.abort();
            error('The uploaded file cannot be greater than 50MB');
        }
    }
}
