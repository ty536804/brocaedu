var swiper = new Swiper('.banner .swiper-container', {
    loop: true,
    autoplay:true,
});
var swiper = new Swiper('.setting_banner .swiper-container', {
    loop: true,
    autoplay:true,
});
getAjax()
//请求数据
function getAjax()
{
    $.ajax({
        type: "GET",
        dataType: "json",
        url: "/wapInfo",
        success: function (result) {
            let _html= "";
            let _dl = "";
            let _dlName = ""
            if (Number(result.code) == 200) {
                //AI智能解决方案
                if (result.data.ai.length > 0) {
                    $.each(result.data.ai, function (k, v) {
                        switch (Number(k)) {
                            case 0:
                                _dlName = 'online';
                                break;
                            case 1:
                                _dlName = 'people';
                                break;
                            case 2:
                                _dlName = 'study';
                                break;
                            case 3:
                                _dlName = 'curr';
                                break;
                            case 4:
                                _dlName = 'hour';
                                break;
                        }
                        _dl += '<dl class="'+_dlName+'"><dt class="small_img"><img src="'+v.imgurl+'"></dt>'
                        _dl += '<dd class="small_con"> <h3 class="small_tit">'+v.bname+'</h3>'
                        _dl += '<p class="small_p">'+v.info+'</p></dd></dl>';
                    })
                    $('.ai_con_ul').empty().append(_dl);
                }
                //品牌介绍
                if (result.data.brandBanner.length > 0) {
                    let _brand = '<dt><img src="'+result.data.brandBanner[0].imgurl+'"></dt><dd>'+result.data.brandBanner[0].info+'</dd>';
                    $('.brand_con dt img').empty().html(_brand)
                }
                //你正在面临的学习问题
                if (result.data.learn.length > 0) {
                    $('.elearn_img img').attr("src",result.data.learn[0].imgurl)
                }
                //选择布罗卡斯的理由
                let _checkDl = "";
                let _check = "";
                if (result.data.reasonBanner.length > 0) {
                    $.each(result.data.reasonBanner, function (k, v) {
                        if (k<2) {
                            _checkDl += '<dl><dt><img src="'+v.imgurl+'"></dt>'
                            _checkDl += '<dd><h3 class="small_tit">'+v.bname+'</h3>'
                            _checkDl += '<p>'+v.info+'</p></dd></dl>';
                        } else {
                            _check += '<dl><dt><img src="'+v.imgurl+'"></dt>'
                            _check += '<dd><h3 class="small_tit">'+v.bname+'</h3>'
                            _check += '<p>'+v.info+'</p></dd></dl>';
                        }
                    })
                    $('.checked_con_first').empty().html(_checkDl);
                    $('.checked_con_s').empty().html(_check);
                }
                //7-15岁少儿英语
                if (result.data.sevenBanner.length > 0) {
                    let _img = "";
                    $.each(result.data.sevenBanner, function (k, v) {
                        _img += '<div class="swiper-slide"><img src="'+v.imgurl+'" alt=""></div>'
                    })
                    $('.english .swiper-wrapper').empty().html(_img);
                }
                //3-6岁幼儿英语
                if (result.data.threeBanner.length > 0) {
                    let _img = "";
                    $.each(result.data.threeBanner, function (k, v) {
                        _img += '<div class="swiper-slide"><img src="'+v.imgurl+'" alt=""></div>'
                    })
                    $('.setting_setting_con .swiper-wrapper').empty().html(_img);
                }
            }
        }
    });
}