{{define "charts"}}
{{$start := .Start}}
{{$end := .End}}

const axisOptions = {
	labels: {
		show: false
	},
	crosshairs: {
		show: false
	},
	lines: {
		show: false
	},
	tooltip: {
		enabled: false
	},
	axisTicks: {
		show: false
	},
	grid: {
		show: false
	},
	marker: {
		show: false
	}
};

	const annotationColor = {
		strokeDashArray: 0,
		borderColor: "#d0222d",
		label: {
			show: false,
		}
	};

	let annotation = {
		annotations: {
			xaxis: [
				{
					// in a datetime series, the x value should be a timestamp, just like it is generated below
					x: new Date("01/29/2019").getTime(),
					...annotationColor
				}]
		}
	};

let options = {
	chart: {
		height: 210,
		width: "100%",
		type: "area",
		animations: {
			enabled: false,
			initialAnimation: {
				enabled: false
			}
		},
		selection: {
			enabled: false
		},
		zoom: {
			enabled: false
		},
		toolbar: {
			show: false
		}
	},
	grid: {
		show: false,
		padding: {
			top: 0,
			right: 0,
			bottom: 0,
			left: 0,
		},
	},
	tooltip: {
		enabled: false,
		marker: {
			show: false,
		},
		x: {
			show: false,
		}
	},
	legend: {
		show: false,
	},
	dataLabels: {
		enabled: false
	},
	floating: true,
	axisTicks: {
		show: false
	},
	axisBorder: {
		show: false
	},
	fill: {
		colors: ["#48d338"],
		opacity: 1,
		type: 'solid'
	},
	stroke: {
		show: true,
		curve: 'smooth',
		lineCap: 'butt',
		colors: ["#3aa82d"],
	},
	series: [
		{
			name: "Series 1",
			data: [
				{
					x: "02-10-2017 GMT",
					y: 34
				},
				{
					x: "02-11-2017 GMT",
					y: 43
				},
				{
					x: "02-12-2017 GMT",
					y: 31
				},
				{
					x: "02-13-2017 GMT",
					y: 43
				},
				{
					x: "02-14-2017 GMT",
					y: 33
				},
				{
					x: "02-15-2017 GMT",
					y: 52
				}
			]
		}
	],
	xaxis: {
		type: "datetime",
		...axisOptions
	},
	yaxis: {
		...axisOptions
	},
	...annotation
};


{{ range .Services }}

let chart{{.Id}} = new ApexCharts(document.querySelector("#service_{{js .Id}}"), options);

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
	{{ range .Services }}AjaxChart(chart{{js .Id}}, {{js .Id}}, 0, 9999999999);
	{{end}}
	});
{{end}}
