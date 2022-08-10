export function findStatus(data) {
  if (!Array.isArray(data)) return null;
  if (data.length === 0) return null;
  const uptime = data.every((d) => d.online === true);
  const degraded = data.some((d) => d.online === false);
  const downtime = data.every((d) => d.online === false);
  if (uptime) return "uptime";
  if (downtime) return "downtime";
  if (degraded) return "degraded";
  return "";
}

export function getIncidentTextType(type) {
  switch (type.toLowerCase()) {
    case "resolved":
      return "text-success";
    case "issue summary":
      return "text-danger";
    case "update":
      return "text-warning";
    default:
      return "";
  }
}

export const isObject = (obj) => {
  if (Object.prototype.toString.call(obj) === "[object Object]") {
    return true;
  }

  return false;
};

export const isObjectEmpty = (obj) => {
  if (Object.keys(obj).length === 0) {
    return true;
  }
  return false;
};

export const calcPer = (uptime, downtime) => {
  const percentage = ((uptime / (uptime + downtime)) * 100).toFixed(2);
  return percentage;
};

// export function formatString(arr) {
//   const arrayStr = arr.map((d) => {
//     let start_dt = DateUtils.parseISO(d.start);
//     let end_dt = DateUtils.parseISO(d.end);
//     let duration = DateUtils.duration(
//       DateUtils.parseISO(d.start),
//       DateUtils.parseISO(d.end)
//     );
//     return `${start_dt.toLocaleDateString()} - ${
//       STATUS_TEXT[d.sub_status]
//     } for ${duration}
//       (${format(start_dt, "hh:mm aaa")} - ${format(end_dt, "hh:mm aaa")})`;
//   });
//   return arrayStr.join("<br/>");
// }

export const setUerId = (id) => {
  localStorage.setItem("stat_user_id", id);
};

export const getUserId = () => {
  return localStorage.getItem("stat_user_id");
};

export const generateUUID = (length) => {
  return Array.from(Array(length), () =>
    Math.floor(Math.random() * 36).toString(36)
  ).join("");
};
