import React, { useState, useEffect } from "react";
import { Page, Block } from "framework7-react";
import { getCustomersApi } from "../services/custapi";

export const ListCustomersPage = (props) => {
  const [customers, setCustomers] = useState([]);
  const [pageCount, setPageCount] = useState(1);
  const [pageSize, setPageSize] = useState(3);
  useEffect(async () => {
    const response = await getCustomersPaginated(props.fetch, props.f7route.query["page"]);
    const content = await response.json();
    setCustomers(content.customers);
    if ('pageSize' in props) {
      setPageSize(props.pageSize);
    }
    if (content.totalsize > 0)
    {
      setPageCount(content.totalsize % pageSize === 0 ? 
        content.totalsize / pageSize : 
        Math.floor(content.totalsize / pageSize) + 1);
    }
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
                        href={"/customers/" + customer.customerid}
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
