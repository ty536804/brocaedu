$('.com').val(window.location.href);
$('.lj_btn').on('click',function () {
    if ($('.banner_form .c-area').val()=="") {
        layer.tips('姓名不能为空!', '.banner_form .c-area', {
            tips: [1, '#3595CC'],
            time: 4000
        });
        return false;
    }
    if ($('.banner_form .c-tel').val()=="") {
        layer.tips('电话不能为空', '.banner_form .c-tel', {
            tips: [1, '#3595CC'],
            time: 4000
        });
        return false;
    }
    if ($('.banner_form .c-tel').val().length < 11) {
        layer.tips('手机号码格式不正确', '.banner_form .c-tel', {
            tips: [1, '#3595CC'],
            time: 4000
        });
        return false;
    }

    if ($('.banner_form .c-city').val()=="") {
        layer.tips('地区不能为空┖', '.banner_form .c-city', {
            tips: [1, '#3595CC'],
            time: 4000
        });
        return false;
    }

    $.ajax({
        type: "POST",
        dataType: "json",
        url: "/AddMessage",
        data:$('#banner_form').serialize(),
        success: function (result) {
            layer.alert(result.msg);
            return false
        }
    })
    return false;
})

getAjax()
//请求数据
function getAjax()
{
    $.ajax({
        type: "GET",
        dataType: "json",
        url: "/joinData",
        success: function (result) {

        }
    });
}