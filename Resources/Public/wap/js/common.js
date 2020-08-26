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
    let _actClass = '/'+$(this).attr('class');
    if (_actClass == currentUrl) {
        $(this).addClass('nav_active').siblings().removeClass('nav_active')
    } else if (currentUrl == '/') {
        $('.home_nav a:eq(0)').addClass('nav_active').siblings().removeClass('nav_active')
    }
})

$('.c-com').val(window.location.href);

$('.bottom_btn').on('click',function () {
    if ($('.footer_right .c-name').val()=="") {
        alert('姓名不能为空');
        return false;
    }
    if ($('.footer_right .c-tel').val()=="") {
        alert('电话不能为空');
        return false;
    }
    if ($('.footer_right .c-tel').val().length < 11) {
        alert('手机号码格式不正确');
        return false;
    }
    if ($('.footer_right .c-area').val()=="") {
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