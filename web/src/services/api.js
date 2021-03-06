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

export function createCustomer(name, email, phonenumber, gstin, showError , showSuccess) {
  const data = {
    name : name,
    email: email,
    phone : phonenumber,
    gstin: gstin
  }
  const params = {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      "Authorization" :  getUserToken()
    },
    body: JSON.stringify(data),
  };
  fetch("/api/createCustomer", params)
  .then(response => response.text())
  .then(data => {
    showSuccess();
  })
  .catch((error) => {
    showError(error)
  });
}

export function getCustomers() {
  // const params = {
  //   method: "GET",
  //   headers: {
  //     "Content-Type": "application/text",
  //     "Authorization" :  getUserToken()
  //   },
  // };
  // return fetch("/api/customers", params);
  return ["Customer1", "Customer2", "Customer3"];
}