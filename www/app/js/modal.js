function modalInit() {
    var modal = document.getElementById('modal');
    var content = document.getElementById('modal-content');
    var background = document.getElementById('modal-background');

    function hide() {
        modal.hidden = true;
    }

    function show() {
        modal.hidden = false;
    }

    // register onClick of background
    background.addEventListener("click", function(e) {
        if (e.target.id === 'modal-background') {
            hide();
        }
    });

    // register escape key
    window.addEventListener("keyup", function(e) {
        if (!modal.hidden && e.keyCode == 27) {
            hide();
        }
    }, false);


    return {
        show: show,
        hide: hide
    };
}

