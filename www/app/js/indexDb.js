function dbInit(cb) {
    var indexedDB = window.indexedDB || window.mozIndexedDB || window.webkitIndexedDB || window.msIndexedDB || window.shimIndexedDB;
    var db;
// Open (or create) the database
    var req = indexedDB.open("www.nafue.com", 1);

// Create the schema
    req.onupgradeneeded = function (evt) {
        evt.currentTarget.result.createObjectStore("fileHeaders");
        evt.currentTarget.result.createObjectStore("fileChunks");

        //store.createIndex('biblioid', 'biblioid', { unique: true });

    };

    req.onsuccess = function () {
        db = this.result;
        cb(db)
    };
}