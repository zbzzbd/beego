function initMenu () {
    var curUrl = window.location.pathname;
    if (curUrl.indexOf('?') > -1) {
        curUrl = curUrl.substr(0, curUrl.indexOf('?'))
    }
    $(".main-menu .ui.menu .item").each(function(){
        if (curUrl === '/') {
            if ($(this).attr('href') === '/') $(this).addClass("active");
        }
        else if (curUrl.indexOf($(this).attr('href')) === 0 && $(this).attr('href') !== '/') {
            $(this).addClass("active");
        }
        else if (curUrl.indexOf($(this).attr('href-tag')) === 0 ) {
            $(this).addClass("active");
        }
        if (curUrl.indexOf("/project/edit")===0 && $(this).attr('href') === '/project/list'){
            $(this).addClass("active");
        }
        if (curUrl.indexOf("/job/view")===0 && $(this).attr('href') === '/job/progress'){
            $(this).addClass("active");
        }
    });
}

(function($){
    initMenu();
})(jQuery);
