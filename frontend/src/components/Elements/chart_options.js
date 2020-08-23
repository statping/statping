

const serviceSparkLine = {
    chart: {
      type: 'bar',
      height: 50,
      sparkline: {
        enabled: true
      },
    },
    stroke: {
      curve: 'straight'
    },
    fill: {
      opacity: 0.3,
    },
    yaxis: {
      min: 0
    },
    colors: ['#b3bdc3'],
    tooltip: {
      theme: false,
      enabled: false,
    },
    title: {
      text: this.title,
      offsetX: 0,
      style: {
        fontSize: '28px',
        cssClass: 'apexcharts-yaxis-title'
      }
    },
    subtitle: {
      text: this.subtitle,
      offsetX: 0,
      style: {
        fontSize: '14px',
        cssClass: 'apexcharts-yaxis-title'
      }
    }
}


export default {
  ServiceList: serviceSparkLine
}
