import { getUserToken } from "../services/authsvc";

let pingAbout = null;
export async function aboutInfoApi(prom) {
  pingAbout = prom;
  await getUserToken(pingApi);
}

function pingApi(userToken) {
  const params = {
    method: "GET",
    headers: {
      "Authorization": userToken
    },
  };
  pingAbout(fetch("/api/ping", params));
}

let custData = null;
let showError = null;
let showSuccess = null;
export async function createCustomer(id, name, email, phonenumber, gstin, error , success) {
  const data = {
    customerid : id,
    name : name,
    email: email,
    phone : phonenumber,
    gstin: gstin
  }
  showSuccess = success;
  showError = error;
  custData = data;

  await getUserToken(createNewCust);
}

async function createNewCust(userToken) {
  const params = {
    method: "PUT",
    headers: {
      "Content-Type": "application/json",
      "Authorization" : userToken
    },
    body: JSON.stringify(custData),
  };

  let response =  await fetch("/api/customers", params).catch(err => showError(err) );
  if(response.ok) {
    showSuccess();
  } else {
    showError(response.status);
  }
}

let custList= null;
export async function getCustomers(prom) {
  custList = prom;
  await getUserToken(updateCustList);
}

function updateCustList(userToken) {
  const params = {
    method: "GET",
    headers: {
      "Authorization" :  userToken
    },
  };
  custList(fetch("/api/customers", params));
}