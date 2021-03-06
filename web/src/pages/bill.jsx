import React, { useState, useEffect } from "react";
import { Page, Block } from "framework7-react";
import { getCustomers } from "../services/api";

export default (props) => {
  const [customers, setCustomers] = useState([]);

  let custPromise = getCustomers();

  useEffect(() => {
    // setCustomers(getCustomers);
    custPromise
      .then((res) => res.json())
      .then((res) => {
        console.log(res);
        setCustomers(res);
      });
  }, []);

  return (
    <Page>
      <Block strong>
        <div className="list">
          <ul>
            <li>
              <a className="button" href="/customer/">
                Add new customer
              </a>
            </li>
            <li>
              <div className="list">
                <ul style={{ paddingLeft: "0" }}>
                  <li>
                    {customers.map((customer) => (
                      <a className="list-button" href="/customer/">
                        {customer.name}
                      </a>
                    ))}
                  </li>
                </ul>
              </div>
            </li>
          </ul>
        </div>
      </Block>
    </Page>
  );
};
