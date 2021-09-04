import React from "react";
// import { groups } from "../utils/data";
import GroupItem from "./GroupItem";
import { isObject, isObjectEmpty } from "../utils/helper";

function showPlus(service) {
  let show = false;
  if (
    isObject(service.sub_services_detail) &&
    !isObjectEmpty(service.sub_services_detail)
  ) {
    const arr = Object.values(service.sub_services_detail);
    const isPublic = arr.some((a) => a.public === true);
    show = service.type === "collection" && isPublic;
  }
  return show;
}

const Group = ({ services }) => {
  // const data = groups.sort((a, b) => a.order_id - b.order_id);
  // if (!data.length > 0) return <></>;

  return (
    <div className="list-group">
      {services?.map((service) => {
        const showPlusButton = showPlus(service);
        return (
          <GroupItem
            key={service.id}
            service={service}
            showPlusButton={showPlusButton}
          />
        );
      })}
    </div>
  );
};

export default Group;
