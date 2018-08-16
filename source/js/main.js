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

$(".service_li").on('click', function() {
    var id = $(this).attr('data-id');
    var position = $("#service_id_"+id).offset();
    window.scroll(0,position.top-23);
    return false;
});


$('form').submit(function() {
    $(this).find("button[type='submit']").prop('disabled',true);
});

$('select#service_type').on('change', function() {
    var selected = $('#service_type option:selected').val();
    if (selected == "tcp") {
        $("#service_port").parent().parent().removeClass("d-none");
        $("#service_check_type").parent().parent().addClass("d-none");
        $("#service_url").attr("placeholder", "localhost");

        $("#post_data").parent().parent().addClass("d-none");
        $("#service_response").parent().parent().addClass("d-none");
        $("#service_response_code").parent().parent().addClass("d-none");
    } else {
        $("#post_data").parent().parent().removeClass("d-none");
        $("#service_response").parent().parent().removeClass("d-none");
        $("#service_response_code").parent().parent().removeClass("d-none");
        $("#service_check_type").parent().parent().removeClass("d-none");
        $("#service_url").attr("placeholder", "https://google.com");

        $("#service_port").parent().parent().addClass("d-none");
    }

});

$('select#service_check_type').on('change', function() {
    var selected = $('#service_check_type option:selected').val();
    if (selected == "POST") {
        $("#post_data").parent().parent().removeClass("d-none");
    } else {
        $("#post_data").parent().parent().addClass("d-none");
    }
});


$(function() {
    var pathname = window.location.pathname;
    if (pathname=="/logs") {
        var lastline;
        setInterval(function() {
            var xhr = new XMLHttpRequest();
            xhr.open('GET', '/logs/line');
            xhr.onload = function () {
                if (xhr.status === 200) {
                    if (lastline != xhr.responseText) {
                        var curr = $.trim($("#live_logs").text());
                        var line = xhr.responseText.replace(/(\r\n|\n|\r)/gm," ");
                        line = line+"\n";
                        $("#live_logs").text(line+curr);
                        lastline = xhr.responseText;
                    }
                }
            };
            xhr.send();
        }, 200);
    }
});


$(".confirm-btn").on('click', function() {
    var r = confirm("Are you sure you want to delete?");
    if (r == true) {
        return true;
    } else {
        return false;
    }
});


$(".select-input").on("click", function () {
    $(this).select();
});


var ranVar = false;
var ranTheme = false;
$('a[data-toggle="pill"]').on('shown.bs.tab', function (e) {
    var target = $(e.target).attr("href");
    if (target=="#v-pills-style" && !ranVar) {
        var sass_vars = CodeMirror.fromTextArea(document.getElementById("sass_vars"), {
            lineNumbers: true,
            matchBrackets: true,
            mode: "text/x-scss",
            colorpicker : true
        });
        ranVar = true;
    } else if (target=="#pills-theme" && !ranTheme) {
        var theme_css = CodeMirror.fromTextArea(document.getElementById("theme_css"), {
            lineNumbers: true,
            matchBrackets: true,
            mode: "text/x-scss",
            colorpicker : true
        });
        ranTheme = true;
    }
});