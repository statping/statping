import React from "react";
import langs from "../config/langs";
import DateUtils from "../utils/DateUtils";

const ServiceTopStats = ({ service }) => {
  return (
    <div className="row stats_area mt-5 mb-4">
      <div className="col-4">
        <span className="font-5 d-block font-weight-bold">
          {DateUtils.humanTime(service.avg_response)}
        </span>
        <span className="font-1 subtitle">{langs("average_response")}</span>
      </div>
      <div className="col-4">
        <span className="font-5 d-block font-weight-bold">
          {service.online_24_hours} %
        </span>
        <span className="font-1 subtitle">
          {langs("last_uptime")}
          {/* {langs("last_uptime", [24, $tc("hour", 24)])} */}
        </span>
      </div>
      <div className="col-4">
        <span className="font-5 d-block font-weight-bold">
          {service.online_7_days} %
        </span>
        <span className="font-1 subtitle">
          {langs("last_uptime")}
          {/* {langs("last_uptime", [7, $tc("day", 7)])} */}
        </span>
      </div>
    </div>
  );
};

export default ServiceTopStats;
