getCityRes("北京")
function getCityRes(tit) {
    $.ajax({
        type: "POST",
        dataType: "json",
        url: "/campusData",
        data: {"province":tit},
        success: function (result) {
            let _html = "";
            if (Number(result.code) == 200) {
                $.each(result.data,function (k,v) {
                    _html +='<dl><dt><h3>'+v.school_name+'</h3><span class="shcool_line"></span>'
                    _html +='<p class="tel">'+v.school_tel+'</p>'
                    _html +='<p class="work">'+v.worker_time+'</p>'
                    _html +='<p class="address">'+v.address+'</p>'
                    _html +='</dt><dd><img src="/static/upload/'+v.school_img+'"></dd></dl>'
                })
            }
            $('.school_ul').empty().append(_html);
        }
    });
}