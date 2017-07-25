/**
 * Created by cmm on 2017/7/24.
 */

function updateGraphSHT(data, data2, name, div) {
    // 基于准备好的dom，初始化echarts实例
    var myChart = echarts.init(document.getElementById(div));

    var dateList = data.map(function (item) {
        return item[0];
    });

    var valueList = data.map(function (item) {
        return item[1];
    });

    var valueList2 = data2.map(function (item) {
        return item[1];
    });

    // 指定图表的配置项和数据
    var option = {
        // Make gradient line here
        visualMap: [{
            show: false,
            type: 'continuous',
            seriesIndex: 0,
            min: 0,
            max: 400,
            color: ['#71c73e']
        }, {
            show: false,
            type: 'continuous',
            seriesIndex: 1,
            dimension: 0,
            min: 0,
            color: ['#77b7c5']
        }],

        title: [{
            left: 'center',
            text: name
        }],
        tooltip: {
            trigger: 'axis',
            formatter: function (params) {
                return params[0].seriesName + ": " + params[0].data + "&#8451<br/>" +
                    params[1].seriesName + ": " + params[1].data + "%"
            }
        },
        xAxis: [{
            data: dateList
        }],
        yAxis: [{
            splitLine: {show: false}
        }],
        grid: [{
            bottom: '50px'
        }, {
            top: '50px'
        }],
        series: [{
            name: '温度',
            type: 'line',
            showSymbol: false,
            data: valueList
        }, {
            name: '湿度',
            type: 'line',
            showSymbol: false,
            data: valueList2
        }]
    };

    // 使用刚指定的配置项和数据显示图表。
    myChart.setOption(option);
}