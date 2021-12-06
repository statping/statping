import React from "react";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faCircleNotch } from "@fortawesome/free-solid-svg-icons";

const ServiceLoader = ({ text = "Loading..." }) => {
  return (
    <div className="row mt-5 mb-5">
      <div className="col-12 mt-5 mb-2 text-center">
        <FontAwesomeIcon
          icon={faCircleNotch}
          className="text-dim"
          size="2x"
          spin
        />
      </div>
      <div className="col-12 text-center mt-3 mb-3">
        <span className="text-dim">{text}</span>
      </div>
    </div>
  );
};

export default ServiceLoader;
