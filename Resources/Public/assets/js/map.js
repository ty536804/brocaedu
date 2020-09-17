window.onload = function () {
    $('.close').on('click',function(){
        $('.houseList').css("display","none");
    });
    var map = null;
    // 画圈完成
    var isDrawingOk = false;
    // 所有行政边界对象
    var plyAll = {};
    //街道下面归属的小区
    var estate = [];
    // 第三级小区数据marker数组
    var thirdlyMkr = [];
    // 是否处于画圈状态下
    var isInDrawing = false;
    // 是否处于鼠标左键按下状态下
    var isMouseDown = false;
    // 存储画出折线点的数组
    var polyPointArray = [];
    // 上次操作画出的折线
    var lastPolyLine = null;
    // 画圈完成后生成的多边形
    var polygonAfterDraw = null;
    var drawBtn = document.getElementById("draw");
    var exitBtn = document.getElementById("exit");
    // 三级数据
    let thirdlyData ={}
    // 二级数据
    let secondData = {}
    var area = {
        "东城区":{
            "lat":"39.93482727239599",
            "lng":"116.4224009776628",
        },
        "西城区":{
            "lat":"39.91812360584148",
            "lng":"116.37251358116619",
        },
        "朝阳区":{
            "lat":"39.926374523079886",
            "lng":"116.44955872950158",
        },
        "丰台区":{
            "lat":"39.8649371975573",
            "lng":"116.29240188731139",
        },
        "石景山区":{
            "lat":"39.911353808778294",
            "lng":"116.22961266775826",
        },
        "海淀区":{
            "lat":"39.96548984110075",
            "lng":"116.3054340544974",
        },
        "房山区":{
            "lat":"39.75432583977336",
            "lng":"116.14944375184247",
        },
        "通州区":{
            "lat":"39.916017122432365",
            "lng":"116.66341535785384",
        },
        "顺义区":{
            "lat":"40.13635076223076",
            "lng":"116.66142426369096",
        },
        "昌平区":{
            "lat":"40.22641337159427",
            "lng":"116.23761791731043",
        },
        "大兴区":{
            "lat":"39.73255523655448",
            "lng":"116.348625212231",
        },
        "平谷区":{
            "lat":"40.146950735799116",
            "lng":"117.12737910459967",
        },
        "密云区":{
            "lat":"40.38217565813752",
            "lng":"116.84954704426833",
        },
        "延庆区":{
            "lat":"40.46216897375426",
            "lng":"115.98163156901515",
        },
        "门头沟区":{
            "lat":"39.94614672003409",
            "lng":"116.10760355576534",
        }
    }

    initMap()

    getRes();

//初始化画圈找房
    drawing();
    /***
     * 初次加载地图，获取当前城市，各大版块的总数据
     */
    function getRes(lng= "",lat="",hierarchy="",isShow=true) {
        let _url = "https://www.fangpaiwang.com/api/second/areaHouse";
        if ( lng !="" ) {
            _url += "?lng="+ lng+"&lat="+lat
        }
        if ( hierarchy !="" ) {
            _url += "&hierarchy="+ hierarchy
        }
        $.get(_url,function(res) {
            if (hierarchy==3) {
                thirdlyData = res.data;
                addLable(res.data,isShow)
            } else {
                if (hierarchy==2) {
                    estate = res.data.street.estate;
                    secondData = res.data.street;
                    addMarker(res.data.street, isShow);
                } else {
                    addMarker(res.data, isShow);
                }
            }
        });
    }

    /***
     * 初始化地图
     */
    function initMap() {
        map = new BMap.Map("allMap", {
            enableMapClick: false,
            minZoom: 12
        });
        map.centerAndZoom(new BMap.Point(116.403694, 39.916042), 11);
        map.enableScrollWheelZoom(true);

        map.addEventListener("zoomend", function() {
            var zoomLevel = map.getZoom();
            if(zoomLevel <= 13) {
                getRes();
            } else if(zoomLevel > 13 && zoomLevel <= 15) {
                addMarker(secondData, false);
            } else if(zoomLevel > 15) {
                addLable(thirdlyData)
            }
        });
        // 监听地图移动,根据视野动态加载
        map.addEventListener("moveend", function() {
            var zoomLevel = map.getZoom(); // 获取地图缩放级别
            if(zoomLevel > 15) {
                addLable(thirdlyData)
            }
        });
    }


    /**
     * 绑定按钮事件
     */
    function drawing() {
        drawBtn.addEventListener('click', function(e) {
            var zoomLevel = map.getZoom();
            if(zoomLevel <= 13) {
                alert("请放大地图后使用画圈找房");
                return;
            }
            if (zoomLevel==14) {
                map.centerAndZoom(map.dc.lat,map.dc.lng, 14);
                addMarker(secondData, false);
            }
            $("#draw").css("display","none");
            $("#exit").css("display","block");
            // 禁止地图移动点击等操作
            map.clearOverlays()
            map.disableDragging();//禁止拖拽
            map.disablePinchToZoom()//禁止缩放
            map.disableDoubleClickZoom();//禁用双击放大
            // 设置鼠标样式
            map.setDefaultCursor('crosshair');
            // 设置标志位进入画圈状态
            isInDrawing = true;
        });

        // 退出画圈找房
        exitBtn.addEventListener('click', function(e) {
            map.clearOverlays();
            // 恢复地图移动点击等操作
            map.enableDragging();
            map.enablePinchToZoom()
            map.enableDoubleClickZoom();
            map.setDefaultCursor('default');
            $("#draw").css("display","block");
            $("#exit").css("display","none");
            addLable(thirdlyData)
            // 设置标志位退出画圈状态
            isInDrawing = false;
            isDrawingOk = false;
        })
        // 为地图绑定鼠标按下事件(开始画圈)
        map.addEventListener('touchstart', function(e) {
            // 如果处于画圈状态下,清空上次画圈的数据结构,设置isMouseDown进入画圈鼠标按下状态
            if(isInDrawing && !isDrawingOk) {
                // 清空地图上画的折线和圈
                map.removeOverlay(polygonAfterDraw);
                map.removeOverlay(lastPolyLine);
                polyPointArray = [];
                lastPolyLine = null;
                isMouseDown = true;
            }
        });
        // 为地图绑定鼠标抬起事件(画圈完成)
        map.addEventListener('touchend', function(e) {
            // 如果处于画圈状态下 且 鼠标是按下状态
            if(isInDrawing && isMouseDown) {
                // isMouseDown = false;
                // 添加多边形覆盖物,设置为禁止点击
                var polygon = new window.BMap.Polygon(polyPointArray, {
                    strokeColor: '#46ACFF',
                    strokeOpacity: 1,
                    fillColor: '#46ACFF',
                    fillOpacity: 0.3,
                    enableClicking: false
                });
                map.addOverlay(polygon);
                isDrawingOk = true;
                //包含情况
                show(polygon);
            }
        });
        // 为地图绑定鼠标移动事件(触发画图)
        map.addEventListener('touchmove', function(e) {
            // 如果处于鼠标按下状态,才能进行画操作
            if(isMouseDown && !isDrawingOk) {
                // 将鼠标移动过程中采集到的路径点加入数组保存
                polyPointArray.push(e.point);
                // 除去上次的画线
                if(lastPolyLine) {
                    map.removeOverlay(lastPolyLine)
                }
                // 根据已有的路径数组构建画出的折线
                var polylineOverlay = new window.BMap.Polyline(polyPointArray, {
                    strokeColor: '#46ACFF',
                    strokeOpacity: 1,
                    enableClicking: false
                });
                // 添加新的画线到地图上
                map.addOverlay(polylineOverlay);
                // 更新上次画线条
                lastPolyLine = polylineOverlay
            }
        })
    }
    /**
     * 根据画的圈，显示相应的marker
     * @param {Object} polygon
     */
    function show(polygon) {
        if (!isDrawingOk) {
            // 得到多边形的点数组
            var pointArray = polygon.getPath();
            // 获取多边形的外包矩形
            var bound = polygon.getBounds();
            // 在多边形内的点的数组
            var pointInPolygonArray = [];
            // 计算每个点是否包含在该多边形内
            for(var i = 0; i < thirdlyMkr.length; i++) {

                var markerPoint = thirdlyMkr[i].getPosition();
                if(isPointInPolygon(markerPoint, bound, pointArray)) {
                    map.addOverlay(thirdlyMkr[i])
                }
            }
        }
    }

    /**
     * 根据行政区划绘制聚合点位
     * @param {Object} data 数据
     * @param {Object} flag 判断是一级点位还是二级,因为一级目前有行政边界，二级目前没有.
     *  此处如果二级也需要和链家完全一致，想要行政边界，那么就去链家爬取即可
     */
    function addMarker(data, flag) {
        map.clearOverlays();
        $.each(data, function(index, data) {
            if(flag) {
                // 绘画行政边界
                getBoundary(data.name)
            }
            var point = new BMap.Point(data.lng, data.lat);
            var tpl = '<div class="bubble bubble-1" data-longitude="' + data.lng + '"' +
                ' data-latitude="' + data.lat + '">' +
                '<p class="name" title="' + data.name + '">' + data.name + '</p>' +
                '<p class="price" title="' + data.price + '">' + data.price + '\/㎡</p>' +
                '<p class="count"><span>' + data.num + '</span>套</p>' +
                '</div>';
            var myLabel = new BMap.Label(tpl, {
                position: point,
                offset: new BMap.Size(-42, -42)
            });
            myLabel.setStyle({
                width: "90px",
                height: "90px",
                border: "0",
                borderRadius: "90px",
                background: "#D75853",
                opacity: 0.9,
                cursor: "pointer",
                zIndex: 2
            });
            myLabel.setTitle(data.name);
            map.addOverlay(myLabel);

            myLabel.addEventListener("mouseover", function() {
                myLabel.setStyle({
                    background: "#D75853",
                    zIndex: 4
                });
                if(flag) {
                    var regionName = myLabel.getTitle();
                    plyAll[regionName].show();
                }
            });

            myLabel.addEventListener("mouseout", function() {
                myLabel.setStyle({
                    background: "#D75853",
                    zIndex: 2
                });
                if(flag) {
                    var regionName = myLabel.getTitle();
                    plyAll[regionName].hide();
                }
            });
            myLabel.addEventListener("click", function() {
                // map.clearOverlays();
                let point = myLabel.getPosition()
                if(flag) {
                    $('.small_tag').empty().html(small_tag = myLabel.getTitle()+' | 均价'+data.price+'/㎡')
                    getRes(point.lng,point.lat,2,false);
                    map.centerAndZoom(point, 14);
                } else {
                    getRes(point.lng,point.lat,3,false);
                    map.centerAndZoom(point, 16);
                }
            });

        })
    }

    /**
     * 加载第三级小区数据
     * @param {Object} data
     */
    function addLable(data) {
        map.clearOverlays();
        // 首先判断是不是第一次请求
        if(thirdlyMkr.length <= 0) {
            $.each(data, function(index, data) {
                var point = new BMap.Point(data.lng, data.lat,data.title);
                // 自定义label样式
                var tpl = '<div class=" bubble-1 ZLQbubble" data-longitude="' + data.lng + '"' +
                    ' data-latitude="' + data.lat + '">' +
                    '<span class="name" title="' + data.title + '">' + data.title + '</span>&nbsp&nbsp' +
                    '<span class="count"><span>' + data.house_num + '</span>套</span>' +
                    '</div>';
                var myLabel = new BMap.Label(tpl, {
                    position: point,//label 在此处添加点位位置信息
                    offset: new BMap.Size(-42, -42)
                });
                myLabel.setStyle({
                    height: "22px",
                    lineHeight: "22px",
                    border: "0",
                    borderRadius: "20px",
                    background: "#D75853",
                    opacity: 0.9,
                    cursor: "pointer",
                    zIndex: 2
                });
                myLabel.setTitle(data.title);
                // 直接缓存起来
                thirdlyMkr.push(myLabel);
                myLabel.addEventListener("mouseover", function() {
                    myLabel.setStyle({
                        background: "#D75853",
                        zIndex: 4
                    });// 修改覆盖物背景颜色
                });
                myLabel.addEventListener("mouseout", function() {
                    myLabel.setStyle({
                        background: "#D75853",
                        zIndex: 2
                    });// 修改覆盖物背景颜色
                });
                myLabel.addEventListener("click", function() {
                    getSingle(data.id,myLabel.getTitle())
                });
            })
        }
        // 根据视野动态加载
        addViewLabel(thirdlyMkr)
    }

    /**
     * 房源列表
     * @param estate_id int 小区ID
     * @param tit string 区名称
     */
    function getSingle(estate_id,tit) {
        let house = "";
        $('.smallTit').empty().html(tit)
        let _url = "https://www.fangpaiwang.com/api/second/houseList?a=h169h170&estate_id="+estate_id;
        $.get(_url,function(res) {
            if (res.data.lists.data.length > 0 ) {
                let _clName = "";
                $.each(res.data.lists.data,function (k,v) {
                    house+='<a><dl class="houseItemView"><dt class="houseItemImg"><img class="thumb_img" src="https://www.fangpaiwang.com'+v.img+'">'
                    house+='<ul class="tag">'
                    if (Number(v.house_type) != 48) {
                        house+='<li>'+v.jieduan_name+'</li>'
                    }
                    if (v.is_free !="") {
                        house += '<li class="tag_label_2">'+v.自由购+'</li>'
                    }
                    if (Number(v.house_type) ==48) {
                        house += '<li class="tag_label_2">'+v.社会委托+'</li>'
                    }

                    if (isEmpty(v.characteristic_name,v.characteristic_name,v.characteristic_name)!=true) {
                        house += '<li class="tag_label_1">'+v.characteristic_name+'</li>'
                    }
                    switch (Number(v.fcstatus)) {
                        case 169:
                            _clName = "house_status_red";
                            break;
                        case 170:
                            _clName = "house_status_blue";
                            break;
                        default:
                            _clName = "house_status_ash";
                            break;
                    }
                    house+='<p class="house_status"><span class='+_clName+'>'+v.fcstatus_name+'</span></p>'
                    house+='</dt><dd class="houseItem"><h3 class="itemTitle">'+v.title+'</h3>'
                    house+='<p class="itemInfo">'+v.room+'室'+v.living_room+'厅 | '+v.acreage+'㎡ | '+v.orientations_name+' | '+v.types_name+'</p>'
                    house+='<ul class="itemPrice">'
                    if (Number(v.fcstatus)==175){
                        house+='<li>成交价<span class="redPrice">'+v.cjprice+'</span></li>'
                    } else {
                        house+='<li>起拍价<span class="redPrice">'+v.qipai+'</span></li>'
                    }
                    house+='<li>市场价<span class="grayPrice">'+v.price+'</span></li>'
                    house+='</ul><p class="createIime">开拍时间：'+v.kptime.trim()+'</p>'
                })
                $('.house_con').empty().append(house);
                $('.houseList').css("display","block");
            }
        });
    }

    /**
     * 判断是否为空
     */
    function isEmpty(va,v2,v3) {
        if (va == null || v2 == "" || v3 == undefined) {
            return true;
        }
        return false;
    }
    /**
     * 根据行政区划绘制边界
     * @param {Object} regionName
     */
    function getBoundary(regionName) {
        var ply = new BMap.Polygon(area[regionName], {
            strokeWeight: 1,
            strokeColor: "#0A77FB",
            fillColor: "#7EB8FC"
        });
        ply.hide();
        plyAll[regionName] = ply
        map.addOverlay(ply);
    }
    /**
     * 根据地图视野动态加载数据，当数据多时此方法用来提高地图加载性能
     * 本次模拟数据较少，看不出太大效果
     * @param {Object} labels
     */
    function addViewLabel(mkr) {
        map.clearOverlays();
        for(var i = 0; i < mkr.length; i++) {
            var result = isPointInRect(mkr[i].point, map.getBounds());
            if(result == true) {
                map.addOverlay(mkr[i])
            } else {
                map.removeOverlay(mkr[i]);
            }
        }
    }

    /***
     * 判断地图视野包含哪些点
     * @param point
     * @param bounds
     * @returns {boolean}
     */
    function isPointInRect(point, bounds) {
        // 检查类型是否正确
        if(!(point instanceof BMap.Point) ||
            !(bounds instanceof BMap.Bounds)) {
            return false;
        }
        var sw = bounds.getSouthWest();// 西南脚点
        var ne = bounds.getNorthEast();// 东北脚点
        return(point.lng >= sw.lng && point.lng <= ne.lng && point.lat >= sw.lat && point.lat <= ne.lat);
    }

    /***
     * 判定一个点是否包含在多边形内
     * @param point
     * @param bound
     * @param pointArray
     * @returns {boolean}
     */
    function isPointInPolygon(point, bound, pointArray) {
        // 首先判断该点是否在外包矩形内，如果不在直接返回false
        if(!bound.containsPoint(point)) {
            return false;
        }
        // 如果在外包矩形内则进一步判断
        // 该点往右侧发出的射线和矩形边交点的数量,若为奇数则在多边形内，否则在外
        var crossPointNum = 0;
        for(var i = 0; i < pointArray.length; i++) {
            // 获取2个相邻的点
            var p1 = pointArray[i];
            var p2 = pointArray[(i + 1) % pointArray.length];
            // 如果点相等直接返回true
            if((p1.lng === point.lng && p1.lat === point.lat) || (p2.lng === point.lng && p2.lat === point.lat)) {
                return true
            }
            // 如果point在2个点所在直线的下方则continue
            if(point.lat < Math.min(p1.lat, p2.lat)) {
                continue;
            }
            // 如果point在2个点所在直线的上方则continue
            if(point.lat >= Math.max(p1.lat, p2.lat)) {
                continue;
            }
            // 有相交情况:2个点一上一下,计算交点
            // 特殊情况2个点的横坐标相同
            var crossPointLng;
            if(p1.lng === p2.lng) {
                crossPointLng = p1.lng;
            } else {
                // 计算2个点的斜率
                var k = (p2.lat - p1.lat) / (p2.lng - p1.lng);
                // 得出水平射线与这2个点形成的直线的交点的横坐标
                crossPointLng = (point.lat - p1.lat) / k + p1.lng;
            }
            // 如果crossPointLng的值大于point的横坐标则算交点(因为是右侧相交)
            if(crossPointLng > point.lng) {
                crossPointNum++;
            }
        }
        // 如果是奇数个交点则点在多边形内
        return crossPointNum % 2 === 1
    }
}