import React, { useState } from "react";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faCircleNotch } from "@fortawesome/free-solid-svg-icons";
import ReactTooltip from "react-tooltip";
import API from "../config/API";
import langs from "../config/langs";
import GroupServiceFailures from "./GroupServiceFailures";
import SubServiceCard from "./SubServiceCard";
// import IncidentsBlock from "./IncidentsBlock";
// import ServiceLoader from "./ServiceLoader";
// import DateUtils from "../utils/DateUtils";
import infoIcon from "../static/info.svg";

const GroupItem = ({ service, showPlusButton }) => {
  const [collapse, setCollapse] = useState(false);
  const [subServices, setSubServices] = useState([]);
  const [loading, setLoading] = useState(false);
  const [hoverText, setHoverText] = useState("");

  // const groupServices = services
  //   .filter((s) => s.group_id === service.id)
  //   .sort((a, b) => a.order_id - b.order_id);

  const fetchSubServices = async () => {
    const data = await API.fetchSubServices(service.id);
    if (Array.isArray(data)) {
      const sorted_data = data.sort((a, b) => a.order_id - b.order_id);
      setSubServices(sorted_data);
    }
    setCollapse(true);
  };

  const openCollapse = () => {
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
  };

  const closeCollapse = () => {
    setCollapse(false);
  };

  const handleMouseOver = (service) => {
    setHoverText(service.description || service.name);
  };

  const handleMouseOut = () => setHoverText("");

  return (
    <div className="service-card service_item card-bg pb-0">
      {/** TODO: change span to navlink */}
      <div className="service_item--header mb-3">
        <div className="service_item--right">
          {!loading && showPlusButton && (
            <>
              {collapse ? (
                <button className="square-minus" onClick={closeCollapse} />
              ) : (
                <button className="square-plus" onClick={openCollapse} />
              )}
            </>
          )}

          {loading && <FontAwesomeIcon icon={faCircleNotch} spin />}

          <span
            className="subtitle no-decoration font-14 mr-1"
            // to="/service/1"
          >
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
          <span
            className={`badge float-right font-12 ${
              service.online ? "uptime" : "downtime"
            }`}
            style={{ display: collapse ? "none" : "block" }}
          >
            {service.online ? langs("online") : langs("offline")}
          </span>
        </div>
      </div>
      <GroupServiceFailures service={service} collapse={collapse} />
      {/*<IncidentsBlock service={service} /> */}
      <div
        className="list-group online_list"
        style={{ display: collapse ? "block" : "none" }}
      >
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
    </div>
  );
};

export default React.memo(GroupItem);

// import React from "react";
// import langs from "../config/langs";
// import { services } from "../data";
// import GroupServiceFailures from "./GroupServiceFailures";
// import IncidentsBlock from "./IncidentsBlock";
// // import DateUtils from "../utils/DateUtils";

// const GroupItem = ({ group }) => {
//   const groupServices = services
//     .filter((s) => s.group_id === group.id)
//     .sort((a, b) => a.order_id - b.order_id);

//   if (!groupServices.length > 0) return null;

//   return (
//     <div className="col-12 full-col-12">
//       {group.name !== "Empty Group" && (
//         <h4 className="group_header mb-3 mt-4">{group.name}</h4>
//       )}
//       <div className="list-group online_list mb-4">
//         {groupServices.map((service, i) => {
//           return (
//             <div key={i} className="service-card service-card-action">
//               <span
//                 className="no-decoration font-3"
//                 // to={DateUtils.serviceLink(service)}
//               >
//                 {service.name}
//               </span>

//               <span
//                 className={`badge text-uppercase float-right ${
//                   service.online ? "bg-success" : "bg-danger"
//                 }`}
//               >
//                 {service.online ? langs("online") : langs("offline")}
//               </span>
//               <GroupServiceFailures service={service} />
//               <IncidentsBlock service={service} />
//             </div>
//           );
//         })}
//       </div>
//     </div>
//   );
// };

// export default GroupItem;
