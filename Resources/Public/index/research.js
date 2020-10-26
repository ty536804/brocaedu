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
                $.each(result.data.banner,function (k,v) {
                    _div += '<div class="carousel-item active"><img src="'+v.imgurl+'"></div>'
                })
                $('.carousel-inner').empty().html(_div)
            }
        }
    });
}