import React, { useState, useEffect } from "react";
import { Page, Block } from "framework7-react";
import { getCustomersPaginated } from "../services/custapi";

export const ListCustomersPage = (props) => {
  const [customers, setCustomers] = useState([]);
  const [pageSize, setPageSize] = useState(3);
  const [previousPageToken, setPreviousPageToken] = useState(0);
  const [currentPageToken, setCurrentPageToken] = useState(0);
  const [nextPageToken, setNextPageToken] = useState(0);

  useEffect(async () => {
    console.log("listcustomers");
    const currentPage = parseInt(props.f7route.query["page"]);
    setCurrentPageToken(currentPage);

    const response = await getCustomersPaginated(props.fetch, currentPage);
    const content = await response.json();
    setCustomers(content.customers);

    if ('pageSize' in props) {
      setPageSize(props.pageSize);
    }

    const numPages = content.totalsize % pageSize === 0 ? 
      content.totalsize / pageSize : 
      Math.floor(content.totalsize / pageSize) + 1;

    if (currentPage > 1) {
      setPreviousPageToken(currentPage - 1);
    }
    if (currentPage < numPages) {
      setNextPageToken(currentPage + 1);
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
            <li>
              <div >
                <ul style={{ paddingLeft: "0", listStyle: "none", textAlign: "right"}}>
                  {
                    !!previousPageToken &&
                      <li style={{ display: "inline" }}>
                        <a style={{ margin: "5px" }}
                          href={"/customers/?page=" + previousPageToken}>
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
                    !!nextPageToken &&
                      <li style={{ display: "inline" }}>
                        <a style={{ margin: "5px" }}
                          href={"/customers/?page=" + nextPageToken}>
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
