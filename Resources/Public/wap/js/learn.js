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
                                _dl += ' <section class="join"><dl><dt class="con"><img src="'+_checkImg+'" class="smallIcon"><p class="summary">'+v.summary+'</p><p class="content">'+v.content+'</p></dt><dd class="thumbBox"><img src="'+v.thumb_img+'" class="thumb_img"></dd></dl></section>'
                                break;
                            case 1:
                                _checkImg = "/static/wap/img/ico_subscript_r@2x.png";
                                _dl += ' <section class="level"><dl><dt class="thumbBox"><img src="'+v.thumb_img+'" class="thumb_img"><</dt><dd class="con"><img src="'+_checkImg+'" class="smallIcon"><p class="summary">'+v.summary+'</p><p class="content">'+v.content+'</p>/dd></dl></section>'
                                break;
                        }

                    })
                    $('.check .system_list').after(_dl)
                }

            }
        }
    });
}