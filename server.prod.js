var port = 8080;

var http = require('http');
var express = require('express'),
    app = module.exports.app = express();
var server = http.createServer(app);
var request = require('request');

app.set('views', __dirname + 'www/dist/');
app.engine('html', require('ejs').renderFile);


var serviceBase = 'http://localhost:9090';

// static routes for html
app.use("/maps", express.static("./www/maps"));
app.use("/docs", express.static("./www/dist/docs"));
app.use("/", express.static("./www/dist/"));

app.get('/api/*', function(req, res) {
    var url = serviceBase + req.url;
    console.log('[GET]: ' + url);
    var r = request(url);
    req.pipe(r).pipe(res);
});

app.post('/api/*', function(req, res) {
    var url = serviceBase + req.url;
    console.log('[POST]: ' + url);
    var options = {uri: url};
    if (req.body) {
        options.form = req.body;
    }
    var r = request.post(options);
    req.pipe(r).pipe(res);
});

app.put('/api/*', function(req, res) {
    var url = serviceBase + req.url;
    console.log('[PUT]: ' + url);
    var options = {uri: url};
    if (req.body) {
        options.form = req.body;
    }
    var r = request.put(options);
    req.pipe(r).pipe(res);
});

app.delete('/api/*', function(req, res) {
    var url = serviceBase + req.url;
    console.log('[DELETE]: ' + url);
    var options = {uri: url};
    if (req.body) {
        options.form = req.body;
    }
    var r = request.del(options);
    req.pipe(r).pipe(res);
});

app.all("/*", function(req, res) {
    console.log('Could not find route: ' + req.originalUrl);
    res.render('index.html');
});

console.log("listening on: " + port);
server.listen(port);

