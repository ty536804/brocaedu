getAjax()
//请求数据
function getAjax()
{
    $.ajax({
        type: "GET",
        dataType: "json",
        url: "/leInfo",
        success: function (result) {
            let _dl = "";
            let _dlName = ""
            if (Number(result.code) == 200) {
                //AI智能解决方案
                if (result.data.selected.length > 0) {
                    $.each(result.data.selected, function (k, v) {
                        switch (Number(k)) {
                            case 0:
                                _dlName = 'zy';
                                break;
                            case 1:
                                _dlName = 'abc';
                                break;
                            case 2:
                                _dlName = 'sy';
                                break;
                            case 3:
                                _dlName = 'qh';
                                break;
                        }
                        _dl += '<dl class="'+_dlName+'"><dt class="dlL"><img src="'+v.imgurl+'"><span>'+v.bname+'</span></dt><dd class="dlR">'+v.info+'</dd></dl>'
                    })
                    $('.selecte .system_list').after(_dl)

                    let _checkImg = "";
                    _dl = "";
                    $.each(result.data.checkAll, function (k, v) {
                        switch (Number(k)) {
                            case 0:
                                _checkImg = "/static/wap/img/ico_subscript_l@2x.png";
                                _dl += ' <section class="join"><dl><dt class="con"><p class="titImg"><img src="'+_checkImg+'" class="smallIcon"></p><p class="summary">'+v.summary+'</p><p class="content">'+v.content+'</p></dt><dd class="thumbBox"><img src="'+v.thumb_img+'" class="thumb_img"></dd></dl></section>'
                                break;
                            case 1:
                                _checkImg = "/static/wap/img/ico_subscript_r@2x.png";
                                _dl += ' <section class="level"><dl><dt class="thumbBox"><img src="'+v.thumb_img+'" class="thumb_img"></dt><dd class="con"><p class="titImg"><img src="'+_checkImg+'" class="smallIcon"></p><p class="summary">'+v.summary+'</p><p class="content">'+v.content+'</p></dd></dl></section>'
                                break;
                        }

                    })
                    $('.check .system_list').after(_dl)
                    _dl = "";
                    let _au = "";
                    $.each(result.data.appList, function (k, v) {

                        _au += '<dl><dt class="auxiliaryThumb"><img src="'+v.imgurl+'"></dt><dd><p class="bname">'+v.bname+'</p><p class="bname">'+v.info+'</p></dd></dl>'

                    })
                    _dl = '<section class="auxiliary_box">'+_au+'</section>';
                    $('.auxiliary .system_list').after(_dl)
                }

            }
        }
    });
}
new Swiper('.brand_ul .swiper-container',{
    autoplay:true,
    speed:1000,
    autoplayDisableOnInteraction : false,
    loop:true,
    centeredSlides : true,
    slidesPerView:2,
    pagination : '.swiper-pagination',
    paginationClickable:true,
    prevButton:'.swiper-button-prev',
    nextButton:'.swiper-button-next',
    onInit:function(swiper){
        swiper.slides[2].className="swiper-slide swiper-slide-active";
    },
    breakpoints: {
        668: {
            slidesPerView: 1,
        }
    }
});

isAndroid();
function isAndroid() {
    var u = navigator.userAgent;

    if (u.indexOf("Android") > -1 || u.indexOf("Linux") > -1) {

        $('.brand_ul .swiper-slide').css({"transform":"scale(0.9)"})
    }
}