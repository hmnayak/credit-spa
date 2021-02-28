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

export function createInvoice(name, email, phonenumber, gstin, ) {
  const data = {
    username : name,
    email: email,
    phonenumber: phonenumber,
    gstin: gstin
  }
  const params = {
    method: "PUT",
    headers: {
      "Content-Type": "application/text",
      "Authorization" :  getUserToken()
    },
    body: JSON.stringify(data),
  };
  return fetch("/api/createInvoice", params)
  .then(response => response.json())
  .then(data => {
    console.log('Success:', data);
  })
  .catch((error) => {
    console.error('Error:', error);
    showError(error)
  });
}