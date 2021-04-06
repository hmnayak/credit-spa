import React from "react";
import { Page, Block } from "framework7-react";

export const HomePage = (props) => {

  return (
    <Page>
      <Block strong>
        <div className="row">
          <div className="card">
            <div className="card-content card-content-padding">
              <div>
                <a href="/about/" className="link">About</a>
                <br/>
                <a href="/customers/?page=1" className="link">Customers</a>
              </div>
            </div>
          </div>
        </div>
      </Block>
    </Page>
  );
};
