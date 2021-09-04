import React from "react";
import { core } from "../utils/data";

const ContentHeader = () => {
  return (
    <div className="header">
      <h1 className="header-title mt-4 mb-3 font-24 fw-700">{core.name}</h1>
      <h5 className="header-description font-12">{core.description}</h5>
    </div>
  );
};

export default ContentHeader;
