(function () {
    if (window.File && window.FileReader && window.FileList && window.Blob) {

        var dropZone = document.getElementById('drop');
        var list = document.getElementById('list');
        var password = document.getElementById('password');
        var passCont = document.getElementById('passCont');
        var fileSize = document.getElementById('totalSize');

        dropZone.addEventListener('dragover', handleDragOver, false);
        dropZone.addEventListener('drop', handleFileSelect, false);
        document.getElementById('file').addEventListener('change', handleFileSelect, false);

    } else {
        document.getElementById('status').innerHTML = 'Your browser does not support the HTML5 FileReader.';
    }


    function handleFileSelect(e) {
        e.stopPropagation();
        e.preventDefault();

        var files;
        if (!!e.dataTransfer && !!e.dataTransfer.files) {
            files = e.dataTransfer.files
        }

        if (!!e.target && !!e.target.files) {
            files = e.target.files
        }

        // only 1 file at a time
        if (files.length > 1) {
            console.log("only 1 file at a time!");
            return;
        }
        console.log("file: ", files[0]);


    }

    function handleDragOver(e) {
        e.stopPropagation();
        e.preventDefault();
        e.dataTransfer.dropEffect = 'copy'; // Explicitly show this is a copy.
    }

    function encrypt() {
        // Each random "word" is 4 bytes, so 8 would be 32 bytes
        var saltBits = sjcl.random.randomWords(8);
        // eg. [588300265, -1755622410, -533744668, 1408647727, -876578935, 12500664, 179736681, 1321878387]

        // I left out the 5th argument, which defaults to HMAC which in turn defaults to use SHA256
        var derivedKey = sjcl.misc.pbkdf2("password", saltBits, 1000, 256);
        // eg. [-605875851, 757263041, -993332615, 465335420, 1306210159, -1270931768, -1185781663, -477369628]

        // Storing the key is probably easier encoded and not as a bitArray
        // I choose base64 just because the output is shorter, but you could use sjcl.codec.hex.fromBits
        var key = sjcl.codec.base64.fromBits(derivedKey);
        // eg. "2+MRdS0i6sHEyvJ5G7x0fE3bL2+0Px7IuVJoYeOL6uQ="
    }

})();