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


$('.service_li').on('click', function() {
    var id = $(this).attr('data-id');
    var position = $('#service_id_' + id).offset();
    window.scroll(0, position.top - 23);
    return false;
});

$('.test_notifier').on('click', function(e) {
    var form = $(this).parents('form:first');
    var values = form.serialize();
    var notifier = form.find('input[name=notifier]').val();
    $.ajax({
        url: form.attr("action")+"/test",
        type: 'POST',
        data: values,
        success: function(data) {
          if (data === 'ok') {
              $('#'+notifier+'-success').removeClass('d-none');
          } else {
              $('#'+notifier+'-error').removeClass('d-none');
              $('#'+notifier+'-error').html(data);
          }
        }
    });
    e.preventDefault();
});

$('form').submit(function() {
    console.log(this);
    $(this).find('button[type=submit]').prop('disabled', true);
});

$('select#service_type').on('change', function() {
    var selected = $('#service_type option:selected').val();
    if (selected === 'tcp') {
        $('#service_port').parent().parent().removeClass('d-none');
        $('#service_check_type').parent().parent().addClass('d-none');
        $('#service_url').attr('placeholder', 'localhost');

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
    if (r === true) {
        return true;
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