function checkPassword(e) {
    if (skipCheck) {
        console.log("skip pass check");
        return;
    }
    if (e.keyCode == 13) {
        share.click();
        return;
    }

    var pass = password.value;

    var p = {};

    if (pass.length > 7) {
        p.length = true;
        len.className = "present";
    }
    else {
        p.length = false;
        len.className = "missing";
    }

    if (!!pass.match(/[0-9]/)) {
        p.num = true;
        num.className = "present"
    }
    else {
        p.num = false;
        num.className = "missing"
    }

    if (!!pass.match("[a-z]")) {
        p.lower = true;
        low.className = "present";
    }
    else {
        p.lower = false;
        low.className = "missing";
    }

    if (!!pass.match("[A-Z]")) {
        p.upper = true;
        upp.className = "present";
    }
    else {
        p.upper = false;
        upp.className = "missing";
    }

    if (!!pass.match(/[!,@,#,$,%,^,&,*,?,_,~,-,(,),\s]/)) {
        p.special = true;
        spc.className = "present";
    }
    else {
        p.special = false;
        spc.className = "missing";
    }


    // update button when password is good enough
    var strength = 0;
    strength += 1 ? p.num : 0;
    strength += 1 ? p.lower : 0;
    strength += 1 ? p.upper : 0;
    strength += 1 ? p.special : 0;
    if (strength >= 3 && p.length) {
        share.disabled = false;
    }
    else {
        share.disabled = true;
    }
}