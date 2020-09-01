getAjax()
//请求数据
function getAjax()
{
    $.ajax({
        type: "GET",
        dataType: "json",
        url: "/researchData",
        success: function (result) {
            if (Number(result.code) == 200) {
                let  _div = "";
                let _li = "";
                // banner
                $.each(result.data.banner,function (k,v) {
                    _div += '<div class="carousel-item '+(k==0? 'active' : '')+'"><img src="'+v.imgurl+'"></div>'
                    _li +='<li data-target="#myCarousel" data-slide-to="'+k+'" class="'+(k==0? 'active' : '')+'"></li>'
                })
                $('.carousel-inner').empty().html(_div)
                $('.carousel-indicators').empty().html(_li)
                //智能学习，轻松提分
                let _dl = "";
                $.each(result.data.reason,function (k,v) {
                    _dl += '<li><img src="'+v.imgurl+'" class="lazy"/></li>'
                })
                $('.rd_con').empty().html(_dl)

                $('.selectR').empty().html('<img src="'+result.data.teacher.imgurl+'" class="lazy">')
            }
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