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
    <div className="col-12 mb-3 pb-2 border-bottom" role="alert">
      <span
        className={`
          font-weight-bold text-capitalize
          ${update.type.toLowerCase() === "resolved" ? "text-success" : ""}
          ${update.type.toLowerCase() === "investigating" ? "text-danger" : ""}
          ${update.type.toLowerCase() === "update" ? "text-warning" : ""}
        `}
      >
        {update.type}
      </span>
      <span className="text-muted">
        - {update.message}
        {admin && (
          <button
            onClick={deleteUpdate(update)}
            type="button"
            className="close"
          >
            <span aria-hidden="true">&times;</span>
          </button>
        )}
      </span>
      <span className="d-block small">
        {DateUtils.ago(update.created_at)} ago
      </span>
    </div>
  );
};

export default IncidentUpdate;
