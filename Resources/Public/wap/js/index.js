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
            if (Number(result.code) == 200) {
               console.log(result)
            }
        }
    });
}