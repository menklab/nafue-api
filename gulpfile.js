var gulp = require('gulp'),
    path = require('path'),
    plumber = require('gulp-plumber'),
    gulp_watch = require('gulp-watch'),
    gulp_concat = require('gulp-concat'),
    gulp_uglifycss = require('gulp-uglifycss'),
    gulp_uglify = require('gulp-uglify'),
    gulp_copy = require('gulp-copy'),
    gulp_rename = require("gulp-rename"),
    gulp_insert = require('gulp-insert'),
    gulp_open = require('gulp-open'),
    gulp_replace = require('gulp-replace'),
    gulp_sourcemaps = require('gulp-sourcemaps'),
    gulp_apidoc = require('gulp-apidoc'),
    jshint = require('gulp-jshint'),
    jshint_stylish = require('jshint-stylish'),
    minimist = require('minimist'),
    run_sequence = require('run-sequence'),
    fs = require('fs'),
    livereload = require('gulp-livereload');


var options = minimist(process.argv.slice(2), {
    string: 'env',
    default: {env: process.env.NODE_ENV || 'dev'}
});

var sourcemapsWriteOptions = {
    includeContent: false,
    sourceRoot: function (file) {
        return path.dirname(
            path.relative('maps', file.path))
            .replace(/\\/g, '/');
    }
};

var jshintOptions = {
    debug: options.env === 'dev',
    devel: options.env === 'dev'
};

gulp.task('build:css', function () {
    return gulp.src([
            'www/app/styles/**/*.css'
        ])
        .pipe(gulp_sourcemaps.init())
        .pipe(gulp_concat('style.min.css'))
        .pipe(gulp_sourcemaps.write('../maps'))
        .pipe(gulp.dest('www/app/'))
        .pipe(livereload());
});
gulp.task('build:vendor:css', function () {
    return gulp.src([
            'www/bower_components/fullpage.js/dist/jquery.fullpage.min.css'
        ])
        .pipe(gulp_concat('vendor.min.css'))
        .pipe(gulp.dest('www/app/'))
        .pipe(livereload());
});
gulp.task('build:js', function () {
    gulp.src(['www/app/js/**/*.js'], {base: 'js'})
        .pipe(jshint(jshintOptions))
        .pipe(jshint.reporter(jshint_stylish))
        //.pipe(jshint.reporter('fail')) // only enable if build needs to fail on bad jshint
        .pipe(gulp_sourcemaps.init())
        .pipe(gulp_concat('app.min.js'))
        //.pipe(gulp_uglify())
        .pipe(gulp_sourcemaps.write('../maps'))
        .pipe(gulp.dest('www/app/'))
        .pipe(livereload());
});

gulp.task('build:vendor:js', function () {
    gulp.src([
            'www/bower_components/forge/js/forge.min.js',
            'www/bower_components/jquery/dist/jquery.min.js',
            'www/bower_components/braintree-web/dist/braintree.js',
            'www/bower_components/file-saver/FileSaver.js',
            'www/bower_components/fullpage.js/dist/jquery.fullpage.min.js',
            'www/bower_components/zxcvbn/dist/zxcvbn.js'
        ], {base: 'js'})
        .pipe(gulp_sourcemaps.init())
        .pipe(gulp_concat('vendor.min.js'))
        //.pipe(gulp_uglify())
        .pipe(gulp_sourcemaps.write('../maps'))
        .pipe(gulp.dest('www/app/'));
});

gulp.task('package', function () {
    run_sequence('apidoc');
    run_sequence('build:vendor:js');
    run_sequence('build:vendor:css');
    run_sequence('build:js');
    run_sequence('build:css');
    gulp.src(['www/app/app.min.js'], {base: 'js'})
        .pipe(gulp_rename("app.min.js"))
        .pipe(gulp_uglify())
        .pipe(gulp_replace(/api_services=".*?"/, 'api_services="https://api.nafue.com"'))
        .pipe(gulp_replace(/www_services=".*?"/, 'www_services="https://www.nafue.com"'))
        .pipe(gulp.dest('www/dist/'));
    gulp.src(['www/app/vendor.min.js'], {base: 'js'})
        .pipe(gulp_uglify())
        .pipe(gulp_rename("vendor.min.js"))
        .pipe(gulp.dest('www/dist/'));
    gulp.src(['www/app/vendor.min.css'])
        .pipe(gulp_uglifycss())
        .pipe(gulp_rename("vendor.min.css"))
        .pipe(gulp.dest('www/dist/'));
    gulp.src(['www/app/style.min.css'])
        .pipe(gulp_uglifycss())
        .pipe(gulp_rename("style.min.css"))
        .pipe(gulp.dest('www/dist/'));
    gulp.src(['www/app/img/**/*', 'www/app/**/*.html', 'www/app/fonts/**/*'])
        .pipe(gulp_copy('www/dist/', {prefix: 2}));

});

gulp.task('apidoc', function (done) {
    gulp_apidoc({
        src: "controllers/",
        dest: "www/dist/docs",
        includeFilters: [".*\\.go$"]
    }, done);
});

gulp.task('open:dev', function () {
    gulp.src('')
        .pipe(gulp_open({uri: 'http://localhost:8080/'}));
});

gulp.task('dev', [], function () {
    run_sequence('build:vendor:js');
    run_sequence('build:vendor:css');
    run_sequence('build:js');
    run_sequence('build:css');
    run_sequence('watch');
    run_sequence('open:dev');
});


gulp.task('watch', [], function () {
    livereload.listen({port: 35729});
    gulp_watch('www/app/styles/**/*.css', function () {
        run_sequence('build:css');
    });
    gulp_watch('www/app/js/**/*.js', function () {
        run_sequence('build:js');
    });
    gulp.watch(["www/app/**/*.html"]).on('change', function (file) {
        livereload.changed(file.path);
    });
});

gulp.task('default', function () {


});