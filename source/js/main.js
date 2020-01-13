/*
 * Statping
 * Copyright (C) 2018.  Hunter Long and the project contributors
 * Written by Hunter Long <info@socialeck.com> and the project contributors
 *
 * https://github.com/hunterlong/statping
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


$('.service_li').on('click', function() {
    var id = $(this).attr('data-id');
    var position = $('#service_id_' + id).offset();
    window.scroll(0, position.top - 23);
    return false;
});

$('.test_notifier').on('click', function(e) {
    var btn = $(this);
    var form = $(this).parents('form:first');
    var values = form.serialize();
    var notifier = form.find('input[name=method]').val();
    var success = $('#'+notifier+'-success');
    var error = $('#'+notifier+'-error');
		Spinner(btn);
    $.ajax({
        url: form.attr("action")+"/test",
        type: 'POST',
        data: values,
        success: function(data) {
          if (data === 'ok') {
              success.removeClass('d-none');
              setTimeout(function() {
                  success.addClass('d-none');
              }, 5000)
          } else {
              error.removeClass('d-none');
              error.html(data);
              setTimeout(function() {
                  error.addClass('d-none');
              }, 8000)
          }
					Spinner(btn, true);
        }
    });
    e.preventDefault();
});

$('.spin_form').on('submit', function() {
	Spinner($(this).find('button[type=submit]'));
});

function Spinner(btn, off = false) {
	btn.prop('disabled', !off);
	if (off) {
		let pastVal = btn.attr("data-past");
		btn.text(pastVal);
		btn.removeAttr("data-past");
	} else {
		let pastVal = btn.text();
		btn.attr("data-past", pastVal);
		btn.html('<i class="fa fa-spinner fa-spin"></i>');
	}
}

function SaveNotifier(data) {
	let button = data.element.find('button[type=submit]');
	button.text('Saved!')
	button.removeClass('btn-primary')
	button.addClass('btn-success')
}

$('.scrollclick').on('click',function(e) {
	let element = $(this).attr("data-id");
	$('html, body').animate({
		scrollTop: $("#"+element).offset().top - 15
	}, 500);
	e.preventDefault();
});

$('.toggle-service').on('click',function(e) {
	let obj = $(this);
	let serviceId = obj.attr("data-id");
	let online = obj.attr("data-online");
	let d = confirm("Do you want to "+(eval(online) ? "stop" : "start")+" checking this service?");
	if (d) {
		$.ajax({
			url: "api/services/" + serviceId + "/running",
			type: 'POST',
			success: function (data) {
				if (online === "true") {
					obj.removeClass("fa-toggle-on text-success");
					obj.addClass("fa-toggle-off text-black-50");
				} else {
					obj.removeClass("fa-toggle-off text-black-50");
					obj.addClass("fa-toggle-on text-success");
				}
				obj.attr("data-online", online !== "true");
			}
		});
	}
});

$('select#service_type').on('change', function() {
    var selected = $('#service_type option:selected').val();
    var typeLabel = $('#service_type_label');
    if (selected === 'tcp' || selected === 'udp') {
        if (selected === 'tcp') {
            typeLabel.html('TCP Port')
        } else {
            typeLabel.html('UDP Port')
        }
        $('#service_port').parent().parent().removeClass('d-none');
        $('#service_check_type').parent().parent().addClass('d-none');
        $('#service_url').attr('placeholder', '192.168.1.1');
        $('#post_data').parent().parent().addClass('d-none');
        $('#service_response').parent().parent().addClass('d-none');
        $('#service_response_code').parent().parent().addClass('d-none');
				$('#headers').parent().parent().addClass('d-none');
    } else if (selected === 'icmp') {
        $('#service_port').parent().parent().removeClass('d-none');
        $('#headers').parent().parent().addClass('d-none');
        $('#service_check_type').parent().parent().addClass('d-none');
        $('#service_url').attr('placeholder', '192.168.1.1');
        $('#post_data').parent().parent().addClass('d-none');
        $('#service_response').parent().parent().addClass('d-none');
        $('#service_response_code').parent().parent().addClass('d-none');
    } else {
        $('#post_data').parent().parent().removeClass('d-none');
        $('#service_response').parent().parent().removeClass('d-none');
        $('#service_response_code').parent().parent().removeClass('d-none');
        $('#service_check_type').parent().parent().removeClass('d-none');
        $('#service_url').attr('placeholder', 'https://google.com');
        $('#service_port').parent().parent().addClass('d-none');
    }
});


async function RenderChart(chart, service, start=0, end=9999999999, group="hour", retry=true) {
		if (!chart.el) {
			return
		}
    let chartData = await ChartLatency(service, start, end, group, retry);
    if (!chartData) {
        chartData = await ChartLatency(service, start, end, "minute", retry);
    }
    chart.render();
    chart.updateSeries([{
        data: chartData || []
    }]);
}


function UTCTime() {
    var now = new Date();
    now = new Date(now.toUTCString());
    return Math.floor(now.getTime() / 1000);
}

function ChartLatency(service, start=0, end=9999999999, group="hour", retry=true) {
    let url = "api/services/" + service + "/data?start=" + start + "&end=" + end + "&group=" + group;
    return new Promise(resolve => {
        $.ajax({
            url: url,
            type: 'GET',
            success: function (data) {
                resolve(data.data);
            }
        });
    });
}


function ChartHeatmap(service) {
    return new Promise(resolve => {
        $.ajax({
            url: "api/services/" + service + "/heatmap",
            type: 'GET',
            success: function (data) {
                resolve(data);
            }
        });
    });
}


function FailureAnnotations(chart, service, start=0, end=9999999999, group="hour", retry=true) {
    const annotationColor = {
        strokeDashArray: 0,
        borderColor: "#d0222d",
        label: {
            show: false,
        }
    };
    var dataArr = [];
    $.ajax({
        url: "api/services/"+service+"/failures?start="+start+"&end="+end+"&group="+group,
        type: 'GET',
        success: function(data) {
            data.forEach(function (d) {
                dataArr.push({x: d.created_at, ...annotationColor})
            });
            chart.addXaxisAnnotation(dataArr);
        }
    });
}


$('input[id=service_name]').on('keyup', function() {
    var url = $(this).val();
    url = url.replace(/[^\w\s]/gi, '').replace(/\s+/g, '-').toLowerCase();
    $('#permalink').val(url);
});

$('input[type=checkbox]').on('change', function() {
	var element = $(this).attr('id');
	$("#"+element+"-value").val(this.checked ? "true" : "false")
});

function PingAjaxChart(chart, service, start=0, end=9999999999, group="hour") {
  $.ajax({
    url: "api/services/"+service+"/ping?start="+start+"&end="+end+"&group="+group,
    type: 'GET',
    success: function(data) {
      chart.data.labels.pop();
      chart.data.datasets.push({
        label: "Ping Time",
        backgroundColor: "#bababa"
      });
      chart.update();
      data.data.forEach(function(d) {
        chart.data.datasets[1].data.push(d);
      });
      chart.update();
    }
  });
}

$('.confirm_btn').on('click', function() {
	let msg = $(this).attr('data-msg');
	var r = confirm(msg);
	if (r !== true) {
		return false;
	}
	return true;
});

$('.ajax_delete').on('click', function() {
	var r = confirm('Are you sure you want to delete?');
	if (r !== true) {
		return false;
	}
	let obj = $(this);
	let id = obj.attr('data-id');
	let element = obj.attr('data-obj');
	let url = obj.attr('href');
	let method = obj.attr('data-method');
	$.ajax({
		url: url,
		type: method,
		data: JSON.stringify({id: id}),
		success: function (data) {
			if (data.status === 'error') {
				alert(data.error)
			} else {
				console.log(data);
				$('#' + element).remove();
			}
		}
	});
	return false
});


$('form.ajax_form').on('submit', function() {
	const form = $(this);
	let values = form.serializeArray();
	let method = form.attr('method');
	let action = form.attr('action');
	let func = form.attr('data-func');
	let redirect = form.attr('data-redirect');
	let button = form.find('button[type=submit]');
	let alerter = form.find('#alerter');
	var arrayData = [];
	let newArr = {};
	Spinner(button);
	values.forEach(function(k, v) {
		if (k.name === "password_confirm" || k.value === "" || k.name === "enabled-option") {
			return
		}
		if (k.value === "on") {
			k.value = (k.value === "on")
		}
		if (k.value === "false" || k.value === "true") {
			k.value = (k.value === "true")
		}
		if($.isNumeric(k.value)){
			if (k.name !== "password") {
				k.value = parseInt(k.value)
			}
		}
		if (k.name === "var1" || k.name === "var2" || k.name === "host" || k.name === "username" || k.name === "password" || k.name === "api_key" || k.name === "api_secret") {
			k.value = k.value.toString()
		}
		newArr[k.name] = k.value;
		arrayData.push(newArr)
	});
	let sendData = JSON.stringify(newArr);
	$.ajax({
		url: action,
		type: method,
		data: sendData,
		success: function (data) {
			setTimeout(function () {
				if (data.status === 'error') {
					let alerter = form.find('#alerter');
					alerter.html(data.error);
					alerter.removeClass("d-none");
					Spinner(button, true);
				} else {
					Spinner(button, true);
					if (func) {
						let fn = window[func];
						if (typeof fn === "function") fn({element: form, form: newArr, data: data});
					}
					if (redirect) {
						window.location.href = redirect;
					}
				}
			}, 1000);
		}
	});
	return false;
});

function CreateService(output) {
	let form = output.form;
	let data = output.data.output;
	let objTbl = `<tr id="service_${data.id}">
                <td><span class="drag_icon d-none d-md-inline"><i class="fas fa-bars"></i></span> ${form.name}</td>
                <td class="d-none d-md-table-cell">${data.online}<span class="badge badge-success">ONLINE</span></td>
                <td class="text-right">
                    <div class="btn-group">
                        <a href="service/${data.id}" class="btn btn-outline-secondary"><i class="fas fa-chart-area"></i> View</a>
                        <a href="api/services/${data.id}" class="ajax_delete btn btn-danger confirm-btn" data-method="DELETE" data-obj="service_${data.id}" data-id="${data.id}"><i class="fas fa-times"></i></a>
                    </div>
                </td>
            </tr>`;
	$('#services_table').append(objTbl);
}

function CreateUser(output) {
	console.log('creating user', output)
	let form = output.form;
	let data = output.data.output;
	let objTbl = `<tr id="user_${data.id}">
                <td>${form.username}</td>
                <td class="text-right">
                    <div class="btn-group">
                        <a href="user/${data.id}" class="btn btn-outline-secondary"><i class="fas fa-user-edit"></i> Edit</a>
                        <a href="api/users/${data.id}" class="ajax_delete btn btn-danger confirm-btn" data-method="DELETE" data-obj="user_${data.id}" data-id="${data.id}"><i class="fas fa-times"></i></a>
                    </div>
                </td>
            </tr>`;
	$('#users_table').append(objTbl);
}

$('select#service_check_type').on('change', function() {
    var selected = $('#service_check_type option:selected').val();
    if (selected === 'POST') {
        $('#post_data').parent().parent().removeClass('d-none');
    } else {
        $('#post_data').parent().parent().addClass('d-none');
    }
});


$(function() {
    var pathname = window.location.pathname;
    if (pathname === '/logs') {
        var lastline;
        var logArea = $('#live_logs');
        setInterval(function() {
            $.get('/logs/line', function(data, status) {
                if (lastline !== data) {
                    var curr = $.trim(logArea.text());
                    var line = data.replace(/(\r\n|\n|\r)/gm, ' ');
                    line = line + '\n';
                    logArea.text(line + curr);
                    lastline = data;
                }
            });
        }, 200);
    }
});


$('.confirm-btn').on('click', function() {
    var r = confirm('Are you sure you want to delete?');
    let obj = $(this);
    let redirect = obj.attr('data-redirect');
    let href = obj.attr('href');
    let method = obj.attr('data-method');
    let data = obj.attr('data-object');
    if (r === true) {
        $.ajax({
            url: href,
            type: method,
            data: data ? data : null,
            success: function (data) {
                console.log("send to url: ", href);
                if (redirect) {
                    window.location.href = redirect;
                }
                return false;
            }
        });
    } else {
        return false;
    }
});


$('.select-input').on('click', function() {
    $(this).select();
});


// $('input[name=password], input[name=password_confirm]').on('change keyup input paste', function() {
//     var password = $('input[name=password]'),
//         repassword = $('input[name=password_confirm]'),
//         both = password.add(repassword).removeClass('is-valid is-invalid');
//
//     var btn = $(this).parents('form:first').find('button[type=submit]');
//     password.addClass(
//         password.val().length > 0 ? 'is-valid' : 'is-invalid'
//     );
//     repassword.addClass(
//         password.val().length > 0 ? 'is-valid' : 'is-invalid'
//     );
//
//     if (password.val() !== repassword.val()) {
//         both.addClass('is-invalid');
//         btn.prop('disabled', true);
//     } else {
//         btn.prop('disabled', false);
//     }
// });


var ranVar = false;
var ranTheme = false;
var ranMobile = false;
$('a[data-toggle=pill]').on('shown.bs.tab', function(e) {
    var target = $(e.target).attr('href');
    if (target === '#v-pills-style' && !ranVar) {
        var sass_vars = CodeMirror.fromTextArea(document.getElementById('sass_vars'), {
            lineNumbers: true,
            matchBrackets: true,
            mode: 'text/x-scss',
            colorpicker: true
        });
        sass_vars.setSize(null, 900);
        ranVar = true;
    } else if (target === '#pills-theme' && !ranTheme) {
        var theme_css = CodeMirror.fromTextArea(document.getElementById('theme_css'), {
            lineNumbers: true,
            matchBrackets: true,
            mode: 'text/x-scss',
            colorpicker: true
        });
        theme_css.setSize(null, 900);
        ranTheme = true;
    } else if (target === '#pills-mobile' && !ranMobile) {
        var mobile_css = CodeMirror.fromTextArea(document.getElementById('mobile_css'), {
            lineNumbers: true,
            matchBrackets: true,
            mode: 'text/x-scss',
            colorpicker: true
        });
        mobile_css.setSize(null, 900);
        ranMobile = true;
    }
});
