import React, { useState, useEffect } from "react";
import { Page, Block } from "framework7-react";
import { getItemsPaginated } from "../services/itemapi";

export const ListItemsPage = (props) => {
  const currentPageToken = parseInt(props.f7route.query["page"]);
  const [itemsListResponse, setItemsList] = useState({ items: [], totalSize: 0 });

  useEffect(async () => {
    const response = await getItemsPaginated(props.fetch, currentPageToken);
    const content = await response.json();
    setItemsList(content);
  }, []);

  return (
    <Page>
      <Block strong>
        <div className="list">
          <ul>
            <li>
              <a className="button" href="/items/new/">
                Add new item
              </a>
            </li>
            <li>
              <div className="list">
                <ul style={{ paddingLeft: "0" }}>
                  <li>
                    {itemsListResponse.items.map((item) => (
                      <a
                        key={item.itemid}
                        className="list-button"
                        href={"/items/" + item.itemid}
                      >
                        {item.name}
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
                          href={"/items/?page=" + (currentPageToken - 1)}>
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
                    !!itemsListResponse.nextpagetoken &&
                      <li style={{ display: "inline" }}>
                        <a style={{ margin: "5px" }}
                          href={"/items/?page=" + (currentPageToken + 1)}>
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
