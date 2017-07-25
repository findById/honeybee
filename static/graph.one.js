/**
 * Created by cmm on 2017/7/24.
 */

function updateTemplateOne(data, title, divId) {
    // 基于准备好的dom，初始化echarts实例
    var myChart = echarts.init(document.getElementById(divId));

    var dateList = data.map(function (item) {
        return item[0];
    });

    var valueList = data.map(function (item) {
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
        }],

        title: [{
            left: 'center',
            text: title
        }],
        tooltip: {
            trigger: 'axis'
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
            // name: '温度',
            type: 'line',
            showSymbol: false,
            data: valueList
        }]
    };

    // 使用刚指定的配置项和数据显示图表。
    myChart.setOption(option);
}