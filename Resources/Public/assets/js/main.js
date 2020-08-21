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
                let _banner = "";
                let _oli = "";
                if (Number(result.code) == 200) {
                    if (result.data.banner.length >= 1) {
                        $.each(result.data.banner, function (k, v) {
                            _banner += '<div class="carousel-item" ><img src="/static/upload/' + v.imgurl + '"></div>'
                            _oli += '<li data-target="#myCarousel" data-slide-to="' + k + '" class="active"></li>';
                        })
                        $(".carousel-inner").empty().html(_banner);
                        $('.carousel-indicators').empty().html(_oli);
                    }
                }
            }
        });
    }
})