import { getUserToken } from "../services/authsvc";
// "Set-Cookie": "yummy_cookie=choco"

export function aboutInfoApi() {
  const params = {
    method: "GET",
    headers: {
      "Content-Type": "application/text",
    },
  };
  return fetch("/api/ping", params);
}


export function updateUsrCookie() {
  console.log(getUserToken());
  const params = {
    method: "POST",
    headers: {
      "Content-Type": "application/text",
      "Authorization" :  getUserToken()
    },
  };
  return fetch("/api/", params);
}