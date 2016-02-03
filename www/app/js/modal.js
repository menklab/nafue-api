function modalInit() {
    var modal = document.getElementById('modal');
    var content = document.getElementById('modal-content');
    var background = document.getElementById('modal-background');
    var bt;

    function hide_modal() {
        setTimeout(function () {
            if (!!bt) {
                bt.teardown(function () {
                    bt = null;
                });
            }
            //dom.paymentForm.innerHTML = null;
            dom.paymentForm.style.height = "0";
            dom.paymentForm.style.opacity = "0";
            hide(dom.amount);
            hide(dom.donate);
            show(dom.checkout_loading);
            hide(dom.thanks);
            hide(dom.checkout_error);
            modal.hidden = true;
        }, 0);
    }

    function show_modal() {
        setTimeout(function () {
            modal.hidden = false;

            // init payment stuff
            http.get(api_services + "/api/payment")
                .success(function (data) {
                    show(dom.paymentForm);
                    var clientToken = data.token;
                    braintree.setup(clientToken, "dropin", {
                        container: "payment-form",
                        onReady: function (integration) {
                            bt = integration;
                            hide(dom.checkout_loading);
                            show(dom.amount);
                            show(dom.donate);
                            dom.paymentForm.style.height = "auto";
                            dom.paymentForm.style.opacity = "100";
                        },
                        onPaymentMethodReceived: function (obj) {
                            var payload = {
                                amount: dom.amount.value,
                                nonce: obj.nonce
                            };
                            dom.amount.value = null;
                            hide(dom.amount);
                            hide(dom.donate);
                            dom.paymentForm.style.height = "0";
                            dom.paymentForm.style.opacity = "0";


                            console.log("payload: ", payload);
                            http.post(api_services + "/api/payment", JSON.stringify(payload))
                                .success(function (res) {
                                    show(dom.thanks);
                                })
                                .error(function (err) {
                                    show(dom.checkout_error);
                                    hide(dom.checkout_loading);
                                });
                        }
                    });
                })
                .error(function (e) {
                    show(dom.checkout_error);
                    hide(dom.checkout_loading);
                });
        }, 0);
    }

    // register onClick of background
    background.addEventListener("click", function (e) {
        if (e.target.id === 'modal-background' && !document.getElementById('modal-background').hidden) {
            hide_modal();
        }
    });

    // register escape key
    window.addEventListener("keyup", function (e) {
        if (!modal.hidden && e.keyCode == 27) {
            hide_modal();
        }
    }, false);

    function getToken() {
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

    dom.amount.addEventListener('blur', function () {
        if (this.value === '') {
            return;
        }
        this.setAttribute('type', 'text');
        if (this.value.indexOf('.') === -1) {
            this.value = this.value + '.00';
        }
        if (this.value.indexOf('.') === 0) {
            this.value = '1.00';
        }
        while (this.value.indexOf('.') > this.value.length - 3) {
            this.value = this.value + '0';
        }
    });
    dom.amount.addEventListener('focus', function () {
        this.setAttribute('type', 'number');
    });

    return {
        show: show_modal,
        hide: hide_modal,
        getToken: getToken
    };


}

