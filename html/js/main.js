$(".service_li").on('click', function() {
    var id = $(this).attr('data-id');
    var position = $("#service_id_"+id).offset();
    window.scroll(0,position.top-23);
    return false;
});