import React, { useState, useEffect } from "react";
// import { NavLink } from "react-router-dom";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import {
  faCheckCircle,
  faExclamationCircle,
} from "@fortawesome/free-solid-svg-icons";
import DateUtils from "../utils/DateUtils";
import Group from "./Group";
import ContentHeader from "./ContentHeader";
import ServiceLoader from "./ServiceLoader";
// import IncidentService from "./IncidentService";
// import MessageBlock from "./MessageBlock";
// import ServiceBlock from "./ServiceBlock";
// import ServicesList from "./ServicesList";
import API from "../config/API";
import { STATUS_COLOR, STATUS_TEXT } from "../utils/constants";
import { findStatus } from "../utils/helper";

const ServicesPage = () => {
  // const data = messages.filter((m) => inRange(m) && m.service === 0);
  const [services, setServices] = useState([]);
  const [status, setStatus] = useState(true);
  const [loading, setLoading] = useState(true);
  const [poll, setPolling] = useState(1);
  const today = DateUtils.format(new Date(), "d MMMM yyyy, hh:mm aaa");

  useEffect(() => {
    const timer = setInterval(() => {
      setPolling((prev) => (prev += 1));
    }, 120000);
    return () => clearInterval(timer);
  }, [poll]);

  useEffect(() => {
    const fetchServices = async () => {
      try {
        const data = await API.fetchServices();
        const status = findStatus(data);
        const sorted_data = data.sort((a, b) => a.order_id - b.order_id);
        setServices(sorted_data);
        setStatus(status);
      } catch (e) {
        console.log(e.message);
      } finally {
        setLoading(false);
      }
    };
    fetchServices();
  }, [poll]);

  return (
    <div className="container col-md-7 col-sm-12 sm-container">
      <ContentHeader />
      <div className="app-content">
        <div className="service">
          <h2 className="title font-20 fw-700">Razorpay Payments</h2>
          <div className="d-flex align-items-center subtitle font-12 mt-2">
            <FontAwesomeIcon
              icon={status === "up" ? faCheckCircle : faExclamationCircle}
              style={{
                fontSize: "16px",
                color: STATUS_COLOR[status],
              }}
            />
            <span className="mx-1">{STATUS_TEXT[status]}</span>
            <span className="date">{today}</span>
          </div>
        </div>

        {loading && <ServiceLoader text="Loading Services" />}

        {/* <ServicesList loading={loading} services={services} /> */}

        {/* TODO --> Grouped Services to Accordian*/}
        {services && services.length > 0 ? (
          <Group services={services} />
        ) : (
          <div className="description text-align-center">No Services</div>
        )}

        {/* <div>
            {data.map((message) => {
              return <MessageBlock key={message.id} message={message} />;
            })}
          </div>

          <div>
            {services.map((service) => {
              return <ServiceBlock key={service.id} service={service} />;
            })}
          </div> */}

        <div className="app-footer">
          <div className="service-status">
            <span className="service-status-badge uptime"></span>
            <span className="description font-12">100% Uptime</span>
            <span></span>
          </div>
          <div className="service-status">
            <span className="service-status-badge degraded"></span>
            <span className="description font-12">Partial degradation</span>
            <span></span>
          </div>
          <div className="service-status">
            <span className="service-status-badge downtime"></span>
            <span className="description font-12">Downtime</span>
            <span></span>
          </div>
        </div>
      </div>
    </div>
  );
};

export default ServicesPage;
