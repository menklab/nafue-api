class SecureFileChunker {

    static config() {
        return {
            fileSizeLimit: 1024 * 1024 * 50, // 50mb
            chunkSize: 32000 // 32kb
        }
    }

    constructor(prop) {
        this.file = prop.file;
        this.fileHeader = new FileHeader(this.file);
        this.reader = new FileReader();
    }

    sealFile() {
        // make cipher
        var cipher = new Cipher({password: 'password', salt: this.fileHeader.salt});

        // calc total chunks
        var tChunks = Math.ceil(this.fileHeader.secureFileSize / SecureFileChunker.config().chunkSize);
        console.log("num of chunks: ", tChunks);
        console.log("tfilesize: ", this.fileHeader.secureFileSize);

        // add fileHeader to indexddb
        var req = g.db.headers().add(this.fileHeader);
        req.onsuccess = function (event) {
            this.fileHeader.id = event.target.result;
            console.log("file header: ", this.fileHeader.id);
            // read first chunk
            SecureFileChunker.readChunk(cipher, this.fileHeader.id, tChunks, 0, function () {
                console.log("all chunks encrypted!");
            });
        };


    }

    static readChunk(cipher, fileHeaderId, tChunks, curChunk, cb) {
        this.reader.onloadend = function (evt) {
            if (evt.target.readyState == FileReader.DONE) {

                // get data from read
                var buffer = evt.target.result;
                var data = new Uint8Array(buffer);

                // encrypt data from read
                cipher.update(new forge.util.ByteBuffer(data));
                var eData = Utility.bytesToUint8(cipher.output.getBytes());
                var addedChunkEvt = g.db.chunks().add(
                    {
                        fileHeaderId: fileHeaderId,
                        data: eData
                    }
                );

                addedChunkEvt.onsuccess = function () {
                    // if there are more chunks read the next one
                    if (curChunk < (tChunks - 1)) {
                        SecureFileChunker.readChunk(cipher, fileHeaderId, tChunks, (curChunk + 1), cb);
                    }
                    // otherwise finish last block with extra data and close
                    else {
                        var paddedFileName = new Uint8Array(255);
                        paddedFileName.fill(0);
                        paddedFileName.set(this.fileHeader.nameAry);
                        cipher.update(new forge.util.ByteBuffer(paddedFileName));
                        eData = Utility.bytesToUint8(cipher.output.getBytes());
                        g.db.chunks().add(
                            {
                                fileHeaderId: fileHeaderId,
                                data: eData
                            }
                        );
                        cb();
                    }
                };
            }
        }
    }

    checkSizeLimit() {
        if (this.fileHeader.fileSize / 1024 / 1024 > 50) {
            this.reader.abort();
            error('The uploaded file cannot be greater than 50MB');
        }
    }
}
