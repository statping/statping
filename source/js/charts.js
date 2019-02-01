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
			name: "Response Time",
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
};

const startOn = Math.floor(Date.now() / 1000) - (86400 * 14);
const endOn = Math.floor(Date.now() / 1000);


async function RenderCharts() {
	{{ range .Services }}
	let chart{{.Id}} = new ApexCharts(document.querySelector("#service_{{js .Id}}"), options);
	{{end}}

{{ range .Services }}
await RenderChart(chart{{js .Id}}, {{js .Id}}, startOn, endOn);{{end}}
}

$( document ).ready(function() {
	RenderCharts()
});
{{end}}
