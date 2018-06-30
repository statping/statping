$(".service_li").on('click', function() {
    var id = $(this).attr('data-id');
    var position = $("#service_id_"+id).offset();
    window.scroll(0,position.top-23);
    return false;
});


$('form').submit(function() {
    $(this).find("button[type='submit']").prop('disabled',true);
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