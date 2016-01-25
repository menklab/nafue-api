/*
 * http module for making ajax requests
 * methods:
 * http.get('url')
 */
var services = "http://localhost:9090";

var http = (function () {

    var parse = function (req) {
        var result;
        try {
            result = JSON.parse(req.responseText);
        } catch (e) {
            result = req.responseText;
        }
        return result;
    };

    function ajx(method, url, data, config) {
        var methods = {
            success: function () {
            },
            error: function () {
            }
        };

        var XHR = window.XMLHttpRequest || ActiveXObject;
        var request = new XHR('MSXML2.XMLHTTP.3.0');
        request.open(method, url, true);
        var auth = localStorage.getItem('ptau');
        if (auth && typeof config.sendWithAuth !== 'undefined' && config.sendWithAuth) {
            request.setRequestHeader('Authorization', auth);
        }
        if (typeof config.contentType === 'undefined') {
            request.setRequestHeader('Content-type', 'application/json; charset=utf-8');
        }
        request.onreadystatechange = function () {
            if (request.readyState === 4) {
                if (request.status >= 200 && request.status < 300) {
                    methods.success(parse(request));
                } else {
                    methods.error(parse(request));
                }
            }
        };
        request.send(data);
        var callbacks = {
            success: function (callback) {
                methods.success = callback;
                return callbacks;
            },
            error: function (callback) {
                methods.error = callback;
                return callbacks;
            }
        };
        return callbacks;
    }


    return {
        get: function (url, data, config) {
            config = config || {};
            return ajx('GET', url, data, config);
        },
        post: function (url, data, config) {
            config = config || {};
            return ajx('POST', url, data, config);
        },
        put: function (url, data, config) {
            config = config || {};
            return ajx('PUT', url, data, config);
        },
        'delete': function (url, data, config) {
            config = config || {};
            return ajx('DELETE', url, data, config);
        }
    };
})();