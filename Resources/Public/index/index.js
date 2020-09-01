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
                    console.log(result.data);
                   let  _div = "";
                   let _li = "";
                   // banner
                    $.each(result.data.banner,function (k,v) {
                        _div += '<div class="carousel-item '+(k==0? 'active' : '')+'"><img src="'+v.imgurl+'"></div>'
                        _li +='<li data-target="#myCarousel" data-slide-to="'+k+'" class="'+(k==0? 'active' : '')+'"></li>'
                    })
                    $('.carousel-inner').empty().html(_div)
                    $('.carousel-indicators').empty().html(_li)
                    $('.elearn_img').empty().html('<img src="'+result.data.learn.imgurl+'">');
                    //智能学习，轻松提分
                    let _dl = "";
                    $.each(result.data.plan,function (k,v) {
                        console.log(result.data.plan)
                        _dl += '<dl><dt><img src="'+v.imgurl+'"></div></dt><dd><h3>'+v.bname+'</h3>'
                        _dl += '<p>'+v.info+'</p></dd></dl>';
                    })
                    $('.plan_con').empty().html(_dl)
                    $('.ai_con').empty().html('<img src="'+result.data.ai.imgurl+'">');
                }
            }
        });
    }
})