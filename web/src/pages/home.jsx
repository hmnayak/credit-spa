import React from "react";
import { Page, Block } from "framework7-react";

export default (props) => {
  const userInfo = () => {
    return <Block strong>You are logged in as {props.username()}</Block>;
  };

  return (
    <Page>
      <Block strong>
        <div className="row">
          <div className="card">
            <div className="card-content card-content-padding">
              customer billing
              <p>&nbsp;</p>
              <div className="card-footer">
                <a href="/customers/" className="link">
                  Customers!
                </a>
              </div>
            </div>
          </div>
        </div>
      </Block>
      {userInfo()}
    </Page>
  );
};
