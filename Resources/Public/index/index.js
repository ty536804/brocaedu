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
                let _dl = "";
                if (Number(result.code) == 200) {


                }
            }
        });
    }
})