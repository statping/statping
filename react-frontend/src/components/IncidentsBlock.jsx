import { useState, useEffect } from "react";
import API from "../config/API";
import DateUtils from "../utils/DateUtils";
import IncidentUpdate from "./IncidentUpdate";

const IncidentsBlock = ({ service }) => {
  const [incidents, setIncidents] = useState([]);

  useEffect(() => {
    async function fetchData() {
      const data = await API.incidents_service(service.id);
      setIncidents(data);
    }
    fetchData();
  }, [service.id]);

  return (
    <div className="row">
      {incidents?.map((incident, i) => {
        return (
          <div className="col-12 mt-2" key={i}>
            <span className="braker mt-1 mb-3"></span>
            <h6>
              {incident.title}
              <span className="font-2 float-right">
                {DateUtils.niceDate(incident.created_at)}
              </span>
            </h6>
            <div className="font-2 mb-3" v-html="incident.description"></div>
            {incident.updates.map((update, i) => {
              return <IncidentUpdate key={i} update={update} admin={false} />;
            })}
          </div>
        );
      })}
    </div>
  );
};

export default IncidentsBlock;
