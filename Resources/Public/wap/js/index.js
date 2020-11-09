getAjax()
//请求数据
function getAjax()
{
    $.ajax({
        type: "GET",
        dataType: "json",
        url: "/wapInfo",
        success: function (result) {
            if (Number(result.code) == 200) {
                let _html = '';
                $.each(result.data.list,function (key,item) {
                    let _d = item.created_at
                    _html += "<a href='/de?id="+item.id+"'><dl><dt><img src='"+item.thumb_img+"'></dt>" +
                        "<dd><h3>"+item.summary+"</h3><p>"+_d.substring(0,10)+"</p></dd></dl></a>"
                })
                if (_html.length >= 4) {
                    _html += "<a href='/list' class='more'>更多</a>"
                }
                $('.new_ul').empty().append(_html)
            }
        }
    });
}
new Swiper('.banner .swiper-container', {
    loop: true,
    autoplay:true,
});

$('.van-button--info').on('click',function () {
    if ($('.wap_banner_form .c-name').val()=="") {
        alert('姓名不能为空');
        return false;
    }
    if ($('.wap_banner_form .c-tel').val()=="") {
        alert('电话不能为空');
        return false;
    }
    if ($('.wap_banner_form .c-tel').val().length < 11) {
        alert('电话号码格式不正确');
        return false;
    }
    if ($('.wap_banner_form .c-area').val()=="") {
        alert('区域不能为空');
        return false;
    }

    $.ajax({
        type: "POST",
        dataType: "json",
        url: "/AddMessage",
        data:$('#wap_myform').serialize(),
        success: function (result) {
            alert(result.msg);
            return false
        }
    })
    return false;
})