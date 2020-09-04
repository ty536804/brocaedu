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
                   let _li = "";
                   // banner
                    $.each(result.data.banner,function (k,v) {
                        _div += '<div class="carousel-item '+(k==0? 'active' : '')+'"><img src="'+v.imgurl+'"></div>'
                        _li +='<li data-target="#myCarousel" data-slide-to="'+k+'" class="'+(k==0? 'active' : '')+'"></li>'
                    })
                    $('.carousel-inner').empty().html(_div)
                    $('.carousel-indicators').empty().html(_li)
                    $('.elearn_img').empty().html('<img src="'+result.data.learn.imgurl+'" class="lazy">');
                    //智能学习，轻松提分
                    let _dl = "";
                    $.each(result.data.plan,function (k,v) {
                        _dl += '<dl><dt><img src="'+v.imgurl+'" class="lazy"></div></dt><dd><h3>'+v.bname+'</h3>'
                        _dl += '<p>'+v.info+'</p></dd></dl>';
                    })
                    $('.plan_con').empty().html(_dl)
                    $('.ai_con').empty().html('<img src="'+result.data.ai.imgurl+'" class="lazy">');
                    let _three = "";
                    _three += '<section class="top_desc"><h3>'+result.data.small.name+'</h3><h4>'+result.data.small.summary+'</h4><p></p></section>'
                    _three += '<section class="bottom_desc">'+result.data.small.content+'</section>';
                    $('.three').empty().html(_three);
                    $('.six_img').empty().html('<img src="'+result.data.small.thumb_img+'" class="lazy">')

                    let _seven = "";
                    _seven += '<section class="top_desc"><h3>'+result.data.seven.name+'</h3><h4>'+result.data.seven.summary+'</h4><p></p></section>'
                    _seven += '<section class="bottom_desc">'+result.data.seven.content+'</section>';
                    $('.seven').empty().html(_seven);
                    $('.seven_img').empty().html('<img src="'+result.data.seven.thumb_img+'" class="lazy">')
                    $('.brand_con').empty().html('<dt><img src="'+result.data.brand.thumb_img+'"/></dt><dd>'+result.data.brand.content+'</dd>')
                    let _piclist = ""
                    $.each(result.data.reason,function (k,v) {
                        _piclist += '<dl><dt><img src="'+v.imgurl+'" class="lazy"></dt><dd><h6>'+v.bname+'</h6><p>'+v.info+'</p></dd></dl>'
                    })
                    $('.pic_reason_con_list').empty().html(_piclist)

                    let _onlinelist = ""
                    $.each(result.data.online,function (k,v) {
                        _onlinelist += '<dl class="'+(k==1 ? 'mid' : '')+'"><dt style="'+(k==1? 'display:none' : '')+'">'+v.bname+'</dt><dd><img src="'+v.imgurl+'"><p  style="'+(k==1? 'display:none' : '')+'">'+v.info+'</p></dd></dl>'
                    })
                    $('.systemList').empty().html(_onlinelist)

                    if (result.data.list.length >=1) {
                        let _dl = "";
                        $.each(result.data.list, function (k, v) {
                            if (k<4) {
                                _dl += "<dl><a href='/detail?id="+v.id+"'><dt><img src='" + v.thumb_img + "'></dt><dd><h5>" + v.title + "</h5><p>" + v.summary + "</p></dd></a></dl>"
                            }
                        })
                        $('.footer_warp').css("marginTop","-130px")
                        $('.news').css("display","block");
                        $('.six_reason').append(_dl);
                    }
                }
            }
        });
    }
})