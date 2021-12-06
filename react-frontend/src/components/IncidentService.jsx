import { useState, useEffect } from "react";
import { service, up_time_data } from "../data";
import DateUtils from "../utils/DateUtils";
import langs from "../config/langs";
import API from "../config/API";

const IncidentService = () => {
  async function lastDaysFailures() {
    const { beginningOf, endOf, nowSubtract, parseISO, toUnix } = DateUtils;
    const failureData = [];
    // const start = beginningOf("day", nowSubtract(86400 * 90));
    // const end = endOf("tomorrow");
    const data = await API.service_failures_data();
    data.forEach((d) => {
      let date = DateUtils.parseISO(d.timeframe);
      failureData.push({
        month: date.getMonth(),
        day: date.getDate(),
        date: date,
        amount: d.amount,
      });
    });

    return failureData;
  }

  const [hoverText, setHoverText] = useState("");
  const [data, setData] = useState([]);
  const [loaded, setLoaded] = useState(false);

  useEffect(async () => {
    const failureData = await lastDaysFailures();
    setData(failureData);
    setLoaded(true);
  }, []);

  const handleMouseOut = () => {
    setHoverText("");
  };

  const handleMouseOver = (d) => {
    const start = DateUtils.parseISO(d.start);
    const end = DateUtils.parseISO(d.end);
    const duration = DateUtils.humanTime(d.duration);

    setHoverText(
      `${start.toLocaleDateString()} - ${end.toLocaleDateString()}
      ${duration} ${d.online ? "Online" : "Offline"}`
    );
  };

  const service_txt = () => {
    return DateUtils.smallText(service);
  };

  const total_time_frame = up_time_data.uptime + up_time_data.downtime;

  return (
    <div className="list-group online_list mb-4">
      <div className="service-card service-card-action">
        {/** TODO: change span to navlink */}
        <span className="no-decoration font-3" t0="/service/1">
          Status Ping
        </span>
        <span
          className={`badge float-right ${
            service.online ? "bg-success" : "bg-danger"
          }`}
        >
          {service.online ? "ONLINE" : "OFFLINE"}
        </span>
        <div className="d-flex mt-3">
          {up_time_data.series.map((d, i) => {
            const time = ((d.duration * 100) / total_time_frame).toFixed(2);
            const width = time < 0.1 ? 0.15 : time;
            return (
              <div
                className={`service_day ${d.online ? "uptime" : "downtime"}`}
                style={{ width: width + "%" }}
                onMouseOver={() => handleMouseOver(d)}
                onMouseOut={handleMouseOut}
                key={i}
              >
                {d.amount !== 0 && (
                  <span className="d-none d-md-block text-center small"></span>
                )}
              </div>
            );
          })}
        </div>
        <div className="row mt-2">
          <div className="col-12 no-select">
            <p className="divided">
              <span className="font-2 text-muted">
                {DateUtils.duration(up_time_data.uptime, up_time_data.downtime)}{" "}
                {langs("days_ago")}
              </span>
              <span className="divider"></span>
              <span
                className={`text-center font-2 ${
                  service.online ? "text-muted" : "text-danger"
                }`}
              >
                {service_txt()}
              </span>
              <span className="divider"></span>
              <span className="font-2 text-muted">{langs("today")}</span>
            </p>
          </div>
        </div>
        <div className="daily-failures small text-right text-dim">
          {hoverText}
        </div>
      </div>
    </div>
  );
};

export default IncidentService;
