function domInit() {
    return {
        dropZone: document.getElementById('dropZone'),
        list: document.getElementById('list'),
        file: document.getElementById('file'),
        password: document.getElementById('password'),
        passwordReqs: document.getElementById('passwordReqs'),
        passCont: document.getElementById('passCont'),
        busy: document.getElementById('busy'),
        showLink: document.getElementById('showLink'),
        linkToShare: document.getElementById('linkToShare'),
        busyMessage: document.getElementById('busyMessage'),
        share: document.getElementById("share"),
        downloadBtn: document.getElementById("downloadBtn"),
        len: document.getElementById("len"),
        low: document.getElementById("low"),
        upp: document.getElementById("upp"),
        num: document.getElementById("num"),
        spc: document.getElementById("spc"),
        reset: document.getElementById("reset"),
        doneDownloading: document.getElementById("doneDownloading"),
        unsupported: document.getElementById("unsupported"),
    }
}

function reset_ui() {
    g.mode = "upload";
    dom.file.value = "";
    show(dom.dropZone);
    hide(dom.showLink);
    hide(dom.passCont);
    hide(dom.doneDownloading);
    hide(dom.busy);
    hide(dom.unsupported);
    resetPassword();
    g.binData = null;
    dom.file.focus();
}

function hide(e) {
    e.hidden = true;
}
function show(e) {
    e.hidden = false;
}
function setContent(e, c) {
    e.innerHTML = c;
}