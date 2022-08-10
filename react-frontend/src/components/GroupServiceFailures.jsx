import { useState, useEffect } from "react";
// import useIntersectionObserver from "../hooks/useIntersectionObserver";
import DateUtils from "../utils/DateUtils";
import langs from "../config/langs";
import API from "../config/API";
import ServiceLoader from "./ServiceLoader";
import ReactTooltip from "react-tooltip";
import { STATUS_CLASS } from "../utils/constants";
import { calcPer, isObjectEmpty } from "../utils/helper";

const STATUS_TEXT = {
  up: "Uptime",
  down: "Down",
  degraded: "Degraded",
};

const groupByStatus = (arr = []) => {
  const res = arr.reduce((acc, val) => {
    const { duration, sub_status } = val;
    if (acc.hasOwnProperty(sub_status)) {
      acc[sub_status]["duration"] += duration;
      acc[sub_status]["count"] += 1;
    } else {
      acc[sub_status] = {
        duration,
        sub_status,
        count: 1,
      };
    }
    return acc;
  }, {});
  return res;
};

function formatStatusDuration(obj) {
  const arrayStr = Object.values(obj).map((d) => {
    let duration = DateUtils.humanize(d.duration);
    return `${STATUS_TEXT[d.sub_status]} for ${duration}`;
  });

  return arrayStr.join("<br/>");
}

async function fetchFailureSeries(url) {
  const { now, beginningOf, endOf, nowSubtract, toUnix } = DateUtils;
  const start = beginningOf("day", nowSubtract(86400 * 89));
  const end = endOf("day", now());
  const data = await API.service_failures_data(
    url,
    toUnix(start),
    toUnix(end),
    "24h",
    true
  );
  // console.log(data);
  return data;
}

const GroupServiceFailures = ({ group = null, service, collapse }) => {
  const [hoverText, setHoverText] = useState("");
  const [loaded, setLoaded] = useState(true);
  const [failureData, setFailureData] = useState([]);
  const [uptime, setUptime] = useState(0);

  useEffect(() => {
    async function fetchData() {
      let url = "/services";
      try {
        if (group) {
          url += `/${group.id}/sub_services/${service.id}/block_series`;
        } else {
          url += `/${service.id}/block_series`;
        }
        const { series, downtime, uptime } = await fetchFailureSeries(url);
        const failureData = [];
        series.forEach((d) => {
          let date = DateUtils.parseISO(d.timeframe);
          date = DateUtils.format(date, "dd MMMM yyyy");
          failureData.push({
            timeframe: date,
            status: d.status,
            downtimes: groupByStatus(d.downtimes),
          });
        });
        const percentage = calcPer(uptime, downtime);
        setFailureData(failureData);
        setUptime(percentage);
      } catch (e) {
        console.log(e.message);
      } finally {
        setLoaded(false);
      }
    }
    fetchData();
  }, [service, group]);

  const handleTooltip = (d) => {
    let txt = "";
    if (d.status === "up") {
      txt = `<div style="text-align:center;">
      <div>${d.timeframe}</div>
      <div>No Downtime</div>
      </div>`;
    } else if (d.status === "down" && !isObjectEmpty(d.downtimes)) {
      txt = `<div style="text-align:center;">
      <div>${d.timeframe}</div>
      <div>${formatStatusDuration(d.downtimes)}</div>
      </div>`;
    } else if (d.status === "degraded") {
      txt = `<div style="text-align:center;">
      <div>${d.timeframe}</div>
      <div>${formatStatusDuration(d.downtimes)}</div>
      </div>`;
    }
    return txt;
  };

  const handleMouseOver = (d) => {
    // let date = DateUtils.parseISO(d.timeframe);
    // date = date.toLocaleDateString();
    const tooltipText = handleTooltip(d);
    setHoverText(tooltipText);
  };

  const handleMouseOut = () => setHoverText("");

  if (loaded) return <ServiceLoader text="Loading series.." />;

  return (
    <div name="fade" style={{ display: collapse ? "none" : "block" }}>
      <div className="block-chart">
        <ReactTooltip
          effect="solid"
          place="bottom"
          backgroundColor="#344A6C"
          html={true}
        />
        {failureData?.length > 0 ? (
          failureData.map((d, i) => {
            return (
              <div
                className={`flex-fill service_day ${STATUS_CLASS[d.status]}`}
                onMouseOver={() => handleMouseOver(d)}
                onMouseOut={handleMouseOut}
                key={i}
                data-tip={hoverText}>
                {d.status !== 0 && (
                  <span className="d-none d-md-block text-center small"></span>
                )}
              </div>
            );
          })
        ) : (
          <span className="description font-10">
            Service does not have any successful hits
          </span>
        )}
      </div>
      <div className="timeline">
        <div className="no-select">
          <p className="divided justify-content-between">
            <span className="timeline-text font-12">
              90 {langs("days_ago")}
            </span>
            {/* <span className="timeline-divider"></span> */}
            {/* <span className="timeline-text font-12">{service_txt()}</span> */}
            {/* <span className="timeline-text font-12">{uptime}% uptime</span> */}
            <span className="timeline-divider"></span>
            <span className="timeline-text font-12">{langs("today")}</span>
          </p>
        </div>
      </div>
      {/* <div className="daily-failures small text-right text-dim">
        {hoverText}
      </div> */}
    </div>
  );
};

export default GroupServiceFailures;
