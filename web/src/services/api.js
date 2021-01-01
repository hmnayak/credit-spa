export function aboutinfoApi() {
  const params = {
    method: "GET",
    headers: {
      "Content-Type": "application/text",
    },
  };
  return fetch("/api/ping", params);
}
