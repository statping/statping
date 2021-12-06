import React from "react";
import langs from "../config/langs";
import GroupServiceFailures from "./GroupServiceFailures";
import IncidentsBlock from "./IncidentsBlock";

const ServiceCard = ({ service }) => {
  return (
    <div className="service-card service_item card-bg">
      {/** TODO: change span to navlink */}
      <div className="service_item--header">
        <div className="service_item--right">
          {service.type === "collection" && (
            <span className="square-plus"></span>
          )}
          <span
            className="subtitle no-decoration font-14"
            // to="/service/1"
          >
            {service.name}
          </span>
          {/* <span className="info">i</span> */}
        </div>
        <div className="service_item--left">
          <span
            className={`badge float-right font-12 ${
              service.online ? "status-green" : "status-red"
            }`}
          >
            {service.online ? langs("online") : langs("offline")}
          </span>
        </div>
      </div>

      <GroupServiceFailures service={service} />
      {/* <IncidentsBlock service={service} /> */}
    </div>
  );
};

export default ServiceCard;
