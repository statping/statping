$('select#database_type').on('change', function(){
    var selected = $('#database_type option:selected').val();
    if (selected=="sqlite") {
        $("#db_host").hide();
        $("#db_password").hide();
        $("#db_port").hide();
        $("#db_user").hide();
        $("#db_database").hide();
    } else {
        $("#db_host").show();
        $("#db_password").show();
        $("#db_port").show();
        $("#db_user").show();
        $("#db_database").show();
    }
});