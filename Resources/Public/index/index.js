$(function () {
    getAjax()
//请求数据
    function getAjax()
    {
        $.ajax({
            type: "GET",
            dataType: "json",
            url: "/index",
            success: function (result) {
                if (Number(result.code) == 200) {
                    let  _div = "";
                    if (result.data.list.length >=1) {
                        $.each(result.data.banner,function (k,v) {
                            _div += '<div class="carousel-item '+(k==0? 'active' : '')+'"><img src="'+v.imgurl+'"></div>'
                        })
                        $('.carousel-inner').empty().html(_div)

                        let _dl = "";
                        $.each(result.data.list, function (k, v) {
                            if (k<4) {
                                _dl += "<dl><a href='/detail?id="+v.id+"'><dt class='weui-desktop-mass-appmsg__thumb' style='background-image: url("+v.thumb_img+")'></dt><dd><h5>" + v.title + "</h5><p>" + v.summary + "</p></dd></a></dl>"
                            }
                        })
                        $('.news').css("display","block");
                        $('.six_reason').append(_dl);
                    }
                }
            }
        });
    }

    $('.join_con dl').on("mouseover",function () {
        let _key =  Number($(this).index())
        $('.join_con dl').eq(_key).css({'transform':'scale(1)','transition':'all 0.7s ease 0.1s'})
        $('.join_con dl').eq(Number(_key) == 1 ? 0 : 1).css({"transform":" scale(0.8)","transition":" all 0.7s ease 0.1s"})
    })

    $('.headBtn').on('click',function () {
        if ($('.header_from .c-area').val()=="") {
            layer.tips('姓名不能为空', '.header_from .c-area', {
                tips: [1, '#3595CC'],
                time: 4000
            });
            return false;
        }
        if ($('.header_from .c-tel').val()=="") {
            layer.tips('电话号码不能为空┖', '.header_from .c-tel', {
                tips: [1, '#3595CC'],
                time: 4000
            });
            return false;
        }
        if ($('.header_from .c-city').val()=="") {
            layer.tips('地区不能为空', '.header_from .c-city', {
                tips: [1, '#3595CC'],
                time: 4000
            });
            return false;
        }

        $.ajax({
            type: "POST",
            dataType: "json",
            url: "/AddMessage",
            data:$('.header_from').serialize(),
            success: function (result) {
                layer.alert(result.msg);
                return false
            }
        })
        return false;
    })
    var wow = new WOW({
        boxClass: 'wow',
        animateClass: 'animated',
        offset: 0,
        mobile: true,
        live: true
    });
    wow.init();
})