import React, { useState, useEffect } from "react";
import { services } from "../utils/data";
import ServiceCard from "./ServiceCard";

const ServicesList = ({ services }) => {
  // const [state, setstate] = useState([]);

  // useEffect(() => {
  //   const data = services
  //     .filter((g) => g.group_id === 0)
  //     .sort((a, b) => a.order_id - b.order_id);
  //   setstate(data);
  // }, []);

  return (
    <div className="list-group online_list">
      {services?.map((service) => {
        return <ServiceCard key={service.id} service={service} />;
      })}
    </div>
  );
};

export default ServicesList;
