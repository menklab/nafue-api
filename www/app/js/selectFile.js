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
    var file = files[0];

    var reader = new FileReader();

    // closure to capture file
    reader.onload = (function (f) {
        return function (e) {
            // Render thumbnail.
            binStr = e.target.result;
            //console.log("binStr: ", binStr);
            dropZone.hidden = true;
            passCont.hidden = false;
            password.focus();
        };
    })(file);

    // Read in the image file as a data URL.

    reader.readAsBinaryString(file);

}

function handleDragOver(e) {
    e.stopPropagation();
    e.preventDefault();
    e.dataTransfer.dropEffect = 'copy'; // Explicitly show this is a copy.
}