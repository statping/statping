import React, { useState } from "react";
import ReactTooltip from "react-tooltip";
import GroupServiceFailures from "./GroupServiceFailures";
import IncidentsBlock from "./IncidentsBlock";
import infoIcon from "../static/info.svg";
import { getServiceStatus } from "../utils/helper";

const SubServiceCard = ({ group, service }) => {
  const [hoverText, setHoverText] = useState("");

  const handleMouseOver = (service) => {
    setHoverText(service.description || service.name);
  };

  const handleMouseOut = () => setHoverText("");

  return (
    <div className="service-card service_item border-radius-0 border-right-0 border-left-0">
      {/** TODO: change span to navlink */}

      <div className="service_item--header">
        <div className="service_item--right">
          <span className="subtitle no-decoration font-14 mr-1">
            {service.name}
          </span>
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
        <div className="service_item--left">
          {getServiceStatus(service.online)}
        </div>
      </div>

      <GroupServiceFailures group={group} service={service} />

      <IncidentsBlock group={group} service={service} />
    </div>
  );
};

export default SubServiceCard;
