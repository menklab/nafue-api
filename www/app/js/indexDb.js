function dbInit(cb) {
    var indexedDB = window.indexedDB || window.mozIndexedDB || window.webkitIndexedDB || window.msIndexedDB || window.shimIndexedDB;
    var db;
// Open (or create) the database
    var req = indexedDB.open("www.nafue.com", 1);

// Create the schema
    req.onupgradeneeded = function (evt) {
        evt.currentTarget.result.createObjectStore("fileHeaders", {autoIncrement: true});
        var chunkStore = evt.currentTarget.result.createObjectStore("fileChunks", {autoIncrement: true});
        chunkStore.createIndex("fileChunks", "fileHeaderId", {unique: false});

    };

    req.onsuccess = function () {
        db = this.result;
        cb(db)
    };
}