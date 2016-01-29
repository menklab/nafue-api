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
            'app/styles/**/*.css'
        ])
        .pipe(gulp_sourcemaps.init())
        .pipe(gulp_concat('style.min.css'))
        //.pipe(gulp_uglifycss())
        .pipe(gulp_sourcemaps.write('../maps'))
        .pipe(gulp.dest('app/'))
        .pipe(livereload());
});

gulp.task('build:js', function () {
        gulp.src(['app/js/**/*.js'], {base: 'js'})
        .pipe(jshint(jshintOptions))
        .pipe(jshint.reporter(jshint_stylish))
        //.pipe(jshint.reporter('fail')) // only enable if build needs to fail on bad jshint
        .pipe(gulp_sourcemaps.init())
        .pipe(gulp_concat('app.min.js'))
        //.pipe(gulp_uglify())
        .pipe(gulp_sourcemaps.write('../maps'))
        .pipe(gulp.dest('app/'))
        .pipe(livereload());
});

gulp.task('build:vendor:js', function () {
    // pre build sjcl with needed wrapper
    gulp.src('bower_components/sjcl/sjcl.js', {base: 'js'})
        .pipe(gulp_insert.prepend('var sjcl =(function() {'))
        .pipe(gulp_insert.append('return sjcl;})();'))
        .pipe(gulp_rename('sjcl_wrapped.js'))
        .pipe(gulp.dest('bower_components/sjcl/'));
    gulp.src([
            'bower_components/sjcl/sjcl_wrapped.js',
            'bower_components/file-saver/FileSaver.js'
        ], {base: 'js'})
        .pipe(gulp_sourcemaps.init())
        .pipe(gulp_concat('vendor.min.js'))
        //.pipe(gulp_uglify())
        .pipe(gulp_sourcemaps.write('../maps'))
        .pipe(gulp.dest('app/'));
});

gulp.task('package', function () {
    run_sequence('build:vendor:js');
    run_sequence('build:js');
    run_sequence('build:css');
    gulp.src(['app/app.min.js'], {base: 'js'})
        .pipe(gulp_rename("app.min.js"))
        .pipe(gulp_uglify())
        .pipe(gulp_replace(/api_services=".*?"/, 'api_services="https://api.nafue.com"'))
        .pipe(gulp_replace(/www_services=".*?"/, 'www_services="https://www.nafue.com"'))
        .pipe(gulp.dest('dist/app/'));
    gulp.src(['app/vendor.min.js'], {base: 'js'})
        .pipe(gulp_uglify())
        .pipe(gulp_rename("vendor.min.js"))
        .pipe(gulp.dest('dist/app/'));
    gulp.src(['app/style.min.css'])
        .pipe(gulp_uglifycss())
        .pipe(gulp_rename("style.min.css"))
        .pipe(gulp.dest('dist/app'));
    gulp.src(['app/img/**/*', 'app/**/*.html'])
        .pipe(gulp_copy('dist/'));
});

//gulp.task('build:vendor:css', function () {
//    gulp.src([
//    ])
//        .pipe(gulp_concat('vendor.css'))
//        .pipe(gulp.dest('app/src/styles/'));
//});

gulp.task('open:dev', function () {
    gulp.src('')
        .pipe(gulp_open({uri: 'http://localhost:8080/'}));
});

gulp.task('dev', [], function () {
    run_sequence('build:vendor:js');
    //run_sequence('build:vendor:css');
    run_sequence('build:js');
    run_sequence('build:css');
    run_sequence('watch');
    run_sequence('open:dev');
});


gulp.task('watch', [], function () {
    livereload.listen({port: 35729});
    gulp_watch('app/styles/**/*.css', function () {
        run_sequence('build:css');
    });
    gulp_watch('app/js/**/*.js', function () {
        run_sequence('build:js');
    });
    gulp.watch(['app/index.html', "app/templates/**/*.html"]).on('change', function (file) {
        livereload.changed(file.path);
    });
});

gulp.task('default', function () {


});