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

var currentLocation = window.location;
var domain = $("#domain_input");
if (domain.val() === "") {
	domain.val(currentLocation.origin);
}

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
    $("#setup_button").text("Creating Statping...");
    return true;
});


$('form').submit(function() {
    $(this).find("button[type='submit']").prop('disabled',true);
    $(this).find("button[type='submit']").text('Loading...');
});