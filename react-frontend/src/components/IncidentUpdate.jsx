import React from "react";
// import API from "../config/API";
import DateUtils from "../utils/DateUtils";

const IncidentUpdate = ({ update, admin }) => {
  // async loadUpdates() {
  //   this.updates = await Api.incident_updates(this.incident)
  // }

  const deleteUpdate = (update) => {
    alert("Delete Incident:", update.incident);
    // const res = await API.incident_update_delete(update);
    // if (res.status === "success") {
    //   this.onUpdate();
    // }
  };

  return (
    <div className="incident-wrapper mb-3 pb-2 d-flex" role="alert">
      <div className="time-line mr-2">
        <span class="dot"></span>
      </div>

      <div>
        <span className="font-14">
          {update.message}
          {admin && (
            <button
              onClick={deleteUpdate(update)}
              type="button"
              className="close">
              <span aria-hidden="true">&times;</span>
            </button>
          )}
        </span>
        <span className="d-block small text-muted">
          Posted {DateUtils.ago(update.created_at)} ago.{" "}
          {DateUtils.format(
            DateUtils.parseISO(update.created_at),
            "MMM d, yyyy - HH:mm"
          )}
        </span>
      </div>
    </div>
  );
};

export default IncidentUpdate;
