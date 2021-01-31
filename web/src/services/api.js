import { getUserToken } from "../services/authsvc";

export function aboutInfoApi() {
  const params = {
    method: "GET",
    headers: {
      "Content-Type": "application/text",
      "Authorization" :  getUserToken()
    },
  };
  return fetch("/api/ping", params);
}


export function updateUsrToken() {
  const params = {
    method: "POST",
    headers: {
      "Content-Type": "application/text",
      "Authorization" :  getUserToken()
    },
  };
  return fetch("/api/usrloggedin", params);
}