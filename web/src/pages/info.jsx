import React from "react";
import { Page, Link, Block, Button } from "framework7-react";

export default (props) => {
  return (
    <Page>
      <div>
        <h1>GSTR-1 - Return Filing, Format, Eligibility & Rules</h1>
        <div style={{ padding: 40, margin: 10, border: 2 }}>
          <p>In this article, we discuss the following topics in detail-</p>
          <ol>
            <li>
              <a href="#basics">Basics of GSTR </a>
              <ul>
                <li>
                  <a href="#what"> What is Gstr1</a>
                </li>
                <li>
                  <a href="#who"> Who has to pay Gstr1</a>
                </li>
                <li>
                  <a href="#how"> How to pay Gstr1</a>
                </li>
              </ul>
            </li>
            <li>
              <a href="#basics"> Pick our software to file the taxes </a>
            </li>
          </ol>
        </div>
      </div>
    </Page>
  );
};
