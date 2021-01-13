$(function() {
    $("img.lazy").lazyload({
        effect : "fadeIn"
    });
});

$('.c-com').val(window.location.href);
$('.home_nav').css({"height":$(document).height()+"px"});
$('.home_icon').on('click',function () {
    let flag = $(".home_nav").is(":hidden");
    if(flag){
        $('.navicon').attr("src","/static/wap/img/clone.png");
        $(".home_nav").show();
    }else{
        $('.navicon').attr("src","/static/wap/img/home_icon_breadnav_nor.png");
        $(".home_nav").hide();
    }
})
let _currentUrl = 'http://'+window.location.host
let currentUrl = window.location.href.substr(_currentUrl.length,window.location.href.length)
$('.home_nav a').each(function () {
    let _actClass = $(this).attr('href');
    if (_actClass == currentUrl) {
        $(this).addClass('nav_active').siblings().removeClass('nav_active')
    }
})

$('.c-com').val(window.location.href);

$('.bottom_btn').on('click',function () {
    var reg =/[^u4e00-u9fa5]/
    let cname = $.trim($('.footer_right .c-name').val())
    if (cname=="" || !reg.test(cname)) {
        alert('姓名不能为空');
        return false;
    }
    var pattern = /^1\d{10}$/;
    let phone = $.trim($('.footer_right .c-tel').val())
    if (phone=="") {
        alert('电话不能为空');
        return false;
    }
    if (!pattern.test(phone)) {
        alert('手机号码格式不正确');
        return false;
    }
    let carea = $.trim($('.footer_right .c-area').val())
    if (carea=="" || !reg.test(carea)) {
        alert('地区不能为空');
        return false;
    }

    $.ajax({
        type: "POST",
        dataType: "json",
        url: "/AddMessage",
        data:$('#myform').serialize(),
        success: function (result) {
            alert(result.msg);
            return false
        }
    })
    return false;
})
new Swiper('.banner .swiper-container', {
    speed:1000,
    autoplayDisableOnInteraction : false,
    centeredSlides : true,
});
$('.banner .swiper-slide').css({"transform":"scale(1)"})