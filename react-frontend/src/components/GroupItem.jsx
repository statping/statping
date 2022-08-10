import React, { useState } from "react";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faCircleNotch } from "@fortawesome/free-solid-svg-icons";
import ReactTooltip from "react-tooltip";
import API from "../config/API";
import langs from "../config/langs";
import GroupServiceFailures from "./GroupServiceFailures";
import SubServiceCard from "./SubServiceCard";
import infoIcon from "../static/info.svg";
import { analyticsTrack } from "../utils/trackers";
import IncidentsBlock from "./IncidentsBlock";

const GroupItem = ({ service, showPlusButton }) => {
  const [collapse, setCollapse] = useState(false);
  const [subServices, setSubServices] = useState([]);
  const [loading, setLoading] = useState(false);
  const [hoverText, setHoverText] = useState("");

  const fetchSubServices = async () => {
    const data = await API.fetchSubServices(service.id);
    if (Array.isArray(data)) {
      const sorted_data = data.sort((a, b) => a.order_id - b.order_id);
      setSubServices(sorted_data);
    }
    setCollapse(true);
  };

  const openCollapse = (event) => {
    if (subServices.length === 0) {
      setLoading(true);
      try {
        fetchSubServices();
      } catch (e) {
        console.log(e.message);
      } finally {
        setLoading(false);
      }
    } else {
      setCollapse(true);
    }

    analyticsTrack({
      objectName: "Service Expand",
      actionName: "clicked",
      screen: "Home page",
      properties: {
        serviceName: event.target.name,
      },
    });
  };

  const closeCollapse = (event) => {
    setCollapse(false);

    analyticsTrack({
      objectName: "Service Collapse",
      actionName: "clicked",
      screen: "Home page",
      properties: {
        serviceName: event.target.name,
      },
    });
  };

  const handleMouseOver = (service) => {
    setHoverText(service.description || service.name);
  };

  const handleMouseOut = () => setHoverText("");

  return (
    <div className="service-parent service-card service_item card-bg">
      {/** TODO: change span to navlink */}
      <div className="service_item--header mb-3">
        <div className="service_item--right">
          {!loading && showPlusButton && (
            <>
              {collapse ? (
                <button
                  className="square-minus"
                  name={service.name}
                  onClick={closeCollapse}
                />
              ) : (
                <button
                  className="square-plus"
                  name={service.name}
                  onClick={openCollapse}
                />
              )}
            </>
          )}

          {loading && <FontAwesomeIcon icon={faCircleNotch} spin />}

          <span className="subtitle no-decoration mr-1">{service.name}</span>
          {service?.description && (
            <>
              <ReactTooltip
                id={`tooltip-${service.name}`}
                effect="solid"
                place="right"
                backgroundColor="#344A6C"
                className="tooltip"
              />
              <img
                onMouseOver={() => handleMouseOver(service)}
                onMouseOut={handleMouseOut}
                src={infoIcon}
                alt="info-icon"
                data-for={`tooltip-${service.name}`}
                data-tip={hoverText}
              />
            </>
          )}
        </div>
        {!collapse && (
          <div className="service_item--left">
            <span
              className={`badge float-right font-12 ${
                service.online ? "uptime" : "downtime"
              }`}>
              {service.online ? langs("online") : langs("offline")}
            </span>
          </div>
        )}
      </div>

      {!collapse && (
        <GroupServiceFailures service={service} collapse={collapse} />
      )}

      {!collapse && <IncidentsBlock service={service} />}

      {collapse && (
        <div className="sub-service-wrapper list-group online_list">
          {subServices && subServices?.length > 0 ? (
            subServices.map((sub_service, i) => {
              return (
                <SubServiceCard
                  key={i}
                  group={service}
                  service={sub_service}
                  collapse={collapse}
                />
              );
            })
          ) : (
            <div className="subtitle text-align-center">No Services</div>
          )}
        </div>
      )}
    </div>
  );
};

export default React.memo(GroupItem);
