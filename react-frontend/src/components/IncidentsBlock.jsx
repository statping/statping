import { useState, useEffect, Fragment } from "react";
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

      if(Array.isArray(data)) {
        setIncidents(data);
      }
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
            const { id, title, description, updates, updated_at } = incident;
            const latestUpdate =
              updates?.length > 0 && updates[0];
            const updatedAt = latestUpdate
              ? latestUpdate.created_at
              : updated_at;

            return (
              <Fragment key={id}>
                <span className="braker mt-1 mb-3"></span>

                <div
                  className={`incident-title col-12 ${
                    incidentsShow[id] && "mb-3"
                  }`}>
                  {updates?.length > 0 && (
                    <>
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
                    </>
                  )}

                  <div className="title-wrapper">
                    <span className="subtitle no-decoration">{title}</span>
                    <span className="d-block small text-dark">
                      {description}
                    </span>
                    <span className="d-block small text-muted">
                      Updated {DateUtils.ago(updatedAt)} ago.{" "}
                      {DateUtils.format(
                        DateUtils.parseISO(updatedAt),
                        "MMM d, yyyy - HH:mm"
                      )}
                    </span>
                  </div>
                </div>

                {incidentsShow[id] && (
                  <div className="incident-updates-wrapper col-12">
                    {updates?.map((update) => {
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
              </Fragment>
            );
          })
        ) : (
          <div className="col-12">
            <span className="font-14 text-muted">No recent incidents</span>
          </div>
        )}
      </div>
    </div>
  );
};

export default IncidentsBlock;
