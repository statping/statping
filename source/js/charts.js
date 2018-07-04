{{ range . }}{{ if .AvgTime }}var ctx = document.getElementById("service_{{.Id}}").getContext('2d');

var chartdata = new Chart(ctx, {
        type: 'line',
        data: {
            datasets: [{
                label: 'Response Time (Milliseconds)',
                data: {{safe .GraphData}},
    backgroundColor: [
    'rgba(47, 206, 30, 0.92)'
],
    borderColor: [
    'rgb(47, 171, 34)'
],
    borderWidth: 1
}]
},
options: {
    maintainAspectRatio: false,
        scaleShowValues: true,
        layout: {
        padding: {
            left: 0,
                right: 0,
                top: 0,
                bottom: -10
        }
    },
    hover: {
        animationDuration: 0,
    },
    responsiveAnimationDuration: 0,
        animation: {
        duration: 3500,
            onComplete: function() {
            var chartInstance = this.chart,
                ctx = chartInstance.ctx;

            var controller = this.chart.controller;
            var xAxis = controller.scales['x-axis-0'];
            var yAxis = controller.scales['y-axis-0'];

            ctx.font = Chart.helpers.fontString(Chart.defaults.global.defaultFontSize, Chart.defaults.global.defaultFontStyle, Chart.defaults.global.defaultFontFamily);
            ctx.textAlign = 'center';
            ctx.textBaseline = 'bottom';

            var	numTicks = xAxis.ticks.length;
            var yOffsetStart = xAxis.width / numTicks;
            var halfBarWidth = (xAxis.width / (numTicks * 2));

            xAxis.ticks.forEach(function(value, index) {
                var xOffset = 20;
                var yOffset = (yOffsetStart * index) + halfBarWidth;
                ctx.fillStyle = '#e2e2e2';
                ctx.fillText(value, yOffset, xOffset);
            });

            this.data.datasets.forEach(function(dataset, i) {
                var meta = chartInstance.controller.getDatasetMeta(i);
                var hxH = 0;
                var hyH = 0;
                var hxL = 0;
                var hyL = 0;
                var highestNum = 0;
                var lowestnum = 999999999999;
                meta.data.forEach(function(bar, index) {
                    var data = dataset.data[index];

                    if (data.y {{safe "<"}} lowestnum) {
                        lowestnum = data.y;
                        hxL = bar._model.x;
                        hyL = bar._model.y;
                    }

                    if (data.y > highestNum) {
                        highestNum = data.y;
                        hxH = bar._model.x;
                        hyH = bar._model.y;
                    }
                });

                if (hxH {{safe ">"}}= 820) {
                    hxH = 820;
                } else if (hxH {{safe "<"}}= 50) {
                    hxH = 50;
                }

                if (hxL {{safe ">"}}= 820) {
                    hxL = 820;
                } else if (hxL {{safe "<"}}= 70) {
                    hxL = 70;
                }

                ctx.fillStyle = '#ffa7a2';
                ctx.fillText(highestNum+"ms", hxH - 40, hyH + 15);
                ctx.fillStyle = '#45d642';
                ctx.fillText(lowestnum+"ms", hxL, hyL + 10);

                console.log("done service_id_{{.Id}}")

            });
        }
    },
    legend: {
        display: false
    },
    tooltips: {
        "enabled": false
    },
    scales: {
        yAxes: [{
            display: false,
            ticks: {
                fontSize: 20,
                display: false,
                beginAtZero: false
            },
            gridLines: {
                display:false
            }
        }],
            xAxes: [{
            type: 'time',
            distribution: 'series',
            autoSkip: false,
            gridLines: {
                display:false
            },
            ticks: {
                stepSize: 1,
                min: 0,
                fontColor: "white",
                fontSize: 20,
                display: false,
            }
        }]
    },
    elements: {
        point: {
            radius: 0
        }
    }
}
});
{{ end }}
{{ end }}