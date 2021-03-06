import React from "react";
import { Page, Block } from "framework7-react";
import { getCurUser } from "../services/authsvc";

export default (props) => {
  const userInfo = () => {
    return <Block strong>You are logged in as {getCurUser()}</Block>;
  };

  return (
    <Page>
      <Block strong>
        <div className="row">
          <div className="col">
            <div className="card">
              <div className="card-content card-content-padding">
                Lorem ipsum dolor sit amet, consectetur adipiscing elit. Quisque
                ac diam ac quam euismod porta vel a nunc. Quisque sodales
                scelerisque est, at porta justo cursus ac. Lorem ipsum dolor sit
                amet, consectetur adipiscing elit. Quisque ac diam ac quam
                euismod porta vel a nunc. Quisque sodales scelerisque est, at
                porta justo cursus ac.
                <p>&nbsp;</p>
                <div className="card-footer">
                  <a href="#" className="link">
                    File tax!
                  </a>
                  <a href="/info" className="link">
                    Read more
                  </a>
                </div>
              </div>
            </div>
          </div>
          <div className="col">
            <div className="card">
              <div className="card-content card-content-padding">
                Lorem ipsum dolor sit amet, consectetur adipiscing elit. Quisque
                ac diam ac quam euismod porta vel a nunc. Quisque sodales
                scelerisque est, at porta justo cursus ac. Lorem ipsum dolor sit
                amet, consectetur adipiscing elit. Quisque ac diam ac quam
                euismod porta vel a nunc. Quisque sodales scelerisque est, at
                porta justo cursus ac.
                <p>&nbsp;</p>
                <div className="card-footer">
                  <a href="/bill" className="link">
                    Billing!
                  </a>
                  <a href="/info" className="link">
                    Read more
                  </a>
                </div>
              </div>
            </div>
          </div>
        </div>
      </Block>
      {userInfo()}
    </Page>
  );
};
