var currentLocation = window.location;
$("#domain_input").val(currentLocation.origin);

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
    if (selected=="mysql") {
        $("#db_port_in").val('3306');
    } else if (selected=="postgres") {
        $("#db_port_in").val('5432');
    }

});

$("#setup_form").submit(function() {
    $("#setup_button").prop("disabled", true);
    $("#setup_button").text("Creating Statup...");
    return true;
});


$('form').submit(function() {
    $(this).find("button[type='submit']").prop('disabled',true);
    $(this).find("button[type='submit']").text('Loading...');
});