import { useState, useEffect } from "react";
import API from "../config/API";
import DateUtils from "../utils/DateUtils";
import IncidentUpdate from "./IncidentUpdate";

const IncidentsBlock = ({ service, group }) => {
  const [incidents, setIncidents] = useState([]);
  const [incidentsShow, setIncidentsShow] = useState(false);

  useEffect(() => {
    async function fetchData() {
      let data = [];

      if (group?.id) {
        data = await API.sub_incidents_service(group.id, service.id);
      } else {
        data = await API.incidents_service(service.id);
      }

      setIncidents(data || []);
    }
    fetchData();
  }, [service.id, group?.id]);

  const handleIncidentShow = (event) => {
    const { id } = event.target;

    setIncidentsShow({ ...incidentsShow, [id]: !incidentsShow[id] });
  };

  return (
    <div className="incidents-wrapper row">
      <div className="col-12 mt-2">
        {incidents?.length > 0 ? (
          incidents?.map((incident) => {
            const { id, title, description, updated_at } = incident;

            return (
              <>
                <span className="braker mt-1 mb-3"></span>
                <div
                  className={`incident-title col-12 ${
                    incidentsShow[id] && "mb-3"
                  }`}>
                  {incidentsShow[id] ? (
                    <button
                      className="square-minus"
                      type="button"
                      id={id}
                      onClick={handleIncidentShow}
                    />
                  ) : (
                    <button
                      className="square-plus"
                      type="button"
                      id={id}
                      onClick={handleIncidentShow}
                    />
                  )}
                  <div className="title-wrapper">
                    <span class="subtitle no-decoration">{title}</span>
                    <span className="d-block small text-dark">
                      {description}
                    </span>
                    <span className="d-block small text-muted">
                      Updated {DateUtils.ago(updated_at)} ago.{" "}
                      {DateUtils.format(
                        DateUtils.parseISO(updated_at),
                        "MMM d, yyyy - HH:mm"
                      )}
                    </span>
                  </div>
                </div>
                {incidentsShow[id] && (
                  <div className="incident-updates-wrapper col-12">
                    {incident?.updates.map((update) => {
                      return (
                        <IncidentUpdate
                          key={update.id}
                          update={update}
                          admin={false}
                        />
                      );
                    })}
                  </div>
                )}
              </>
            );
          })
        ) : (
          <div className="col-12">
            <span class="font-14 text-muted">No recent incidents</span>
          </div>
        )}
      </div>
    </div>
  );
};

export default IncidentsBlock;
