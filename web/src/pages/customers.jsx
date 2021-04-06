import React, { useState, useEffect } from "react";
import { Page, Block } from "framework7-react";
import { getCustomersPaginated } from "../services/custapi";

export const ListCustomersPage = (props) => {
  const [customers, setCustomers] = useState([]);
  const [pageCount, setPageCount] = useState([]);
  const [pageSize, setPageSize] = useState(3);
  useEffect(async () => {
    const response = await getCustomersPaginated(props.fetch, props.f7route.query["page"]);
    setCustomers(response.customers);
    if ('pageSize' in props) {
      setPageSize(props.pageSize);
    }
    setPageCount(response.totalsize % pageSize === 0 ? response.totalsize / pageSize : Math.floor(response.totalsize / pageSize) + 1);
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
                    {Array.from({length: pageCount}, (_, i) => ++i).map((pageNum) => (
                      <li style={{ display: "inline" }} key={pageNum}>
                        <a style={{ marginLeft: "10px" }}
                          href={"/customers/?page=" + pageNum}
                        >
                          {pageNum}
                        </a>
                      </li>
                    ))}
                </ul>
              </div>
            </li>
          </ul>
        </div>
      </Block>
    </Page>
  );
};
