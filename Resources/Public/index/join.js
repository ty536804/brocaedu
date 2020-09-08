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
        data:$('#top_form').serialize(),
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

var $slider = $('.slider ul');
var $slider_child_l = $('.slider ul li').length;
var $slider_width = $('.slider ul li').width();
var speed =4;
$slider.width($slider_child_l * $slider_width);

var slider_count = 0;

if ($slider_child_l < 4) {
    $('#btn-right').css({cursor: 'auto'});
}

$('#btn-right').click(function() {
    if ($slider_child_l < 4 || slider_count >= $slider_child_l - 4) {
        return false;
    }
    console.log(slider_count);
    slider_count++;
    $slider.animate({left: '-=' + $slider_width + 'px'}, 'slow');
    slider_pic();
});

$('#btn-left').click(function() {
    if (slider_count <= 0) {
        return false;
    }
    slider_count--;
    $slider.animate({left: '+=' + $slider_width + 'px'}, 'slow');
    slider_pic();
});

function slider_pic() {
    if (slider_count >= $slider_child_l - 4) {
        $('#btn-right').css({cursor: 'auto'});
    }
    else if (slider_count > 0 && slider_count <= $slider_child_l - 4) {
        $('#btn-left').css({cursor: 'pointer'});
        $('#btn-right').css({cursor: 'pointer'});
    }
    else if (slider_count <= 0) {
        $('#btn-left').css({cursor: 'auto'});
    }
}