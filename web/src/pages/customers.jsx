import React, { useState, useEffect } from "react";
import { Page, Block } from "framework7-react";
import { getCustomersApi } from "../services/custapi";

export const ListCustomersPage = (props) => {
  const [customers, setCustomers] = useState([]);

  useEffect(() => {
    getCustomersApi(props.fetch).then((res) => {
      setCustomers(res);
    });
  }, []);

  return (
    <Page>
      <Block strong>
        <div className="list">
          <ul>
            <li>
              <a className="button" href="/customers/new/">
                Add new customer
              </a>
            </li>
            <li>
              <div className="list">
                <ul style={{ paddingLeft: "0" }}>
                  <li>
                    {customers.map((customer) => (
                      <a
                        key={customer.customerid}
                        className="list-button"
                        href="/customer/"
                      >
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
