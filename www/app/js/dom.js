function domInit() {
    return {
        preview: document.getElementById('preview'),
        dropZone: document.getElementById('dropZone'),
        list: document.getElementById('list'),
        file: document.getElementById('file'),
        password: document.getElementById('password'),
        passCont: document.getElementById('passCont'),
        busy: document.getElementById('busy'),
        showLink: document.getElementById('showLink'),
        linkToShare: document.getElementById('linkToShare'),
        busyMessage: document.getElementById('busyMessage'),
        share: document.getElementById("share"),
        downloadBtn: document.getElementById("downloadBtn"),
        reset: document.getElementById("reset"),
        doneDownloading: document.getElementById("doneDownloading"),
        unsupported: document.getElementById("unsupported"),
        badPass: document.getElementById("badPass"),
        error: document.getElementById("error"),
        errMsg: document.getElementById("errMsg"),
        paymentForm: document.getElementById("payment-form"),
        checkout_loading: document.getElementById("checkout_loading"),
        checkout_error: document.getElementById("checkout_error"),
        amount: document.getElementById("amount"),
        donate: document.getElementById("donate"),
        passwordStrength: document.getElementById("passwordStrength"),
        passwordSuggestions: document.getElementById("passwordSuggestions"),
        terrible: document.getElementById("terrible"),
        weak: document.getElementById("weak"),
        mediocre: document.getElementById("mediocre"),
        strong: document.getElementById("strong"),
        excellent: document.getElementById("excellent"),
        thanks: document.getElementById("thanks"),
        passwordStrengthBar: document.getElementById("passwordStrengthBar")

    };
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
    hide(dom.doneDownloading);
    hide(dom.busy);
    hide(dom.unsupported);
    hide(dom.error);
    hide(dom.badPass);
    hide(dom.passwordStrength);
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

function updatePasswordStrength(results) {
    var s = 0;
    if (!!results && !!results.score) {
        s = results.score; // score
    }

    // enable button if password has value
    if (!!dom.password.value) {
        dom.share.disabled = false;
    }
    else {
        dom.share.disabled = true;
    }

    // calculate percent
    var perc = ((s + 1)/5) * 100;
    if (!dom.password.value || dom.password.value === '') {
        perc = 0;
    }

    setBarPercent(dom.passwordStrengthBar, perc);
}
