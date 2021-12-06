import React from "react";
import Chart from "react-apexcharts";

const ServiceChart = () => {
  const state = {
    options: {
      chart: {
        type: "area",
        height: 350,
      },
      stroke: { curve: "straight" },
      xaxis: {
        type: "datetime",
        categories: [1991, 1992, 1993, 1994, 1995, 1996, 1997, 1998, 1999],
      },
    },
    series: [
      {
        name: "series-1",
        data: [30, 40, 45, 50, 49, 60, 70, 91],
      },
    ],
  };

  return (
    <div className="app">
      <div className="row">
        <div className="mixed-chart">
          <Chart options={state.options} series={state.series} width="500" />
        </div>
      </div>
    </div>
  );
};

export default ServiceChart;

// import React, { useState, useEffect } from "react";
// import ReactApexChart from "react-apexcharts";
// import API from "../config/API";
// import DateUtils from "../utils/DateUtils";

// const timeoptions = {
//   weekday: "long",
//   year: "numeric",
//   month: "long",
//   day: "numeric",
//   hour: "numeric",
//   minute: "numeric",
// };

// const axisOptions = {
//   labels: {
//     show: false,
//   },
//   crosshairs: {
//     show: true,
//   },
//   lines: {
//     show: false,
//   },
//   tooltip: {
//     enabled: true,
//   },
//   axisTicks: {
//     show: false,
//   },
//   grid: {
//     show: false,
//   },
// };

// const ServiceChart = ({ service, visible, chartTimeframe }) => {
//   const [ready, setReady] = useState(false);
//   const [showing, setShowing] = useState(null);
//   const [data, setData] = useState(null);
//   const [pingData, setPingData] = useState(null);
//   const [series, setSeries] = useState(null);

//   const state = {
//     options: {
//       noData: {
//         text: "Loading...",
//       },
//       chart: {
//         id: "ping-chart",
//         height: "100%",
//         width: "100%",
//         type: "area",
//         animations: {
//           enabled: true,
//           easing: "easeinout",
//           speed: 800,
//           animateGradually: {
//             enabled: false,
//             delay: 400,
//           },
//           dynamicAnimation: {
//             enabled: true,
//             speed: 500,
//           },
//           hover: {
//             animationDuration: 0, // duration of animations when hovering an item
//           },
//           responsiveAnimationDuration: 0,
//         },
//         selection: {
//           enabled: false,
//         },
//         zoom: {
//           enabled: false,
//         },
//         toolbar: {
//           show: false,
//         },
//       },
//       grid: {
//         show: false,
//         padding: {
//           top: 0,
//           right: 0,
//           bottom: 0,
//           left: -10,
//         },
//       },
//       dropShadow: {
//         enabled: false,
//       },
//       xaxis: {
//         type: "datetime",
//         labels: {
//           show: false,
//         },
//         tooltip: {
//           enabled: false,
//         },
//       },
//       yaxis: {
//         labels: {
//           show: false,
//         },
//       },
//       markers: {
//         size: 0,
//         strokeWidth: 0,
//         hover: {
//           size: undefined,
//           sizeOffset: 0,
//         },
//       },
//       tooltip: {
//         theme: false,
//         enabled: true,
//         custom: ({ series, seriesIndex, dataPointIndex, w }) => {
//           let ts = w.globals.seriesX[seriesIndex][dataPointIndex];
//           const dt = new Date(ts).toLocaleDateString("en-us", timeoptions);
//           let val = series[0][dataPointIndex];
//           let pingVal = series[1][dataPointIndex];
//           return `<div class="chartmarker">
// <span>Average Response Time: ${DateUtils.humanTime(val)}/${
//             chartTimeframe.interval
//           }</span>
// <span>Average Ping: ${DateUtils.humanTime(pingVal)}/${
//             chartTimeframe.interval
//           }</span>
// <span>${dt}</span>
// </div>`;
//         },
//         fixed: {
//           enabled: true,
//           position: "topRight",
//           offsetX: -30,
//           offsetY: 0,
//         },
//         x: {
//           show: false,
//         },
//         y: {
//           formatter: (value) => {
//             return value + " %";
//           },
//         },
//       },
//       legend: {
//         show: false,
//       },
//       dataLabels: {
//         enabled: false,
//       },
//       floating: true,
//       axisTicks: {
//         show: false,
//       },
//       axisBorder: {
//         show: false,
//       },
//       fill: {
//         colors: service.online
//           ? ["#3dc82f", "#48d338"]
//           : ["#c60f20", "#dd3545"],
//         opacity: 1,
//         type: "solid",
//       },
//       stroke: {
//         show: false,
//         curve: "smooth",
//         lineCap: "butt",
//         colors: service.online
//           ? ["#38bc2a", "#48d338"]
//           : ["#c60f20", "#dd3545"],
//       },
//     },
//   };

//   const chartHits = async (val) => {
//     setReady(false);
//     const end = DateUtils.endOf("hour", DateUtils.now());
//     const start = DateUtils.beginningOf(
//       "hour",
//       DateUtils.fromUnix(val.start_time)
//     );
//     const hits_data = await API.service_hits(
//       service.id,
//       DateUtils.toUnix(start),
//       DateUtils.toUnix(end),
//       val.interval,
//       false
//     );
//     const ping_data = await API.service_ping(
//       service.id,
//       DateUtils.toUnix(start),
//       DateUtils.toUnix(end),
//       val.interval,
//       false
//     );

//     const series = [
//       { name: "Latency", ...DateUtils.convertToChartData(hits_data) },
//       { name: "Ping", ...DateUtils.convertToChartData(ping_data) },
//     ];
//     setSeries(series);
//     setReady(true);
//   };

//   useEffect(() => {
//     chartHits();
//   }, []);

//   return (
//     <div className="app">
//       <div className="row">
//         <div className="mixed-chart">
//           {ready ? (
//             <ReactApexChart
//               options={state.options}
//               series={series}
//               type="bar"
//               width="500"
//               class="service-chart"
//               width="100%"
//               height="100%"
//               type="area"
//             />
//           ) : (
//             <span>Loading Chart...</span>
//           )}
//         </div>
//       </div>
//     </div>
//   );
// };

// export default ServiceChart;
