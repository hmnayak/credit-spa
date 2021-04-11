import React, { useState, useEffect } from "react";
import { Page, Block } from "framework7-react";
import { getCustomersPaginated } from "../services/custapi";

export const ListCustomersPage = (props) => {
  const currentPageToken = parseInt(props.f7route.query["page"]);
  const [customersListResponse, setCustomersList] = useState({ customers: [], totalSize: 0 });

  useEffect(async () => {
    const response = await getCustomersPaginated(props.fetch, currentPageToken);
    const content = await response.json();
    setCustomersList(content);
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
                    {customersListResponse.customers.map((customer) => (
                      <a
                        key={customer.customerid}
                        className="list-button"
                        href={"/customers/" + customer.customerid}
                      >
                        {customer.name}
                      </a>
                    ))}
                  </li>
                </ul>
              </div>
            </li>
            <li>
              <div >
                <ul style={{ paddingLeft: "0", listStyle: "none", textAlign: "right"}}>
                  {
                    !!(currentPageToken > 1) &&
                      <li style={{ display: "inline" }}>
                        <a style={{ margin: "5px" }}
                          href={"/customers/?page=" + (currentPageToken - 1)}>
                          prev
                        </a>
                      </li>
                  }
                  <li style={{ display: "inline" }}>
                    <label style={{ margin: "5px" }}>
                      {currentPageToken}
                    </label>
                  </li>
                  {
                    !!customersListResponse.nextpagetoken &&
                      <li style={{ display: "inline" }}>
                        <a style={{ margin: "5px" }}
                          href={"/customers/?page=" + (currentPageToken + 1)}>
                          next
                        </a>
                      </li>
                  }
                </ul>
              </div>
            </li>
          </ul>
        </div>
      </Block>
    </Page>
  );
};
