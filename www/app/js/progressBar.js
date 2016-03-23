// on browser resize...
$(window).resize(function () {
    moveProgressBar();
});

// SIGNATURE PROGRESS
function moveProgressBar() {
    console.log("moveProgressBar");
    var getPercent = ($('.progress-wrap').data('progress-percent') / 100);
    var getProgressWrapWidth = $('.progress-wrap').width();
    var progressTotal = getPercent * getProgressWrapWidth;
    var animationLength = 100;

    // on page load, animate percentage bar to data percentage length
    // .stop() used to prevent animation queueing
    $('.progress-bar').stop().animate({
        left: progressTotal
    }, animationLength);
}

function setBarPercent(obj, percent) {
    console.log("Set percent: ", percent);
    $('#passwordStrengthBar').data('progress-percent', percent);
    moveProgressBar();
}
