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

export function createCustomer(name, email, phonenumber, gstin, showError ) {
  const data = {
    username : name,
    email: email,
    phonenumber: phonenumber,
    gstin: gstin
  }
  const params = {
    method: "POST",
    headers: {
      "Content-Type": "application/text",
      "Authorization" :  getUserToken()
    },
    body: JSON.stringify(data),
  };
  return fetch("/api/createCustomer", params)
  .then(response => response.json())
  .then(data => {
    console.log('Success:', data);
  })
  .catch((error) => {
    console.error('Error:', error);
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