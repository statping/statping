import {
  faCheckCircle,
  faExclamationCircle,
} from "@fortawesome/free-solid-svg-icons";

export const STATUS_COLOR = {
  uptime: "#5ebe5b",
  downtime: "#d50a0a",
  degraded: "#e4951d",
};

export const STATUS_ICON = {
  uptime: faCheckCircle,
  downtime: faExclamationCircle,
  degraded: faExclamationCircle,
};

export const STATUS_TEXT = {
  uptime: "All services are up and running",
  downtime: "Multiple services are down",
  degraded: "Some services are impacted",
};

export const STATUS_CLASS = {
  up: "uptime",
  degraded: "degraded",
  down: "downtime",
};
