import React, { useState } from "react";
import { useHistory } from "react-router-dom";
import langs from "../config/langs";
import DateUtils from "../utils/DateUtils";
import ServiceChart from "./ServiceChart";
import ServiceTopStats from "./ServiceTopStats";

const timeset = (seconds) => {
  return DateUtils.toUnix(DateUtils.nowSubtract(seconds));
};

const timeframes = [
  { value: timeset(1800), text: "30 Minutes", set: 1 },
  { value: timeset(3600), text: "1 Hour", set: 2 },
  { value: timeset(21600), text: "6 Hours", set: 3 },
  { value: timeset(43200), text: "12 Hours", set: 4 },
  { value: timeset(86400), text: "1 Day", set: 5 },
  { value: timeset(259200), text: "3 Days", set: 6 },
  { value: timeset(604800), text: "7 Days", set: 7 },
  { value: timeset(1209600), text: "14 Days", set: 8 },
  { value: timeset(2592000), text: "1 Month", set: 9 },
  { value: timeset(7776000), text: "3 Months", set: 10 },
  { value: 0, text: "All Records" },
];

const intervals = [
  { value: "1m", text: "1/min", set: 1 },
  { value: "5m", text: "5/min", set: 2 },
  { value: "15m", text: "15/min", set: 3 },
  { value: "30m", text: "30/min", set: 4 },
  { value: "60m", text: "1/hr", set: 5 },
  { value: "180m", text: "3/hr", set: 6 },
  { value: "360m", text: "6/hr", set: 7 },
  { value: "720m", text: "12/hr", set: 8 },
  { value: "1440m", text: "1/day", set: 9 },
  { value: "4320m", text: "3/day", set: 10 },
  { value: "10080m", text: "7/day", set: 11 },
];

const stats = {
  total_failures: {
    title: "Total Failures",
    subtitle: "Last 7 Days",
    value: 0,
  },
  high_latency: {
    title: "Highest Latency",
    subtitle: "Last 7 Days",
    value: 0,
  },
  lowest_latency: {
    title: "Lowest Latency",
    subtitle: "Last 7 Days",
    value: 0,
  },
  high_ping: {
    title: "Highest Ping",
    subtitle: "Last 7 Days",
    value: 0,
  },
  low_ping: {
    title: "Lowest Ping",
    subtitle: "Last 7 Days",
    value: 0,
  },
};

const ServiceBlock = ({ service }) => {
  const history = useHistory();
  const [expanded, setExpanded] = useState(false);
  const [visible, setVisible] = useState(false);
  const [dropDownMenu, setDropDownMenu] = useState(false);
  const [intervalMenu, setIntervalMenu] = useState(false);
  const [intervalVal, setIntervalVal] = useState("60m");
  const [timeframeVal, setTimeframeVal] = useState(timeset(259200));
  // const [service, setService] = useState(null);

  const timeframepick = timeframes.find((s) => s.value === timeframeVal);

  const intervalpick = intervals.find((s) => s.value === intervalVal);

  const chartTimeframe = {
    start_time: timeframeVal,
    interval: intervalVal,
  };

  const disabled_interval = (interval) => {
    let min = timeframepick.set - interval.set - 1;
    return min >= interval.set;
  };

  const openMenu = (tm) => {
    if (tm === "interval") {
      setIntervalMenu(!intervalMenu);
      setDropDownMenu(false);
    } else if (tm === "timeframe") {
      setDropDownMenu(!dropDownMenu);
      setIntervalMenu(false);
    }
  };

  const changeInterval = (interval) => {
    setIntervalVal(interval.value);
    setIntervalMenu(false);
    setDropDownMenu(false);
  };

  const changeTimeframe = (timeframe) => {
    setTimeframeVal(timeframe.value);
    setIntervalMenu(false);
    setDropDownMenu(false);
  };

  const handlesetService = () => {
    // setService(service);
    history.push("/service/" + service.id, {
      props: { service: service },
    });
  };

  const visibleChart = (isVisible, entry) => {
    if (isVisible && !visible) {
      setVisible(true);
    }
  };

  return (
    <div className="mb-md-4 mb-4">
      <div className={`card index-chart ${expanded ? "expanded-service" : ""}`}>
        <div className="card-body">
          <div className="col-12">
            <h4 className="mt-2">
              <span
                className="d-inline-block text-truncate font-4"
                style={{ maxWidth: "65vw" }}
                // to={serviceLink(service)}
                // in_service={service}
              >
                {service.name}
              </span>
              <span
                className={`badge float-right ${
                  service.online ? "bg-success" : "bg-danger"
                }`}
              >
                {service.online ? "ONLINE" : "OFFLINE"}
              </span>
            </h4>

            <ServiceTopStats service={service} />
          </div>
        </div>

        {/* <div
          // v-observe-visibility="{ callback: visibleChart, throttle: 200 }"
          className="chart-container"
        >
          <ServiceChart
            service={service}
            visible={visible}
            chartTimeframe={chartTimeframe}
          />
        </div> */}

        <div
          className={`row lower_canvas full-col-12 text-white ${
            service.online ? "bg-success" : "bg-danger"
          }`}
        >
          <div className="col-md-10 col-6">
            <div className={`dropup ${dropDownMenu ? "show" : ""}`}>
              <button
                style={{ fontSize: "10pt" }}
                onClick={() => openMenu("timeframe")}
                type="button"
                className="col-4 float-left btn btn-sm float-right btn-block text-white dropdown-toggle service_scale pr-2"
              >
                {timeframepick.text}
              </button>
              <div
                className={`service-tm-menu ${!dropDownMenu ? "d-none" : ""}`}
              >
                {timeframes.map((timeframe, i) => {
                  return (
                    <a
                      href="#"
                      onClick={() => changeTimeframe(timeframe)}
                      className={`dropdown-item ${
                        timeframepick === timeframe ? "active" : ""
                      }`}
                      key={i}
                    >
                      {timeframe.text}
                    </a>
                  );
                })}
              </div>
            </div>

            <div className={`dropup ${intervalMenu ? "show" : ""}`}>
              <button
                style={{ fontSize: "10pt" }}
                onClick={() => openMenu("interval")}
                type="button"
                className="col-4 float-left btn btn-sm float-right btn-block text-white dropdown-toggle service_scale pr-2"
              >
                {intervalpick.text}
              </button>
              <div
                className={`service-tm-menu ${!intervalMenu ? "d-none" : ""}`}
              >
                {intervals.map((interval, i) => {
                  return (
                    <a
                      href="#"
                      onClick={() => changeInterval(interval)}
                      className={`dropdown-item ${
                        intervalpick === interval ? "active" : ""
                      }`}
                      disabled={disabled_interval(interval)}
                      key={i}
                    >
                      {interval.text}
                    </a>
                  );
                })}
              </div>

              <span className="d-none float-left d-md-inline">
                {DateUtils.smallText(service)}
              </span>
            </div>
          </div>

          {/* <div className="col-md-2 col-6 float-right">
            <button
              onClick={() => setExpanded(!expanded)}
              // onClick={handlesetService}
              className={`btn btn-sm float-right dyn-dark text-white ${
                service.online ? "bg-success" : "bg-danger"
              }`}
            >
              {langs("view")}
            </button>
          </div> */}
        </div>
      </div>
    </div>
  );
};

export default ServiceBlock;
