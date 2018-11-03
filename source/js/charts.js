{{define "charts"}}
{{$start := .Start}}
{{$end := .End}}
{{ range .Services }}
var ctx_{{js .Id}} = document.getElementById("service_{{js .Id}}").getContext('2d');
var chartdata_{{js .Id}} = new Chart(ctx_{{js .Id}}, {
  type: 'line',
    data: {
    datasets: [{
      label: 'Response Time (Milliseconds)',
      data: [],
      backgroundColor: ['{{if .Online}}rgba(47, 206, 30, 0.92){{else}}rgb(221, 53, 69){{end}}'],
      borderColor: ['{{if .Online}}rgb(47, 171, 34){{else}}rgb(183, 32, 47){{end}}'],
      borderWidth: 1
    }]
  },
  options: {
    maintainAspectRatio: false,
      scaleShowValues: false,
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
        onComplete: onChartComplete
    },
    legend: {
      display: false
    },
    tooltips: {
      enabled: false
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
          display: false
        }
      }],
        xAxes: [{
        type: 'time',
        distribution: 'series',
        autoSkip: false,
        time: {
          displayFormats: {
            'hour': 'MMM DD hA'
          },
          source: 'auto'
        },
        gridLines: {
          display: false
        },
        ticks: {
          source: 'auto',
          stepSize: 1,
          min: 0,
          fontColor: "white",
          fontSize: 20,
          display: false
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
{{end}}

function onChartComplete(chart) {
	var chartInstance = chart.chart,
  ctx = chartInstance.ctx;
	var controller = chart.chart.controller;
	var xAxis = controller.scales['x-axis-0'];
	var yAxis = controller.scales['y-axis-0'];
	ctx.font = Chart.helpers.fontString(Chart.defaults.global.defaultFontSize, Chart.defaults.global.defaultFontStyle, Chart.defaults.global.defaultFontFamily);
	ctx.textAlign = 'center';
	ctx.textBaseline = 'bottom';
	var numTicks = xAxis.ticks.length;
	var yOffsetStart = xAxis.width / numTicks;
	var halfBarWidth = (xAxis.width / (numTicks * 2));
	xAxis.ticks.forEach(function(value, index) {
		var xOffset = 20;
		var yOffset = (yOffsetStart * index) + halfBarWidth;
		ctx.fillStyle = '#e2e2e2';
		ctx.fillText(value, yOffset, xOffset)
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
			if (lowestnum > data.y) {
				lowestnum = data.y;
				hxL = bar._model.x;
				hyL = bar._model.y
			}
			if (data.y > highestNum) {
				highestNum = data.y;
				hxH = bar._model.x;
				hyH = bar._model.y
			}
		});
		if (hxH >= 820) {
			hxH = 820
		} else if (50 >= hxH) {
			hxH = 50
		}
		if (hxL >= 820) {
			hxL = 820
		} else if (70 >= hxL) {
			hxL = 70
		}
		ctx.fillStyle = '#ffa7a2';
		ctx.fillText(highestNum + "ms", hxH - 40, hyH + 15);
		ctx.fillStyle = '#45d642';
		ctx.fillText(lowestnum + "ms", hxL, hyL + 10);
	})
}

$( document ).ready(function() {
{{ range .Services }}
AjaxChart(chartdata_{{js .Id}},{{js .Id}},{{$start}},9999999999,"hour");{{end}}
});
{{end}}
