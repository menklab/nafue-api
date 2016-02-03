function modalInit() {
    var modal = document.getElementById('modal');
    var content = document.getElementById('modal-content');
    var background = document.getElementById('modal-background');

    function hide_modal() {
        dom.paymentForm.innerHTML = null;
        modal.hidden = true;
    }

    function show_modal() {
        modal.hidden = false;

        // init payment stuff
        http.get(api_services + "/api/clientToken")
            .success(function (data) {
                show(dom.paymentForm);
                var clientToken = data.token;
                braintree.setup(clientToken, "dropin", {
                    container: "payment-form"
                });
                setTimeout(function() {
                    hide(dom.checkout_loading);
                }, 3000);
            })
            .error(function (e) {
                show(dom.checkout_error);
                hide(dom.checkout_loading);
            });
    }

    // register onClick of background
    background.addEventListener("click", function (e) {
        if (e.target.id === 'modal-background') {
            hide();
        }
    });

    // register escape key
    window.addEventListener("keyup", function (e) {
        if (!modal.hidden && e.keyCode == 27) {
            hide();
        }
    }, false);

    function getToken() {
        console.log("GetToken");
        // init payment stuff
        http.get(api_services + "/api/clientToken")
            .success(function (data) {
                show(dom.paymentForm);
                hide(dom.checkout_loading);
                var clientToken = data.token;
                braintree.setup(clientToken, "dropin", {
                    container: "payment-form"
                });
            })
            .error(function (e) {
                show(dom.checkout_error);
                hide(dom.checkout_loading);
            });
    }

    return {
        show: show_modal,
        hide: hide_modal,
        getToken: getToken
    };


}

