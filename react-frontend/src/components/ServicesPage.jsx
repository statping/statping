import React, { useState, useEffect } from "react";
import Group from "./Group";
import ContentHeader from "./ContentHeader";
import ServiceLoader from "./ServiceLoader";
import API from "../config/API";
import { findStatus } from "../utils/helper";
import { analyticsTrack } from "../utils/trackers";

const ServicesPage = () => {
  const [services, setServices] = useState([]);
  const [status, setStatus] = useState("uptime");
  const [loading, setLoading] = useState(true);
  const [poll, setPolling] = useState(1);

  useEffect(() => {
    if (!loading) {
      analyticsTrack({
        objectName: "Status Page",
        actionName: "displayed",
        screen: "Home page",
      });
    }
  }, [loading]);

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
        </div>

        {loading && <ServiceLoader text="Loading Services" />}

        {services && services.length > 0 ? (
          <Group services={services} />
        ) : (
          <div className="description text-align-center">No Services</div>
        )}

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
