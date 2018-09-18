/*
 * Statup
 * Copyright (C) 2018.  Hunter Long and the project contributors
 * Written by Hunter Long <info@socialeck.com> and the project contributors
 *
 * https://github.com/hunterlong/statup
 *
 * The licenses for most software and other practical works are designed
 * to take away your freedom to share and change the works.  By contrast,
 * the GNU General Public License is intended to guarantee your freedom to
 * share and change all versions of a program--to make sure it remains free
 * software for all its users.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

{{$d := .Data}}{{ range $i, $s := .Services }}{{ if $s.AvgTime }}var ctx_{{$s.Id}}=document.getElementById("service_{{$s.Id}}").getContext('2d');var chartdata=new Chart(ctx_{{$s.Id}},{type:'line',data:{datasets:[{label:'Response Time (Milliseconds)',data:{{safe (index $d $i)}},backgroundColor:['rgba(47, 206, 30, 0.92)'],borderColor:['rgb(47, 171, 34)'],borderWidth:1}]},options:{maintainAspectRatio:!1,scaleShowValues:!0,layout:{padding:{left:0,right:0,top:0,bottom:-10}},hover:{animationDuration:0,},responsiveAnimationDuration:0,animation:{duration:3500,onComplete:function(){var chartInstance=this.chart,ctx=chartInstance.ctx;var controller=this.chart.controller;var xAxis=controller.scales['x-axis-0'];var yAxis=controller.scales['y-axis-0'];ctx.font=Chart.helpers.fontString(Chart.defaults.global.defaultFontSize,Chart.defaults.global.defaultFontStyle,Chart.defaults.global.defaultFontFamily);ctx.textAlign='center';ctx.textBaseline='bottom';var numTicks=xAxis.ticks.length;var yOffsetStart=xAxis.width/numTicks;var halfBarWidth=(xAxis.width/(numTicks*2));xAxis.ticks.forEach(function(value,index){var xOffset=20;var yOffset=(yOffsetStart*index)+halfBarWidth;ctx.fillStyle='#e2e2e2';ctx.fillText(value,yOffset,xOffset)});this.data.datasets.forEach(function(dataset,i){var meta=chartInstance.controller.getDatasetMeta(i);var hxH=0;var hyH=0;var hxL=0;var hyL=0;var highestNum=0;var lowestnum=999999999999;meta.data.forEach(function(bar,index){var data=dataset.data[index];if(lowestnum>data.y){lowestnum=data.y;hxL=bar._model.x;hyL=bar._model.y}
    if(data.y>highestNum){highestNum=data.y;hxH=bar._model.x;hyH=bar._model.y}});if(hxH>=820){hxH=820}else if(50>=hxH){hxH=50}
    if(hxL>=820){hxL=820}else if(70>=hxL){hxL=70}
    ctx.fillStyle='#ffa7a2';ctx.fillText(highestNum+"ms",hxH-40,hyH+15);ctx.fillStyle='#45d642';ctx.fillText(lowestnum+"ms",hxL,hyL+10);console.log("done service_id_{{.Id}}")})}},legend:{display:!1},tooltips:{"enabled":!1},scales:{yAxes:[{display:!1,ticks:{fontSize:20,display:!1,beginAtZero:!1},gridLines:{display:!1}}],xAxes:[{type:'time',distribution:'series',autoSkip:!1,gridLines:{display:!1},ticks:{stepSize:1,min:0,fontColor:"white",fontSize:20,display:!1,}}]},elements:{point:{radius:0}}}})
{{ end }}
{{ end }}
