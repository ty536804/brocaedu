window.onload = function () {
    //返回上一页
    $('.left').on('click',function () {
        window.history.back();
    })
    //监听表单是否输入内容
    $('.searchInput').on('input',function () {
        let keyword = $('.searchInput').val().trim();
        if (keyword == "") {
            $('.close').hide();
        } else {
            $('.close').show();
        }
    });
    //清空表单
    $('.close').on('click',function () {
        $('.searchInput').val("");
        $('.close').hide();
    });
    //请求数据
    $('.searchBtn').on('click',function () {
        let keyword = $('.searchInput').val().trim();
        if (keyword == "") {
            layer.open({
                content: '输入小区名称'
                ,skin: 'msg'
                ,time: 2 //2绉掑悗鑷姩鍏抽棴
            });
            return;
        }
        $.get("https://www.fangpaiwang.com/api/estate/estate_list?keyword="+keyword,function (res) {
            if (res.lists.data.length >=1) {
                let _areaHouse = "";
                $.each(res.lists.data,function (k,con) {
                    _areaHouse += '<a class="search_item" href="/map?lng=111&lat=222">小区：'+con.title+'</a>'
                })
                $('.houseList').empty().html(_areaHouse)
            }
        })
    })
}