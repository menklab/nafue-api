function domInit() {
    return {
        preview: document.getElementById('preview'),
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
        badPass: document.getElementById("badPass"),
        error: document.getElementById("error"),
        errMsg: document.getElementById("errMsg")
    }
}

function error(errMsg) {
    reset_ui();
    hide(dom.dropZone);

    show(dom.error);
    resetPassword();
    g.binData = null;
    dom.errMsg.innerHTML = errMsg;
}

function badPassword() {
    show(dom.badPass);
    dom.password.value = "";
}

function reset_ui() {
    g = {};
    g.mode = "upload";
    dom.file.value = "";
    show(dom.dropZone);
    hide(dom.showLink);
    hide(dom.passCont);
    show(dom.passwordReqs);
    hide(dom.doneDownloading);
    hide(dom.busy);
    hide(dom.unsupported);
    hide(dom.error);
    hide(dom.badPass);
    dom.share.disabled = true;
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