import React from "react";
import DateUtils from "../utils/DateUtils";

const MessageBlock = ({ message }) => {
  return (
    <div className="card shadow mb-4" role="alert">
      <div className="card-body pb-2">
        <h3 className="mb-3 font-weight-bold">{message.title}</h3>
        <span className="mb-2">{message.description}</span>
        <div className="col-12 mb-0">
          <div className="dates">
            <div className="start">
              <strong>STARTS</strong> {DateUtils.niceDate(message.start_on)}
              <span></span>
            </div>
            <div className="ends">
              <strong>ENDS</strong> {DateUtils.niceDate(message.end_on)}
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default MessageBlock;
